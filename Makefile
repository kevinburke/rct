.PHONY: test serve install experiment

install:
	go get github.com/jmhodges/justrun
	go install github.com/jmhodges/justrun
	go get ./...
	go install ./...

test:
	go test ./...

serve:
	find . -name '*.go' -o -name '*.html' | justrun -stdin -v=true -c 'go run server/main.go --template-directory="$(PWD)/server/templates" --static-directory="$(PWD)/server/static"'

experiment:
	go run genetic/runners/main.go
