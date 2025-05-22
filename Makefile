.PHONY: build clean deploy

build:
	GOOS=linux GOARCH=amd64 go build -o app/entrypoints/lambda/main app/entrypoints/lambda/main.go

clean:
	rm -rf ./app/entrypoints/lambda/main

deploy: build
	sam deploy --guided

local-invoke: build
	sam local invoke VehicleFunction -e events/event.json

local-api: build
	sam local start-api