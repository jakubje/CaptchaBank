package captcha

import (
	"context"
	http "github.com/bogdanfinn/fhttp"
	"github.com/gin-gonic/gin"
	"github.com/jakubje/captcha_bank/clibar"
)

const (
	CAPTCHA_TYPE = "NoCaptchaTaskProxyless"
)

type getCaptchaRequest struct {
	WebsiteUrl string `json:"website_url"`
}

func (server *Server) getCaptcha(ctx *gin.Context) {
	var req getCaptchaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	captcha, ok := server.localStore.GetCaptcha(req.WebsiteUrl)

	if ok {
		ctx.JSON(http.StatusOK, gin.H{"response": captcha})
		clibar.CaptchaCount--
	} else {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"response": "no captcha available"})
	}
}

type addCaptchaToBankRequest struct {
	WebsiteUrl     string `json:"website_url"`
	CaptchaSiteKey string `json:"captcha_site_key"`
}

func (server *Server) addCaptchaToBank(ctx *gin.Context) {
	var req addCaptchaToBankRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx1, cancel := context.WithCancel(context.Background())
	defer cancel()

	go server.captchaService.createCaptchaTask(ctx1, req.WebsiteUrl, req.CaptchaSiteKey)

	ctx.JSON(http.StatusOK, gin.H{"message": "captcha requested"})
	return
}
