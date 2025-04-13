package lib

import (
	"github.com/brunowang/gframe/gfs3"
	"os"
)

func NewS3Mgr() *gfs3.S3Mgr {
	s3mgr := gfs3.NewS3Mgr(gfs3.AwsS3Config{
		AccessKey: os.Getenv("S3_AK"),
		SecretKey: os.Getenv("S3_SK"),
		Region:    os.Getenv("S3_REGION"),
	})
	return s3mgr
}
