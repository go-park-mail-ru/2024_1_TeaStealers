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
        "/adverts/": {
            "post": {
                "description": "Create a new advert",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "adverts"
                ],
                "summary": "Create a new advert",
                "parameters": [
                    {
                        "description": "Advert data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AdvertCreateData"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Advert"
                        }
                    },
                    "400": {
                        "description": "Incorrect data format",
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
        "/adverts/list/": {
            "get": {
                "description": "Get list of adverts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "adverts"
                ],
                "summary": "Get list of adverts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Advert"
                            }
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
        "/adverts/{id}": {
            "get": {
                "description": "Get advert by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "adverts"
                ],
                "summary": "Get advert by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Advert ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Advert"
                        }
                    },
                    "400": {
                        "description": "Invalid ID parameter",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Advert not found",
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
            },
            "post": {
                "description": "Update advert by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "adverts"
                ],
                "summary": "Update advert by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Advert ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Advert data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "UPDATED advert",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid ID parameter or incorrect data format",
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
            },
            "delete": {
                "description": "Delete advert by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "adverts"
                ],
                "summary": "Delete advert by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Advert ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "DELETED advert",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid ID parameter",
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
        "/auth/login": {
            "post": {
                "description": "User login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User login data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLoginData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Incorrect password or login",
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
        "/auth/logout": {
            "get": {
                "description": "User logout",
                "tags": [
                    "auth"
                ],
                "summary": "User logout",
                "responses": {
                    "200": {
                        "description": "Logged out",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLoginData"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Incorrect data format",
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
        "/buildings/": {
            "post": {
                "description": "Create a new building",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Create a new building",
                "parameters": [
                    {
                        "description": "Building data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BuildingCreateData"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Building"
                        }
                    },
                    "400": {
                        "description": "Incorrect data format",
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
        "/buildings/list/": {
            "get": {
                "description": "Get list of buildings",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Get list of buildings",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Building"
                            }
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
        "/buildings/{id}": {
            "get": {
                "description": "Get building by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Get building by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Building ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Building"
                        }
                    },
                    "400": {
                        "description": "Invalid ID parameter",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Building not found",
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
            },
            "post": {
                "description": "Update building by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Update building by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Building ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Building data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "UPDATED building",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid ID parameter or incorrect data format",
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
            },
            "delete": {
                "description": "Delete building by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buildings"
                ],
                "summary": "Delete building by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Building ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "DELETED building",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid ID parameter",
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
        "/companies": {
            "post": {
                "description": "Create a new company",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "companies"
                ],
                "summary": "Create a new company",
                "parameters": [
                    {
                        "description": "Company data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CompanyCreateData"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Company"
                        }
                    },
                    "400": {
                        "description": "Incorrect data format",
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
        "/companies/list/": {
            "get": {
                "description": "Get list of companies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "companies"
                ],
                "summary": "Get list of companies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Company"
                            }
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
        "/companies/{id}": {
            "get": {
                "description": "Get company by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "companies"
                ],
                "summary": "Get company by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Company ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Company"
                        }
                    },
                    "400": {
                        "description": "Invalid ID parameter",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Company not found",
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
            },
            "post": {
                "description": "Update company by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "companies"
                ],
                "summary": "Update company by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Company ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Company data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "UPDATED company",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid ID parameter or incorrect data format",
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
            },
            "delete": {
                "description": "Delete company by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "companies"
                ],
                "summary": "Delete company by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Company ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "DELETED company",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid ID parameter",
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
        "models.Advert": {
            "type": "object",
            "properties": {
                "buildingid": {
                    "description": "BuildingId is the id of the building to which the advert belongs.",
                    "type": "string"
                },
                "companyid": {
                    "description": "CompanyId is the id of the company to which the advert belongs.",
                    "type": "string"
                },
                "datacreation": {
                    "description": "DataCreation is the time of adding a record to the database.",
                    "type": "string"
                },
                "description": {
                    "description": "Description is the description of the advert.",
                    "type": "string"
                },
                "id": {
                    "description": "ID uniquely identifies the advert.",
                    "type": "string"
                },
                "isdeleted": {
                    "description": "isDeleted means is the advert deleted?.",
                    "type": "boolean"
                },
                "location": {
                    "description": "Location is the location of the object in advert.",
                    "type": "string"
                },
                "phone": {
                    "description": "Phone is the phone of the owner advert.",
                    "type": "integer"
                },
                "price": {
                    "description": "Price is the price of the object in advert.",
                    "type": "number"
                },
                "userid": {
                    "description": "UserId uniquely identifies who owns the advert.",
                    "type": "string"
                }
            }
        },
        "models.AdvertCreateData": {
            "type": "object",
            "properties": {
                "buildingid": {
                    "description": "BuildingId is the id of the building to which the advert belongs.",
                    "type": "string"
                },
                "companyid": {
                    "description": "CompanyId is the id of the company to which the advert belongs.",
                    "type": "string"
                },
                "description": {
                    "description": "Description is the description of the advert.",
                    "type": "string"
                },
                "location": {
                    "description": "Location is the location of the object in advert.",
                    "type": "string"
                },
                "phone": {
                    "description": "Phone is the phone of the owner advert.",
                    "type": "integer"
                },
                "price": {
                    "description": "Price is the price of the object in advert.",
                    "type": "number"
                },
                "userid": {
                    "description": "UserId uniquely identifies who owns the advert.",
                    "type": "string"
                }
            }
        },
        "models.Building": {
            "type": "object",
            "properties": {
                "datacreation": {
                    "description": "DataCreation is the time of adding a record to the database.",
                    "type": "string"
                },
                "description": {
                    "description": "Description is the description of the building.",
                    "type": "string"
                },
                "id": {
                    "description": "ID uniquely identifies the building.",
                    "type": "string"
                },
                "isdeleted": {
                    "description": "isDeleted means is the building deleted?.",
                    "type": "boolean"
                },
                "location": {
                    "description": "Location is the location of the building.",
                    "type": "string"
                }
            }
        },
        "models.BuildingCreateData": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "Description is the description of the building.",
                    "type": "string"
                },
                "location": {
                    "description": "Location is the location of the building.",
                    "type": "string"
                }
            }
        },
        "models.Company": {
            "type": "object",
            "properties": {
                "datacreation": {
                    "description": "DataCreation is the time of adding a record to the database.",
                    "type": "string"
                },
                "description": {
                    "description": "Description is the description of the company.",
                    "type": "string"
                },
                "id": {
                    "description": "ID uniquely identifies the company.",
                    "type": "string"
                },
                "isdeleted": {
                    "description": "isDeleted means is the company deleted?.",
                    "type": "boolean"
                },
                "name": {
                    "description": "Name is the name of the company.",
                    "type": "string"
                },
                "phone": {
                    "description": "Phone is the phone of the company.",
                    "type": "integer"
                }
            }
        },
        "models.CompanyCreateData": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "Descpription stands for company description",
                    "type": "string"
                },
                "name": {
                    "description": "Name stands for company name",
                    "type": "string"
                },
                "phone": {
                    "description": "Phone stands for company phone",
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID uniquely identifies the user.",
                    "type": "string"
                },
                "login": {
                    "description": "Login is the username of the user.",
                    "type": "string"
                }
            }
        },
        "models.UserLoginData": {
            "type": "object",
            "properties": {
                "login": {
                    "description": "Login stands for users nickname",
                    "type": "string"
                },
                "password": {
                    "description": "Password stands for users password",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "0.0.0.0:8080",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "Sample Project API",
	Description:      "This is a sample server Tean server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
