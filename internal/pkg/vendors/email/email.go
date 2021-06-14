package email

import (
	"crypto/tls"
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/util"
	"net"
	"net/smtp"
)

const (
	ID       = "email"
	Host     = "host"
	Port     = "port"
	Username = "username"
	Password = "password"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
}

func SendEmail(config *Config, toMail string, title, content string) error {
	header := make(map[string]string)
	header["From"] = "Assistant" + "<" + config.Username + ">"
	header["To"] = toMail
	header["Subject"] = title
	header["Content-Type"] = "text/html; charset=UTF-8"
	body := content
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	auth := smtp.PlainAuth(
		"",
		config.Username,
		config.Password,
		config.Host,
	)
	return sendMailUsingTLS(
		fmt.Sprintf("%s:%d", config.Host, config.Port),
		auth,
		config.Username,
		[]string{toMail},
		util.StringToByte(message),
	)
}

func dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		return nil, err
	}
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	return smtp.NewClient(conn, host)
}

func sendMailUsingTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
	c, err := dial(addr)
	if err != nil {
		return err
	}
	defer func() { _ = c.Close() }()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
