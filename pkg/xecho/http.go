package xecho

import (
	"context"
	"errors"
	logger "github.com/Nav1Cr0ss/s-user-storage/pkg/logger/zerolog"
	"github.com/Nav1Cr0ss/s-user-storage/pkg/xecho/middleware"
	"github.com/labstack/echo/v4"
	mwEcho "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"reflect"
	//libC "github.com/acretrader/lib/configurator"
	//"github.com/acretrader/lib/logger"
	//"github.com/acretrader/lib/server"
	//"github.com/acretrader/lib/server/xecho/healthcheck"
	//"github.com/acretrader/lib/server/xecho/helpers"
	//"github.com/acretrader/lib/server/xecho/middleware"
	//"github.com/labstack/echo-contrib/prometheus"
	//"github.com/labstack/echo/v4"
	//mwEcho "github.com/labstack/echo/v4/middleware"
	//"github.com/labstack/gommon/log"
	//echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/fx"
)

type HTTPServerConfig interface {
	GetServerInitString() string
}

type HttpServer struct {
	Echo     *echo.Echo
	log      logger.Logger
	c        HTTPServerConfig
	authTool AuthTool
}

type AuthTool interface {
	JWTValidation() echo.MiddlewareFunc
	OptionalJWTValidation() echo.MiddlewareFunc
}

type Route struct {
	Method        string
	Path          string
	Handle        echo.HandlerFunc
	NoCORS        bool
	NoCORSHandler bool
	NoAuth        bool
	NoBasePath    bool
	OptionalAuth  bool
	Middlewares   []echo.MiddlewareFunc
}

type Handler interface {
	GetRoutes() []Route
}

type HandlerResult struct {
	fx.Out

	Handler Handler `group:"server"`
}

type ServerParams struct {
	fx.In

	Handlers []Handler `group:"server"`
}

func (s *HttpServer) InitRoutes(h Handler) {
	for _, route := range h.GetRoutes() {
		s.AddRoute(route)
	}
}

// AddRoute adds route to the router.
func (s *HttpServer) AddRoute(route Route) {
	var (
		handle      = route.Handle
		middlewares = []echo.MiddlewareFunc{}
		//versionedPath string
		//err           error
	)

	//if !route.NoBasePath {
	//	versionedPath, err = helpers.AddAdditionalBasePath(route.Path)
	//	if err != nil {
	//		s.log.Error().Err(err).Msgf(err.Error())
	//	}
	//}

	if !route.NoCORS {
		middlewares = append(middlewares, middleware.CORS)
	}

	//if !route.NoCORSHandler {
	//	s.Echo.OPTIONS(route.Path, middleware.CORSHandler)
	//
	//	if versionedPath != "" {
	//		s.Echo.OPTIONS(versionedPath, middleware.CORSHandler)
	//	}
	//}

	if s.authTool != nil {
		if !route.NoAuth && !route.OptionalAuth {
			middlewares = append(middlewares, s.authTool.JWTValidation())
		}

		if route.OptionalAuth {
			middlewares = append(middlewares, s.authTool.OptionalJWTValidation())
		}
	}

	//if route.Method != http.MethodOptions {
	//	middlewares = append(middlewares, middleware.RequestLogger)
	//}

	middlewares = append(middlewares, route.Middlewares...)
	//middlewares = append(route.Middlewares, middlewares...)
	s.Echo.Add(route.Method, route.Path, handle, middlewares...)

	//if versionedPath != "" {
	//	s.Echo.Add(route.Method, versionedPath, handle, middlewares...)
	//}
}

// SetAuthMiddleware sets auth middleware to the router.
func (s *HttpServer) SetAuthMiddleware(authTool AuthTool) {
	s.authTool = authTool
}

