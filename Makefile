BINARY=freya-rest-api
TESTS=go test -race $$(go list ./... | grep -v /vendor/)

build:
	go build -o ${BINARY}

test:
	${TESTS}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

