.PHONY: test serve install experiment

# need to skip the wip directory


install:
	go get github.com/jmhodges/justrun
	go install github.com/jmhodges/justrun
	go get ./bits/... ./genetic/... ./geo/... ./image/... ./rle/... ./server/... ./td6/... ./tracks/...
	go install ./bits/... ./genetic/... ./geo/... ./image/... ./rle/... ./server/... ./td6/... ./tracks/...

test:
	go test ./bits/... ./genetic/... ./geo/... ./image/... ./rle/... ./server/... ./td6/... ./tracks/...

serve:
	find . -name '*.go' -o -name '*.html' | justrun -stdin -v=true -c 'go run server/main.go --template-directory="$(PWD)/server/templates" --static-directory="$(PWD)/server/static"'

experiment:
	go run genetic/runners/main.go
