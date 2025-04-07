i: install

install: ./main.go
	go install .

run: install
	dj-sandbox

