// Code generated by swaggo/swag. DO NOT EDIT
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
            "email": "fiber@swagger.io"
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
        "/attendances": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get list attendance",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Attendance"
                ],
                "summary": "Get list attendance details api",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page Size",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "default": "desc",
                        "description": "Sort direction",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Class id",
                        "name": "classId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Time range",
                        "name": "period",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/attendances.Attendance"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
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
                "description": "Insert attendance",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Attendance"
                ],
                "summary": "Create attendance api",
                "parameters": [
                    {
                        "description": "Create attendance body",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/attendances.CreateAttendanceRequest"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/attendances/{id}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Patch attendance",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Attendance"
                ],
                "summary": "Patch attendance api",
                "parameters": [
                    {
                        "description": "Patch attendance body",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/attendances.UpdateAttendanceRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Attendance ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/attendances.Attendance"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/registrations": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get list registration",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Registration"
                ],
                "summary": "Get list registration details api",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page Size",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "default": "desc",
                        "description": "Sort direction",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search term",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/registrations.FindRegistrationResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
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
                "description": "Insert registration",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Registration"
                ],
                "summary": "Create registration api",
                "parameters": [
                    {
                        "description": "Create registration body",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/registrations.WriteRegistrationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/registrations.Registration"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
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
                "description": "Delete registration",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Registration"
                ],
                "summary": "Delete registration api",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "description": "Registration IDs",
                        "name": "ids",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/registrations/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get one registration",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Registration"
                ],
                "summary": "Get registration details api",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Registration ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/registrations.Registration"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
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
                "description": "Update registration",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Registration"
                ],
                "summary": "Update registration api",
                "parameters": [
                    {
                        "description": "Update registration body",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/registrations.WriteRegistrationRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Registration ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/registrations.Registration"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Mark a registration as processed",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Registration"
                ],
                "summary": "Mark as done registration api",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Registration ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "attendances.Attendance": {
            "type": "object",
            "properties": {
                "attendedAt": {
                    "type": "string"
                },
                "classId": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isAttended": {
                    "type": "boolean"
                },
                "studentId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "attendances.CreateAttendanceRequest": {
            "type": "object",
            "properties": {
                "attendedAt": {
                    "type": "string"
                },
                "classId": {
                    "type": "integer"
                },
                "isAttended": {
                    "type": "boolean"
                },
                "studentId": {
                    "type": "integer"
                }
            }
        },
        "attendances.UpdateAttendanceRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isAttended": {
                    "type": "boolean"
                }
            }
        },
        "registrations.FindRegistrationResp": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/registrations.Registration"
                    }
                },
                "page": {
                    "type": "integer"
                }
            }
        },
        "registrations.Registration": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isProcessed": {
                    "type": "boolean"
                },
                "note": {
                    "type": "string"
                },
                "parentName": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "studentClass": {
                    "description": "Class type:\n* buds - Children who is 3 yo.\n* seed - Children who is 4 yo.\n* leaf - Children who is 5 yo.",
                    "type": "string",
                    "enum": [
                        "buds",
                        "seed",
                        "leaf"
                    ]
                },
                "studentDob": {
                    "type": "string"
                },
                "studentName": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "registrations.WriteRegistrationRequest": {
            "type": "object",
            "properties": {
                "note": {
                    "type": "string"
                },
                "parentName": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "studentClass": {
                    "description": "Class type:\n* buds - Children who is 3 yo.\n* seed - Children who is 4 yo.\n* leaf - Children who is 5 yo.",
                    "type": "string",
                    "enum": [
                        "buds",
                        "seed",
                        "leaf"
                    ]
                },
                "studentDob": {
                    "type": "string"
                },
                "studentName": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Apply \"bearer \" before token in authorization",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Fiber Example API",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
