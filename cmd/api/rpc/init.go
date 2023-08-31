package rpc

// InitRPC init rpc client
func InitRPC() {
	initUserRpc()
	initFeedRPC()
	initPublishRpc()
	initRelationRpc()
	initMessageRpc()
	initCommentRPC()
	initFavoriteRPC()
}
