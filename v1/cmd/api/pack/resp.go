package pack

type User struct {
	Id              int64  `json:"id"`               // 用户id
	Name            string `json:"name"`             // 用户名称
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  int64  `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数量
	FavoriteCount   int64  `json:"favorite_count"`   // 点赞数量
}

type FriendUser struct {
	Id              int64  `json:"id"`               // 用户id
	Name            string `json:"name"`             // 用户名称
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  int64  `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数量
	FavoriteCount   int64  `json:"favorite_count"`   // 点赞数量
	Message         string `json:"message"`          // 和该好友的最新聊天消息
	MsgType         int64  `json:"msgType"`          // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
}

type Video struct {
	Id            int64  `json:"id"`             // 视频唯一标识
	Author        *User  `json:"author"`         // 视频作者信息
	PlayUrl       string `json:"play_url"`       // 视频播放地址
	CoverUrl      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	Title         string `json:"title"`          // 视频标题
}

type Comment struct {
	Id         int64  `json:"id"`          // 视频评论id
	User       *User  `json:"user"`        // 评论用户信息
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
}

type Message struct {
	Id         int64  `json:"id"`           // 消息id
	FromUserId int64  `json:"from_user_id"` // 该消息发送者的id
	ToUserId   int64  `json:"to_user_id"`   // 该消息接收者的id
	Content    string `json:"content"`      // 消息内容
	CreateTime string `json:"create_time"`  // 消息创建时间
}

type RelationActionResponse struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type RelationFollowListResponse struct {
	StatusCode int32   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	UserList   []*User `json:"user_list"`   // 用户信息列表
}

type RelationFollowerListResponse struct {
	StatusCode int32   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	UserList   []*User `json:"user_list"`   // 用户列表
}

type RelationFriendListResponse struct {
	StatusCode int32         `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string        `json:"status_msg"`  // 返回状态描述
	UserList   []*FriendUser `json:"user_list"`   // 用户列表
}

type MessageChatResponse struct {
	StatusCode  int32      `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string     `json:"status_msg"`   // 返回状态描述
	MessageList []*Message `json:"message_list"` // 消息列表
}

type MessageActionResponse struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type FavoriteActionResponse struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type FavoriteListResponse struct {
	StatusCode int32    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string   `json:"status_msg"`  // 返回状态描述
	VideoList  []*Video `json:"video_list"`  // 用户点赞视频列表
}

type UserRegisterResponse struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	UserId     int64  `json:"user_id"`     // 用户id
	Token      string `json:"token"`       // 用户鉴权token,token放在api层
}

type UserInfoResponse struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	User       *User  `json:"user"`        // 用户信息
}

type FeedResponse struct {
	StatusCode int32    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string   `json:"status_msg"`  // 返回状态描述
	VideoList  []*Video `json:"video_list"`  // 视频列表
	NextTime   int64    `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

type PublishActionResponse struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type PublishListResponse struct {
	StatusCode int32    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string   `json:"status_msg"`  // 返回状态描述
	VideoList  []*Video `json:"video_list"`  // 用户发布的视频列表
}

type CommentActionResponse struct {
	StatusCode int32    `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string   `json:"status_msg"`  // 返回状态描述
	Comment    *Comment `json:"comment"`     // 评论id
}

type CommentListResponse struct {
	StatusCode  int32      `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string     `json:"status_msg"`   // 返回状态描述
	CommentList []*Comment `json:"comment_list"` // 用户发布的视频列表
}
