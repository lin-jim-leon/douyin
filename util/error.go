package util

// CommonError 抖音返回的通用错误.
type CommonError struct {
	ErrCode int64  `json:"error_code"`
	ErrMsg  string `json:"description"`
}
