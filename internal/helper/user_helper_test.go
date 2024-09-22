package helper_test

import (
	"ecommerce-cloning-app/internal/helper"
	"fmt"
	_ "image/png"
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
	// Panggil fungsi GetImage
	img := helper.GetImage("account_profile.png")

	fmt.Println(img)
}
