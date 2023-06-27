package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"main/lib"
)

func main() {

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		log.Fatal(err)
	}

	userid, err := lib.GetUserId(sess, "yunjie.xiao@pingcap.com")
	if err != nil {
		log.Fatal(err)
	}

	PermissionSetArn, err := lib.GetPermissionSetArn(sess, "DBaaS-Prod-ViewOnly-Role")
	targetId := "316218510314"

	err = lib.Auth(sess, targetId, PermissionSetArn, userid)
	// err = lib.Revoke(sess, targetId, PermissionSetArn, userid)
	if err != nil {
		log.Fatal(err)
	}
}
