package di

import (
	"order-service/service/http/handler"
	"order-service/service/http/router"
	"order-service/service/repository"
	"order-service/service/usecase"

	"go.uber.org/fx"
)

// Module provided to fx
var Module = fx.Provide(
	// provide usecase
	ProvideUserUsecase,
	provideOrderUsecase,
	// handler
	provideUserHandler,
	provideProductHandler,
	// router
	provodeUserRouter,
	provideOrderRouter,

	// provide repositories
	provideUserRepo,
	provideProducRepo,
)

// Handler
func provideUserHandler(userUsecase usecase.IuserUsecase) handler.UserHandler {
	return handler.NewUserhandler(userUsecase)
}

func provideProductHandler(
	productUsecase usecase.IproductUsecase,
) handler.ProductHandler {
	return handler.NewProductHandler(productUsecase)
}

// Router
func provodeUserRouter(
	userHandlerr handler.UserHandler,
) router.UserRouter {
	return router.NewUserRouter(userHandlerr)
}
func provideOrderRouter(
	productHandler handler.ProductHandler,
) router.ProductRouter {
	return router.NewProductRouter(productHandler)
}

// Usecases
func ProvideUserUsecase(
	userRepo repository.IuserRepo,
) usecase.IuserUsecase {
	return usecase.NewUserUsecase(userRepo)
}

func provideOrderUsecase(
	productRepo repository.IproductRepo,
) usecase.IproductUsecase {
	return usecase.NewProductUsecase(productRepo)
}

// Repositories
func provideUserRepo() repository.IuserRepo {
	return repository.NewUserRepo()
}

// Repositories
func provideProducRepo() repository.IproductRepo {
	return repository.NewProductRepo()
}
