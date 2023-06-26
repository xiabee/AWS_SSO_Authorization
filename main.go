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

	// lib.Auth(sess)

	// lib.ListUser(sess)

	lib.GetUserId("yunjie.xiao@pingcap.com", sess)
}
