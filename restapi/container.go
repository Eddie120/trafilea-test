package restapi

import (
	"api.trafilea.com/controllers"
	"api.trafilea.com/services/cart"
	"github.com/patrickmn/go-cache"
	"go.uber.org/dig"
	"time"
)

func buildContainer() (container *dig.Container) {
	container = dig.New()

	provide := func(constructor interface{}, opts ...dig.ProvideOption) {
		err := container.Provide(constructor, opts...)
		if err != nil {
			panic(err)
		}
	}

	provide(func() cache.Cache {
		return *cache.New(24*time.Hour, 24*time.Hour)
	})

	provide(cart.NewShoppingCartService)

	provide(func(svc *cart.ServiceImpl) cart.ShoppingCartService {
		return svc
	})

	provide(controllers.NewShoppingController)

	return container
}
