package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// get请求
func HTTPGet(uri string) ([]byte, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

// post请求
func HTTPPost(uri string, data string, contenttype string) ([]byte, error) {
	body := bytes.NewBuffer([]byte(data))
	response, err := http.Post(uri, contenttype, body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

// getwithheader
func HTTPGetWithHeaders(uri string, headers map[string]string, params map[string]string) ([]byte, error) {
	// 创建一个请求对象
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	// 添加请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 添加请求参数
	query := req.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	// 发送请求
	client := http.DefaultClient
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// 读取响应内容
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

// httppost post json
// PostJSONWithToken post json 数据请求，并附带 access_token 请求头
func PostJSON(uri string, obj interface{}, ctype string, accessToken string) ([]byte, error) {
	// 将对象转换为 JSON 格式的数据
	jsonData, err := json.Marshal(obj)
	fmt.Println(string(jsonData))
	if err != nil {
		return nil, err
	}

	// 替换特殊字符的 Unicode 编码
	jsonData = bytes.Replace(jsonData, []byte("\\u003c"), []byte("<"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u003e"), []byte(">"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u0026"), []byte("&"), -1)

	// 创建一个字节缓冲区并写入 JSON 数据
	body := bytes.NewBuffer(jsonData)

	// 创建一个 HTTP POST 请求
	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		fmt.Println("创建请求失败")
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("access-token", accessToken)

	// 使用默认的 HTTP 客户端执行请求
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败")
		return nil, err
	}
	defer response.Body.Close()

	// 检查响应状态码
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error: uri=%v, statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}
