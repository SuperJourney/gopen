package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/SuperJourney/gopen/infra"
)

var translateURL = "http://api.fanyi.baidu.com/api/trans/vip/translate"

// Translate 调用百度翻译 API 进行文本翻译
func Translate(message string, fromLang, toLang string) (string, error) {
	var appID = infra.Setting.BaiduAppId // 替换为您的 appid
	var apiKey = infra.Setting.BaiduApiKey
	// 生成随机 salt
	salt := strconv.FormatInt(time.Now().Unix(), 10)

	// 计算签名
	sign := md5Sign(appID, message, salt, apiKey)

	// 构建请求参数
	params := url.Values{}
	params.Add("q", message)
	params.Add("from", fromLang)
	params.Add("to", toLang)
	params.Add("appid", appID)
	params.Add("salt", salt)
	params.Add("sign", sign)

	// 发送 GET 请求
	resp, err := http.Get(translateURL + "?" + params.Encode())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 解析响应 JSON
	var result struct {
		ErrorMsg string `json:"error_msg"`
		Trans    []struct {
			Src string `json:"src"`
			Dst string `json:"dst"`
		} `json:"trans_result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	// 检查是否有错误
	if result.ErrorMsg != "" {
		return "", fmt.Errorf("translation failed: %s", result.ErrorMsg)
	}

	// 提取翻译结果
	if len(result.Trans) > 0 {
		return result.Trans[0].Dst, nil
	}

	return "", fmt.Errorf("no translation result found")
}

// md5Sign 计算签名
func md5Sign(appID, message, salt, appKey string) string {
	signStr := appID + message + salt + appKey
	signBytes := md5.Sum([]byte(signStr))
	sign := hex.EncodeToString(signBytes[:])
	return sign
}
