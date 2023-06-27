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

func Revoke(sess *session.Session, targetId string, permissionSetArn string, principalId string) error {
	svc := ssoadmin.New(sess)

	params := &ssoadmin.DeleteAccountAssignmentInput{
		InstanceArn:      cfg.InstanceArn,
		TargetId:         aws.String(targetId),
		TargetType:       cfg.TargetType,
		PermissionSetArn: aws.String(permissionSetArn),
		PrincipalType:    cfg.PrincipalType,
		PrincipalId:      aws.String(principalId),
	}

	_, err := svc.DeleteAccountAssignment(params)
	if err != nil {
		return err
	}

	listInput := &ssoadmin.ListAccountAssignmentDeletionStatusInput{
		InstanceArn: cfg.InstanceArn,
	}

	listOutput, err := svc.ListAccountAssignmentDeletionStatus(listInput)
	if len(listOutput.AccountAssignmentsDeletionStatus) > 0 {
		requestId := listOutput.AccountAssignmentsDeletionStatus[0].RequestId
		fmt.Println("Request ID:", *requestId)

		describeInput := &ssoadmin.DescribeAccountAssignmentDeletionStatusInput{
			InstanceArn:                        cfg.InstanceArn,
			AccountAssignmentDeletionRequestId: requestId,
		}

		// Polling to check the status of the request
		for {
			describeOutput, err := svc.DescribeAccountAssignmentDeletionStatus(describeInput)
			if err != nil {
				return err
			}

			Status := *describeOutput.AccountAssignmentDeletionStatus.Status
			FailureReason := describeOutput.AccountAssignmentDeletionStatus.FailureReason
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
		err := errors.New("No Account Assignment Deletion Status found.")
		return err
	}

	fmt.Println("Successfully Revoked!")
	return nil
}
