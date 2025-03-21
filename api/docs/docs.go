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
            "name": "AGPL 3.0",
            "url": "https://www.gnu.org/licenses/agpl-3.0.en.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/apikeys": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create apikey",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "apikeys"
                ],
                "summary": "Create apikey",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/io.CreateAPIKeyResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/io.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete apikey",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "apikeys"
                ],
                "summary": "Delete apikey",
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/io.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/apikeys/status": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Check if the user has an apikey",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "apikeys"
                ],
                "summary": "Check apikey Status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/io.APIKeyStatusResponse"
                        }
                    }
                }
            }
        },
        "/projects": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "List projects",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "List projects",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Project"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Create a project",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/io.CreateProjectRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Project"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/io.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/projects/{project_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get a project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Get a project",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "project_id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Project"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/io.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Update a project",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "project_id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/io.UpdateProjectRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Project"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/io.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/io.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete a projects",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Delete a projects",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "project_id",
                        "name": "project_id",
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
                            "$ref": "#/definitions/io.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/projects/{project_id}/sprints": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "List sprints",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sprints"
                ],
                "summary": "List sprints",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "project_id",
                        "name": "project_id",
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
                                "$ref": "#/definitions/domain.Sprint"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/io.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/projects/{project_id}/sprints/{sprint_id}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a sprint",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sprints"
                ],
                "summary": "Update a sprint",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "project_id",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "sprint_id",
                        "name": "sprint_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/io.UpdateSprintRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Sprint"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/io.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/sign_in": {
            "post": {
                "description": "Sign in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign in",
                "parameters": [
                    {
                        "description": "sign in request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/io.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/io.SignInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/io.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/sign_up": {
            "post": {
                "description": "Sign up",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign up",
                "parameters": [
                    {
                        "description": "sign up request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/io.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/io.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Project": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "sprint_count": {
                    "type": "integer"
                },
                "sprint_duration": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "total_sp": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "domain.Sprint": {
            "type": "object",
            "properties": {
                "actual_sp": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "ideal_sp": {
                    "type": "integer"
                },
                "project_id": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "description": "always must be hashed",
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "io.APIKeyStatusResponse": {
            "type": "object",
            "properties": {
                "exists": {
                    "description": "APIキーが存在するかどうか",
                    "type": "boolean"
                }
            }
        },
        "io.CreateAPIKeyResponse": {
            "type": "object",
            "properties": {
                "raw_key": {
                    "type": "string"
                }
            }
        },
        "io.CreateProjectRequest": {
            "type": "object",
            "required": [
                "sprint_count",
                "sprint_duration",
                "start_date",
                "title",
                "total_sp"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 500
                },
                "sprint_count": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 1
                },
                "sprint_duration": {
                    "type": "integer",
                    "enum": [
                        1,
                        2,
                        3
                    ]
                },
                "start_date": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 100
                },
                "total_sp": {
                    "type": "integer",
                    "maximum": 1000
                }
            }
        },
        "io.ErrorDetail": {
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
        "io.ErrorResponse": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/io.ErrorDetail"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "io.SignInRequest": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8
                }
            }
        },
        "io.SignInResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "io.SignUpRequest": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8
                }
            }
        },
        "io.UpdateProjectRequest": {
            "type": "object",
            "required": [
                "sprint_count",
                "title",
                "total_sp"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 500
                },
                "sprint_count": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 1
                },
                "title": {
                    "type": "string",
                    "maxLength": 100
                },
                "total_sp": {
                    "type": "integer",
                    "maximum": 1000
                }
            }
        },
        "io.UpdateSprintRequest": {
            "type": "object",
            "properties": {
                "actual_sp": {
                    "type": "integer",
                    "maximum": 1000
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "value must be \"ApiKey {API_KEY}\"\"",
            "type": "apiKey",
            "name": "Authorization",
            "in": "Header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Burny API",
	Description:      "Burny Backend API Doc. \\nNOTE: JWT Auhtentication is also supported but not displayed due to swagger generaton tool limitation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
