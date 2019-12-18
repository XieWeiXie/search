package seapi

// API...
type API struct {
	WeiBo  *WeiBo
	ZhiHu  *ZhiHu
	WeChat *WeChat
}

// WeiBo...
type WeiBo struct {
	Synthetically *WeiBoSynthetically
	User          WeiBoUser
	Passage       *WeiBoPassage
	Video         *WeiBoVideo
	Picture       WeiBoPicture
	Topic         WeiBoTopic
}

// ZhiHu...
type ZhiHu struct {
	Topic *ZhiHuTopic
}

// WeChat...
type WeChat struct {
	Article *WeChatArticle
	Account *WeChatAccount
}

func New(t Transport) *API {
	return &API{
		WeiBo: &WeiBo{
			User:    newWeiBoUser(t),
			Topic:   newWeiBoTopic(t),
			Picture: newWeiBoPicture(t),
		},
	}
}
