// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones",
        "contact": {
            "name": "API Support",
            "url": "https://developers.mercadolibre.com.ar/support"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users": {
            "get": {
                "description": "get users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "List users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "store user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Store user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User to store",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "put": {
                "description": "update user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User to update",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.RequestUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "update lastname and age user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update lastname and age user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "lastname and age to update",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.RequestUpdateLastNameAndAge"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Request": {
            "type": "object",
            "required": [
                "age",
                "email",
                "height",
                "lastName",
                "name"
            ],
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handler.RequestUpdate": {
            "type": "object",
            "required": [
                "active",
                "age",
                "email",
                "height",
                "lastName",
                "name"
            ],
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handler.RequestUpdateLastNameAndAge": {
            "type": "object",
            "required": [
                "age",
                "lastName"
            ],
            "properties": {
                "age": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                }
            }
        },
        "web.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {},
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "MELI Bootcamp API",
	Description:      "This API Handle MELI Users.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
