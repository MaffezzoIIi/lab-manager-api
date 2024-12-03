# build_and_run.ps1

go mod tidy

go build -o ./bin/ .

swag init

./bin/lab-manager-api.exe