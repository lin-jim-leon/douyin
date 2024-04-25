package video

import (
	"encoding/json"
	"fmt"
	"github.com/lin-jim-leon/douyin/util"
	"net/http"
	"strings"
)

const (
	videoDataurl string = "https://open.douyin.com/api/douyin/v1/video/video_data/?open_id=%s"
	videoList    string = "https://open.douyin.com/api/douyin/v1/video/video_list/?open_id=%s&cursor=%t&count=%t"
)

// GetShortUrl短链解长链
func GetShortUrl(longUrl string) (shortUrl string) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	Respon, err := client.Get(shortUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	redirectedURL := Respon.Header.Get("Location")
	// 从重定向链接中提取视频ID
	parts := strings.Split(redirectedURL, "/")
	videoID := parts[len(parts)-2]
	// 拼接完整的视频链接
	fullURL := "https://www.douyin.com/video/" + videoID
	return fullURL
}

type VideoResponse struct {
	Extra VideoExtra `json:"extra"`
	Data  VideoData  `json:"data"`
}

type VideoExtra struct {
	LogID          string `json:"logid"`
	Now            int64  `json:"now"`
	ErrorCode      int    `json:"error_code"`
	Description    string `json:"description"`
	SubErrorCode   int    `json:"sub_error_code"`
	SubDescription string `json:"sub_description"`
}

type VideoData struct {
	ErrorCode   int         `json:"error_code"`
	Description string      `json:"description"`
	List        []VideoItem `json:"list"`
}

type VideoItem struct {
	Title       string     `json:"title"`
	CreateTime  int64      `json:"create_time"`
	VideoStatus int        `json:"video_status"`
	ShareURL    string     `json:"share_url"`
	Cover       string     `json:"cover"`
	IsTop       bool       `json:"is_top"`
	Statistics  VideoStats `json:"statistics"`
	ItemID      string     `json:"item_id"`
	IsReviewed  bool       `json:"is_reviewed"`
	MediaType   int        `json:"media_type"`
}

type VideoStats struct {
	DiggCount     int `json:"digg_count"`
	DownloadCount int `json:"download_count"`
	PlayCount     int `json:"play_count"`
	ShareCount    int `json:"share_count"`
	ForwardCount  int `json:"forward_count"`
	CommentCount  int `json:"comment_count"`
}

// DataReq .
type DataReq struct {
	VideoIDS []string `json:"video_ids"`
}

// GetVideoInfo查询单条视频数据
func GetVideoInfo(accesstoken string, openid string, video_id string) (videoInfo VideoResponse, err error) {
	videoids := []string{video_id}
	req := &DataReq{
		VideoIDS: videoids,
	}
	var response []byte
	const_type := "application/json"
	Accesstoken := accesstoken
	url := fmt.Sprintf(videoDataurl, openid)
	response, err = util.PostJSON(url, req, const_type, Accesstoken)
	if err != nil {
		return VideoResponse{}, err
	}
	var result VideoResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		return VideoResponse{}, err
	}
	if result.Data.ErrorCode != 0 {
		return VideoResponse{}, fmt.Errorf("Data error : errcode=%v , description=%v", result.Data.ErrorCode, result.Data.Description)
	}
	return result, nil
}

type VideoListResponse struct {
	Extra VideoListExtra `json:"extra"`
	Data  VideoListData  `json:"data"`
}

type VideoListExtra struct {
	LogID          string `json:"logid"`
	Now            int64  `json:"now"`
	ErrorCode      int    `json:"error_code"`
	Description    string `json:"description"`
	SubErrorCode   int    `json:"sub_error_code"`
	SubDescription string `json:"sub_description"`
}

type VideoListData struct {
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
	HasMore     bool            `json:"has_more"`
	List        []VideoListItem `json:"list"`
	Cursor      int             `json:"cursor"`
}

type VideoListItem struct {
	Title       string     `json:"title"`
	IsTop       bool       `json:"is_top"`
	CreateTime  int64      `json:"create_time"`
	IsReviewed  bool       `json:"is_reviewed"`
	Videostatus int        `json:"video_status"`
	ShareURL    string     `json:"share_url"`
	ItemID      string     `json:"item_id"`
	MediaType   int        `json:"media_type"`
	Cover       string     `json:"cover"`
	Statistics  VideoStats `json:"statistics"`
}

type Videostats struct {
	ForwardCount  int `json:"forward_count"`
	CommentCount  int `json:"comment_count"`
	DiggCount     int `json:"digg_count"`
	DownloadCount int `json:"download_count"`
	PlayCount     int `json:"play_count"`
	ShareCount    int `json:"share_count"`
}

// GetVideoListInfo查询视频列表数据
func GetVideolistInfo(accesstoken string, openid string, count int32, cursor int64) (videoInfo VideoListResponse, err error) {
	uri := fmt.Sprintf(videoList, openid, count, cursor)
	headers := map[string]string{
		"access-token": accesstoken,
		"Content-Type": "application/json",
	}
	params := map[string]string{
		"count":   fmt.Sprintf("%d", count),
		"open_id": openid,
		"cursor":  fmt.Sprintf("%d", cursor),
	}
	var response []byte
	response, err = util.HTTPGetWithHeaders(uri, headers, params)
	if err != nil {
		return VideoListResponse{}, err
	}
	var result VideoListResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		return VideoListResponse{}, err
	}
	if result.Data.ErrorCode != 0 {
		return VideoListResponse{}, fmt.Errorf("GetVideoListResponses error : errcode=%v , errmsg=%v", result.Data.ErrorCode, result.Data.Description)
	}
	return result, nil
}
