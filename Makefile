BINARY=freya-rest-api
FORMAT=go fmt $$(go list ./... | grep -v /vendor/)
CHECK=go vet $$(go list ./... | grep -v /vendor/)
TESTS=go test -race $$(go list ./... | grep -v /vendor/)

build:
	go build -o ${BINARY}

test:
	${FORMAT}
	${CHECK}
	${TESTS}

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

