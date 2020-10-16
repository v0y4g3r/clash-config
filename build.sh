#! /bin/bash

[[ ! -d ./bin ]] && mkdir bin

for f in $(ls ./cmd); do
  echo "Building cmd $f"
  go build -o ./bin/$f  cmd/$f/main.go
done
