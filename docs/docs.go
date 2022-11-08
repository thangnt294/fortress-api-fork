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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/auth": {
            "post": {
                "description": "Authorise user when login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Authorise user when login",
                "parameters": [
                    {
                        "description": "Google login code",
                        "name": "code",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Google redirect url",
                        "name": "redirectUrl",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.AuthData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/view.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/view.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/view.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/employee/{id}": {
            "get": {
                "description": "Get employee by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Get employee by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.EmployeeListData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/view.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/view.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/view.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/metadata/countries": {
            "get": {
                "description": "Get all countries",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Metadata"
                ],
                "summary": "Get all countries",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.CountriesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/view.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/view.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/metadata/countries/:country_id/cities": {
            "get": {
                "description": "Get list cities by country",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Metadata"
                ],
                "summary": "Get list cities by country",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.CitiesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/view.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/view.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/metadata/positions": {
            "get": {
                "description": "Get list values for positions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Metadata"
                ],
                "summary": "Get list values for positions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.PositionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/view.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/view.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/metadata/working-status": {
            "get": {
                "description": "Get list values for working status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Metadata"
                ],
                "summary": "Get list values for working status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/view.WorkingStatusData"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/view.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/view.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Country": {
            "type": "object",
            "properties": {
                "cities": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "code": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.Position": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "view.ApiError": {
            "description": "validation error details",
            "type": "object",
            "properties": {
                "enums": {
                    "description": "available options incase of field's payload is enums",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "field": {
                    "description": "the field cause the error",
                    "type": "string"
                },
                "msg": {
                    "description": "error message",
                    "type": "string"
                }
            }
        },
        "view.AuthData": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "employee": {
                    "$ref": "#/definitions/view.EmployeeListData"
                }
            }
        },
        "view.CitiesResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "view.CountriesResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Country"
                    }
                }
            }
        },
        "view.EmployeeListData": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "avatar": {
                    "type": "string"
                },
                "birthday": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "displayName": {
                    "type": "string"
                },
                "fullName": {
                    "description": "basic info",
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "horoscope": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "joinedDate": {
                    "type": "string"
                },
                "leftDate": {
                    "type": "string"
                },
                "mbti": {
                    "type": "string"
                },
                "personalEmail": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "status": {
                    "description": "working info",
                    "type": "string"
                },
                "teamEmail": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "view.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.ApiError"
                    }
                }
            }
        },
        "view.PositionResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Position"
                    }
                }
            }
        },
        "view.WorkingStatusData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
