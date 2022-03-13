package kinesisclient

import (
	"encoding/json"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

var (
	config *kinesisConfig
	client *kinesis.Kinesis
)

type kinesisConfig struct {
	stream   *string
	endpoint *string
}

func init() {

	config = &kinesisConfig{
		stream:   aws.String(os.Getenv("KINESIS_STREAM_NAME")),
		endpoint: aws.String(os.Getenv("AWS_ENDPOINT")),
	}

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewEnvCredentials(),
		Endpoint:    config.endpoint,
	})

	if err != nil {
		panic(err)
	}

	client = kinesis.New(sess)

	_, err = client.DescribeStream(
		&kinesis.DescribeStreamInput{
			StreamName: config.stream,
		},
	)

	if err != nil {
		panic(err)
	}

}

func Produce(key string, recordType string, data map[string]interface{}) error {

	data["timestamp"] = time.Now().UnixMilli()
	data["type"] = recordType

	recordData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	_, err = client.PutRecord(&kinesis.PutRecordInput{
		Data:         recordData,
		StreamName:   config.stream,
		PartitionKey: aws.String(key),
	})

	if err != nil {
		return err
	}

	return nil

}
