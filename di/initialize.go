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
	provideCartUsecase,
	// handler
	provideUserHandler,
	provideProductHandler,
	provideCartHandler,
	// router
	provodeUserRouter,
	provideOrderRouter,
	provideCartRouter,

	// provide repositories
	provideUserRepo,
	provideProducRepo,
	provideCartRepo,
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

func provideCartHandler(
	cartUsecase usecase.IcartUsercase,
) handler.CartHandler {
	return handler.NewCartHandler(cartUsecase)
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

func provideCartRouter(
	cartHandler handler.CartHandler,
) router.CartRouter {
	return router.NewCartRouter(cartHandler)
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

func provideCartUsecase(
	cartRepo repository.IcartRepo,
	productRepo repository.IproductRepo,
) usecase.IcartUsercase {
	return usecase.NewCartUsercase(cartRepo, productRepo)
}

// Repositories
func provideUserRepo() repository.IuserRepo {
	return repository.NewUserRepo()
}

func provideCartRepo() repository.IcartRepo {
	return repository.NewCartRepo()
}

// Repositories
func provideProducRepo() repository.IproductRepo {
	return repository.NewProductRepo()
}
