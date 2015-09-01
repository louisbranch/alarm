package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const (
	osError = "Platform not supported. Alarm requires Linux or OS X >= 10.7"
	help    = `Usage: alarm title [description] [time] &

examples:
alarm Pizza 15m &
alarm "Pomodoro Timer" "Long break" 1h30m &
`
)

func main() {
	var title, desc, timer string
	var delay time.Duration
	var cmd *exec.Cmd

	switch len(os.Args) {
	case 2:
		title = os.Args[1]
	case 3:
		title, timer = os.Args[1], os.Args[2]
	case 4:
		title, desc, timer = os.Args[1], os.Args[2], os.Args[3]
	default:
		fmt.Print(help)
		os.Exit(1)
	}

	if timer != "" {
		var err error
		delay, err = time.ParseDuration(timer)
		if err != nil {
			fmt.Printf("Time parsing error, %s. Valid examples: 1h40m, 5m, 30s\n", err)
			os.Exit(1)
		}
	}

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("notify-send", title, desc, "--expire-time=0")
	case "darwin":
		msg := fmt.Sprintf(`display notification "%s" with title "%s"`, desc, title)
		cmd = exec.Command("osascript", "-e", msg)
	default:
		fmt.Println(osError)
		os.Exit(1)
	}

	if delay > 0 {
		time.Sleep(delay)
	}
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
