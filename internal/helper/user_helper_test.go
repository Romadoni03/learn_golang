package helper_test

import (
	"ecommerce-cloning-app/internal/helper"
	"fmt"
	"testing"
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

func TestEncodeDecodeImage(t *testing.T) {
	image := helper.EncodeImageName("account_profile.png")
	image2 := helper.EncodeImageName("account_profile.png")

	imageDecoded := helper.DecodeImageName("YWNjb3VudF9wcm9maWxlLnBuZw==")

	fmt.Printf("Hasil encode: %s \n", image)
	fmt.Printf("Hasil encode 2: %s \n", image2)
	fmt.Printf("Hasil decoded: %s", imageDecoded)
}
