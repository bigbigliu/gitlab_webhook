# gitlab_webhook

## introduce
gitlab code changes enterprise WeChat notification, you can configure conf/conf.json to use

gitlab 代码变更企业微信消息通知, 配置conf/conf.json即可使用

## build docker image
``` bash
make docker-build
```

## docker run
``` bash
docker run -d -p [port]:8080 webhook
```

