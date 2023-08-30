package captcha

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router         *gin.Engine
	ctx            gin.Context
	localStore     *LocalStore
	captchaService *CaptchaService
}

func NewServer(store *LocalStore, captchaService *CaptchaService) (*Server, error) {

	server := &Server{
		localStore:     store,
		captchaService: captchaService,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/captchaBank", server.addCaptchaToBank)
	router.POST("/getCaptcha/", server.getCaptcha)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
