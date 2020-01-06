SOURCE_PATH = /go/src/github.com/20k-ultra/readability
PACKAGE = ./...
TIMEOUT = 10
REPO = 20k-ultra/readability

test:
	go test ./... -timeout $(TIMEOUT)s -v

fmt:
	go fmt ./...

dbuild:
	docker build -t $(REPO) .

dtest:
	docker run $(REPO) go test ./...