package main

import (
	//"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"main/lib"
	//"github.com/aws/aws-sdk-go/service/sso"
	//"github.com/aws/aws-sdk-go/service/identitystore"
	//"github.com/aws/aws-sdk-go/service/ssoadmin"
	//"fmt"
)

func main() {

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		panic(err)
	}
	lib.Auth(sess)

	lib.ListUser(sess)
}
