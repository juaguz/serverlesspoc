.PHONY: build clean deploy

build:
	go mod tidy
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/lambda/shopifywebhook cmd/lambda/shopify/webhook/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/lambda/toairtable cmd/lambda/shopify/toairtable/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
