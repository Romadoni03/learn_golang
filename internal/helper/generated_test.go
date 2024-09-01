package helper_test

import (
	"ecommerce-cloning-app/internal/helper"
	"fmt"
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
	var times = time.Now()
	time1 := times.Local().Format("2006-01-02 15:04:05")

	timeSecond, _ := time.Parse(time.RFC3339, time1)
	fmt.Printf("time1 %v\n", time1)
	fmt.Println("hasil parse ", timeSecond)
}
