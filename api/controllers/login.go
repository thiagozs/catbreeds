package controllers

import "github.com/gin-gonic/gin"

// @Summary Refesh Token
// @Description Get a new fesh token for your account
// @Accept json
// @Produce json
// @Success 200 {object} models.Login
// @Router /login/refresh_token [get]
// @Param login body models.ReqLogin true "Login"
func (ctl *CtlRepo) RefreshHandler() func(c *gin.Context) {
	return ctl.GetJwt().RefreshHandler
}

// @Summary Login
// @Description Authorize your account with jwt token
// @Accept json
// @Produce json
// @Success 201 {object} models.Login
// @Router /login [post]
// @Param login body models.ReqLogin true "Login"
func (ctl *CtlRepo) LoginHandler() func(c *gin.Context) {
	return ctl.GetJwt().LoginHandler
}

func (ctl *CtlRepo) MiddlewareFunc() gin.HandlerFunc {
	return ctl.GetJwt().MiddlewareFunc()
}
