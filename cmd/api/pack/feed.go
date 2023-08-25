package pack

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"tiktok-backend/pkg/errno"
)

func SendFeedResponse(c *app.RequestContext, err error, videoList interface{}, nextTime int64) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, FeedResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		VideoList:  videoList,
		NextTime:   nextTime,
	})
}
