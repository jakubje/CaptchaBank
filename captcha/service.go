package captcha

import (
	"context"
	"errors"
	"github.com/LouisBillaut/capMonsterTool"
	"github.com/jakubje/captcha_bank/clibar"
	"github.com/jakubje/captcha_bank/util"
	"github.com/rs/zerolog/log"
	"time"
)

type CaptchaService struct {
	config     util.Config
	localStore *LocalStore
}

func NewCaptchaService(config util.Config, store *LocalStore) (*CaptchaService, error) {
	return &CaptchaService{
		config:     config,
		localStore: store,
	}, nil
}

func (captchaService *CaptchaService) createCaptchaTask(ctx context.Context, webSiteUrl string, captchaSiteKey string) {

	capTask := capMonsterTool.NoCaptchaTaskProxyless{Type: CAPTCHA_TYPE, WebsiteURL: webSiteUrl, WebsiteKey: captchaSiteKey}
	taskId, err := capMonsterTool.CreateTask(captchaService.config.CapMonsterApiKey, capTask)

	if err != nil {
		log.Error().Err(err).Msg("Captcha Error")
		return
	}

	captchaRes, err := capMonsterTool.GetTaskResult(captchaService.config.CapMonsterApiKey, taskId)
	for err != nil {
		if !errors.Is(capMonsterTool.CapMonsterProcessing{}, err) {
			log.Error().Err(err).Msg("Solve Captcha - Retrying")
		}

		time.Sleep(2 * time.Second)
		captchaRes, err = capMonsterTool.GetTaskResult(captchaService.config.CapMonsterApiKey, taskId)
	}
	captchaService.localStore.AddCaptcha(webSiteUrl, captchaRes)
	clibar.CaptchaCount++
	return
}
