.PHONY: test serve install experiment

test:
	go test ./...

install:
	go get github.com/jmhodges/justrun
	go install github.com/jmhodges/justrun
	go install ./...

serve:
	find . -name '*.go' -o -name '*.html' | justrun -stdin -v=true -c 'go run server/main.go --template-directory="$(PWD)/server/templates" --static-directory="$(PWD)/server/static"'

experiment:
	go run genetic/runners/main.go
