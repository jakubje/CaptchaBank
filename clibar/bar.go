package clibar

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os/exec"
	"runtime"
	"time"
)

// Call clibar.StartCLIBar() inside site module to start CLI bar updating

var (
	CaptchaCount int = 0
)

func InitTitle() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", "title", "Captcha Bank")
		err := cmd.Run()
		if err != nil {
			log.Error().Err(err)
		}
	}
}

func UpdateBar() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", "title", fmt.Sprintf("Captcha: %d", CaptchaCount))
		err := cmd.Run()
		if err != nil {
			log.Error().Err(err)
		}
	}
}

func StartCLIBar() {
	for {
		UpdateBar()

		time.Sleep(250 * time.Millisecond)
	}
}
