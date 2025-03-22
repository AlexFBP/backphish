BIN_NAME=backphish

main:
	go build -o ${BIN_NAME}

clean:
	rm ${BIN_NAME}
