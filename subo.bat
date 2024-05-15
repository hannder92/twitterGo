git add .
git commit -m "Ultimo commit"
git push
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -tags lambda.norpc -o main main.go
%USERPROFILE%\Go\bin\build-lambda-zip.exe -o main.zip main