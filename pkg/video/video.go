package video

import (
	"context"
	"io"
	"os/exec"
	"time"
)

// make an a VideoOptions struct?
func Capture() (result []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// --output - doesn't save a file and just sends the bytes to stdout, which is what we want
	cmd := exec.CommandContext(ctx, "raspivid", "--bitrate", "10000000", "--output", "-")

	// but this captures the whole video, even though --output - streams it to stdout
	// need to stream back to the writer...
	output, err := cmd.Output()
	return output, err
}



func Stream(w io.Writer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// --output - doesn't save a file and just sends the bytes to stdout, which is what we want
	cmd := exec.CommandContext(ctx, "raspivid", "--bitrate", "10000000", "--output", "-")
	cmd.Stdout = w

	return cmd.Run()
}

func Save(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "raspivid", "--bitrate", "10000000", "--output", name)

	return cmd.Run()
}

// wrap in mp4
func MP4Box(input string, output string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// --output - doesn't save a file and just sends the bytes to stdout, which is what we want
	cmd := exec.CommandContext(ctx, "MP4Box", "-add", input, output)
	return cmd.Run()
}


