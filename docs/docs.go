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
            "name": "Erik Rio Setiawan",
            "url": "http://www.swagger.io/support",
            "email": "erikriosetiawan15@gmail.com"
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
        "/addresses/{id}": {
            "put": {
                "description": "Update an address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "addresses"
                ],
                "summary": "Update an Address",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "default",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.UpdateAddress"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/admins": {
            "post": {
                "description": "Admin login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admins"
                ],
                "summary": "Admin Login",
                "parameters": [
                    {
                        "description": "admin credentials",
                        "name": "default",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.Credential"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.loginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/groups": {
            "get": {
                "description": "Get Groups",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Get Groups",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.groupsResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Create a Group",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "default",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.CreateGroup"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controller.createGroupResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/groups/{id}": {
            "get": {
                "description": "Get group by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Get Group by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "group ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.groupResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a group",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Update a Group",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "default",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.UpdateGroup"
                        }
                    },
                    {
                        "type": "string",
                        "description": "group ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete group by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Delete Group by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "group ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/groups/{id}/generate": {
            "get": {
                "description": "Generate QR Code",
                "produces": [
                    "image/png"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Generate QR Code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "group ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/groups/{id}/properties": {
            "post": {
                "description": "Add a Property",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "groups"
                ],
                "summary": "Add a Property",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "default",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.CreateProperty"
                        }
                    },
                    {
                        "type": "string",
                        "description": "group ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controller.createPropertyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.createGroupResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "x-order": "0"
                },
                "message": {
                    "type": "string",
                    "x-order": "1"
                },
                "data": {
                    "x-order": "2",
                    "$ref": "#/definitions/controller.idData"
                }
            }
        },
        "controller.createPropertyResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "x-order": "0"
                },
                "message": {
                    "type": "string",
                    "x-order": "1"
                },
                "data": {
                    "x-order": "2",
                    "$ref": "#/definitions/controller.idData"
                }
            }
        },
        "controller.groupData": {
            "type": "object",
            "properties": {
                "group": {
                    "$ref": "#/definitions/response.Group"
                }
            }
        },
        "controller.groupResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "x-order": "0"
                },
                "message": {
                    "type": "string",
                    "x-order": "1"
                },
                "data": {
                    "x-order": "2",
                    "$ref": "#/definitions/controller.groupData"
                }
            }
        },
        "controller.groupsData": {
            "type": "object",
            "properties": {
                "groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.Group"
                    }
                }
            }
        },
        "controller.groupsResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "x-order": "0"
                },
                "message": {
                    "type": "string",
                    "x-order": "1"
                },
                "data": {
                    "x-order": "2",
                    "$ref": "#/definitions/controller.groupsData"
                }
            }
        },
        "controller.idData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "controller.loginResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "maxLength": 80,
                    "minLength": 2,
                    "x-order": "0"
                },
                "message": {
                    "type": "string",
                    "maxLength": 80,
                    "minLength": 2,
                    "x-order": "1"
                },
                "data": {
                    "x-order": "2",
                    "$ref": "#/definitions/controller.tokenData"
                }
            }
        },
        "controller.tokenData": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "payload.CreateGroup": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 80,
                    "minLength": 2,
                    "x-order": "0"
                },
                "leader": {
                    "type": "string",
                    "maxLength": 80,
                    "minLength": 2,
                    "x-order": "1"
                },
                "address": {
                    "type": "string",
                    "maxLength": 1000,
                    "minLength": 2,
                    "x-order": "2"
                },
                "villageID": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 2,
                    "x-order": "3"
                }
            }
        },
        "payload.CreateProperty": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 80,
                    "minLength": 2,
                    "x-order": "0"
                },
                "description": {
                    "type": "string",
                    "maxLength": 1000,
                    "minLength": 2,
                    "x-order": "1"
                },
                "amount": {
                    "type": "integer",
                    "minimum": 1,
                    "x-order": "2"
                }
            }
        },
        "payload.Credential": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 2,
                    "x-order": "0"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2,
                    "x-order": "1"
                }
            }
        },
        "payload.UpdateAddress": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "maxLength": 1000,
                    "minLength": 2,
                    "x-order": "0"
                },
                "villageID": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 2,
                    "x-order": "1"
                }
            }
        },
        "payload.UpdateGroup": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 80,
                    "minLength": 2,
                    "x-order": "0"
                },
                "leader": {
                    "type": "string",
                    "maxLength": 80,
                    "minLength": 2,
                    "x-order": "1"
                }
            }
        },
        "response.Address": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "x-order": "0"
                },
                "address": {
                    "type": "string",
                    "x-order": "1"
                },
                "villageID": {
                    "type": "string",
                    "x-order": "2"
                },
                "villageName": {
                    "type": "string",
                    "x-order": "3"
                },
                "districtID": {
                    "type": "string",
                    "x-order": "4"
                },
                "districtName": {
                    "type": "string",
                    "x-order": "5"
                },
                "regencyID": {
                    "type": "string",
                    "x-order": "5"
                },
                "regencyName": {
                    "type": "string",
                    "x-order": "6"
                },
                "provinceID": {
                    "type": "string",
                    "x-order": "7"
                },
                "provinceName": {
                    "type": "string",
                    "x-order": "8"
                }
            }
        },
        "response.Group": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "x-order": "0"
                },
                "name": {
                    "type": "string",
                    "x-order": "1"
                },
                "leader": {
                    "type": "string",
                    "x-order": "2"
                },
                "address": {
                    "x-order": "3",
                    "$ref": "#/definitions/response.Address"
                },
                "properties": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.Property"
                    },
                    "x-order": "4"
                }
            }
        },
        "response.Property": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "x-order": "0"
                },
                "name": {
                    "type": "string",
                    "x-order": "1"
                },
                "description": {
                    "type": "string",
                    "x-order": "2"
                },
                "amount": {
                    "type": "integer",
                    "x-order": "3"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Reog Apps API",
	Description:      "API for Reog Group in Ponorogo",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
