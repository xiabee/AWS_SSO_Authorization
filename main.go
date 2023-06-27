package main

import (
	"fmt"
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
	PermissionSetArn := "arn:aws:sso:::permissionSet/ssoins-7758e707bb6ea352/ps-cc8887ee3dd4c3d0" // Dbaas-ViewOnly-Role
	targetId := "316218510314"

	//lib.Auth(sess, targetId, PermissionSetArn, userid)
	lib.Revoke(sess, targetId, PermissionSetArn, userid)

	arn, _ := lib.GetPermissionSetArn(sess, "Dbaas-ViewOnly-Role")
	fmt.Println(arn)
}
