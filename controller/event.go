package controller

import (
	"strings"
	"time"
	"github.com/bigbigliu/gitlab_webhook/conf"
	"github.com/bigbigliu/gitlab_webhook/utils"

	"gopkg.in/go-playground/webhooks.v5/gitlab"
)

// MergeRequest 请求
func MergeRequest(mergeRequest *gitlab.MergeRequestEventPayload) {
	var (
		Action = mergeRequest.ObjectAttributes.Action
	)
	//同意merge合入并且源分支与目标分支不一致
	if Action == "merge" && mergeRequest.ObjectAttributes.SourceBranch != mergeRequest.ObjectAttributes.TargetBranch && strings.Contains(conf.Config.Branch, mergeRequest.ObjectAttributes.TargetBranch) {
		Later8h, _ := time.ParseDuration("1h")
		var sendBody = map[string]interface{}{
			"msgtype": "markdown",
			"markdown": map[string]string{
				"content": "### 😄 " + mergeRequest.ObjectKind + " 😄 \n" +
					"提交人：" + mergeRequest.User.Name + " \n" +
					"项目名称：" + mergeRequest.Project.Name + "\n" +
					"源分支：" + mergeRequest.ObjectAttributes.SourceBranch + "\n" +
					"目标分支：" + mergeRequest.ObjectAttributes.TargetBranch + "\n" +
					"合并描述：" + mergeRequest.ObjectAttributes.Description + "\n" +
					"提交时间：" + mergeRequest.ObjectAttributes.UpdatedAt.Add(8*Later8h).Format("2006-01-02 15:04:05") + "\n" +
					"地址：[" + mergeRequest.ObjectAttributes.URL + "](" + mergeRequest.ObjectAttributes.URL + ")",
			},
		}
		utils.PostWechat(sendBody, mergeRequest.Project.WebURL)
	}

}

// PushRequest push事件
func PushRequest(push *gitlab.PushEventPayload) {
	var (
		message = ""
	)
	refList := strings.Split(push.Ref, "/")
	ref := refList[len(refList)-1]
	// merge_request： open 时会触发一次 目标分支的push ，通过isPush判断，merge 时 会触发一次源分支的push 通过 len(push.Commit) 判断
	if strings.Contains(conf.Config.Branch, ref) && len(push.Commits) > 0 {
		message = push.Commits[len(push.Commits)-1].Message
		isPush := strings.Contains(message, "See merge request")
		// 不是由merge触发的push发送消息通知
		if !isPush {
			refList := strings.Split(push.Ref, "/")
			branch := refList[len(refList)-1]
			var sendBody = map[string]interface{}{
				"msgtype": "markdown",
				"markdown": map[string]string{
					"content": "### 😄 " + push.ObjectKind + " 😄 \n" +
						"提交人：" + push.UserName + " \n" +
						"项目名称：" + push.Repository.Name + "\n" +
						"分支名称：" + branch + "\n" +
						"commmit：" + message + "\n" +
						"提交时间：" + time.Now().Format("2006-01-02 15:04:05") + "\n" +
						"地址：[" + push.Project.WebURL + "](" + push.Project.WebURL + ")",
				},
			}
			utils.PostWechat(sendBody, push.Project.WebURL)
		}
	}
}
