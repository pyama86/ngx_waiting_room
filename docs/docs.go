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
        "license": {
            "name": "MIT",
            "url": "http://github.com/pyama86/waitingroom/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/queues": {
            "get": {
                "description": "get queues",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queues"
                ],
                "summary": "get queues",
                "operationId": "queues#get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Queue Domain",
                        "name": "domain",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "per_page",
                        "name": "per_page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/waitingroom.Queue"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/waitingroom.Queue"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "create queue",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queues"
                ],
                "summary": "create queue",
                "operationId": "queues#post",
                "parameters": [
                    {
                        "description": "Queue Object",
                        "name": "queue",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/waitingroom.Queue"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    }
                }
            }
        },
        "/queues/{domain}": {
            "put": {
                "description": "update queue",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queues"
                ],
                "summary": "update queue",
                "operationId": "queues#put",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Queue Name",
                        "name": "domain",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Queue Object",
                        "name": "queue",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/waitingroom.Queue"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete queue",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queues"
                ],
                "summary": "delete queue",
                "operationId": "queues#delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Queue Name",
                        "name": "domain",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    }
                }
            }
        },
        "/viron": {
            "get": {
                "description": "get global menu",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "viron"
                ],
                "summary": "get global menu",
                "operationId": "viron#get",
                "responses": {}
            }
        },
        "/viron_authtype": {
            "get": {
                "description": "get auth type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "viron"
                ],
                "summary": "get auth type",
                "operationId": "viron_authtype#get",
                "responses": {}
            }
        },
        "/whitelist": {
            "get": {
                "description": "get whiteLists",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "whitelist"
                ],
                "summary": "get whiteLists",
                "operationId": "whitelist#get",
                "parameters": [
                    {
                        "type": "string",
                        "description": "WhiteList Domain",
                        "name": "domain",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "per_page",
                        "name": "per_page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/waitingroom.WhiteList"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/waitingroom.WhiteList"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "create whiteList",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "whitelist"
                ],
                "summary": "create whiteList",
                "operationId": "whitelist#post",
                "parameters": [
                    {
                        "description": "WhiteList Object",
                        "name": "whitelist",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/waitingroom.WhiteList"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    }
                }
            }
        },
        "/whitelist/{domain}": {
            "delete": {
                "description": "delete whiteList",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "whitelist"
                ],
                "summary": "delete whiteList",
                "operationId": "whitelist#delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "WhiteList Domain",
                        "name": "domain",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "waitingroom.Queue": {
            "type": "object",
            "required": [
                "domain"
            ],
            "properties": {
                "current_number": {
                    "type": "integer",
                    "minimum": 0
                },
                "domain": {
                    "type": "string"
                },
                "permitted_number": {
                    "type": "integer",
                    "minimum": 0
                }
            }
        },
        "waitingroom.WhiteList": {
            "type": "object",
            "required": [
                "domain"
            ],
            "properties": {
                "domain": {
                    "type": "string"
                }
            }
        }
    },
    "tags": [
        {
            "name": "queues"
        },
        {
            "name": "whitelist"
        },
        {
            "name": "viron"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "WaitingRoomAPI",
	Description:      "API for WaitingRoom",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
