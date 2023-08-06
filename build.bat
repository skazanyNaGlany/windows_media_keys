REM build without console window, production build
del *.txt
del *.exe
go build -ldflags -H=windowsgui .
