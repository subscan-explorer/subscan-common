package mail

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type Config struct {
	Server   string //smtp server .
	Port     int    // smtp server port
	UserName string // email or username
	APIKey   string // password or apikey
	FrontEnd string // front end landing page host
}

type Request struct {
	c       *Config
	from    string
	to      []string
	cc      []string
	subject string
	body    string
}

func NewRequest(c *Config, from, subject string, to, cc []string) *Request {
	return &Request{
		c:       c,
		from:    from,
		to:      to,
		cc:      cc,
		subject: subject,
	}
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2\n"
)

func (r *Request) parseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}
	r.body = buffer.String()
	return nil
}

func (r *Request) buildMessage() string {
	msg := ""
	msg += fmt.Sprintf("From: %s\r\n", r.from)
	if len(r.to) > 0 {
		msg += fmt.Sprintf("To: %s\r\n", r.to[0])
	}
	msg += "Subject:" + r.subject + "\r\n"
	if len(r.cc) > 0 {
		msg += fmt.Sprintf("Cc: %s\r\n", strings.Join(r.cc, ";"))
	}
	msg += fmt.Sprintf("%s\r\n", MIME)
	msg += fmt.Sprintf("\r\n%s\r\n", r.body)
	return msg
}
