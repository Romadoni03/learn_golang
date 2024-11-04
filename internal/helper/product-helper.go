package helper

import (
	"ecommerce-cloning-app/internal/logger"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func UploadPhotoProduct(imageInput string) (string, error) {
	imageData, errEncoding := base64.StdEncoding.DecodeString(imageInput)
	if errEncoding != nil {
		return "", errEncoding
	}

	folderPath := "D:/dev/portofolio/ecommerce-cloning-app/assets/images/product/"
	fileName := strconv.FormatInt(time.Now().Local().Unix(), 10) + "_.png"
	filePath := filepath.Join(folderPath, fileName)

	err := os.WriteFile(filePath, imageData, 0644)
	if err != nil {
		return "", err
	}
	logger.Logging().Info("success save product's photo")

	return fileName, nil
}

func GetImageProduct(image string) string {
	// imageDcd := DecodeImageName(image)
	imagePath := fmt.Sprintf("D:/dev/portofolio/ecommerce-cloning-app/assets/images/product/%s", image)
	fileImg, err := os.Open(imagePath)
	IfPanicError(err)
	defer fileImg.Close()

	imgData, errImg := io.ReadAll(fileImg)
	IfPanicError(errImg)

	encodedImg := base64.StdEncoding.EncodeToString(imgData)

	return encodedImg
}
