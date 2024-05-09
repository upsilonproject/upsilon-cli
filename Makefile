default:
	go build -v github.com/upsilonproject/upsilon-cli/cmd/upsilon/

install:
	cp upsilon /usr/local/sbin/upsilon
