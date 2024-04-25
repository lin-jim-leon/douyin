package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/lin-jim-leon/douyin/util"
	"net/url"
)

const (
	redirectOauthURL string = "https://open.douyin.com/platform/oauth/connect?client_key=%s&response_type=code&scope=%s&redirect_uri=%s"
	accessTokenUrl   string = "https://open.douyin.com/oauth/access_token?client_key=%s&client_secret=%s&code=%s&grant_type=authorization_code"
	refreshTokenURL  string = "https://open.douyin.com/oauth/oauth/refresh_token?client_key=%s&grant_type=refresh_token&refresh_token=%s"
)

// GetRedirectURL 获取授权码的url地址
func GetRedirectURL(redirecturl string, ClientKey string, Scopes string) string {
	uri := url.QueryEscape(redirecturl)
	return fmt.Sprintf(redirectOauthURL, ClientKey, Scopes, uri)
}

// AccessToken返回结构体
type ainfo struct {
	//util.CommonError
	AccessToken      string `json:"access_token"`
	Description      string `json:"description"`
	ErrorCode        int    `json:"error_code"`
	ExpiresIn        int    `json:"expires_in"`
	OpenId           string `json:"open_id"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
}

// AccessToken返回response
type AccessTokenRes struct {
	Message string `json:"message"`
	Data    ainfo  `json:"data"`
}

// GetAccessToken 通过网页授权的code 换取access_token
func GetAccessToken(ClientKey string, ClientSecret string, code string) (accessToken AccessTokenRes, err error) {
	uri := fmt.Sprintf(accessTokenUrl, ClientKey, ClientSecret, code)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return AccessTokenRes{}, err
	}
	var result AccessTokenRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return AccessTokenRes{}, err
	}
	if result.Data.ErrorCode != 0 {
		return AccessTokenRes{}, fmt.Errorf("GetAccessToken error: error_code=%v, description=%v", result.Data.ErrorCode, result.Data.Description)
	}
	return result, nil
}

// RefreshAccessToken 刷新AccessToken.
// 当access_token过期（过期时间15天）后，可以通过该接口使用refresh_token（过期时间30天）进行刷新
func RefreshAccessToken(refreshkey string, clientkey string) (accessToken AccessTokenRes, err error) {
	req := url.Values{}
	req.Set("client_key", clientkey)
	req.Set("refresh_token", refreshkey)
	req.Set("grant_type", "refresh_token")

	data := req.Encode()
	var response []byte
	response, err = util.HTTPPost(refreshTokenURL, data, "application/x-www-form-urlencoded")
	if err != nil {
		return AccessTokenRes{}, err
	}
	var result AccessTokenRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return AccessTokenRes{}, err
	}
	if result.Data.ErrorCode != 0 {
		return AccessTokenRes{}, fmt.Errorf("GetAccessToken error: error_code=%v, description=%v", result.Data.ErrorCode, result.Data.Description)
	}
	return result, nil
}
