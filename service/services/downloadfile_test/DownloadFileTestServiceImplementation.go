package service_downloadfile_test

type DownloadFileTestServiceImplementation struct {
}

func (s *DownloadFileTestServiceImplementation) GetFile() (*Response, error) {
	fileContent := "Hello World. This is a test content"

	return &Response{FileByte: []byte(fileContent)}, nil
}

func NewDownloadFileTestServiceImplementation() DownloadFileTestService {
	return &DownloadFileTestServiceImplementation{}
}
