package interfaces

import "mime/multipart"

type IS3Service interface {
	UploadFile(*multipart.File, string) (*string, error)
}
