// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "E-commerce platform",
    "title": "Trafilea API",
    "contact": {
      "name": "Eddie Riascos",
      "email": "eriascos@trafilea.com"
    },
    "version": "1.0.0"
  },
  "paths": {
    "/users/{userId}/carts/": {
      "post": {
        "description": "Create a new cart if it does not exist",
        "operationId": "CreateCart",
        "parameters": [
          {
            "type": "string",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "UUID for the new cart",
            "schema": {
              "$ref": "#/definitions/NewCartResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/users/{userId}/carts/{cartId}": {
      "post": {
        "description": "Add a new product to cart",
        "operationId": "AddNewProduct",
        "parameters": [
          {
            "type": "string",
            "name": "userId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "cartId",
            "in": "path",
            "required": true
          },
          {
            "name": "productInput",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ProductInput"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ListProducts"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "Resource not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/users/{userId}/carts/{cartId}/products/{productId}/": {
      "patch": {
        "description": "Modify quantity for a product",
        "operationId": "ModifyQuantity",
        "parameters": [
          {
            "type": "string",
            "name": "userId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "cartId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "productId",
            "in": "path",
            "required": true
          },
          {
            "name": "updateQuantity",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UpdateQuantity"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "Resource not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/users/{userId}/orders/": {
      "post": {
        "description": "Create a new order",
        "operationId": "CreateOrder",
        "parameters": [
          {
            "type": "string",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/OrderResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "Resource not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ListProducts": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Product"
      }
    },
    "NewCartResponse": {
      "description": "cart output",
      "type": "object",
      "properties": {
        "cartId": {
          "type": "string"
        }
      },
      "example": {
        "cartId": "be16910858a41fd19ea5c1b4e9decca9a784d1024cb00b2158defe2f29dc86dd"
      }
    },
    "OrderResponse": {
      "description": "order response",
      "type": "object",
      "properties": {
        "cartId": {
          "type": "string"
        },
        "totals": {
          "type": "object",
          "properties": {
            "discounts": {
              "type": "number"
            },
            "order": {
              "type": "number"
            },
            "products": {
              "type": "number"
            },
            "shipping": {
              "type": "number"
            }
          }
        }
      }
    },
    "Product": {
      "type": "object",
      "properties": {
        "category": {
          "description": "product category",
          "type": "string",
          "enum": [
            "Coffee",
            "Equipment",
            "Accessories"
          ]
        },
        "name": {
          "description": "product name",
          "type": "string"
        },
        "price": {
          "description": "product price",
          "type": "number"
        },
        "productId": {
          "description": "product id",
          "type": "string"
        },
        "quantity": {
          "description": "quantity",
          "type": "number"
        }
      }
    },
    "ProductInput": {
      "type": "object",
      "properties": {
        "productId": {
          "description": "product id",
          "type": "string"
        },
        "quantity": {
          "description": "quantity",
          "type": "number"
        }
      }
    },
    "UpdateQuantity": {
      "type": "object",
      "properties": {
        "quantity": {
          "type": "number"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "E-commerce platform",
    "title": "Trafilea API",
    "contact": {
      "name": "Eddie Riascos",
      "email": "eriascos@trafilea.com"
    },
    "version": "1.0.0"
  },
  "paths": {
    "/users/{userId}/carts/": {
      "post": {
        "description": "Create a new cart if it does not exist",
        "operationId": "CreateCart",
        "parameters": [
          {
            "type": "string",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "UUID for the new cart",
            "schema": {
              "$ref": "#/definitions/NewCartResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/users/{userId}/carts/{cartId}": {
      "post": {
        "description": "Add a new product to cart",
        "operationId": "AddNewProduct",
        "parameters": [
          {
            "type": "string",
            "name": "userId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "cartId",
            "in": "path",
            "required": true
          },
          {
            "name": "productInput",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ProductInput"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ListProducts"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "Resource not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/users/{userId}/carts/{cartId}/products/{productId}/": {
      "patch": {
        "description": "Modify quantity for a product",
        "operationId": "ModifyQuantity",
        "parameters": [
          {
            "type": "string",
            "name": "userId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "cartId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "productId",
            "in": "path",
            "required": true
          },
          {
            "name": "updateQuantity",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UpdateQuantity"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "Resource not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/users/{userId}/orders/": {
      "post": {
        "description": "Create a new order",
        "operationId": "CreateOrder",
        "parameters": [
          {
            "type": "string",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/OrderResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "404": {
            "description": "Resource not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ListProducts": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Product"
      }
    },
    "NewCartResponse": {
      "description": "cart output",
      "type": "object",
      "properties": {
        "cartId": {
          "type": "string"
        }
      },
      "example": {
        "cartId": "be16910858a41fd19ea5c1b4e9decca9a784d1024cb00b2158defe2f29dc86dd"
      }
    },
    "OrderResponse": {
      "description": "order response",
      "type": "object",
      "properties": {
        "cartId": {
          "type": "string"
        },
        "totals": {
          "type": "object",
          "properties": {
            "discounts": {
              "type": "number"
            },
            "order": {
              "type": "number"
            },
            "products": {
              "type": "number"
            },
            "shipping": {
              "type": "number"
            }
          }
        }
      }
    },
    "OrderResponseTotals": {
      "type": "object",
      "properties": {
        "discounts": {
          "type": "number"
        },
        "order": {
          "type": "number"
        },
        "products": {
          "type": "number"
        },
        "shipping": {
          "type": "number"
        }
      }
    },
    "Product": {
      "type": "object",
      "properties": {
        "category": {
          "description": "product category",
          "type": "string",
          "enum": [
            "Coffee",
            "Equipment",
            "Accessories"
          ]
        },
        "name": {
          "description": "product name",
          "type": "string"
        },
        "price": {
          "description": "product price",
          "type": "number"
        },
        "productId": {
          "description": "product id",
          "type": "string"
        },
        "quantity": {
          "description": "quantity",
          "type": "number"
        }
      }
    },
    "ProductInput": {
      "type": "object",
      "properties": {
        "productId": {
          "description": "product id",
          "type": "string"
        },
        "quantity": {
          "description": "quantity",
          "type": "number"
        }
      }
    },
    "UpdateQuantity": {
      "type": "object",
      "properties": {
        "quantity": {
          "type": "number"
        }
      }
    }
  }
}`))
}