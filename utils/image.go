package utils

import (
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// SaveImageToFolder saves an uploaded image file to the specified destination folder with the given filename.
// Parameters:
// - dst: the destination folder path (e.g., "uploads/images")
// - filename: the name to save the file as (e.g., "image123.jpg")
// - file: the uploaded file from a multipart/form request (usually from c.FormFile in Gin)
// Returns:
// - error: nil if success, otherwise an error describing what went wrong
func SaveImageToFolder(ctx *gin.Context, dst string, file *multipart.FileHeader) error {
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		return err
	}
	return nil
}

// GenerateDst create destination file
// Parameters:
// - path: folder upload path
// - fileName: image file name
// Returns: -> string
func GenerateDst(path, fileName string) string {
	return filepath.Join(path, fileName)
}

func createNewUploadFolder() {
	log.Println("Starting to create new folder upload image")
	if err := os.MkdirAll("upload", os.ModePerm); err != nil {
		log.Println("Error during creating new folder upload image")
	}
}

func ImageUploadConfig(path string) error {
	// Make sure the dir upload exists
	if path == "" {
		path = "/upload"
	}
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		// if the folder not exist, create new
		log.Println("Folder upload not exists, start to create a new one")
		createNewUploadFolder()
	}
	return nil
}
