package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"gopkg.in/vansante/go-ffprobe.v2"
)

func main() {
	stdinBytes, err := io.ReadAll(os.Stdin)
	check(err)

	videos := strings.Split(string(stdinBytes), "\n")
	total, err := getTotalDuration(videos)
	check(err)

	fmt.Println(total)
}

func check(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func getVideoDuration(videoPath string) (time.Duration, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	data, err := ffprobe.ProbeURL(ctx, videoPath)
	if err != nil {
		return 0, err
	}

	return data.Format.Duration(), nil
}

func getTotalDuration(videosPath []string) (time.Duration, error) {
	var total time.Duration
	for _, path := range videosPath {
		path := strings.TrimSpace(path)
		if path == "" {
			continue
		}

		duration, err := getVideoDuration(path)
		check(err)

		total += duration
	}

	return total, nil
}
