package pack

type RelationActionRequest struct {
	Token      string `json:"token"`       // 用户鉴权token
	ToUserId   string `json:"to_user_id"`  // 对方用户id
	ActionType string `json:"action_type"` // 1-关注，2-取消关注
}
