// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/labs": {
            "get": {
                "description": "Fetches a list of all labs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "labs"
                ],
                "summary": "Retrieve all labs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/res.CreateLabResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    }
                }
            }
        },
        "/api/v1/labs/create": {
            "post": {
                "description": "Create a new lab",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "labs"
                ],
                "summary": "Create a new lab",
                "parameters": [
                    {
                        "description": "Lab object that needs to be created",
                        "name": "lab",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.CreateLabRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/res.CreateLabResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    }
                }
            }
        },
        "/api/v1/labs/{id}": {
            "get": {
                "description": "Fetches a lab by its unique identifier",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "labs"
                ],
                "summary": "Retrieve a lab by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Lab ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.CreateLabResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates the details of a specific lab by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "labs"
                ],
                "summary": "Update an existing lab",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Lab ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Lab object with updated details",
                        "name": "lab",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.CreateLabRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.CreateLabResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a specific lab by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "labs"
                ],
                "summary": "Delete a lab",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Lab ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    }
                }
            }
        },
        "/api/v1/users/create": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User object that needs to be created",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.CreateUserRequest"
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
                            "$ref": "#/definitions/rest_err.RestErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_type": {
                    "type": "integer"
                }
            }
        },
        "req.CreateLabRequest": {
            "type": "object",
            "properties": {
                "acessible": {
                    "type": "boolean"
                },
                "local": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pcNumbers": {
                    "type": "integer"
                },
                "softwares": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "req.CreateUserRequest": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_type": {
                    "type": "integer"
                }
            }
        },
        "res.CreateLabResponse": {
            "type": "object",
            "properties": {
                "acessible": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "local": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pcNumbers": {
                    "type": "integer"
                },
                "softwares": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "rest_err.Causes": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "rest_err.RestErr": {
            "type": "object",
            "properties": {
                "causes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/rest_err.Causes"
                    }
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Lab Manager API",
	Description:      "This is a sample server for a lab manager.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
