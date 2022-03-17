package mail

import (
	"fmt"
	"net/smtp"

	"github.com/subscan-explorer/subscan-common/core/log"
)

func (r *Request) smtpMail() (err error) {
	SMTP := fmt.Sprintf("%s:%d", r.c.Server, r.c.Port)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", r.c.UserName, r.c.APIKey, r.c.Server), r.c.UserName, r.to, []byte(r.buildMessage())); err != nil {
		log.Errorf("smtp.SendMail error %v", err)
	}
	return
}

func (r *Request) SMTPSend(templateName string, items interface{}) (err error) {
	err = r.parseTemplate(templateName, items)
	if err != nil {
		log.Errorf("parseTemplate error %v", r.to)
		return
	}
	if err = r.smtpMail(); err != nil {
		log.Infof("smtpMail error %v", err)
	}
	return
}
