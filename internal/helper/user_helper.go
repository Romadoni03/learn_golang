package helper

import (
	"encoding/base64"
	"image"
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

func EncodeImageName(data string) string {
	encode := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encode, []byte(data))
	encodeImage := string(encode)

	return encodeImage
}

func DecodeImageName(data string) string {
	decode := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	_, err := base64.StdEncoding.Decode(decode, []byte(data))
	IfPanicError(err)

	decoded := string(decode)

	return decoded
}

func GetImage() image.Image {
	file, err := os.Open(`D:\dev\portofolio\ecommerce-cloning-app\assets\images\photo_profile\account_profile.png`)
	IfPanicError(err)
	defer file.Close()

	img, _, err := image.Decode(file)

	IfPanicError(err)
	// img.Bounds()
	// fmt.Println(img.Bounds())
	return img
}
