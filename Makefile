run:
	go run app/main.go

test:
	go test ./... -v -cover

.PHONY: run test