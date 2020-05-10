package image

import (
	"context"
	"os/exec"
	"strconv"
	"time"
)

func Echo(arg string) (result []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "echo", "cmd echo")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return output, nil
}

// make an ImageOptions struct?
func Capture(quality int) (result []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// --output - doesn't save a file and just sends the bytes to stdout, which is what we want
	cmd := exec.CommandContext(ctx, "raspistill", "--quality", strconv.Itoa(quality), "--output", "-")

	output, err := cmd.Output()
	return output, err
}
