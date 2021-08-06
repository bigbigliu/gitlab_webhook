package main

import (
	"log"
	"net/http"
	"github.com/bigbigliu/gitlab_webhook/common"
	"github.com/bigbigliu/gitlab_webhook/conf"
)

func main() {
	http.HandleFunc("/webhooks", common.Webhook)

	log.Println("服务端口：", conf.Config.Port)
	log.Fatal(http.ListenAndServe(conf.Config.Port, nil))
}
