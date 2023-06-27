package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/identitystore"
)

func GetUserId(sess *session.Session, email string) (string, error) {
	svc := identitystore.New(sess, &aws.Config{
		Region:                        aws.String("ap-northeast-1"),
		Endpoint:                      aws.String("https://identitystore.ap-northeast-1.amazonaws.com"),
		CredentialsChainVerboseErrors: aws.Bool(true),
	})

	params := &identitystore.ListUsersInput{
		IdentityStoreId: aws.String("d-95670a3c83"),
		MaxResults:      aws.Int64(100), // Max request is 100
	}

	for {
		resp, err := svc.ListUsers(params)
		if err != nil {
			return "", nil
		}

		for _, user := range resp.Users {

			if *user.Emails[0].Value == email || *user.UserName == email {
				return *user.UserId, nil
			}
		}

		// check the other results
		if resp.NextToken != nil {
			params.NextToken = resp.NextToken // Next page token
		} else {
			return "", nil
		}
	}
}
