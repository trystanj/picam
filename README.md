# piserver

env GOOS=linux GOARCH=arm GOARM=5 go build cmd/server/main.go && scp main pi@raspberrypi.local:/home/pi
