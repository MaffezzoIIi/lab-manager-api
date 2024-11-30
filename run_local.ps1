# build_and_run.ps1

go mod tidy 

go build -o ./bin/ ./cmd/...

./bin/cmd.exe