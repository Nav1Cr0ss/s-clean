package middleware

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	mwEcho "github.com/labstack/echo/v4/middleware"
)

var (
	DefaultCORSConfig = mwEcho.CORSConfig{
		//Skipper: func(c echo.Context) bool {
		//	return false
		//},
		//AllowOrigins: cfg.CORSOrigins,
		AllowOriginFunc: func(origin string) (bool, error) {
			if origin != "" {
				return true, nil
			}

			return false, nil
		},
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions,
		},
		AllowHeaders: []string{
			echo.HeaderAuthorization, echo.HeaderContentType,
		},
	}
	CORS             = mwEcho.CORSWithConfig(DefaultCORSConfig)
	PreflightOffCORS = preflightOFF(DefaultCORSConfig)
	CORSHandler      = CORS(func(c echo.Context) error { return c.NoContent(http.StatusNoContent) })
)

func preflightOFF(config mwEcho.CORSConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = mwEcho.DefaultCORSConfig.Skipper
	}
	if len(config.AllowOrigins) == 0 {
		config.AllowOrigins = mwEcho.DefaultCORSConfig.AllowOrigins
	}
	if len(config.AllowMethods) == 0 {
		config.AllowMethods = mwEcho.DefaultCORSConfig.AllowMethods
	}

	allowOriginPatterns := []string{}
	for _, origin := range config.AllowOrigins {
		pattern := regexp.QuoteMeta(origin)
		pattern = strings.Replace(pattern, "\\*", ".*", -1)
		pattern = strings.Replace(pattern, "\\?", ".", -1)
		pattern = "^" + pattern + "$"
		allowOriginPatterns = append(allowOriginPatterns, pattern)
	}

	allowMethods := strings.Join(config.AllowMethods, ",")
	allowHeaders := strings.Join(config.AllowHeaders, ",")
	exposeHeaders := strings.Join(config.ExposeHeaders, ",")
	maxAge := strconv.Itoa(config.MaxAge)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			origin := req.Header.Get(echo.HeaderOrigin)
			allowOrigin := ""

			preflight := req.Method == http.MethodOptions
			res.Header().Add(echo.HeaderVary, echo.HeaderOrigin)

			// No Origin provided
			if origin == "" {
				if !preflight {
					return next(c)
				}
				return c.NoContent(http.StatusNoContent)
			}

			if config.AllowOriginFunc != nil {
				allowed, err := config.AllowOriginFunc(origin)
				if err != nil {
					return err
				}
				if allowed {
					allowOrigin = origin
				}
			} else {
				// Check allowed origins
				for _, o := range config.AllowOrigins {
					if o == "*" && config.AllowCredentials {
						allowOrigin = origin
						break
					}
					if o == "*" || o == origin {
						allowOrigin = o
						break
					}
					if matchSubdomain(origin, o) {
						allowOrigin = origin
						break
					}
				}

				// Check allowed origin patterns
				for _, re := range allowOriginPatterns {
					if allowOrigin == "" {
						didx := strings.Index(origin, "://")
						if didx == -1 {
							continue
						}
						domAuth := origin[didx+3:]
						// to avoid regex cost by invalid long domain
						if len(domAuth) > 253 {
							break
						}

						if match, _ := regexp.MatchString(re, origin); match {
							allowOrigin = origin
							break
						}
					}
				}
			}

			// Origin not allowed
			if allowOrigin == "" {
				if !preflight {
					return next(c)
				}
				return c.NoContent(http.StatusNoContent)
			}

			// Simple request
			if !preflight {
				res.Header().Set(echo.HeaderAccessControlAllowOrigin, allowOrigin)
				if config.AllowCredentials {
					res.Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
				}
				if exposeHeaders != "" {
					res.Header().Set(echo.HeaderAccessControlExposeHeaders, exposeHeaders)
				}
				return next(c)
			}

			// Preflight request
			res.Header().Add(echo.HeaderVary, echo.HeaderAccessControlRequestMethod)
			res.Header().Add(echo.HeaderVary, echo.HeaderAccessControlRequestHeaders)
			res.Header().Set(echo.HeaderAccessControlAllowOrigin, allowOrigin)
			res.Header().Set(echo.HeaderAccessControlAllowMethods, allowMethods)
			if config.AllowCredentials {
				res.Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
			}
			if allowHeaders != "" {
				res.Header().Set(echo.HeaderAccessControlAllowHeaders, allowHeaders)
			} else {
				h := req.Header.Get(echo.HeaderAccessControlRequestHeaders)
				if h != "" {
					res.Header().Set(echo.HeaderAccessControlAllowHeaders, h)
				}
			}
			if config.MaxAge > 0 {
				res.Header().Set(echo.HeaderAccessControlMaxAge, maxAge)
			}
			return next(c)
		}
	}
}
