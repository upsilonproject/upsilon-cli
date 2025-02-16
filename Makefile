default:
	go build -v -o upsilon

install:
	cp upsilon /usr/local/sbin/upsilon