// NewHttpServer returns new API instance.
func NewHttpServer(c HTTPServerConfig) *HttpServer {
	var (
		e       = echo.New()
		echoLog = logger.NewDefaultComponentLogger("http server")
	)

	// sets CORS headers if Origin is present
	e.Use(
		mwEcho.CORSWithConfig(mwEcho.CORSConfig{
			Skipper: func(c echo.Context) bool {
				return true
			},
			//AllowOrigins: cfg.CORSOrigins,
			AllowOriginFunc: func(origin string) (bool, error) {
				if origin != "" {
					return true, nil
				}

				return false, nil
			},
			AllowMethods: []string{
				http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete,
			},
			AllowHeaders: []string{
				echo.HeaderAuthorization, echo.HeaderContentType,
			},
		}),
	)
	// Set context logger
	//e.Use(middleware.SetContextLogger)
	//// Set env name
	//e.Use(middleware.NewEnvSetter(cfg.EnvName).SetEnvName())
	//// Trace ID middleware generates a unique id for a request.
	//e.Use(middleware.SetTraceID)
	//// Ip address ID middleware set user's ip address to context.
	//e.Use(middleware.SetIPAddress)
	//// Request logger middleware
	//e.Use(middleware.RequestLogger)
	//// Dump request & response body
	e.Use(middleware.DumpBody)

	// Add the healthcheck endpoint
	//e.GET(`/healthcheck`, healthcheck.Healthcheck)

	//statusPath, err := helpers.AddAdditionalBasePath("/status")
	//if err != nil {
	//	statusPath = "/status"
	//}
	//
	//docsPath, err := helpers.AddAdditionalBasePath("/docs")
	//if err != nil {
	//	statusPath = "/docs"
	//}
	//
	//e.GET(statusPath, healthcheck.Status)

	// avoid any native logging of echo, because we use custom library for logging
	e.HideBanner = true        // don't log the banner on startup
	e.HidePort = true          // hide log about port server started on
	e.Logger.SetLevel(log.OFF) // disable echo#Logger

	//if cfg.Debug {
	//	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	//		return func(c echo.Context) error {
	//			switch c.Request().URL.Path {
	//			case "/docs", "/docs/", "/docs/index.html", docsPath, docsPath + "/":
	//				return c.Redirect(http.StatusTemporaryRedirect, docsPath+"/index.html")
	//			}
	//
	//			return next(c)
	//		}
	//	})
	//
	//	e.GET(docsPath+"/*", echoSwagger.WrapHandler, mwEcho.BasicAuth(
	//		func(username, password string, c echo.Context) (bool, error) {
	//			if username == "dev" && password == "acredev" {
	//				return true, nil
	//			}
	//
	//			return false, nil
	//		}))
	//
	//	e.Debug = true
	//}

	//e.Use(middleware.RequestLogger)

	// Add prometheus metrics
	//p := prometheus.NewPrometheus("echo", nil)
	//p.Use(e)

	return &HttpServer{
		Echo: e,
		c:    c,
		log:  echoLog,
	}
}

func InitHandlers(srv *HttpServer, p ServerParams) {
	srv.log.Info().Interface("handlers", p.Handlers).Msg("Initializing handlers")

	for _, handler := range p.Handlers {
		t := reflect.TypeOf(handler).Elem().PkgPath()

		srv.log.Info().Msgf("Initializing handler: %s", t)

		srv.InitRoutes(handler)
	}
}

// SetAuthMiddleware is function which set auth middleware
func SetAuthMiddleware(srv *HttpServer, authMiddleware AuthTool) {
	srv.SetAuthMiddleware(authMiddleware)
}

// StartServer is function that registers start of http server in lifecycle
func StartServer(lc fx.Lifecycle, srv *HttpServer) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				on := srv.c.GetServerInitString()

				srv.log.Info().Msgf("starting server on %s", on)

				go func() {
					err := srv.Echo.Start(on)

					if errors.Is(err, http.ErrServerClosed) {
						srv.log.Info().Err(err).Msgf("server %s stopped", on)
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				err := srv.Echo.Shutdown(ctx)
				if err != nil {
					srv.log.Info().Err(err).Msg("couldn't stop server")
				}

				return nil
			},
		},
	)
}
