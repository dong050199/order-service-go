// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cart": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Get cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ListProductResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ListProductResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ListProductResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update cart",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Update cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "model",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateCartRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ListProductResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ListProductResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ListProductResponse"
                        }
                    }
                }
            }
        },
        "/product/list": {
            "get": {
                "description": "Get details product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get details product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "size",
                        "name": "size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ListProductResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ListProductResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ListProductResponse"
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "get": {
                "description": "Get details product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get details product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "campaign_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Product"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.Product"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.Product"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "API for get token from user name email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "API for get token from user name email and password",
                "parameters": [
                    {
                        "description": "model",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/wrapper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Response"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Get my record page for non login usersas",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get my record page for non login usersas",
                "parameters": [
                    {
                        "description": "model",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/wrapper.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/wrapper.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Product": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "deleted_by": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "productCart": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.ProductCart"
                    }
                },
                "productOrder": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.ProductOrder"
                    }
                },
                "quantity": {
                    "type": "integer"
                },
                "short_description": {
                    "type": "string"
                },
                "unique_offset": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        },
        "entity.ProductCart": {
            "type": "object",
            "properties": {
                "cartID": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "productID": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "entity.ProductOrder": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "createdBy": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "deletedBy": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "orderID": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "productID": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "updatedBy": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "request.ProductCart": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "request.UpdateCartRequest": {
            "type": "object",
            "properties": {
                "produc_ids": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.ProductCart"
                    }
                }
            }
        },
        "request.UserLogin": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "user_name_or_email": {
                    "type": "string"
                }
            }
        },
        "request.UserRegister": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "response.ListProductResponse": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Product"
                    }
                },
                "total_pages": {
                    "type": "integer"
                }
            }
        },
        "wrapper.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                },
                "trace_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
