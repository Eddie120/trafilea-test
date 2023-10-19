package cart

import (
	swaggerModels "api.trafilea.com/models"
	"api.trafilea.com/services/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"strings"
)

type ShoppingCartService interface {
	CreateShoppingCart(userId string) string
	AddProductToCart(userId string, cartId string, productInput swaggerModels.ProductInput) (swaggerModels.ListProducts, error)
	UpdateQuantity(userId string, cartId string, productId string, payload swaggerModels.UpdateQuantity) (*swaggerModels.Product, error)
	CreateOrder(userId string) (*swaggerModels.OrderResponse, error)
}

type ServiceImpl struct {
	globalStore cache.Cache
}

var (
	ErrCartNotFound    = errors.New("cart not found for given userId")
	ErrInvalidCartId   = errors.New(fmt.Sprintf("provided cartId not match with assigned cart for the user"))
	ErrProductNotFound = errors.New("product not found, please review the productId provided")
	ErrInvalidQuantity = errors.New("invalid quantity, the value must be greater than zero")
	ErrMissingUserID   = errors.New("missing user id")
	ErrMissingCartID   = errors.New("missing cart id")
	ErrCartEmpty       = errors.New("cart empty")
)

const defaultShippingCost = 3

func NewShoppingCartService(cache cache.Cache) *ServiceImpl {
	return &ServiceImpl{
		globalStore: cache,
	}
}

func (svc *ServiceImpl) CreateShoppingCart(userId string) string {
	if value, ok := svc.globalStore.Get(userId); ok {
		return value.(string)
	}

	cartId := fmt.Sprintf("CART-%s", uuid.New().String())
	svc.globalStore.Set(userId, cartId, cache.NoExpiration)

	return cartId
}

func (svc *ServiceImpl) AddProductToCart(userId string, cartId string, productInput swaggerModels.ProductInput) (swaggerModels.ListProducts, error) {
	if strings.TrimSpace(userId) == "" {
		return nil, ErrMissingUserID
	}

	if strings.TrimSpace(cartId) == "" {
		return nil, ErrMissingCartID
	}

	storedCartId, found := svc.globalStore.Get(userId)
	if !found {
		return nil, ErrCartNotFound
	}

	if storedCartId.(string) != cartId {
		return nil, ErrInvalidCartId
	}

	productRecord, found := models.HasProduct(productInput.ProductID)
	if !found {
		return nil, ErrProductNotFound
	}

	if productInput.Quantity == 0 {
		return nil, ErrInvalidQuantity
	}

	productInCart, found := svc.globalStore.Get(cartId)
	if !found {
		newProductList := swaggerModels.ListProducts{
			&swaggerModels.Product{
				ProductID: productInput.ProductID,
				Category:  string(productRecord.Category),
				Name:      productRecord.Name,
				Price:     productRecord.Price,
				Quantity:  productInput.Quantity,
			},
		}

		svc.globalStore.Set(cartId, newProductList, cache.NoExpiration)

		return newProductList, nil
	}

	oldProductList := productInCart.(swaggerModels.ListProducts)
	product := findProductInList(oldProductList, productInput.ProductID)
	if product == nil {
		newProduct := &swaggerModels.Product{
			ProductID: productInput.ProductID,
			Category:  string(productRecord.Category),
			Name:      productRecord.Name,
			Price:     productRecord.Price,
			Quantity:  productInput.Quantity,
		}

		oldProductList = append(oldProductList, newProduct)
	} else {
		product.Quantity += productInput.Quantity
	}

	svc.globalStore.Set(cartId, oldProductList, cache.NoExpiration)

	return oldProductList, nil
}

func (svc *ServiceImpl) UpdateQuantity(userId string, cartId string, productId string, payload swaggerModels.UpdateQuantity) (*swaggerModels.Product, error) {
	if strings.TrimSpace(userId) == "" {
		return nil, ErrMissingUserID
	}

	if strings.TrimSpace(cartId) == "" {
		return nil, ErrMissingCartID
	}

	storedCartId, found := svc.globalStore.Get(userId)
	if !found {
		return nil, ErrCartNotFound
	}

	if storedCartId.(string) != cartId {
		return nil, ErrInvalidCartId
	}

	productInCart, found := svc.globalStore.Get(cartId)
	if !found {
		return nil, ErrCartEmpty
	}

	oldProductList := productInCart.(swaggerModels.ListProducts)
	product := findProductInList(oldProductList, productId)

	if product == nil {
		return nil, ErrProductNotFound
	}

	product.Quantity = payload.Quantity

	svc.globalStore.Set(cartId, oldProductList, cache.NoExpiration)

	return product, nil
}

func (svc *ServiceImpl) CreateOrder(userId string) (*swaggerModels.OrderResponse, error) {
	if len(strings.TrimSpace(userId)) == 0 {
		return nil, ErrMissingUserID
	}

	storedCartId, found := svc.globalStore.Get(userId)
	if !found {
		return nil, ErrCartNotFound
	}

	productInCart, found := svc.globalStore.Get(storedCartId.(string))
	if !found {
		return nil, ErrCartEmpty
	}

	oldProductList := productInCart.(swaggerModels.ListProducts)
	if len(oldProductList) == 0 {
		return nil, ErrCartEmpty
	}

	return svc.createOrder(oldProductList, storedCartId.(string)), nil
}

func (svc *ServiceImpl) createOrder(productList swaggerModels.ListProducts, cartId string) *swaggerModels.OrderResponse {
	coffeeCount := 0
	equipmentCount := 0
	totalCoffee := float64(0)
	totalEquipment := float64(0)
	totalAccessories := float64(0)

	for _, item := range productList {
		switch item.Category {
		case string(models.CoffeeCategory):
			coffeeCount++
			totalCoffee += item.Price * item.Quantity
		case string(models.EquipmentCategory):
			equipmentCount++
			totalEquipment += item.Price * item.Quantity
		case string(models.AccessoryCategory):
			totalAccessories += item.Price * item.Quantity
		}
	}

	total := totalCoffee + totalEquipment + totalAccessories
	newOrder := &swaggerModels.OrderResponse{
		CartID: cartId,
		Totals: &swaggerModels.OrderResponseTotals{
			Shipping: defaultShippingCost,
			Products: total,
		},
	}

	if coffeeCount >= 2 { // add extra item with the same category as free
		productId := "PROD-9eab9lf2-6d54-11ee-b962-0242ac120002"
		gift, _ := models.HasProduct(productId)
		productList = append(productList, &swaggerModels.Product{
			ProductID: gift.ProductId,
			Category:  string(gift.Category),
			Name:      gift.Name,
			Price:     0,
			Quantity:  1,
		})

		svc.globalStore.Set(cartId, productList, cache.NoExpiration)
	}

	if equipmentCount > 3 { // free shipping
		newOrder.Totals.Shipping = 0
	}

	if totalAccessories > 70 { // apply 10% discount on products
		newOrder.Totals.Discounts = total * 0.10
	}

	newOrder.Totals.Order = (newOrder.Totals.Products - newOrder.Totals.Discounts) + newOrder.Totals.Shipping

	return newOrder
}

func findProductInList(productList swaggerModels.ListProducts, productId string) *swaggerModels.Product {
	for _, record := range productList {
		if record.ProductID == productId {
			return record
		}
	}

	return nil
}
