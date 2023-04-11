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
        "/v1/apps": {
            "get": {
                "description": "Get all apps",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App"
                ],
                "summary": "Get all apps",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.App"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new app with the provided data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App"
                ],
                "summary": "Create a new app",
                "parameters": [
                    {
                        "description": "App data",
                        "name": "app",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.App"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.App"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/apps/{appID}/attrs/{attrID}": {
            "put": {
                "description": "Update an attribute by providing the app ID, attribute ID, and updated attribute information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attr"
                ],
                "summary": "Update an attribute",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int32",
                        "description": "App ID",
                        "name": "appID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "format": "int32",
                        "description": "Attribute ID",
                        "name": "attrID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated attribute information",
                        "name": "updatedAttr",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.Attr"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/common.SuccResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an attribute by providing the app ID, attribute ID, and updated attribute information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attr"
                ],
                "summary": "Delete an attribute",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int32",
                        "description": "App ID",
                        "name": "appID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "format": "int32",
                        "description": "Attribute ID",
                        "name": "attrID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated attribute information",
                        "name": "updatedAttr",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.Attr"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/common.SuccResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/apps/{app_id}": {
            "delete": {
                "description": "Delete an app by app ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App"
                ],
                "summary": "Delete an app",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "App ID",
                        "name": "app_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted app",
                        "schema": {
                            "$ref": "#/definitions/common.SuccResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid app ID",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/apps/{app_id}/attrs": {
            "get": {
                "description": "Get all attrs",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attr"
                ],
                "summary": "Get all Attrs",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "App ID",
                        "name": "app_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Attr"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new Attr with the provided data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attr"
                ],
                "summary": "Create a new Attr",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "App ID",
                        "name": "app_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Attr data",
                        "name": "app",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.Attr"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Attr"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/apps/{app_id}/attrs/{attr_id}": {
            "get": {
                "description": "Get Attr by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attr"
                ],
                "summary": "Get Attr by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "App ID",
                        "name": "app_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Attr ID",
                        "name": "attr_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Attr"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/apps/{id}": {
            "get": {
                "description": "Get app by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App"
                ],
                "summary": "Get app by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "App ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.App"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "根据提供的应用ID和更新的应用数据更新应用信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App"
                ],
                "summary": "更新应用",
                "parameters": [
                    {
                        "type": "string",
                        "description": "应用ID",
                        "name": "app_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新的应用数据",
                        "name": "appData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.App"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.App"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/gpt/chat-completion": {
            "get": {
                "description": "基于OpenAI的Chat Completion API，生成对话文本。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "使用OpenAI生成对话文本",
                "parameters": [
                    {
                        "type": "string",
                        "description": "对话角色",
                        "name": "role",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "对话内容",
                        "name": "content",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功生成对话文本",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "生成对话文本失败",
                        "schema": {}
                    }
                }
            }
        },
        "/v1/gpt/text-completion": {
            "get": {
                "description": "Generate text completion based on prompt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Generate Text Completion",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Prompt for text completion",
                        "name": "prompt",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Maximum number of tokens to generate",
                        "name": "max_tokens",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Generated text",
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
        "/v1/prompt/{id}": {
            "get": {
                "description": "Retrieves a prompt by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Retrieve a prompt by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Prompt ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/api.Prompt"
                        }
                    }
                }
            }
        },
        "/v2/prompt/{id}": {
            "get": {
                "description": "Retrieves a prompt by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v2"
                ],
                "summary": "Retrieve a prompt by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Prompt ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/api.Prompt"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.App": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "api.Attr": {
            "type": "object",
            "properties": {
                "context": {
                    "description": "内容",
                    "type": "string"
                },
                "name": {
                    "description": "Tab",
                    "type": "string"
                },
                "type": {
                    "description": "1 chat completion 2 img",
                    "type": "integer"
                }
            }
        },
        "api.Prompt": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "common.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "common.SuccResponse": {
            "type": "object",
            "properties": {
                "succ": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "10.11.28.73:7211",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "OPENAI",
	Description:      "This is a API documentation for OPENAI.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
