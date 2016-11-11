package awsecs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func getClient() {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}

	region, err := ec2metadata.New(session).Region()
	if err != nil {
		return err
	}

	return ecs.New(sess, &aws.Config{Region: aws.String(region)})
}
