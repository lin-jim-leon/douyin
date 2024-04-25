package user

import (
	"encoding/json"
	"fmt"
	"github.com/lin-jim-leon/douyin/util"
	"net/url"
)

const (
	UserInfoUrl     string = "https://open.douyin.com/oauth/userinfo/"
	UserFansUrl     string = "https://open.douyin.com/data/external/user/fans/"
	UserLikeUrl     string = "https://open.douyin.com/data/external/user/like/"
	UserCommentsUrl string = "https://open.douyin.com/data/external/user/comment/"
	UserShareUrl    string = "https://open.douyin.com/data/external/user/share/"
)

// Info .
type Info struct {
	util.CommonError

	Avatar       string `json:"avatar"`
	AvatarLarger string `json:"avatar_larger"`
	ClientKey    string `json:"client_key"`
	EAccountRole string `json:"e_account_role"`
	ErrorCode    int    `json:"error_code"`
	LogId        string `json:"log_id"`
	Nickname     string `json:"nickname"`
	OpenId       string `json:"open_id"`
	UnionId      string `json:"union_id"`
}

type UserInfoRes struct {
	Message string `json:"message"`
	Data    Info   `json:"data"`
}

// GetUserInfo获取用户基本信息
func GetUserInfo(openId string, accesstoken string) (info UserInfoRes, err error) {
	req := url.Values{}
	req.Set("access_token", accesstoken)
	req.Set("open_id", openId)

	data := req.Encode()
	var response []byte
	response, err = util.HTTPPost(UserInfoUrl, data, "application/x-www-form-urlencoded")
	if err != nil {
		return UserInfoRes{}, err
	}
	var result UserInfoRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return UserInfoRes{}, err
	}
	if result.Data.ErrorCode != 0 {
		return UserInfoRes{}, fmt.Errorf("GetUserInfo error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
	}
	return result, nil
}

type UserfanRes struct {
	Data  UserfanData  `json:"data"`
	Extra UserfanExtra `json:"extra"`
}

type UserfanData struct {
	Description string          `json:"description"`
	ErrorCode   int             `json:"error_code"`
	ResultList  []UserfanResult `json:"result_list"`
}

type UserfanResult struct {
	Date      string `json:"date"`
	NewFans   string `json:"new_fans"`
	TotalFans string `json:"total_fans"`
}

type UserfanExtra struct {
	Description    string `json:"description"`
	ErrorCode      string `json:"error_code"`
	LogID          string `json:"logid"`
	Now            string `json:"now"`
	SubDescription string `json:"sub_description"`
	SubErrorCode   string `json:"sub_error_code"`
}

// GetUserFans
func GetUserFans(accesstoken string, date string, open_id string) (info UserfanRes, err error) {
	headers := map[string]string{
		"access-token": accesstoken,
		"Content-Type": "application/json",
	}
	params := map[string]string{
		"date_type": date, //7/15/30
		"open_id":   open_id,
	}
	var response []byte
	response, err = util.HTTPGetWithHeaders(UserFansUrl, headers, params)
	if err != nil {
		return UserfanRes{}, err
	}
	var result UserfanRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return UserfanRes{}, err
	}
	if result.Data.ErrorCode != 0 {
		return UserfanRes{}, fmt.Errorf("GetUserfan error : errcode=%v , errmsg=%v", result.Data.ErrorCode, result.Data.Description)
	}
	return result, nil
}

type UserLikeRes struct {
	Data  UserLikeData  `json:"data"`
	Extra UserLikeExtra `json:"extra"`
}

type UserLikeData struct {
	Description string           `json:"description"`
	ErrorCode   int              `json:"error_code"`
	ResultList  []UserLikeResult `json:"result_list"`
}

type UserLikeResult struct {
	Date    string `json:"date"`
	NewLike string `json:"new_like"`
}

