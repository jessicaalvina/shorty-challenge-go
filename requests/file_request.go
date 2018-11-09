package requests

import "mime/multipart"

type FileRequest struct {
	UploadFile struct {
		File     *multipart.FileHeader `form_file:"file"`
		FileName string                `form:"file_name"`
		SavePath string                `form:"path"`
	}
}
