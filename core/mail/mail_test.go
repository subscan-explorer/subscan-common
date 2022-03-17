package mail

import (
	"fmt"
	"testing"
)

var cfg *Config

func init() {
	cfg = &Config{
		Server:   "smtp.qq.com",
		Port:     587,
		UserName: "",
		APIKey:   "",
	}
}

// go test -v -test.run TestSMTPSend
func TestSMTPSend(t *testing.T) {
	req := NewRequest(cfg, "cort-xie@foxmail.com", "Subscan的测试邮件", []string{"i@xiequan.info"}, []string{})
	err := req.SMTPSend("./template.html", map[string]string{"username": "Xiequan"})
	if err != nil {
		fmt.Println(err)
	}
}

// go test -v -test.run TestMailchimpSend
func TestMailchimpSend(t *testing.T) {
	req := NewRequest(cfg, "cort-xie@foxmail.com", "Subscan的测试邮件", []string{"i@xiequan.info"}, []string{})
	err := req.MailchimpSend("./template.html", map[string]string{"username": "Xiequan"})
	if err != nil {
		fmt.Println(err)
	}
}
