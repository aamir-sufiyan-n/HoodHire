package utils

import (
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(file any) (string, error) {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err!=nil{
		return "",err
	}
	result,err:=cld.Upload.Upload(context.Background(),file,uploader.UploadParams{
		Folder: "hoodhire/profiles",
	})
	if err !=nil{
		return "",err
	}
	return result.SecureURL,nil
}
