package common

import (
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

func HashPassword(password string) (hashedPassword string, err error) {
	byteHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(byteHashed), nil
}

func CheckPassword(hashedPassword, password string) (matches bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func HashPin(pin int64) (hashedPassword string, err error) {
	sPin := strconv.FormatInt(pin, 10)
	byteHashed, err := bcrypt.GenerateFromPassword([]byte(sPin), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(byteHashed), nil
}

func CheckPin(hashedPin string, pin int64) (matches bool) {
	sPin := strconv.FormatInt(pin, 10)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPin), []byte(sPin))
	if err != nil {
		return false
	}

	return true
}
