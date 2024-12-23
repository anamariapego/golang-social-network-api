package docs

import "github.com/swaggo/swag"

const docTemplate = `
{
    "swagger": "2.0",
    "info": {
        "version": "1.0.0",
        "title": "Golang Social Network API",
        "description": "API simula funcionalidades básicas de uma rede social, como criação de publicações e interação entre usuários (seguir, parar de seguir, curtir publicações)"
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
     "paths": {
     "/login": {
            "post": {
                "summary": "Realiza o login de um usuário",
                "description": "Autentica o usuário com email e senha e retorna um token JWT.",
                "parameters": [
                    {
                        "name": "login",
                        "in": "body",
                        "description": "Dados de login do usuário",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login bem-sucedido",
                        "schema": {
                            "type": "object",
                            "properties": {
                                "id": {
                                    "type": "string",
                                    "example": "1"
                                },
                                "token": {
                                    "type": "string",
                                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
                                }
                            }
                        }
                    },
                    "401": {
                        "description": "Credenciais inválidas"
                    }
                }
            }
        },
        "/users": {
            "post": {
                "summary": "Cria um novo usuário",
                "description": "Cria um usuário e armazena no banco de dados PostgreSQL",
                "parameters": [
                    {
                        "name": "user",
                        "in": "body",
                        "description": "Dados do usuário",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Usuário criado com sucesso",
                        "schema": {
                            "type": "string",
                            "example": "Usuário criado com sucesso: id 1"
                        }
                    }
                }
            },
            "get": {
                "summary": "Lista todos os usuários",
                "description": "Retorna uma lista de todos os usuários cadastrados no banco de dados",
                "responses": {
                    "200": {
                        "description": "Lista de usuários retornada com sucesso",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/UserGet"
                            }
                        }
                    }
                },
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            }
        },
        "/users/{id}": {
            "get": {
                "summary": "Obtém detalhes de um usuário",
                "description": "Retorna os detalhes de um usuário específico com base no Id. Requer autenticação com token JWT.",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Id do usuário",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Usuário retornado com sucesso",
                        "schema": {
                            "$ref": "#/definitions/UserGet"
                        }
                    },
                    "401": {
                        "description": "Token inválido ou ausente"
                    }
                },
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            },
            "put": {
                "summary": "Atualiza as informações de um usuário",
                "description": "Atualiza o usuário com base no Id fornecido.",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Id do usuário a ser atualizado",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "user",
                        "in": "body",
                        "description": "Dados atualizados do usuário",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UserPut"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Informações do usuário atualizadas com sucesso",
                        "schema": {
                            "$ref": "#/definitions/User"
                        }
                    },
                    "401": {
                        "description": "Token inválido ou ausente"
                    }
                },
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            },
            "delete": {
                "summary": "Deleta as informações de um usuário",
                "description": "Deleta o usuário com base no Id fornecido. Requer autenticação com token JWT.",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Id do usuário a ser deletado",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Usuário deletado com sucesso"
                    },
                    "401": {
                        "description": "Token inválido ou ausente"
                    },
                    "404": {
                        "description": "Usuário não encontrado"
                    }
                },
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            }
        },
        "/users/{id}/follower": {
            "post": {
                "summary": "Segue um usuário",
                "description": "Permite que um usuário siga outro usuário. O Id especificado corresponde ao usuário a ser seguido.",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Id do usuário a ser seguido",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Usuário seguido com sucesso"
                    },
                    "403": {
                        "description": "Não é possível seguir você mesmo"
                    },
                    "404": {
                        "description": "Erro ao seguir o usuário"
                    }
                },
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            }
        },
        "/users/{id}/stop-follower": {
            "post": {
                "summary": "Deixar de seguir um usuário",
                "description": "Permite que um usuário pare de seguir outro usuário. O Id especificado corresponde ao usuário que será deixado de seguir..",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Id do usuário a ser deixado de seguir",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Usuário deixado de seguir com sucesso"
                    },
                    "403": {
                        "description": "Não é possível deixar de seguir você mesmo"
                    },
                    "404": {
                        "description": "Erro ao deixar de seguir o usuário"
                    }
                },
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            }
        },
        "/users/{id}/followers": {
            "get": {
                "summary": "Buscar os seguidores de um usuário",
                "description": "Permite buscar todos os seguidores de um usuário pelo Id.",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Id do usuário para buscar seus seguidores",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lista de seguidores retornada com sucesso"
                    },
                    "404": {
                        "description": "Usuário não encontrado"
                    },
                    "500": {
                        "description": "Erro interno ao buscar seguidores"
                    }
                },
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            }
        },
        "/users/{id}/following": {
            "get": {
                "summary": "Buscar retornar os usuários que um usuário específico segue",
                "description": "pelo Id.",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Id do usuário para buscar os usuários que ele segue",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Lista de usuários seguidos retornada com sucesso"
                    },
                    "404": {
                        "description": "Usuário não encontrado"
                    },
                    "500": {
                        "description": "Erro interno ao buscar os usuários seguidos"
                    }
                },
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            }
        },
        "/users/{id}/update-password": {
            "post": {
                "summary": "Atualizar a senha do usuário",
                "description": "Atualiza a senha de um usuário pelo Id.",
                "parameters": [
                {
                    "name": "id",
                    "in": "path",
                    "description": "Id do usuário que vai atualizar a senha",
                    "required": true,
                    "type": "string"
                },
                {
                    "name": "body",
                    "in": "body",
                    "description": "Estrutura para atualizar a senha do usuário",
                    "required": true,
                    "schema": {
                    "$ref": "#/definitions/Password"
                    }
                }
                ],
                "responses": {
                "200": {
                    "description": "Senha atualizada com sucesso",
                    "content": {
                    "application/json": {
                        "example": {
                        "message": "Senha atualizada com sucesso"
                        }
                    }
                    }
                },
                "403": {
                    "description": "Não é possível atualizar senha de outro usuário"
                },
                "404": {
                    "description": "Usuário não encontrado"
                },
                "500": {
                    "description": "Erro interno ao atualizar a senha"
                }
                },
                "security": [
                {
                    "BearerAuth": []
                }
                ]
}},
            "/users/{id}/publications": {
            "get": {
                "summary": "Retorna todas as publicações do usuário",
                "description": "Retorna as publicações de um usuário específico pelo seu Id",
                "parameters": [
                {
                    "name": "id",
                    "in": "path",
                    "description": "Id do usuário que terá suas publicações retornadas",
                    "required": true,
                    "type": "string"
                },
                ],
                "responses": {
                "200": {
                "description": "Publicações retornadas com sucesso"
                },
                "403": {
                "description": "Usuário não autorizado a visualizar publicações"
            },
            "404": {
                "description": "Usuário não encontrado"
            },
            "500": {
                "description": "Erro interno ao recuperar publicações"
            }
                },
                "security": [
                {
                    "BearerAuth": []
                }
                ]
            }
        },
        "/publications": {
            "post": {
                "summary": "Cria uma publicação ",
                "description": "",
                "parameters": [
                    {
                        "name": "publications",
                        "in": "body",
                        "description": "",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Publications"
                        }
                    }
                ],
                "responses": {
                    "200": {
                    "description": "Publicação criada com sucesso",
                },
                "401": {
                        "description": "Credenciais inválidas"
                    }
                },
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            },
            "get": {
                "summary": "Retorna as publicações do usário e de seus seguidores",
                "description": "Retorna uma lista com todas as publicações do usuário e de seus seguidores",
                "parameters": [
                ],
                "responses": {
                    "200": {
                        "description": "Publicações retornadas com sucesso"
                    },
                    "401": {
                        "description": "Usuário não possui publicações"
                    }
                },
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            }
        },
        "/publications/{id}": {
            "get": {
                "summary": "Retorna uma publicação",
                "description": "Busca um publicação com base no Id fornecido",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Id da publicaçãos",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Publicações retornadas com sucesso"
                    },
                    "401": {
                        "description": "Não encontrou nenhuma publicação"
                    }
                },
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            },
            "put": {
                "summary": "Atualizar uma publicação",
                "description": "Atualiza uma publicação com base no Id fornecido",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Id da publicação a ser atualizada",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "publication",
                        "in": "body",
                        "description": "Publicação atualizada",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/PublicationPut"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Publicação atualizada com sucesso"
                    },
                    "401": {
                        "description": "Não encontrou nenhuma publicação"
                    }
                },
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            },
            "delete": {
                "summary": "Deleta uma publicação",
                "description": "Deleta uma publicação com base no Id fornecido",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Id da publicação a ser deletada",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Publicação deletada com sucesso"
                    },
                    "401": {
                        "description": "Não encontrou nenhuma publicação"
                    }
                },
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            }
        },
        "/publications/{id}/like": {
            "post": {
                "summary": "Curtir uma publicação",
                "description": "Permite ao usuário curtir uma publicação específica",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Id da publicação que será descurtida",
                        "required": true,
                        "type": "string"
                    }
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            }
        },
        "/publications/{id}/dislike": {
            "post": {
                "summary": "Descutir uma publicação ",
                "description": "Permite ao usuário remover o like de uma publicação específica",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Id da publicação que será descurtida",
                        "required": true,
                        "type": "string"
                    }
                ],
                "security": [
                    {
                        "BearerAuth": []
                    }
                ]
            }
        },
    },
    "definitions": {
        "User": {
            "type": "object",
            "properties": {
                "nameUser": {
                    "type": "string",
                    "example": "Jose Oliveira"
                },
                "nick": {
                    "type": "string",
                    "example": "jose_o"
                },
                "email": {
                    "type": "string",
                    "example": "jose.oliveira@gmail.com"
                },
                "passwordUser": {
                    "type": "string",
                    "example": "4321"
                }
            }
        },
        "UserGet": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "number",
                    "example": 1
                },
                "nameUser": {
                    "type": "string",
                    "example": "Jose Oliveira"
                },
                "nick": {
                    "type": "string",
                    "example": "jose_o"
                },
                "email": {
                    "type": "string",
                    "example": "jose.oliveira@gmail.com"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2024-12-08T19:17:15.210043Z"
                }
            }
        },
        "UserPut": {
            "type": "object",
            "properties": {
                "nameUser": {
                    "type": "string",
                    "example": "Jose Oliveira"
                },
                "nick": {
                    "type": "string",
                    "example": "jose_o"
                },
                "email": {
                    "type": "string",
                    "example": "jose.oliveira@gmail.com"
                },
                "passwordUser": {
                    "type": "string",
                    "example": "4321"
                }
            }
        },
        "Login": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "jose.oliveira@gmail.com"
                },
                "passwordUser": {
                    "type": "string",
                    "example": "4321"
                }
            },
            "required": ["email", "passwordUser"]
        },
        "Password": {
            "type": "object",
            "properties": {
            "current": {
                "type": "string",
                "example": "1234"
            },
            "new": {
                "type": "string",
                "example": "4321"
            }
            },
            "required": ["current", "new"]
        },
        "Publications": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string",
                    "example": "O Pequeno Princípe"
                },
                "text": {
                    "type": "string",
                    "example": "O livro trata de uma história infantil"
                }
            }
        },
        "PublicationPut": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string",
                    "example": "O Pequeno Princípe"
                },
                "text": {
                    "type": "string",
                    "example": "O livro trata de uma história infantil"
                }
            }
        },
    }
}`

var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:2468",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Golang Social Network API",
	Description:      "API simula funcionalidades básicas de uma rede social, como criação de publicações e interação entre usuários (seguir, parar de seguir, curtir publicações)",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
