package platform

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SnsBroker struct {
}

func NewSqsBroker() *SnsBroker {
	return &SnsBroker{}
}

func (s *SnsBroker) Publish(topic string, message []byte) error {
	//TODO implement me
	sess := session.Must(session.NewSession())

	m := string(message)

	svc := sns.New(sess)
	result, err := svc.Publish(&sns.PublishInput{
		Message:  &m,
		TopicArn: &topic,
	})

	if err != nil {
		return err
	}

	fmt.Println(*result.MessageId)

	return nil
}
