package main

import (
	"os"
	"fmt"
	"strconv"
)

const fifoName = "/tmp/pidi-spotify.fifo"

func main() {
	event := os.Getenv("PLAYER_EVENT")
	id := os.Getenv("TRACK_ID")
	pos, _ := strconv.Atoi(os.Getenv("POSITION_MS"))
	vol, _ := strconv.Atoi(os.Getenv("VOLUME"))

	var cmd string

	switch event {
	case "start", "started":
		cmd = fmt.Sprintf("track:%s:%d\n", id, pos)
	case "change", "changed":
		cmd = fmt.Sprintf("track:%s:%d\n", id, pos)
	case "playing":
		cmd = fmt.Sprintf("track:%s:%d\n", id, pos)
	case "paused":
		cmd = fmt.Sprintf("pause:%s\n", id)
	case "stop", "stopped":
		cmd = fmt.Sprintf("track:%s:%d\n", id, pos)
	case "volume_set":
		cmd = fmt.Sprintf("volume:%d\n", 100 * vol / 65535)

	default:
		fmt.Println("Unknown PLAYER_EVENT: ", event)
	}

	if cmd != "" {
		f, _ := os.OpenFile(fifoName, os.O_RDWR, 0600)
		defer f.Close()

		f.Write([]byte(cmd))
	}
}
