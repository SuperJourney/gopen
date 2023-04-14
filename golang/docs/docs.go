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
        "/v1/apps/attrs": {
            "get": {
                "description": "Get all attrs",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attr"
                ],
                "summary": "Get all Attrs",
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
                "summary": "Get  all Attrs by App ID",
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
                "description": "用户通用属性创建 类型为chat请使用  CreateChatAttr",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attr"
                ],
                "summary": "通用Attr创建 (1 Chat 2 Edit )",
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
        "/v1/apps/{app_id}/chat_attrs": {
            "post": {
                "description": "使用提供的数据创建新的 对话Attr",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attr"
                ],
                "summary": "对话Attr创建",
                "parameters": [
                    {
                        "type": "integer",
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
                            "$ref": "#/definitions/api.ChatAttr"
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
        "/v1/apps/{app_id}/chat_attrs/{attr_id}": {
            "put": {
                "description": "使用提供的数据创建新的 对话Attr",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attr"
                ],
                "summary": "对话Attr更新",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "应用ID",
                        "name": "app_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "AttrID",
                        "name": "attr_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新的应用数据",
                        "name": "appData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.ChatAttr"
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
                            "$ref": "#/definitions/api.App_S"
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
        "/v1/gpt/{attrID}/chat-completion": {
            "post": {
                "description": "Generate chat completion text based on input messages.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ChatGpt"
                ],
                "summary": "Generate Chat Completion",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Attr ID",
                        "name": "attrID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User Messages",
                        "name": "userMessage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UserMessage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ChatCompletionResponse"
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
        "/v1/sd/{attr_id}/img2img": {
            "post": {
                "description": "将一张图片文件上传并转换成另一张图片",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "image/jpeg"
                ],
                "tags": [
                    "SD"
                ],
                "summary": "图片转换",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Attr ID",
                        "name": "attrID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "待上传的图片文件",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "图片文件",
                        "schema": {
                            "type": "file"
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
        "/v1/sd/{attr_id}/txt2img": {
            "post": {
                "description": "将文本转换为图片",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "image/jpeg"
                ],
                "tags": [
                    "SD"
                ],
                "summary": "文本转图片",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int32",
                        "description": "Attr ID",
                        "name": "attr_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新的应用数据",
                        "name": "appData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.TextToImgMessage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "图片文件",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "500": {
                        "description": "错误信息",
                        "schema": {
                            "$ref": "#/definitions/common.ErrorResponse"
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
                    "type": "string",
                    "example": "商城商品"
                }
            }
        },
        "api.App_S": {
            "type": "object",
            "properties": {
                "attrs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Attr"
                    }
                },
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
                    "type": "string",
                    "example": "按照stable diffusion的关键字要求，给出主题鲜明的prompt,并使用英文文回答"
                },
                "context_type": {
                    "description": "1 chat 2 edit",
                    "type": "integer",
                    "example": 1
                },
                "id": {
                    "description": "Example ID",
                    "type": "integer",
                    "example": 0
                },
                "name": {
                    "description": "Tab",
                    "type": "string",
                    "example": "商城商品"
                },
                "type": {
                    "description": "1 chat completion 2 img",
                    "type": "integer",
                    "enum": [
                        1,
                        2
                    ],
                    "example": 1
                }
            }
        },
        "api.ChatAttr": {
            "type": "object",
            "properties": {
                "context": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.ChatCompletionMessage"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "description": "属性名称",
                    "type": "string",
                    "example": "商城商品"
                },
                "type": {
                    "description": "1 纯文字 2 img",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "api.ChatCompletionMessage": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "api.ChatCompletionResponse": {
            "type": "object",
            "properties": {
                "context": {
                    "type": "string"
                }
            }
        },
        "api.TextToImgMessage": {
            "type": "object",
            "properties": {
                "negative_prompt": {
                    "type": "string"
                },
                "prompt": {
                    "type": "string"
                },
                "userMessage": {
                    "description": "用户输入",
                    "type": "string"
                }
            }
        },
        "api.UserMessage": {
            "type": "object",
            "properties": {
                "content": {
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
	Host:             "localhost:8080",
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
