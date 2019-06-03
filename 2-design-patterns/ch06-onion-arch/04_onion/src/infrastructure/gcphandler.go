package infrastructure

func (handler *GcpHandler) FileExists(fileName string) (fileExists bool, err error) {
	br, err := handler.Client.Bucket(buckketName).Object(fullPath).NewReader(ctx)
}
