package user_routes

import (
	auth_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/auth"
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	web_assembler "github.com/mariojuniortrab/hauling-api/internal/presentation/web/assembler"
)

type router struct {
	userAssembler       *web_assembler.UserAssembler
	middlewareAssembler *web_assembler.MiddlewareAssembler
}

func NewRouter(
	userAssembler *web_assembler.UserAssembler,
	middlewareAssembler *web_assembler.MiddlewareAssembler,
) *router {
	return &router{
		userAssembler,
		middlewareAssembler,
	}
}

func (r *router) Route(route protocol_application.Router) protocol_application.Router {

	r.routeLogin(route)
	r.routeSignup(route)

	route.Route("/user", func(rr protocol_application.Router) {
		protected := r.middlewareAssembler.GetAssembledProtectedMiddleware()
		rr.Use(protected)

		r.routeListUser(rr)
		r.routeCreateUser(rr)

		rr.Group(func(rrr protocol_application.Router) {
			uuidParser := r.middlewareAssembler.GetAssembledUuidParserMiddleware()
			rrr.Use(uuidParser)

			r.routeGetUser(rrr)
			r.routeDeleteUser(rrr)
			r.routeUpdateUser(rrr)
		})
	})

	return route
}

func (r *router) routeSignup(route protocol_application.Router) {
	var signupInputDto auth_entity.SignupInputDto
	bodyValidator := r.middlewareAssembler.GetAssmbledBodyValidatorMiddleware(&signupInputDto)
	route.With(bodyValidator).Post("/signup", r.userAssembler.GetAssembledSignupHandle())
}

func (r *router) routeLogin(route protocol_application.Router) {
	var loginInputDto auth_entity.LoginInputDto
	bodyValidator := r.middlewareAssembler.GetAssmbledBodyValidatorMiddleware(&loginInputDto)
	route.With(bodyValidator).Post("/login", r.userAssembler.GetAssembledLoginHandle())
}

func (r *router) routeListUser(route protocol_application.Router) {
	paginate := r.middlewareAssembler.GetAssembledPaginatedMiddleware()
	route.With(paginate).Get("/", r.userAssembler.GetAssembledListUserHandle())
}

func (r *router) routeGetUser(route protocol_application.Router) {
	route.Get("/{id}", r.userAssembler.GetAssembledDetailUserHandle())
}

func (r *router) routeDeleteUser(route protocol_application.Router) {
	route.Delete("/{id}", r.userAssembler.GetAssembledRemoveUserHandle())
}

func (r *router) routeUpdateUser(route protocol_application.Router) {
	var userUpdateInputDto user_entity.UserUpdateInputDto
	updateBodyValidation := r.middlewareAssembler.GetAssmbledBodyValidatorMiddleware(&userUpdateInputDto)
	route.With(updateBodyValidation).Patch("/{id}", r.userAssembler.GetAssembledUpdateUserHandle())
}

func (r *router) routeCreateUser(route protocol_application.Router) {
	var userCreateInputDto user_entity.CreateUserInputDto
	createBodyValidation := r.middlewareAssembler.GetAssmbledBodyValidatorMiddleware(&userCreateInputDto)
	route.With(createBodyValidation).Post("/", r.userAssembler.GetAssembledCreateUserHandle())
}
