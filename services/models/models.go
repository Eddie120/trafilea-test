package models

type Product struct {
	ProductId string   `json:"product_id"`
	Name      string   `json:"name"`
	Category  Category `json:"category"`
	Price     float64  `json:"price"`
}

const (
	CoffeeCategory    Category = "Coffee"
	AccessoryCategory Category = "Accessories"
	EquipmentCategory Category = "Equipment"
)

type Category string

var Products = []Product{
	{
		ProductId: "PROD-9eab9df2-6d54-11ee-b962-0242ac120002",
		Name:      "A",
		Category:  CoffeeCategory,
		Price:     5,
	},
	{
		ProductId: "PROD-9eab9zf2-6d54-11ee-b962-0242ac120002",
		Name:      "JJJJJJJ",
		Category:  CoffeeCategory,
		Price:     5,
	},
	{
		ProductId: "PROD-9eab9lf2-6d54-11ee-b962-0242ac120002",
		Name:      "BBBBBBB",
		Category:  CoffeeCategory,
		Price:     1.5,
	},
	{
		ProductId: "PROD-9eab9xf2-6d54-11ee-b962-0242ac120002",
		Name:      "QQQQQQQ",
		Category:  CoffeeCategory,
		Price:     5,
	},
	{
		ProductId: "PROD-9eaba09a-6d54-11ee-b962-0242ac120002",
		Name:      "B",
		Category:  AccessoryCategory,
		Price:     5,
	},
	{
		ProductId: "PROD-9eaba1da-6d54-11ee-b962-0242ac120002",
		Name:      "C",
		Category:  EquipmentCategory,
		Price:     5,
	},
	{
		ProductId: "PROD-9eaba842-6d54-11ee-b962-0242ac120002",
		Name:      "D",
		Category:  AccessoryCategory,
		Price:     5,
	},
	{
		ProductId: "PROD-9eaba842-6d54-11ee-b962-0242ac120002",
		Name:      "E",
		Category:  AccessoryCategory,
		Price:     5,
	},
	{
		ProductId: "PROD-9eabaab8-6d54-11ee-b962-0242ac120002",
		Name:      "F",
		Category:  EquipmentCategory,
		Price:     5,
	},
	{
		ProductId: "PROD-9eababee-6d54-11ee-b962-0242ac120002",
		Name:      "G",
		Category:  EquipmentCategory,
		Price:     5,
	},
	{
		ProductId: "PROD-9eabad06-6d54-11ee-b962-0242ac120002",
		Name:      "X",
		Category:  EquipmentCategory,
		Price:     5,
	},
	{
		ProductId: "PROD-9eabae28-6d54-11ee-b962-0242ac120002",
		Name:      "Y",
		Category:  AccessoryCategory,
		Price:     5,
	},
	{
		ProductId: "PROD-9eabaf4a-6d54-11ee-b962-0242ac120002",
		Name:      "Z",
		Category:  AccessoryCategory,
		Price:     5,
	},
}

func HasProduct(productId string) (*Product, bool) {
	for _, product := range Products {
		if product.ProductId == productId {
			return &product, true
		}
	}

	return nil, false
}
