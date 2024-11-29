package auth

import (
	"errors"
	"fmt"
	"golang-social-network-api/src/config"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken retorna um token assinado com as permissões do usuário
func CreateToken(userId uint64) (string, error) {

	// Permissões
	permissions := jwt.MapClaims{}
	permissions["authorized"]=true

	// Tempo de expiração
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId

	// Gera uma chave secret para fazer a assinatura do token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey)) // secret para assinar o token
}

// ValideteToken verifica se o token passado na requisição é válido
func ValideteToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnKeyValidate)
	if err != nil {
		return err 
	}

	// Verificação do token e se está expirado
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

// extractToken extrai o token
func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	
	// o retorno é Bearer e o que precisamos é o token
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnKeyValidate(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

// ExtractUserId extrai o id do usuário dentro do token
func ExtractUserId(r *http.Request) (uint64, error) {
	// Extrair o token do cabeçalho
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnKeyValidate)
	if err != nil {
		return 0, err 
	}

	fmt.Println(token)

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if err != nil {
			return 0, nil
		}

		return userId, nil
	}

	return 0, errors.New("token inválido")

}