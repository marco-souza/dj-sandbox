i: install

install: ./main.go
	go install .

test: install
	go test ./...

run: install
	djc

