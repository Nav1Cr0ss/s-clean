package ports

import (
	"github.com/Nav1Cr0ss/s-user-storage/internal/ports/user"
	"github.com/Nav1Cr0ss/s-user-storage/pkg/xecho"
	"go.uber.org/fx"
)

func NewServer() fx.Option {
	//xechoHelpers.SetAPIBasePath("/v1/user")

	return fx.Options(
		fx.Provide(
			//jwt.New,
			// bind jwt middleware to the auth tool interface
			//func(middleware *jwt.Middleware) xecho.AuthTool { return middleware },
			// Init data validator
			//validator.New,
			// Init handlers
			//validate.NewHandler,
			//temptoken.NewHandler,
			//signup.NewHandler,
			//debug.NewHandler,
			//webhook.NewHandler,
			//login.NewHandler,
			//profile.NewHandler,
			//portfolio.NewHandler,
			//other.NewHandler,
			user.NewHandler,
			//wallet.NewHandler,
			//admin.NewHandler,
			//beneficiary.NewHandler,
			//advisor.NewHandler,
			//// Init http server
			xecho.NewHttpServer,
		),
		//fx.Invoke(
		//	// setup middlewares
		//	xecho.SetAuthMiddleware,
		//	// setup validator
		//	validator.RegisterExistingValidator,
		//	// register custom validation rules
		//	func(srv *xecho.HttpServer, vl *validator.Validator) error {
		//		err := vl.RegisterValidation(
		//			"dwolla_zipcode",
		//			validators.IsValidZipCode,
		//			"Invalid zip code",
		//		)
		//		if err != nil {
		//			return err
		//		}
		//		err = vl.RegisterValidation(
		//			"required_ein",
		//			validators.ValidateRequiredEIN,
		//			"The field is required",
		//		)
		//		if err != nil {
		//			return err
		//		}
		//		err = vl.RegisterValidation(
		//			"required_entity_name",
		//			validators.ValidateRequiredEntityName,
		//			"The field is required",
		//		)
		//		if err != nil {
		//			return err
		//		}
		//		err = vl.RegisterValidation(
		//			"required_title_entity",
		//			validators.ValidateRequiredTitleEntity,
		//			"The field is required",
		//		)
		//		if err != nil {
		//			return err
		//		}
		//
		//		return vl.RegisterValidation(
		//			"beneficiary_subtype",
		//			validators.IsValidTypeSubType,
		//			"The selected subtype does not match the type",
		//		)
		//	},
		//	// Registration routes and handlers for http server
		//	xecho.InitHandlers,
		//	// set log function in context
		//	func(repo adapters.Repository, srv *xecho.HttpServer) {
		//		srv.Echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		//			return func(c echo.Context) error {
		//				c.SetRequest(c.Request().WithContext(keys.SetLogFunc(c.Request().Context(), repo.LogRequest)))
		//
		//				return next(c)
		//			}
		//		})
		//	},
		//	// Run HTTP server

		fx.Invoke(
			xecho.InitHandlers,
			xecho.StartServer,
		),
		//),
	)
}
