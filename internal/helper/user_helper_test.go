package helper_test

import (
	"ecommerce-cloning-app/internal/helper"
	"encoding/base64"
	"fmt"
	_ "image/png"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

func TestHashing(t *testing.T) {
	password := "admin"
	fmt.Println(password)

	hashpw := helper.HashingPassword(password)
	fmt.Println("hashing", hashpw)

	error := helper.CompiringPassword(hashpw, password)

	if error != nil {
		fmt.Println("gagal")
	} else {
		fmt.Println("Berhasil")
	}
}

func TestGeneratedId(t *testing.T) {
	id1 := helper.GenerateId()
	id2 := helper.GenerateId()
	id31 := helper.GenerateId()

	fmt.Println("id 1 : ", id1)
	fmt.Println("id 2 : ", id2)
	fmt.Println("id 31 : ", id31)
}

func TestGeneratedUsername(t *testing.T) {
	user1 := helper.GeneratedUsername()
	user2 := helper.GeneratedUsername()
	user3 := helper.GeneratedUsername()

	fmt.Println("user 1 : ", user1)
	fmt.Println("user 2 : ", user2)
	fmt.Println("user 3 : ", user3)
}

func TestTime(t *testing.T) {
	times := helper.GeneratedTimeNow()

	fmt.Printf("%v", times)
}

func TestExpiredAt(t *testing.T) {
	myTime := time.Now().Local()

	mili := myTime.UnixMilli()

	fmt.Println(mili)
	fmt.Println("data 5 menit  120000 : ", 2*60*1000)
}

func TestGetImage(t *testing.T) {
	img := helper.GetImage("bg.jpg")

	fmt.Println(img)
}

func TestUploadImage(t *testing.T) {
	imagePath := "D:/dev/portofolio/ecommerce-cloning-app/assets/images/testImage/bg.jpg"
	fileImg, err := os.Open(imagePath)
	if err != nil {
		fmt.Println(err)
	}
	defer fileImg.Close()

	imgData, errImg := io.ReadAll(fileImg)
	if errImg != nil {
		fmt.Println(errImg)
	}

	encodedImg := base64.StdEncoding.EncodeToString(imgData)
	name, errUpload := helper.UploadPhotoProfile(encodedImg)
	if errUpload != nil {
		fmt.Println(errUpload)
	}

	fmt.Println(encodedImg)
	fmt.Println(name)
}

func TestEncoded(t *testing.T) {
	// Buka file gambar
	file, err := os.Open("D:/dev/portofolio/ecommerce-cloning-app/assets/images/photo_profile/1730622800_.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Baca isi file gambar
	imgData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Encode ke format Base64
	base64Str := base64.StdEncoding.EncodeToString(imgData)

	// Cetak hasil encode
	fmt.Println("Base64 Encoded Image:")
	fmt.Println(base64Str)
}

func TestCoba(t *testing.T) {
	// Buka file gambar
	file, err := os.Open("D:/dev/portofolio/ecommerce-cloning-app/assets/images/testImage/me.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Baca isi file gambar
	imgData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	base64Str := base64.StdEncoding.EncodeToString(imgData)
	fmt.Println(base64Str)
}
