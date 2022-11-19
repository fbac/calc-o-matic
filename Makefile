BINARY = calc-o-matic
BIN_FOLDER = bin

build: clean test
	@mkdir ${BIN_FOLDER}
	@go build -o ${BIN_FOLDER}

clean:
	@rm -f ${BIN_FOLDER}/${BINARY}

test:
	@go test -v ./...

coverage:
	@go test -cover -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
	
doc:
	@go doc github.com/fbac/calc-o-matic