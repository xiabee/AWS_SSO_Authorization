package lib

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/identitystore"
)

func ListUser(sess *session.Session) {
	svc := identitystore.New(sess, &aws.Config{
		Region:                        aws.String("ap-northeast-1"),
		Endpoint:                      aws.String("https://identitystore.ap-northeast-1.amazonaws.com"),
		CredentialsChainVerboseErrors: aws.Bool(true),
	})

	params := &identitystore.ListUsersInput{
		IdentityStoreId: aws.String("d-95670a3c83"),
		MaxResults:      aws.Int64(100), // 设置每次返回的最大结果数
	}

	for {
		resp, err := svc.ListUsers(params)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// 处理 resp.Users 中的用户结果
		for _, user := range resp.Users {
			fmt.Println(*user.UserName, *user.UserId)
		}

		// 检查是否还有更多结果
		if resp.NextToken != nil {
			params.NextToken = resp.NextToken // 设置下一页请求的令牌
		} else {
			break // 已获取所有结果，退出循环
		}
	}

	fmt.Println("All users retrieved!")
}
