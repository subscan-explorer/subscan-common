package mail

import (
	"github.com/mattbaird/gochimp"
	"github.com/subscan-explorer/subscan-common/core/log"
)

func (r *Request) mailchimp() (err error) {
	mandrillApi, err := gochimp.NewMandrill(r.c.APIKey)
	if err != nil {
		log.Errorf("Error instantiating client %v", err)
		return
	}
	recipients := []gochimp.Recipient{}
	for _, v := range r.to {
		recipient := gochimp.Recipient{Email: v}
		recipients = append(recipients, recipient)
	}
	message := gochimp.Message{
		Subject:   r.subject,
		FromEmail: r.from,
		FromName:  r.from,
		To:        recipients,
		Html:      r.body,
		AutoHtml:  true,
	}
	_, err = mandrillApi.MessageSend(message, false)
	if err != nil {
		log.Errorf("mandrillApi.MessageSend error %v", err)
		return
	}
	return
}

func (r *Request) MailchimpSend(templateName string, items interface{}) (err error) {
	err = r.parseTemplate(templateName, items)
	if err != nil {
		log.Errorf("parseTemplate error %v", r.to)
		return
	}
	if err = r.mailchimp(); err != nil {
		log.Errorf("mailchimp error %v", err)
	}
	return
}
