#!/bin/bash

compress_directory() {
    filename="$1"
    if [[ -z "$filename" ]]; then
        echo "filename not set. exiting"
        exit 1
    fi
    if [[ -d "$filename/iterations" ]]; then
        echo "compressing $filename/iterations"
        tar -zcf "$filename/iterations.tar.gz" "$filename/iterations"
        if [[ -n "$filename" ]]; then
            rm -rf "$filename/iterations"
        fi
    fi

}

unzip_directory() {
    filename="$1"
    if [[ -z "$filename" ]]; then
        echo "filename not set. exiting"
        exit 1
    fi
    if [[ -f "$filename/iterations.tar.gz" ]]; then
        tar -xf "$filename/iterations.tar.gz" -C /
    fi
}

main() {
    while read filename; do
        compress_directory "$filename"
    done
}

main
