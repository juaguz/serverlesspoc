package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"lambdapoc/internal/platform"
	"lambdapoc/internal/shopify"
)

type WebHookHandler interface {
	Handler(order *shopify.WebHookOrder) error
}

type shopifyWebHook struct {
	w WebHookHandler
}

func (s *shopifyWebHook) Handler(ctx context.Context, order shopify.WebHookOrder) error {
	return s.w.Handler(&order)
}

func main() {
	s := shopifyWebHook{
		w: shopify.NewWebHookToBroker("", platform.NewSqsBroker()),
	}

	lambda.Start(s.Handler)
}
