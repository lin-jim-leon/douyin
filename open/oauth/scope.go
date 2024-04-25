package oauth

import "strings"

const (
	//Scopetrialwhitelist获取测试权限上线删
	Scopetrialwhitelist = "trial.whitelist"
	//ScopeUserInfo获取用户公开信息
	ScopeUserInfo = "user_info"
	//ScopeUserdata获取用户基础信息
	ScopeUserdata = "data.external.user"
	//ScopeUserVideo获取用户视频信息
	ScopeUserVideo = "video.data.bind"
)

// GetUserScope 获取用户相关scope
func GetUserScope() string {
	scopes := []string{ScopeUserInfo, Scopetrialwhitelist}
	return strings.Join(scopes, ",")
}

// GetVideoScope 获取视频相关Scope.
func GetVideoScope() string {
	scopes := []string{ScopeUserVideo}
	return strings.Join(scopes, ",")
}

// GetAllScope 获取所有Scope.
func GetAllScope() string {
	scopes := []string{GetUserScope(), GetVideoScope()}
	return strings.Join(scopes, ",")
}
