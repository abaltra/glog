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
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v2/tenant/{tenantID}/posts": {
            "get": {
                "description": "Create a new post with an auto-generated ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new Post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tenant ID",
                        "name": "tenantID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Post to create",
                        "name": "{object}",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/post.CreatePostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/post.Post"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsehandler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responsehandler.Error"
                        }
                    }
                }
            }
        },
        "/v2/tenant/{tenantID}/posts/{slug}/publish": {
            "put": {
                "description": "Publshes an existing Post, adding the proper timestamps",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Publish a Post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tenant ID",
                        "name": "tenantID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Unique slug of the post",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsehandler.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responsehandler.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "post.CreatePostRequest": {
            "type": "object",
            "properties": {
                "Abstract": {
                    "type": "string"
                },
                "ContentRaw": {
                    "type": "string"
                },
                "Title": {
                    "type": "string"
                }
            }
        },
        "post.Post": {
            "type": "object",
            "properties": {
                "Abstract": {
                    "type": "string"
                },
                "AuthorID": {
                    "type": "string"
                },
                "ContentRaw": {
                    "type": "string"
                },
                "CreatedAt": {
                    "type": "string"
                },
                "ID": {
                    "type": "string"
                },
                "IsPublished": {
                    "type": "boolean"
                },
                "LastEditedBy": {
                    "type": "string"
                },
                "PublishedAt": {
                    "type": "string"
                },
                "Slug": {
                    "type": "string"
                },
                "Title": {
                    "type": "string"
                },
                "UpdatedAt": {
                    "type": "string"
                },
                "Version": {
                    "type": "integer"
                }
            }
        },
        "responsehandler.Error": {
            "description": "Error message with HTTP status code and error message",
            "type": "object",
            "properties": {
                "Message": {
                    "type": "string"
                },
                "StatusCode": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "blog.abaltra.me/api",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Glog - A Go Blogging backend using Mongo",
	Description:      "Very simple implementation of a bloggin platform using Mongo as a data store and Go with gorilla/mux.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}