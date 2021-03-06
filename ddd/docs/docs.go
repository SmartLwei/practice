// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ddd/api/v1/user": {
            "post": {
                "description": "add a user to the system so that he will have an identity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "add a user to the system",
                "operationId": "AddUser",
                "parameters": [
                    {
                        "description": "AddUserReq",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/dto.AddUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.JSONResult"
                        }
                    }
                }
            }
        },
        "/ddd/api/v1/users": {
            "get": {
                "description": "find users from the given condition",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "find users",
                "operationId": "GetUsers",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "demo_name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 0,
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.GetUserResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AddUserReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "demo_name"
                }
            }
        },
        "dto.GetUserResp": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "demos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.User"
                    }
                }
            }
        },
        "dto.JSONResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "string": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
