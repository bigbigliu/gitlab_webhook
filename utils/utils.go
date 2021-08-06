package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/bigbigliu/gitlab_webhook/conf"

	"github.com/labstack/gommon/log"
)

// PostWechat  企业微信发送消息
func PostWechat(sendBody map[string]interface{}, webURL string) {
	b, err := json.Marshal(sendBody)
	if err != nil {
		log.Error("err:", err)
	}
	//根据不同的项目地址 发送到不同的webhook
	var address string
	for k, v := range conf.Config.PushAddress {
		ok := strings.HasPrefix(webURL, k)
		if ok {
			address = v.(string)
			break
		}
	}
	if address != "" {
		payload := bytes.NewBuffer(b)
		client := &http.Client{}
		req, err := http.NewRequest("POST", address, payload)
		if err != nil {
			log.Error("err", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Error("err", err)
		}
		defer resp.Body.Close()
		fmt.Printf("== %+v \n", resp)
		if resp.StatusCode != 200 {
			log.Error("err", err)

		}
	}

}
