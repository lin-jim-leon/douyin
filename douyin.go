package douyin

import (
	"github.com/lin-jim-leon/douyin/open/oauth"
	"github.com/lin-jim-leon/douyin/open/user"
	"github.com/lin-jim-leon/douyin/open/video"
)

// GetAllScope获取授权范围
func GetAllScope() string {
	return oauth.GetAllScope()
}

// GetOAuthRedirectURL 获取授权链接的函数
func GetOAuthRedirectURL(redirectURL string, clientKey string, scopes string) string {
	return oauth.GetRedirectURL(redirectURL, clientKey, scopes)
}

// GetAccessToken 通过授权码换取 access token 的函数
func GetAccessToken(clientKey string, clientSecret string, code string) (oauth.AccessTokenRes, error) {
	return oauth.GetAccessToken(clientKey, clientSecret, code)
}

// 刷新 AccessToken 的函数
func RefreshAccessToken(refreshKey string, clientKey string) (oauth.AccessTokenRes, error) {
	return oauth.RefreshAccessToken(refreshKey, clientKey)
}

// GetUserInfo用户头像昵称
func GetUserInfo(openId string, accesstoken string) (user.UserInfoRes, error) {
	return user.GetUserInfo(openId, accesstoken)
}

// GetUserFans粉丝数获取
func GetUserFans(accesstoken string, date string, open_id string) (user.UserfanRes, error) {
	return user.GetUserFans(accesstoken, date, open_id)
}

// GetUserLike点赞数
func GetUserLike(accesstoken string, date string, open_id string) (user.UserLikeRes, error) {
	return user.GetUserLike(accesstoken, date, open_id)
}

// GetUserComments评论数
func GetUserComments(accesstoken string, date string, open_id string) (user.UserCommentRes, error) {
	return user.GetUserComments(accesstoken, date, open_id)
}

// GetUserShare分享数
func GetUserShare(accesstoken string, date string, open_id string) (user.UserShareRes, error) {
	return user.GetUserShare(accesstoken, date, open_id)
}

// 短链解长链
func GetShortUrl(longUrl string) (shortUrl string) {
	return video.GetShortUrl(longUrl)
}

// GetVideoInfo视频信息
func GetVideoInfo(accesstoken string, openid string, video_id string) (video.VideoResponse, error) {
	return video.GetVideoInfo(accesstoken, openid, video_id)
}

// GetVideolistInfo视频列表信息
func GetVideolistInfo(accesstoken string, openid string, count int32, cursor int64) (video.VideoListResponse, error) {
	return video.GetVideolistInfo(accesstoken, openid, count, cursor)
}
