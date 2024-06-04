package web_assembler

import (
	protocol_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/protocol"
	auth_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/auth"
	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	web_middleware "github.com/mariojuniortrab/hauling-api/internal/presentation/web/middleware"
)

type MiddlewareAssembler struct {
	validator protocol_application.Validator
	urlParser protocol_application.URLParser
	tokenizer protocol_application.Tokenizer
}

func NewMiddleWareAssembler(
	validator protocol_application.Validator,
	urlParser protocol_application.URLParser,
	tokenizer protocol_application.Tokenizer,
) *MiddlewareAssembler {
	return &MiddlewareAssembler{
		validator,
		urlParser,
		tokenizer,
	}
}

func (a *MiddlewareAssembler) GetAssembledProtectedMiddleware() protocol_application.Middleware {
	return a.assembleProtectedMiddleware().Middleware
}

func (a *MiddlewareAssembler) GetAssembledPaginatedMiddleware() protocol_application.Middleware {
	return a.assemblePaginatedMiddleware().Middleware
}

func (a *MiddlewareAssembler) GetAssembledUuidParserMiddleware() protocol_application.Middleware {
	return a.assembleUuidParserMiddleware().Middleware
}

func (a *MiddlewareAssembler) GetAssmbledBodyValidatorMiddleware(fieldsToValidate protocol_entity.Emptyable) protocol_application.Middleware {
	return a.assembleBodyValidationMiddleware(fieldsToValidate).Middleware
}

func (a *MiddlewareAssembler) assembleProtectedMiddleware() protocol_application.MiddlewareGetter {
	authUsecase := auth_usecase.NewAuthorization(a.tokenizer)
	return web_middleware.NewProtectedMiddleware(authUsecase)
}

func (a *MiddlewareAssembler) assemblePaginatedMiddleware() protocol_application.MiddlewareGetter {
	return web_middleware.NewPaginateMiddleware(a.validator, a.urlParser)
}

func (a *MiddlewareAssembler) assembleUuidParserMiddleware() protocol_application.MiddlewareGetter {
	return web_middleware.NewUuidParser(a.urlParser)
}

func (a *MiddlewareAssembler) assembleBodyValidationMiddleware(fieldsToValidate protocol_entity.Emptyable) protocol_application.MiddlewareGetter {
	return web_middleware.NewBodyValidator(fieldsToValidate)
}
