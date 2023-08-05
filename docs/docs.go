// Code generated by swaggo/swag. DO NOT EDIT.

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
            "url": "https://durp.info",
            "email": "developerdurp@durp.info"
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
        "/health/getHealth": {
            "get": {
                "description": "Get the health of the API",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Generate Health status",
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/jokes/dadjoke": {
            "get": {
                "description": "get a dad joke",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DadJoke"
                ],
                "summary": "Get dadjoke",
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            },
            "post": {
                "description": "create a dad joke",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DadJoke"
                ],
                "summary": "Generate dadjoke",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dad Joke you wish to enter into database",
                        "name": "joke",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            },
            "delete": {
                "description": "create a dad joke",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DadJoke"
                ],
                "summary": "Generate dadjoke",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dad joke you wish to delete from the database",
                        "name": "joke",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/openai/general": {
            "get": {
                "description": "Ask ChatGPT a general question",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "openai"
                ],
                "summary": "Gerneral ChatGPT",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ask ChatGPT a general question",
                        "name": "message",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/openai/travelagent": {
            "get": {
                "description": "Ask ChatGPT for suggestions as if it was a travel agent",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "openai"
                ],
                "summary": "Travel Agent ChatGPT",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ask ChatGPT for suggestions as a travel agent",
                        "name": "message",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/unraid/powerusage": {
            "get": {
                "description": "Gets the PSU Data from unraid",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "unraid"
                ],
                "summary": "Unraid PSU Stats",
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "$ref": "#/definitions/model.PowerSupply"
                        }
                    },
                    "412": {
                        "description": "error",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "message"
                }
            }
        },
        "model.PowerSupply": {
            "type": "object",
            "properties": {
                "12v_load": {
                    "type": "integer"
                },
                "12v_watts": {
                    "type": "integer"
                },
                "3v_load": {
                    "type": "integer"
                },
                "3v_watts": {
                    "type": "integer"
                },
                "5v_load": {
                    "type": "integer"
                },
                "5v_watts": {
                    "type": "integer"
                },
                "capacity": {
                    "type": "string"
                },
                "efficiency": {
                    "type": "integer"
                },
                "fan_rpm": {
                    "type": "integer"
                },
                "load": {
                    "type": "integer"
                },
                "poweredon": {
                    "type": "string"
                },
                "poweredon_raw": {
                    "type": "string"
                },
                "product": {
                    "type": "string"
                },
                "temp1": {
                    "type": "integer"
                },
                "temp2": {
                    "type": "integer"
                },
                "uptime": {
                    "type": "string"
                },
                "uptime_raw": {
                    "type": "string"
                },
                "vendor": {
                    "type": "string"
                },
                "watts": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "DurpAPI",
	Description:      "API for Durp's needs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
