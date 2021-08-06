package common

import (
	"log"
	"net/http"
	"github.com/bigbigliu/gitlab_webhook/controller"

	"gopkg.in/go-playground/webhooks.v5/gitlab"
)

// Webhook 事件
func Webhook(w http.ResponseWriter, r *http.Request) {
	hook, _ := gitlab.New()
	//事件类型 mergerRequest
	payload, err := hook.Parse(r, gitlab.MergeRequestEvents, gitlab.PushEvents)
	if err != nil {
		if err == gitlab.ErrEventNotFound {
			log.Println("err:", err)
		}
	}
	switch payload.(type) {

	case gitlab.MergeRequestEventPayload:
		merge := payload.(gitlab.MergeRequestEventPayload)
		controller.MergeRequest(&merge)
	case gitlab.PushEventPayload:
		push := payload.(gitlab.PushEventPayload)
		controller.PushRequest(&push)

	}
}
