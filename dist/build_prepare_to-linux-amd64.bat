@echo on
go env CGO_ENABLED
go env GOOS
go env GOARCH
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go env CGO_ENABLED
go env GOOS
go env GOARCH
