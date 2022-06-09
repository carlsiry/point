package ssh

import (
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

type Config struct {
	User string
	Pwd  string
	Addr string
}

type Cli struct {
	conf       Config
	client     *ssh.Client
	session    *ssh.Session
	LastResult string
}

func InitCliWith(conf Config) (*Cli, error) {
	cli := &Cli{conf: conf}
	if err := cli.connect(); err != nil {
		return nil, err
	}

	return cli, nil
}

func (c *Cli) connect() error {
	config := &ssh.ClientConfig{}
	config.SetDefaults()
	config.User = c.conf.User
	config.Auth = []ssh.AuthMethod{ssh.Password(c.conf.Pwd)}
	config.Timeout = time.Second * 5
	config.HostKeyCallback = func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	var err error
	c.client, err = ssh.Dial("tcp", c.conf.Addr, config)
	if err != nil {
		return err
	}

	return err
}

func (c *Cli) RunShell(cmd string) (string, error) {
	if c.client != nil {
		if err := c.connect(); err != nil {
			return "", err
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	c.LastResult = string(output)
	return c.LastResult, nil
}
