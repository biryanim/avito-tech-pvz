package auth

import "golang.org/x/crypto/bcrypt"

func (s *serv) verifyPassword(hashedPassword string, candidatePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
	return err == nil
}

func (s *serv) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), err
}
