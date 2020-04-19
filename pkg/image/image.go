package image

import (
	"context"
	"os/exec"
	"time"
)

func Image(arg string) (result []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "echo", "cmd echo")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return output, nil
}
