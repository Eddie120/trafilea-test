package controllers

import (
	"api.trafilea.com/models"
	"api.trafilea.com/restapi/operations"
	"api.trafilea.com/services/cart"
	"github.com/go-openapi/runtime/middleware"
	"strings"
)

type ShoppingController struct {
	cartService cart.ShoppingCartService
}

func NewShoppingController(cartService cart.ShoppingCartService) *ShoppingController {
	return &ShoppingController{
		cartService: cartService,
	}
}

func (controller *ShoppingController) CreateCart(params operations.CreateCartParams) middleware.Responder {
	if strings.TrimSpace(params.UserID) == "" {
		return operations.NewCreateCartBadRequest().WithPayload("missing user id")
	}

	cartID := controller.cartService.CreateShoppingCart(params.UserID)

	return operations.NewCreateCartOK().WithPayload(&models.NewCartResponse{
		CartID: cartID,
	})
}

func (controller *ShoppingController) AddProductToCart(params operations.AddNewProductParams) middleware.Responder {
	list, err := controller.cartService.AddProductToCart(params.UserID, params.CartID, *params.ProductInput)
	if err != nil {
		switch err {
		case cart.ErrCartNotFound,
			cart.ErrProductNotFound:
			return operations.NewAddNewProductNotFound().WithPayload(err.Error())
		case cart.ErrInvalidCartId,
			cart.ErrMissingCartID,
			cart.ErrMissingUserID,
			cart.ErrInvalidQuantity:
			return operations.NewAddNewProductBadRequest().WithPayload(err.Error())
		default:
			return operations.NewAddNewProductInternalServerError().WithPayload(err.Error())
		}
	}

	return operations.NewAddNewProductOK().WithPayload(list)
}

func (controller *ShoppingController) UpdateQuantity(params operations.ModifyQuantityParams) middleware.Responder {
	product, err := controller.cartService.UpdateQuantity(params.UserID, params.CartID, params.ProductID, *params.UpdateQuantity)
	if err != nil {
		switch err {
		case cart.ErrCartNotFound,
			cart.ErrProductNotFound:
			return operations.NewModifyQuantityNotFound().WithPayload(err.Error())
		case cart.ErrInvalidCartId,
			cart.ErrMissingCartID,
			cart.ErrMissingUserID,
			cart.ErrInvalidQuantity:
			return operations.NewModifyQuantityBadRequest().WithPayload(err.Error())
		default:
			return operations.NewModifyQuantityInternalServerError().WithPayload(err.Error())
		}
	}

	return operations.NewModifyQuantityOK().WithPayload(product)
}

func (controller *ShoppingController) CreateOrder(params operations.CreateOrderParams) middleware.Responder {
	order, err := controller.cartService.CreateOrder(params.UserID)
	if err != nil {
		switch err {
		case cart.ErrMissingUserID:
			return operations.NewCreateOrderBadRequest().WithPayload(err.Error())
		case cart.ErrCartNotFound:
			return operations.NewCreateOrderNotFound().WithPayload(err.Error())
		default:
			return operations.NewCreateOrderInternalServerError().WithPayload(err.Error())
		}
	}

	return operations.NewCreateOrderOK().WithPayload(order)
}