type UserLikeExtra struct {
	Description    string `json:"description"`
	ErrorCode      string `json:"error_code"`
	LogID          string `json:"logid"`
	Now            string `json:"now"`
	SubDescription string `json:"sub_description"`
	SubErrorCode   string `json:"sub_error_code"`
}

// GetUserLike
func GetUserLike(accesstoken string, date string, open_id string) (info UserLikeRes, err error) {
	headers := map[string]string{
		"access-token": accesstoken,
		"Content-Type": "application/json",
	}
	params := map[string]string{
		"date_type": date, //7/15/30
		"open_id":   open_id,
	}
	var response []byte
	response, err = util.HTTPGetWithHeaders(UserLikeUrl, headers, params)
	if err != nil {
		return UserLikeRes{}, err
	}
	var result UserLikeRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return UserLikeRes{}, err
	}
	if result.Data.ErrorCode != 0 {
		return UserLikeRes{}, fmt.Errorf("GetUserlike error : errcode=%v , errmsg=%v", result.Data.ErrorCode, result.Data.Description)
	}
	return result, nil
}

type UserCommentRes struct {
	Data  UserCommentData  `json:"data"`
	Extra UserCommentExtra `json:"extra"`
}

type UserCommentData struct {
	Description string             `json:"description"`
	ErrorCode   int                `json:"error_code"`
	ResultList  []UserCommentEntry `json:"result_list"`
}

type UserCommentEntry struct {
	Date       string `json:"date"`
	NewComment string `json:"new_comment"`
}

type UserCommentExtra struct {
	Description    string `json:"description"`
	ErrorCode      string `json:"error_code"`
	LogID          string `json:"logid"`
	Now            string `json:"now"`
	SubDescription string `json:"sub_description"`
	SubErrorCode   string `json:"sub_error_code"`
}

func GetUserComments(accesstoken string, date string, open_id string) (info UserCommentRes, err error) {
	headers := map[string]string{
		"access-token": accesstoken,
		"Content-Type": "application/json",
	}
	params := map[string]string{
		"date_type": date, //7/15/30
		"open_id":   open_id,
	}
	var response []byte
	response, err = util.HTTPGetWithHeaders(UserCommentsUrl, headers, params)
	if err != nil {
		return UserCommentRes{}, err
	}
	var result UserCommentRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return UserCommentRes{}, err
	}
	if result.Data.ErrorCode != 0 {
		return UserCommentRes{}, fmt.Errorf("GetUsercomments error : errcode=%v , errmsg=%v", result.Data.ErrorCode, result.Data.Description)
	}
	return result, nil
}

type UserShareRes struct {
	Data  UserShareData  `json:"data"`
	Extra UserShareExtra `json:"extra"`
}

type UserShareData struct {
	Description string          `json:"description"`
	ErrorCode   int             `json:"error_code"`
	ResultList  []UserShareItem `json:"result_list"`
}

type UserShareItem struct {
	Date     string `json:"date"`
	NewShare string `json:"new_share"`
}

type UserShareExtra struct {
	Description    string `json:"description"`
	ErrorCode      string `json:"error_code"`
	LogID          string `json:"logid"`
	Now            string `json:"now"`
	SubDescription string `json:"sub_description"`
	SubErrorCode   string `json:"sub_error_code"`
}

func GetUserShare(accesstoken string, date string, open_id string) (info UserShareRes, err error) {
	headers := map[string]string{
		"access-token": accesstoken,
		"Content-Type": "application/json",
	}
	params := map[string]string{
		"date_type": date, //7/15/30
		"open_id":   open_id,
	}
	var response []byte
	response, err = util.HTTPGetWithHeaders(UserShareUrl, headers, params)
	if err != nil {
		return UserShareRes{}, err
	}
	var result UserShareRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return UserShareRes{}, err
	}
	if result.Data.ErrorCode != 0 {
		return UserShareRes{}, fmt.Errorf("GetUsercomments error : errcode=%v , errmsg=%v", result.Data.ErrorCode, result.Data.Description)
	}
	return result, nil
}
