package cart

import (
	"api.trafilea.com/models"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestServiceImpl_CreateShoppingCart(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	cartId := svc.CreateShoppingCart("fb0eee40-6d4b-11ee-b962-0242ac120002")
	assert.True(t, strings.HasPrefix(cartId, "CART"))
}

func TestServiceImpl_AddProductToCart_Success(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"

	cartId := svc.CreateShoppingCart(userId)
	productInput := models.ProductInput{
		ProductID: "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002",
		Quantity:  5,
	}

	productList, err := svc.AddProductToCart(userId, cartId, productInput)
	assert.Nil(t, err)
	assert.Len(t, productList, 1)
}

func TestServiceImpl_AddProductToCart_InvalidCartID(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"

	_ = svc.CreateShoppingCart(userId)
	productInput := models.ProductInput{
		ProductID: "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002",
		Quantity:  5,
	}

	_, err := svc.AddProductToCart(userId, "fakeCartId", productInput)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrInvalidCartId)
}

func TestServiceImpl_AddProductToCart_ProductNotFound(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"

	cartId := svc.CreateShoppingCart(userId)
	productInput := models.ProductInput{
		ProductID: "PROD-9eabaf4a-6d54-11ee-b962-0242ac12002",
		Quantity:  5,
	}

	_, err := svc.AddProductToCart(userId, cartId, productInput)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrProductNotFound)
}

func TestServiceImpl_AddProductToCart_MissingUserID(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"

	cartId := svc.CreateShoppingCart(userId)
	productInput := models.ProductInput{
		ProductID: "PROD-9eabaf4a-6d54-11ee-b962-0242ac12002",
		Quantity:  5,
	}

	_, err := svc.AddProductToCart("", cartId, productInput)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrMissingUserID)
}

func TestServiceImpl_AddProductToCart_MissingCartID(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"
	productInput := models.ProductInput{
		ProductID: "PROD-9eabaf4a-6d54-11ee-b962-0242ac12002",
		Quantity:  5,
	}

	_, err := svc.AddProductToCart(userId, "", productInput)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrMissingCartID)
}

func TestServiceImpl_AddProductToCart_CartNotFound(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac12000"
	cartId := svc.CreateShoppingCart(userId)

	productInput := models.ProductInput{
		ProductID: "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002",
		Quantity:  5,
	}

	_, err := svc.AddProductToCart("fakeUserID", cartId, productInput)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrCartNotFound)
}

func TestServiceImpl_AddProductToCart_InvalidQuantity(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"

	cartId := svc.CreateShoppingCart(userId)
	productInput := models.ProductInput{
		ProductID: "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002",
		Quantity:  0,
	}

	_, err := svc.AddProductToCart(userId, cartId, productInput)
	assert.NotNil(t, err)
	assert.Error(t, err, ErrInvalidQuantity)
}

func TestServiceImpl_AddProductToCart_II(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"

	cartId := svc.CreateShoppingCart(userId)
	newProductList := models.ListProducts{
		&models.Product{
			ProductID: "PROD-9eabae28-6d54-11ee-b962-0242ac120002",
			Quantity:  10,
		},
	}

	globalStore.Set(cartId, newProductList, cache.NoExpiration)

	productInput := models.ProductInput{
		ProductID: "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002",
		Quantity:  5,
	}

	productList, err := svc.AddProductToCart(userId, cartId, productInput)
	assert.Nil(t, err)
	assert.Len(t, productList, 2)
}

func TestServiceImpl_AddProductToCart_UpdateQuantity(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"

	cartId := svc.CreateShoppingCart(userId)
	newProductList := models.ListProducts{
		&models.Product{
			ProductID: "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002",
			Quantity:  10,
		},
	}

	globalStore.Set(cartId, newProductList, cache.NoExpiration)

	productInput := models.ProductInput{
		ProductID: "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002",
		Quantity:  5,
	}

	productList, err := svc.AddProductToCart(userId, cartId, productInput)
	assert.Nil(t, err)
	assert.Len(t, productList, 1)
}

func TestServiceImpl_UpdateQuantity_MissingUserID(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"
	productId := "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002"

	cartId := svc.CreateShoppingCart(userId)
	payload := models.UpdateQuantity{
		Quantity: 5,
	}

	_, err := svc.UpdateQuantity("", cartId, productId, payload)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrMissingUserID)
}

func TestServiceImpl_UpdateQuantity_MissingCartID(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"
	productId := "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002"

	payload := models.UpdateQuantity{
		Quantity: 5,
	}

	_, err := svc.UpdateQuantity(userId, "", productId, payload)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrMissingCartID)
}

func TestServiceImpl_UpdateQuantity_CartNotFound(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"
	productId := "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002"

	cartId := svc.CreateShoppingCart(userId)
	payload := models.UpdateQuantity{
		Quantity: 5,
	}

	_, err := svc.UpdateQuantity("fakeUserID", cartId, productId, payload)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrCartNotFound)
}

