// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mohammed Hattab",
            "url": "mtahatb.com",
            "email": "mohammedhattab97@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api-key": {
            "get": {
                "description": "Responds with a simple JSON message that has the api key",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ApiKey"
                ],
                "summary": "Get API Key",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/health-check": {
            "get": {
                "description": "Responds with a simple JSON message indicating the service status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HealthCheck"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/mileage": {
            "get": {
                "description": "Get a list of all mileage records.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mileage"
                ],
                "summary": "Retrieve all mileage records",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.IrsStandardMileageRate"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/mileage/{id}": {
            "get": {
                "description": "Get a mileage record by its unique ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mileage"
                ],
                "summary": "Retrieve mileage record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Mileage Record ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.IrsStandardMileageRate"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/travel-tracker": {
            "get": {
                "description": "Retrieve all records from the travel tracker",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TravelTracker"
                ],
                "summary": "Get all travel records",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.TaxTravelTracker"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new record to the travel tracker",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TravelTracker"
                ],
                "summary": "Create a new travel record",
                "parameters": [
                    {
                        "description": "Travel Record",
                        "name": "record",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TaxTravelTracker"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.TaxTravelTracker"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/travel-tracker/{id}": {
            "get": {
                "description": "Fetch a single travel record using its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TravelTracker"
                ],
                "summary": "Get a travel record by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Travel Record ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TaxTravelTracker"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "put": {
                "description": "Modify an existing record in the travel tracker",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TravelTracker"
                ],
                "summary": "Update an existing travel record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Travel Record ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated Travel Record",
                        "name": "record",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TaxTravelTracker"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TaxTravelTracker"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove a record from the travel tracker",
                "tags": [
                    "TravelTracker"
                ],
                "summary": "Delete a travel record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Travel Record ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Add a new user record to the users table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
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
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gin.H": {
            "type": "object",
            "additionalProperties": {}
        },
        "models.IrsStandardMileageRate": {
            "type": "object",
            "properties": {
                "centsPerMile": {
                    "type": "integer"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "models.TaxTravelTracker": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "createdBy": {
                    "type": "string"
                },
                "estimatedTaxDeductions": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "travelDistance": {
                    "type": "number"
                },
                "travelFrom": {
                    "type": "string"
                },
                "travelTime": {
                    "type": "string"
                },
                "travelTo": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
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
	Title:            "Travel Tracker API",
	Description:      "This is API will be used to track travel for users for the purposes of tax savings.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
