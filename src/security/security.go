package security

import "golang.org/x/crypto/bcrypt"

// FuncHash coloca um hash em uma string
func FuncHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ValidatePassword compara a senha do usuário com a senha hash do banco de dados
func ValidatePassword(passwordString, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}