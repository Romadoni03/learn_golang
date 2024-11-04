package helper

import (
	"database/sql"
	"ecommerce-cloning-app/internal/logger"
	"encoding/base64"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func UploadLogoStore(imageInput string) (string, error) {
	imageData, errEncoding := base64.StdEncoding.DecodeString(imageInput)
	if errEncoding != nil {
		return "", errEncoding
	}

	folderPath := "D:/dev/portofolio/ecommerce-cloning-app/assets/images/logo_store/"
	fileName := strconv.FormatInt(time.Now().Local().Unix(), 10) + "_.png"
	filePath := filepath.Join(folderPath, fileName)

	err := os.WriteFile(filePath, imageData, 0644)
	if err != nil {
		return "", err
	}
	logger.Logging().Info("success save image store")

	return fileName, nil
}
