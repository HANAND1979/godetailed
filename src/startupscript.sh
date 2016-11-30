#!/bin/bash
sudo apt-get install -y git golang
#Create go folders 
mkdir go
cd go
mkdir src pkg bin 
#set go environment variables 
export GOPATH=/home/anand_hariharan/go
export GOBIN=/home/anand_hariharan/go/bin
echo "go environment setup completed"
#copy go source files 
cd src
gsutil cp gs://goconcurrency/web-jsonupd.go ./
echo "file copy successful"
#get  go packages 
go get github.com/gorilla/mux
echo "git package fetched"
#run go  api server code 
echo "starting web server"
go run web-jsonupd.go 