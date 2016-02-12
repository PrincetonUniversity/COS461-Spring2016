CC=g++
GO=go
CFLAGS= -g -Wall -Werror

all: proxy http_server

http_server: http_server.c
	$(CC) $(CFLAGS) -o proxy_parse.o -c proxy_parse.c
	$(CC) $(CFLAGS) -o http_server.o -c http_server.c
	$(CC) $(CFLAGS) -o http_server proxy_parse.o http_server.o

proxy: proxy.go
	$(GO) build proxy.go

#go_http_server: go_http_server.go
#	$(GO) build go_http_server.go

clean:
	rm -f proxy http_server *.o

tar:
	tar -cvzf ass1.tgz proxy.go http_server.c README Makefile proxy_parse.c proxy_parse.h
