@echo off

start cmd /c "cd TeamsService && go run main.go"

start cmd /c "cd predictions && py grpc_server.py"

start cmd /c "cd api-gateway && go run main.go"

echo Back to the original window

pause