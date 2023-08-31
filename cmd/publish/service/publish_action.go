package service

import (
	"bytes"
	"context"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"os"
	"strconv"
	"tiktok-backend/dal/db"
	"tiktok-backend/kitex_gen/publish"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/jwt"
	"tiktok-backend/pkg/minio"
	"time"
)

type PublishActionService struct {
	ctx context.Context
}

// NewPublishActionService new PublishActionService
func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{
		ctx: ctx,
	}
}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishActionService) PublishAction(req *publish.DouyinPublishActionRequest) error {
	// 登录id
	claims, err := jwt.GetClaimsFromTokenStr(req.Token)
	if err != nil {
		return err
	}
	loginId := int64(claims[constants.IdentityKey].(float64))

	// 上传视频
	videoData := req.Data
	videoReader := bytes.NewReader(videoData)
	randomStr := strconv.Itoa(int(time.Now().Unix()))
	fileName := randomStr + ".mp4"
	if err := minio.UploadFile(s.ctx, constants.MinioVideoBucketName, fileName, videoReader, int64(len(videoData))); err != nil {
		return err
	}

	// 获取视频链接
	playUrl, err := minio.GetFileUrl(s.ctx, constants.MinioVideoBucketName, fileName, 0)
	if err != nil {
		return err
	}

	// 获取封面
	coverPath := randomStr + ".jpg"
	coverData, err := readFrameAsJpeg(playUrl)
	if err != nil {
		return err
	}

	// 上传封面
	coverReader := bytes.NewReader(coverData)
	if err := minio.UploadFile(s.ctx, constants.MinioVideoBucketName, coverPath, coverReader, int64(len(coverData))); err != nil {
		return err
	}

	// 获取封面链接
	coverUrl, err := minio.GetFileUrl(s.ctx, constants.MinioVideoBucketName, coverPath, 0)
	if err != nil {
		return err
	}

	// 添加一个video记录
	if err := db.CreateVideo(s.ctx, &db.Video{
		UserId:        loginId,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.Title,
	}, loginId); err != nil {
		return err
	}

	return nil
}

// ReadFrameAsJpeg
// 从视频流中截取一帧并返回 需要在本地环境中安装ffmpeg并将bin添加到环境变量
func readFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return nil, err
	}

	return buf.Bytes(), err
}
