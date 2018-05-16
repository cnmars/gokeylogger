keylogger golang

install : go get -u -v github.com/mehrati/gokeylogger

run : gokeylogger

kill : ./kill.sh or kill $(cat pid)