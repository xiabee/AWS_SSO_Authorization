package lib

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/identitystore"
	"main/cfg"
)

func GetUserId(sess *session.Session, email string) (string, error) {
	svc := identitystore.New(sess, &aws.Config{
		Region:                        cfg.Region,
		Endpoint:                      cfg.Endpoint,
		CredentialsChainVerboseErrors: aws.Bool(true),
	})

	params := &identitystore.ListUsersInput{
		IdentityStoreId: cfg.IdentityStoreId,
		MaxResults:      cfg.MaxResults,
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
			err := errors.New("Error: User email not found")
			return "", err
		}
	}
}
