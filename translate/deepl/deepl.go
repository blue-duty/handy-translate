package deepl

import (
	"encoding/json"
	"fmt"
	"handy-translate/config"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
)

// https://www.deepl.com/docs-api
const Way = "deepl"

type DeepL struct {
	config.Translate
}

const (
	endpoint = "https://api-free.deepl.com"
	path     = "/v2/translate"
)

type translateResult struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Translations           []struct {
		Text string `json:"text"`
	} `json:"translations"`
}

func (d *DeepL) GetName() string {
	return Way
}

func (d *DeepL) PostQuery(query, fromLang, toLang string) ([]string, error) {
	logrus.WithFields(logrus.Fields{
		"query": query, "fromLang": fromLang, "toLang": toLang,
	}).Info("PostQuery")

	uri := endpoint + path
	authKey := d.Key
	fmt.Println("authKey:", authKey)

	form := url.Values{}
	form.Add("text", query)
	form.Add("target_lang", strings.ToUpper(toLang))
	if fromLang != "" && fromLang != "auto" {
		form.Add("source_lang", strings.ToUpper(fromLang))
	}
	// 避免被按照句号分割
	form.Add("split_sentences", "0")
	// form.Add("tag_handling", "html")

	fmt.Println("form:", form)

	// Send request
	client := &http.Client{}
	req, err := http.NewRequest("POST", uri, strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "DeepL-Auth-Key "+authKey)

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

	if len(result.Translations) > 0 {
		var res []string
		for _, v := range result.Translations {
			res = append(res, v.Text)
		}
		return res, nil
	}
	return nil, err
}
