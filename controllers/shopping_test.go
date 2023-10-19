package controllers

import (
	"api.trafilea.com/models"
	"api.trafilea.com/restapi/operations"
	"api.trafilea.com/services/cart"
	"api.trafilea.com/services/cart/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShoppingController_CreateCart_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShoppingService := mocks.NewMockShoppingCartService(ctrl)
	controller := NewShoppingController(mockShoppingService)

	params := operations.CreateCartParams{
		UserID: "",
	}

	result := controller.CreateCart(params)
	assert.IsType(t, &operations.CreateCartBadRequest{}, result)
}

func TestShoppingController_CreateCart_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShoppingService := mocks.NewMockShoppingCartService(ctrl)
	mockShoppingService.EXPECT().CreateShoppingCart(gomock.Any()).Return("CART-12345")

	controller := NewShoppingController(mockShoppingService)

	params := operations.CreateCartParams{
		UserID: "123456789",
	}

	result := controller.CreateCart(params)
	assert.IsType(t, &operations.CreateCartOK{}, result)
}

func TestShoppingController_AddProductToCart_CartNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShoppingService := mocks.NewMockShoppingCartService(ctrl)
	mockShoppingService.EXPECT().AddProductToCart(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, cart.ErrCartNotFound)

	controller := NewShoppingController(mockShoppingService)

	params := operations.AddNewProductParams{
		ProductInput: &models.ProductInput{},
	}

	result := controller.AddProductToCart(params)
	assert.IsType(t, &operations.AddNewProductNotFound{}, result)
}

func TestShoppingController_AddProductToCart_InvalidCartID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShoppingService := mocks.NewMockShoppingCartService(ctrl)
	mockShoppingService.EXPECT().AddProductToCart(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, cart.ErrInvalidCartId)

	controller := NewShoppingController(mockShoppingService)

	params := operations.AddNewProductParams{
		ProductInput: &models.ProductInput{},
	}

	result := controller.AddProductToCart(params)
	assert.IsType(t, &operations.AddNewProductBadRequest{}, result)
}

func TestShoppingController_AddProductToCart_InternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShoppingService := mocks.NewMockShoppingCartService(ctrl)
	mockShoppingService.EXPECT().AddProductToCart(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

	controller := NewShoppingController(mockShoppingService)

	params := operations.AddNewProductParams{
		ProductInput: &models.ProductInput{},
	}

	result := controller.AddProductToCart(params)
	assert.IsType(t, &operations.AddNewProductInternalServerError{}, result)
}

func TestShoppingController_AddProductToCart_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShoppingService := mocks.NewMockShoppingCartService(ctrl)

	newProductList := models.ListProducts{
		&models.Product{},
	}

	mockShoppingService.EXPECT().AddProductToCart(gomock.Any(), gomock.Any(), gomock.Any()).Return(newProductList, nil)

	controller := NewShoppingController(mockShoppingService)

	params := operations.AddNewProductParams{
		ProductInput: &models.ProductInput{},
	}

	result := controller.AddProductToCart(params)
	assert.IsType(t, &operations.AddNewProductOK{}, result)
}

func TestShoppingController_UpdateQuantity_CartNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShoppingService := mocks.NewMockShoppingCartService(ctrl)

	mockShoppingService.EXPECT().UpdateQuantity(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, cart.ErrCartNotFound)
	controller := NewShoppingController(mockShoppingService)

	params := operations.ModifyQuantityParams{
		UpdateQuantity: &models.UpdateQuantity{},
	}

	result := controller.UpdateQuantity(params)
	assert.IsType(t, &operations.ModifyQuantityNotFound{}, result)
}

func TestShoppingController_UpdateQuantity_InvalidCartID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShoppingService := mocks.NewMockShoppingCartService(ctrl)

	mockShoppingService.EXPECT().UpdateQuantity(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, cart.ErrInvalidCartId)
	controller := NewShoppingController(mockShoppingService)

	params := operations.ModifyQuantityParams{
		UpdateQuantity: &models.UpdateQuantity{},
	}

	result := controller.UpdateQuantity(params)
	assert.IsType(t, &operations.ModifyQuantityBadRequest{}, result)
}

func TestShoppingController_UpdateQuantity_InternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShoppingService := mocks.NewMockShoppingCartService(ctrl)

	mockShoppingService.EXPECT().UpdateQuantity(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, assert.AnError)
	controller := NewShoppingController(mockShoppingService)

	params := operations.ModifyQuantityParams{
		UpdateQuantity: &models.UpdateQuantity{},
	}

	result := controller.UpdateQuantity(params)
	assert.IsType(t, &operations.ModifyQuantityInternalServerError{}, result)
}

func TestShoppingController_UpdateQuantity_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShoppingService := mocks.NewMockShoppingCartService(ctrl)

	product := &models.Product{}

	mockShoppingService.EXPECT().UpdateQuantity(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(product, nil)
	controller := NewShoppingController(mockShoppingService)

	params := operations.ModifyQuantityParams{
		UpdateQuantity: &models.UpdateQuantity{},
	}

	result := controller.UpdateQuantity(params)
	assert.IsType(t, &operations.ModifyQuantityOK{}, result)
}

func TestShoppingController_CreateOrder_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShoppingService := mocks.NewMockShoppingCartService(ctrl)

	mockShoppingService.EXPECT().CreateOrder(gomock.Any()).Return(nil, cart.ErrMissingUserID)
	controller := NewShoppingController(mockShoppingService)

	params := operations.CreateOrderParams{}

	result := controller.CreateOrder(params)
	assert.IsType(t, &operations.CreateOrderBadRequest{}, result)
}

func TestShoppingController_CreateOrder_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShoppingService := mocks.NewMockShoppingCartService(ctrl)

	mockShoppingService.EXPECT().CreateOrder(gomock.Any()).Return(nil, cart.ErrCartNotFound)
	controller := NewShoppingController(mockShoppingService)

	params := operations.CreateOrderParams{}

	result := controller.CreateOrder(params)
	assert.IsType(t, &operations.CreateOrderNotFound{}, result)
}

func TestShoppingController_CreateOrder_InternalServerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShoppingService := mocks.NewMockShoppingCartService(ctrl)

	mockShoppingService.EXPECT().CreateOrder(gomock.Any()).Return(nil, assert.AnError)
	controller := NewShoppingController(mockShoppingService)

	params := operations.CreateOrderParams{}

	result := controller.CreateOrder(params)
	assert.IsType(t, &operations.CreateOrderInternalServerError{}, result)
}

func TestShoppingController_CreateOrder_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShoppingService := mocks.NewMockShoppingCartService(ctrl)
	newOrder := &models.OrderResponse{}

	mockShoppingService.EXPECT().CreateOrder(gomock.Any()).Return(newOrder, nil)
	controller := NewShoppingController(mockShoppingService)

	params := operations.CreateOrderParams{}

	result := controller.CreateOrder(params)
	assert.IsType(t, &operations.CreateOrderOK{}, result)
}
