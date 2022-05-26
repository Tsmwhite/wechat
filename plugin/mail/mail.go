package mail

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net"
	"net/smtp"
	"os"
	"strings"
	"time"
)

//邮件
type Mail struct {
	Server Server
	Msg    Msg
}

//服务配置
type Server struct {
	host     string
	account  string
	password string
	port     string
	ssl      bool
}

//邮件信息
type Msg struct {
	to          string
	subject     string
	body        string
	attachment  Attachment
	contentType string
}

//附件
type Attachment struct {
	name        string
	contentType string
	withFile    bool
}

//设置邮件内容
func (m *Mail) SetMessage(to, sub, body string) *Mail {
	m.Msg.to = to
	m.Msg.subject = sub
	m.Msg.body = body
	return m
}

//添加附件
func (m *Mail) AddFile(filename string) *Mail {
	if _, e := os.Stat(filename); e == nil {
		m.Msg.attachment.withFile = true
		m.Msg.attachment.name = filename
		m.Msg.attachment.contentType = strings.Split(filename, ".")[1]
	}
	return m
}

//返回一个默认邮件结构
func ResMail() *Mail {
	return &Mail{
		Server: Server{
			host:     "smtp.163.com",
			account:  "xxxxxxx@163.com",
			password: "xxxxxxxx",
			port:     "465",
			ssl:      true,
		},
		Msg: Msg{
			contentType: "html",
		},
	}
}

func (s Server) Auth() smtp.Auth {
	return smtp.PlainAuth("", s.account, s.password, s.host)
}

//发送邮件
func (m *Mail) Send() error {
	server := m.Server
	auth := server.Auth()
	message := m.Msg
	if message.contentType == "html" {
		message.contentType = "text/html;charset=UTF-8"
	} else {
		message.contentType = "text/plain;charset=utf-8"
	}

	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	Header := make(map[string]string)
	Header["From"] = server.account
	Header["To"] = message.to
	Header["Subject"] = message.subject
	Header["Content-Type"] = "multipart/mixed;boundary=" + boundary
	Header["Mime-Version"] = "1.0"
	Header["Date"] = time.Now().String()
	writeHeader(buffer, Header)

	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type:" + message.contentType + "\r\n"
	body += "\r\n" + message.body + "\r\n"
	buffer.WriteString(body)

	if message.attachment.withFile {
		attachment := "\r\n--" + boundary + "\r\n"
		attachment += "Content-Transfer-Encoding:base64\r\n"
		attachment += "Content-Disposition:attachment\r\n"
		attachment += "Content-Type:" + message.attachment.contentType + ";name=\"" + message.attachment.name + "\"\r\n"
		buffer.WriteString(attachment)
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln(err)
			}
		}()
		writeFile(buffer, message.attachment.name)
	}

	buffer.WriteString("\r\n--" + boundary + "--")
	send_to := strings.Split(message.to, ";")
	if server.ssl {
		return SendMailUsingTLS(server.host+":"+server.port, auth, server.account, send_to, buffer.Bytes())
	} else {
		return smtp.SendMail(server.host+":"+server.port, auth, server.account, send_to, buffer.Bytes())
	}
}

func writeHeader(buffer *bytes.Buffer, Header map[string]string) string {
	header := ""
	for key, value := range Header {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)
	return header
}

// read and write the file to buffer
func writeFile(buffer *bytes.Buffer, fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(payload, file)
	buffer.WriteString("\r\n")
	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
}

//return a smtp client
func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
//len(to)>1时,to[1]开始提示是密送
func SendMailUsingTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) (err error) {
	//create smtp client
	c, err := Dial(addr)
	if err != nil {
		log.Println("Create smpt client error:", err)
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
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
