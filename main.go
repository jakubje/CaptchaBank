package main

import (
	"github.com/jakubje/captcha_bank/captcha"
	"github.com/jakubje/captcha_bank/clibar"
	"github.com/jakubje/captcha_bank/util"
	"github.com/rs/zerolog/log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config:")
	}

	go clibar.StartCLIBar()
	runCaptchaBank(config)
}

func runCaptchaBank(config util.Config) {

	localStore := captcha.NewLocalStore()
	captchaService, err := captcha.NewCaptchaService(config, localStore)

	server, err := captcha.NewServer(localStore, captchaService)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot create captcha bank: ")
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start captcha bank: ")
	}
}
