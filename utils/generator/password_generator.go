package generator

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
