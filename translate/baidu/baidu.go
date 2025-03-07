package baidu

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"handy-translate/config"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const Way = "baidu"

type Baidu struct {
	config.Translate
}

const (
	fromLang = "auto"
	endpoint = "http://api.fanyi.baidu.com"
	path     = "/api/trans/vip/translate"
)

func (b *Baidu) GetName() string {
	return Way
}

type translateResult struct {
	From        string        `json:"from"`
	To          string        `json:"to"`
	TransResult []TransResult `json:"trans_result"`
}
type TransResult struct {
	Dst string `json:"dst"`
	Src string `json:"src"`
}

func (b *Baidu) PostQuery(query, fromLang, toLang string) ([]string, error) {
	logrus.WithFields(logrus.Fields{
		"query": query, "fromLang": fromLang, "toLang": toLang,
	}).Info("PostQuery")
	endpoint := "http://api.fanyi.baidu.com"
	path := "/api/trans/vip/translate"
	uri := endpoint + path
	appKey := b.Key
	appID := b.AppID

	// Generate salt and sign
	salt := strconv.Itoa(rand.Intn(32768) + 32768)
	sign := makeMD5(appID + query + salt + appKey)

	// Build request
	form := url.Values{}
	form.Add("appid", appID)
	form.Add("q", query)
	form.Add("from", fromLang)
	form.Add("to", toLang)
	// form.Add("from", "en")
	// form.Add("to", "zh")
	form.Add("salt", salt)
	form.Add("sign", sign)

	// Send request
	client := &http.Client{}
	req, err := http.NewRequest("POST", uri, strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var result translateResult
	if err := json.Unmarshal(body, &result); err != nil {
		logrus.Error("Error:", err)
		return nil, err
	}

	prettyResult, _ := json.MarshalIndent(result, "", "    ")
	logrus.Println(string(prettyResult))

	if len(result.TransResult) > 0 {
		if result.TransResult[0].Dst == result.TransResult[0].Src {
			return nil, nil
		}
		var res []string
		for _, v := range result.TransResult {
			res = append(res, v.Dst)
		}
		return res, nil
	}
	return nil, err
}

func makeMD5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
