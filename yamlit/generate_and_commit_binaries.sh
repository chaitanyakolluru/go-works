#1 /bin/bash

GOOS=linux GOARCH=amd64 go build .
mv yamlit  ./binaries/linux/amd64
                                        
GOOS=windows GOARCH=amd64 go build .
mv yamlit.exe  ./binaries/windows/amd64
                                        
GOOS=darwin GOARCH=amd64 go build .
mv yamlit  ./binaries/mac/amd64

git add ./binaries
git commit -m "creating executables for linux/windows/mac Oses"
