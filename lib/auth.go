package lib

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssoadmin"
)

// Auth :an example of auth
func Auth(sess *session.Session, targetId string, permissionSetArn string, principalId string) {
	svc := ssoadmin.New(sess)

	params := &ssoadmin.CreateAccountAssignmentInput{
		InstanceArn:      aws.String("arn:aws:sso:::instance/ssoins-7758e707bb6ea352"),
		TargetId:         aws.String(targetId),
		TargetType:       aws.String("AWS_ACCOUNT"),
		PermissionSetArn: aws.String(permissionSetArn),
		PrincipalType:    aws.String("USER"),
		PrincipalId:      aws.String(principalId),
	}

	_, err := svc.CreateAccountAssignment(params)
	if err != nil {
		fmt.Println("Failed:", err)
		return
	}

	fmt.Println("Success!")
}
