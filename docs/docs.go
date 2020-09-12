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
        "contact": {
            "name": "jameshwc",
            "url": "https://jameshsu.csie.org",
            "email": "jameshwc@gmail.com"
        },
        "license": {
            "name": "GPL v3.0",
            "url": "https://www.gnu.org/licenses/gpl-3.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/game/levels/new": {
            "post": {
                "description": "Add a new level",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "game"
                ],
                "summary": "Add a new level",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of the song, should have many",
                        "name": "songs",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "title of the level",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/game/levels/{id}": {
            "post": {
                "description": "Get a level",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "game"
                ],
                "summary": "Get a level",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of the level",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/game/lyrics/{song_id}": {
            "post": {
                "description": "Get lyrics with a song's ID; normally it is only for internal use.",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "game"
                ],
                "summary": "Get lyrics with a song's ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of the song",
                        "name": "song_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/game/songs/new": {
            "post": {
                "description": "Add a song to database",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "game"
                ],
                "summary": "Add a song",
                "parameters": [
                    {
                        "type": "file",
                        "description": "subtitle file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "youtube url",
                        "name": "url",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name of the song",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "singer of the song",
                        "name": "singer",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "miss lyrics index, seperated by commas",
                        "name": "miss_lyrics",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "genre of the song",
                        "name": "genre",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "language of the song, needs to be short, and seperated by commas if multiple languages",
                        "name": "language",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/game/songs/{id}": {
            "get": {
                "description": "Get a song instance, with a miss lyrics id randomly generated",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "game"
                ],
                "summary": "Get a song Instance",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of the song",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/game/tours/{id}": {
            "get": {
                "description": "Add a new level",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "game"
                ],
                "summary": "Add a new level",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of the game",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string",
                    "example": "success"
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
	Version:     "1.0",
	Host:        "ms.csie.org",
	BasePath:    "/api",
	Schemes:     []string{"http"},
	Title:       "Add a level API",
	Description: "To add a level into database",
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
