package middleware

import (
	"fmt"
	"time"

	"github.com/lachimere/blog-service/global"
	"github.com/lachimere/blog-service/pkg/app"
	"github.com/lachimere/blog-service/pkg/email"
	"github.com/lachimere/blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		Username: global.EmailSetting.Username,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)

				err := defaultMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("exception throws, occurring time: %d", time.Now().Unix()),
					fmt.Sprintf("error info: %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
