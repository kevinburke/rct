.PHONY: test

test:
	go test ./...

install:
	go get github.com/jmhodges/justrun
	go install github.com/jmhodges/justrun

serve:
	find . -name '*.go' -o -name '*.html' | justrun -stdin -v=true -c 'go run server/main.go --template-directory="$(PWD)/server/templates"'
