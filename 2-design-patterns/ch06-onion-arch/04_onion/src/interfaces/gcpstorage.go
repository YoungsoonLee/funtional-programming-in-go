package interfaces

import (
	"github.com/YoungsoonLee/funtional-programming-in-go/2-design-patterns/ch06-onion-arch/04_onion/src/domain"
)

type GcpHandler interface {
	ListBuckets(flowType domain.FlowType, projectID string) (buckets []domain.Bucket, err error)
	FileExists(fileName string) (fileExists bool, err error)
	DownloadFile(fileName string) (success bool, err error)
	UploadFile(fileName string) (success bool, err error)
}

type GcpRepo struct {
	gcpHandlers map[string]GcpHandler
	gcpHandler  GcpHandler
}

type SourceBucketRepo GcpRepo
type SinkBucketRepo GcpRepo

func NewSourceBucketRepo(gcpHandendlers map[string]GcpHandler) *SourceBucketRepo {
	bucketRepo := new(SourceBucketRepo)
	bucketRepo.gcpHandlers = gcpHandendlers
	bucketRepo.gcpHandler = gcpHandendlers["SourceBucketRepo"]
	return bucketRepo
}
