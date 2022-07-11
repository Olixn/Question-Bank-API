RD /s /Q %cd%\server

echo "Create server"
MD server

echo "Start..."

go build main.go

MOVE main.exe %cd%\server
COPY config.yaml %cd%\server
COPY notice.json %cd%\server

echo "Finished!"

pause