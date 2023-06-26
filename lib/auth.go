package lib

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssoadmin"
)

// Auth :an example of auth
func Auth(sess *session.Session) {
	svc := ssoadmin.New(sess)

	params := &ssoadmin.CreateAccountAssignmentInput{
		InstanceArn:      aws.String("arn:aws:sso:::instance/ssoins-7758e707bb6ea352"),
		TargetId:         aws.String("316218510314"),
		TargetType:       aws.String("AWS_ACCOUNT"),
		PermissionSetArn: aws.String("arn:aws:sso:::permissionSet/ssoins-7758e707bb6ea352/ps-cc8887ee3dd4c3d0"), // Dbaas-ViewOnly-Role
		PrincipalType:    aws.String("USER"),
		PrincipalId:      aws.String("95670a3c83-4c2e0132-6440-4040-8dd2-aa4d994ca926"), // yunjie.xiao@pingcap.com
	}

	_, err := svc.CreateAccountAssignment(params)
	if err != nil {
		fmt.Println("Failed:", err)
		return
	}

	fmt.Println("Success!")
}
