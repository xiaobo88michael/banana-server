package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/utils"
)

type EmailService struct{}

func (e *EmailService) EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = utils.EmailTest(subject, body)
	return err
}

func (e *EmailService) SendEmail(to, subject, body string) (err error) {
	err = utils.Email(to, subject, body)
	return err
}
