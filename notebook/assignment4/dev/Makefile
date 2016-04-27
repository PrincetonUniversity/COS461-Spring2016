GO=go

all: proxy4 http_client

http_client: http_client.go
	$(GO) build http_client.go

proxy4: proxy4.go
	$(GO) build proxy4.go

clean:
	rm -f proxy4 http_client
