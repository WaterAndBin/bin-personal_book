package email

import (
	"github.com/go-mail/mail/v2"
)

func SendCode(to string, code string) error {
	// 创建一封邮件对象
	// 后面会往里面设置：发件人、收件人、标题、正文等
	m := mail.NewMessage()

	// 设置邮件头 From
	m.SetHeader(
		"From",
		"waterandbin@163.com",
	)

	// 设置收件人
	m.SetHeader(
		"To",
		to,
	)

	// 设置邮件标题
	m.SetHeader(
		"Subject",
		"记账本注册验证码",
	)

	// 设置邮件正文
	m.SetBody(
		"text/html",
		"您的验证码："+code,
	)

	// 创建 SMTP 拨号器
	d := mail.NewDialer(
		"smtp.163.com",
		465,
		"waterandbin@163.com",
		"",
	)

	// 连接 SMTP服务器，并发送邮件
	return d.DialAndSend(m)
}
