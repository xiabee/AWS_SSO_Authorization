package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"main/lib"
)

func main() {

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}

	userid, _ := lib.GetUserId(sess, "yunjie.xiao@pingcap.com")
	PermissionSetArn, _ := lib.GetPermissionSetArn(sess, "DBaaS-Prod-ViewOnly-Role")
	targetId := "316218510314"

	//lib.Auth(sess, targetId, PermissionSetArn, userid)
	lib.Revoke(sess, targetId, PermissionSetArn, userid)
}
