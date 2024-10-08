package helper

import (
	"encoding/base64"
	"io"
	"math/rand"
	"os"
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

func GeneratedTimeNow() time.Time {
	t := time.Now().Local().Format("2006-01-02 15:04:05")
	layoutFormat := "2006-01-02 15:04:05"

	date, err := time.Parse(layoutFormat, t)

	IfPanicError(err)

	return date
}

func GetImage(image string) string {
	// imageDcd := DecodeImageName(image)
	fileImg, err := os.Open(`D:\dev\portofolio\ecommerce-cloning-app\assets\images\photo_profile\` + image)
	IfPanicError(err)
	defer fileImg.Close()

	imgData, errImg := io.ReadAll(fileImg)
	IfPanicError(errImg)

	encodedImg := base64.StdEncoding.EncodeToString(imgData)

	return encodedImg
}

func LastUpdateUsername(data time.Time) int64 {

	dataString := data.String()

	layout := "2006-01-02 15:04:05"

	t, err := time.Parse(layout, dataString)
	IfPanicError(err)

	milliseconds := t.UnixMilli()

	return milliseconds
}
