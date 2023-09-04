cd ..
@REM linux
go build -o .\dist\server\user-server  .\internal\mircoservice\user\
go build -o .\dist\server\video-server  .\internal\mircoservice\video\
go build -o .\dist\server\interaction-server  .\internal\mircoservice\interaction\


@REM windows
go build -o .\dist\server\user-server.exe  .\internal\mircoservice\user\