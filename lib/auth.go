package lib

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssoadmin"
	"main/cfg"
	"time"
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

	listInput := &ssoadmin.ListAccountAssignmentCreationStatusInput{
		InstanceArn: cfg.InstanceArn,
	}

	listOutput, err := svc.ListAccountAssignmentCreationStatus(listInput)
	if len(listOutput.AccountAssignmentsCreationStatus) > 0 {
		requestId := listOutput.AccountAssignmentsCreationStatus[0].RequestId
		fmt.Println("Request ID:", *requestId)

		describeInput := &ssoadmin.DescribeAccountAssignmentCreationStatusInput{
			InstanceArn:                        cfg.InstanceArn,
			AccountAssignmentCreationRequestId: requestId,
		}

		// Polling to check the status of the request
		for {
			describeOutput, err := svc.DescribeAccountAssignmentCreationStatus(describeInput)
			if err != nil {
				return err
			}

			Status := *describeOutput.AccountAssignmentCreationStatus.Status
			FailureReason := describeOutput.AccountAssignmentCreationStatus.FailureReason
			fmt.Println("Status: " + Status)
			if FailureReason != nil {
				err := errors.New(*FailureReason)
				return err
			}

			if Status == ssoadmin.StatusValuesSucceeded || Status == ssoadmin.StatusValuesFailed {
				break
			}
			time.Sleep(time.Second * (time.Duration(cfg.SleepTime)))
		}

	} else {
		err := errors.New("No Account Assignment Creation Status found.")
		return err
	}

	fmt.Println("Successfully Authorized!")
	return nil
}
