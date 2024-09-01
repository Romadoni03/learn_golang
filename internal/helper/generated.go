package helper

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

func GenerateId() string {
	b := make([]rune, 20)
	for i := range b {
		b[i] = letter[randomizer.Intn(len(letter))]
	}
	return string(b)
}

func GeneratedUsername() string {
	b := make([]rune, 8)
	for i := range b {
		b[i] = letter[randomizer.Intn(len(letter))]
	}
	return string(b)
}

func HashingPassword(pw string) string {
	password := []byte(pw)
	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	IfPanicError(err)

	return string(hashed)
}

func CompiringPassword(hashPassword, password string) error {
	hash := []byte(hashPassword)
	pw := []byte(password)
	err := bcrypt.CompareHashAndPassword(hash, pw)
	if err != nil {
		return err
	}
	return nil
}
