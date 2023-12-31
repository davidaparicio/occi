#!/usr/bin/env bash

FuzzFUNC="Fuzz" #"FuzzReverse"

echo "Let's Test"
go test -v ./... -coverprofile=coverage.out

echo "Let's Fuzz"
go test -fuzz ${FuzzFUNC} -fuzztime 15s

echo "Let's Mutate"
# go get -t -v github.com/avito-tech/go-mutesting/...
go-mutesting github.com/avito-tech/go-mutesting/...

echo "Let's Bench"
go test -v -run=^$ -bench . -benchmem -benchtime=3s ./

if ! command -v gosec &> /dev/null
then
  echo "gosec required but it's not installed. Skipping."
  exit
else
  echo "Let's Gosec"
  gosec ./...
  #gosec -no-fail -fmt sarif -out results.sarif ./...
fi