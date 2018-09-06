@echo off
REM FOR /F "delims=" %%I IN ('go run src\jump.go %1') DO SET V=%%I
FOR /F "delims=" %%I IN ('%~dp0jump-walker %1') DO SET V=%%I
echo chdir %V%
chdir %V%
dir
