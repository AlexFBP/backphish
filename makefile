BIN_NAME=backphish

main:
	go build -o ${BIN_NAME}

test:
	go test -v ./... -coverprofile=coverage.out; RET=$$?; \
	go tool cover -html=coverage.out -o coverage.html; exit $$RET

clean:
	rm ${BIN_NAME}
