// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/user": {
            "post": {
                "description": "Save User data in Db.",
                "produces": [
                    "application/json"
                ],
                "summary": "Create User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/user/{userId}": {
            "get": {
                "description": "get user data in Db.",
                "produces": [
                    "application/json"
                ],
                "summary": "Create tags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Save tags data in Db.",
                "produces": [
                    "application/json"
                ],
                "summary": "Create tags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Save tags data in Db.",
                "produces": [
                    "application/json"
                ],
                "summary": "Create tags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Save tags data in Db.",
                "produces": [
                    "application/json"
                ],
                "summary": "Create tags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "responses.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Tag Service API",
	Description:      "A Tag service API in Go using Gin framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
