package lib

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssoadmin"
)

// GetPermissionSetName : Get Permission set Arn from Permission set name
func GetPermissionSetName(sess *session.Session, roleArn string) (string, error) {
	client := ssoadmin.New(sess)

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
	client := ssoadmin.New(sess)

	params := &ssoadmin.ListPermissionSetsInput{
		InstanceArn: aws.String("arn:aws:sso:::instance/ssoins-7758e707bb6ea352"),
	}

	for {
		resp, err := client.ListPermissionSets(params)
		if err != nil {
			return "", err
		}

		for _, permissionSet := range resp.PermissionSets {

			name, err := GetPermissionSetName(sess, aws.StringValue(permissionSet))
			fmt.Println(name, permissionSet)
			if err != nil {
				return "", err
			}

			if name == permissionSetName {
				fmt.Println(aws.StringValue(permissionSet))
				break
			}

			// fmt.Println("Permission Set ARN:", aws.StringValue(permissionSet))
		}

		// check the other results
		if resp.NextToken != nil {
			params.NextToken = resp.NextToken // // Next page token
		} else {
			break
		}
	}
	return "", nil
}
