install:
	go get -v -t -d ./...
	go build -o ohmygosh src/*.go
	mv ohmygosh /usr/bin