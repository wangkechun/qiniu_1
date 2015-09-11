test:
	DEBUG=* go test -v

install:
	go get -u -v -x github.com/bmizerany/assert
	go get -u -v -x github.com/k0kubun/pp
	go get -u -v -x github.com/tj/go-debug

.PHONY: test install
