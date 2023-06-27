package lib

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssoadmin"
	"main/cfg"
)

// Auth :an example of auth
func Auth(sess *session.Session, targetId string, permissionSetArn string, principalId string) error {
	svc := ssoadmin.New(sess)

	params := &ssoadmin.CreateAccountAssignmentInput{
		InstanceArn:      cfg.InstanceArn,
		TargetId:         aws.String(targetId),
		TargetType:       cfg.TargetType,
		PermissionSetArn: aws.String(permissionSetArn),
		PrincipalType:    cfg.PrincipalType,
		PrincipalId:      aws.String(principalId),
	}

	_, err := svc.CreateAccountAssignment(params)
	if err != nil {
		return err
	}

	fmt.Println("Successfully Authorized!")
	return nil
}
