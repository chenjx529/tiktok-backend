package constants

import "time"

const (
	// jwt
	SecretKey   = "secret key"
	IdentityKey = "id"

	// rpc service
	ApiServiceName      = "api"
	FeedServiceName     = "feed"
	PublishServiceName  = "publish"
	UserServiceName     = "user"
	FavoriteServiceName = "favorite"
	CommentServiceName  = "comment"
	RelationServiceName = "relation"
	MessageServiceName  = "message"

	// rpc addr
	ApiServicePort      = ":8881"
	FeedServicePort     = ":8882"
	PublishServicePort  = ":8883"
	UserServicePort     = ":8884"
	FavoriteServicePort = ":8885"
	CommentServicePort  = ":8886"
	RelationServicePort = ":8887"
	MessageServicePort  = ":8888"

	// limit
	CPURateLimit = 80.0
	DefaultLimit = 10

	// MySQL
	MySQLMaxIdleConns    = 10        //空闲连接池中连接的最大数量
	MySQLMaxOpenConns    = 100       //打开数据库连接的最大数量
	MySQLConnMaxLifetime = time.Hour //连接可复用的最大时间
	MySQLDefaultDSN      = "root:123456@tcp(119.23.67.36:3307)/tiktok?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress          = "119.23.67.36:2379"
)
