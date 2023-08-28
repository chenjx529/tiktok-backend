package minio

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"net/url"
	"tiktok-backend/pkg/constants"
	"time"
)

var minioClient *minio.Client

// InitMinio 对象存储初始化
func InitMinio() {
	client, err := minio.New(constants.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinioAccessKeyId, constants.MinioSecretAccessKey, ""),
		Secure: constants.MinioUseSSL,
	})
	if err != nil {
		klog.Errorf("minio client init failed: %v", err)
	}

	minioClient = client

	if err := CreateBucket(context.Background(), constants.MinioVideoBucketName); err != nil {
		klog.Errorf("bucket create failed: %v", err)
	}
}

// CreateBucket 创建桶
func CreateBucket(ctx context.Context, bucketName string) error {
	if len(bucketName) <= 0 {
		return errors.New("bucketName invalid")
	}

	if err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "beijing"}); err != nil {
		if exists, errBucketExists := minioClient.BucketExists(ctx, bucketName); errBucketExists == nil && exists {
			klog.Infof("bucket %s already exists", bucketName)
			return nil
		}
		return err
	}

	klog.Info("bucket create successfully")
	return nil
}

// UploadLocalFile 上传本地文件（提供文件路径）至 minio
func UploadLocalFile(ctx context.Context, bucketName string, objectName string, filePath string, contentType string) error {
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return errors.New("UploadLocalFile failed")
	}
	klog.Infof("upload %s of size %d successfully", objectName, info.Size)
	return nil
}

// UploadFile 上传文件（提供reader）至 minio
func UploadFile(ctx context.Context, bucketName string, objectName string, reader io.Reader, objectsize int64) error {
	info, err := minioClient.PutObject(ctx, bucketName, objectName, reader, objectsize, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return errors.New("UploadFile fail")
	}
	klog.Infof("upload %s of bytes %d successfully", objectName, info.Size)
	return nil
}

// GetFileUrl 从 minio 获取文件Url
func GetFileUrl(ctx context.Context, bucketName string, fileName string, expires time.Duration) (*url.URL, error) {
	reqParams := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 * 24
	}
	presignedUrl, err := minioClient.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
	if err != nil {
		return nil, errors.New("get url fail")
	}
	return presignedUrl, nil
}
