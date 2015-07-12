.PHONY: test serve install experiment

# need to skip the wip directory

all: install test

install:
	go get github.com/jmhodges/justrun
	go install github.com/jmhodges/justrun
	go get ./bits/... ./genetic/... ./geo/... ./image/... ./rle/... ./server/... ./td6/... ./tracks/...
	go install ./bits/... ./genetic/... ./geo/... ./image/... ./rle/... ./server/... ./td6/... ./tracks/...

test:
	go test -timeout 1s \
		./bits/... \
		./genetic/... \
		./geo/... \
		./image/... \
		./rle/... \
		./server/... \
		./td6/... \
		./tracks/...

serve:
	find . -name '*.go' -o -name '*.html' | justrun -stdin -v=true -c 'go run server/main.go --template-directory="$(PWD)/server/templates" --static-directory="$(PWD)/server/static"'

experiment:
	run_experiment

compress: install
	get_old_experiments | bash scripts/compress.bash
