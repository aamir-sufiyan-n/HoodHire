package utils

import (
	"context"
	"io"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var cld *cloudinary.Cloudinary

func InitCloudinary() error {
    var err error
    cld, err = cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
    return err
}

func UploadImage(file any) (string, error) {
	result,err:=cld.Upload.Upload(context.Background(),file,uploader.UploadParams{
		Folder: "hoodhire/profiles",
	})
	if err !=nil{
		return "",err
	}
	return result.SecureURL,nil
}


func UploadPDF(src io.Reader) (string, error) {
    result, err := cld.Upload.Upload(context.Background(), src, uploader.UploadParams{
        Folder:       "hoodhire/resumes",
        ResourceType: "raw",
    })
    if err != nil {
        return "", err
    }
    return result.SecureURL, nil
}