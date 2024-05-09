default:
	go build -v github.com/upsilonproject/upsilon-cli/cmd/upsilon/

grpc:
	buf generate

install:
	cp upsilon /usr/local/sbin/upsilon