func TestServiceImpl_UpdateQuantity_InvalidCartID(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"
	productId := "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002"

	_ = svc.CreateShoppingCart(userId)
	payload := models.UpdateQuantity{
		Quantity: 5,
	}

	_, err := svc.UpdateQuantity(userId, "fakeCartID", productId, payload)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrInvalidCartId)
}

func TestServiceImpl_UpdateQuantity_CartEmpty(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"
	productId := "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002"

	cartId := svc.CreateShoppingCart(userId)
	payload := models.UpdateQuantity{
		Quantity: 5,
	}

	_, err := svc.UpdateQuantity(userId, cartId, productId, payload)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrCartEmpty)
}

func TestServiceImpl_UpdateQuantity_ProductNotFound(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"
	productId := "PROD-9eabaf4a-6d54-11ee-b962-0242ac12000"

	cartId := svc.CreateShoppingCart(userId)

	oldProductList := models.ListProducts{
		&models.Product{
			ProductID: "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002",
			Quantity:  10,
		},
	}

	globalStore.Set(cartId, oldProductList, cache.NoExpiration)

	payload := models.UpdateQuantity{
		Quantity: 5,
	}

	_, err := svc.UpdateQuantity(userId, cartId, productId, payload)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrProductNotFound)
}

func TestServiceImpl_UpdateQuantity_Success(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"
	productId := "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002"

	cartId := svc.CreateShoppingCart(userId)

	oldProductList := models.ListProducts{
		&models.Product{
			ProductID: "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002",
			Quantity:  10,
		},
	}

	globalStore.Set(cartId, oldProductList, cache.NoExpiration)

	payload := models.UpdateQuantity{
		Quantity: 5,
	}

	product, err := svc.UpdateQuantity(userId, cartId, productId, payload)
	assert.Nil(t, err)
	assert.Equal(t, float64(5), product.Quantity)
}

func TestServiceImpl_CreateOrder_MissingUserId(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := ""

	_, err := svc.CreateOrder(userId)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrMissingUserID)
}

func TestServiceImpl_CreateOrder_CartNotFound(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"

	_, err := svc.CreateOrder(userId)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrCartNotFound)
}

func TestServiceImpl_CreateOrder_CartEmpty(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"
	_ = svc.CreateShoppingCart(userId)

	_, err := svc.CreateOrder(userId)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrCartEmpty)
}

func TestServiceImpl_CreateOrder_CartEmpty_I(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"
	cartId := svc.CreateShoppingCart(userId)

	oldProductList := models.ListProducts{}
	globalStore.Set(cartId, oldProductList, cache.NoExpiration)

	_, err := svc.CreateOrder(userId)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrCartEmpty)
}

func TestServiceImpl_CreateOrder_Success(t *testing.T) {
	var globalStore = cache.New(5*time.Minute, 10*time.Minute)
	svc := NewShoppingCartService(*globalStore)

	userId := "fb0eee40-6d4b-11ee-b962-0242ac120002"
	cartId := svc.CreateShoppingCart(userId)

	oldProductList := models.ListProducts{
		&models.Product{
			Category: "Coffee",
			Name:     "XX",
			Price:    float64(5),
			Quantity: 10,
		},
		&models.Product{
			Category: "Accessories",
			Name:     "Y",
			Price:    float64(5),
			Quantity: 10,
		},
		&models.Product{
			Category: "Accessories",
			Name:     "B",
			Price:    float64(5),
			Quantity: 21,
		},
		&models.Product{
			Category: "Accessories",
			Name:     "Z",
			Price:    float64(5),
			Quantity: 1,
		},
		&models.Product{
			Category: "Coffee",
			Name:     "A",
			Price:    float64(5),
			Quantity: 21,
		},
		&models.Product{
			Category: "Equipment",
			Name:     "tttt",
			Price:    float64(5),
			Quantity: 1,
		},
		&models.Product{
			Category: "Equipment",
			Name:     "qqqq",
			Price:    float64(5),
			Quantity: 1,
		},
		&models.Product{
			Category: "Equipment",
			Name:     "wwwww",
			Price:    float64(5),
			Quantity: 1,
		},
	}

	globalStore.Set(cartId, oldProductList, cache.NoExpiration)

	order, err := svc.CreateOrder(userId)
	assert.Nil(t, err)
	assert.Equal(t, float64(33), order.Totals.Discounts)
	assert.Equal(t, float64(300), order.Totals.Order)
	assert.Equal(t, float64(330), order.Totals.Products)
	assert.Equal(t, float64(3), order.Totals.Shipping)

	items, _ := globalStore.Get(cartId)
	productList := items.(models.ListProducts)
	assert.Len(t, productList, 9)
}
