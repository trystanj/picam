# piserver

env GOOS=linux GOARCH=arm GOARM=5 go build cmd/main.go && scp main pi@raspberrypi.local:/home/pi
