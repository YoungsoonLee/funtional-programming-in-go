package domain

type File struct {
	ID       int
	Name     string  `json:"name"`
	ErrorMsg string  `json:"error"`
	Contents LogFile `json:"logFile"`
	Bytes    []byte  `json:"bytes"`
}

type CloudFile struct {
	Name string `json:"name"`
}

type CloudFiles struct {
	Names []CloudFile
}

type CloudPath struct {
	Path string `json:"path"`
}

type CloudPaths struct {
	Paths []CloudPath
}
