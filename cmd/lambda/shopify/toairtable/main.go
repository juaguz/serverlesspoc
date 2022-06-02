package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"lambdapoc/internal/shopify"
	"log"
)

type syncAirtable struct {
}

func (s *syncAirtable) Handler(ctx context.Context, order shopify.WebHookOrder) error {
	log.Println(order)
	return nil
}

func main() {
	s := syncAirtable{}

	lambda.Start(s.Handler)
}
