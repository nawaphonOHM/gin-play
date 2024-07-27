package service_downloadfile_test

type Response struct {
	FileByte []byte
}

type DownloadFileTestService interface {
	GetFile() (*Response, error)
}
