package lib

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssoadmin"
	"main/cfg"
)

// GetPermissionSetName : Get Permission set Arn from Permission set name
func GetPermissionSetName(sess *session.Session, roleArn string) (string, error) {
	client := ssoadmin.New(sess, &aws.Config{
		Region: cfg.Region})

	params := &ssoadmin.DescribePermissionSetInput{
		InstanceArn:      aws.String("arn:aws:sso:::instance/ssoins-7758e707bb6ea352"),
		PermissionSetArn: aws.String(roleArn),
	}

	describePermissionSetOutput, err := client.DescribePermissionSet(params)
	if err != nil {
		return "", err
	}

	permissionSetName := *describePermissionSetOutput.PermissionSet.Name
	return permissionSetName, nil
}

// GetPermissionSetArn : Get Permission set Arn from Permission set name
func GetPermissionSetArn(sess *session.Session, permissionSetName string) (string, error) {
	client := ssoadmin.New(sess, &aws.Config{
		Region: cfg.Region})

	params := &ssoadmin.ListPermissionSetsInput{
		InstanceArn: cfg.InstanceArn,
		MaxResults:  cfg.MaxResults,
	}

	for {
		resp, err := client.ListPermissionSets(params)
		if err != nil {
			return "", err
		}

		for _, currentArn := range resp.PermissionSets {
			Arn := *currentArn
			curName, err := GetPermissionSetName(sess, Arn)
			if err != nil {
				return "", err
			}
			if curName == permissionSetName {
				return Arn, nil
			}
		}

		// check the other results
		if resp.NextToken != nil {
			params.NextToken = resp.NextToken // // Next page token
		} else {
			break
		}
	}
	err := errors.New("PermissionSet Not Found!")
	return "", err
}
