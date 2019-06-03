package domain

type (
	HostProvider int
	FlowType     int
)

type CloudStorage struct {
	HostProvider HostProvider
	ProjectID    string
	FlowType     FlowType
}

type LocalRepository interface {
	FileExists(fileName string) (fileExists bool, err error)
}

type BucketRepository interface {
	List(projectID string) (buckets []Bucket, err error)
	FileExists(fileName string) (fileExists bool, err error)
	DownloadFile(fileName string) (success bool, err error)
	UploadFile(fileName string) (success bool, err error)
}

type FileRepository interface {
	Store(file File)
	FindById(id int) File
}

type Bucket struct {
	Name string `json:"name"`
}

type Buckets struct {
	Buckets []Bucket `json:"buckets"`
}
