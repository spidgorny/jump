@echo off
set asd=test
echo %asd%
FOR /F "delims=" %%I IN ('go run src\jump.go %1') DO SET V=%%I
rem go run src\jump.go %1
echo %V%
chdir %V%
