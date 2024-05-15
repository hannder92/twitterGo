git add .
git commit -m "Ultimo commit"
git push
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -tags lambda.norpc -o bootstrap main.go
%USERPROFILE%\Go\bin\linux_amd64\build-lambda-zip.exe -o myFunction.zip bootstrap