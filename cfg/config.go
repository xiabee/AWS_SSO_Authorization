package cfg

import "github.com/aws/aws-sdk-go/aws"

// InstanceArn SSO instance
var InstanceArn *string = aws.String("arn:aws:sso:::instance/ssoins-7758e707bb6ea352")
var IdentityStoreId *string = aws.String("d-95670a3c83")
var Endpoint *string = aws.String("https://identitystore.ap-northeast-1.amazonaws.com")

// Region SSO Region
var Region *string = aws.String("ap-northeast-1")

var TargetType *string = aws.String("AWS_ACCOUNT")
var PrincipalType *string = aws.String("USER")

var MaxResults *int64 = aws.Int64(100) // Max request is 100
var SleepTime int = 1                  // Seconds of polling check
