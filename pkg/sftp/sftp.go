// Package sftp packaging some api of "github.com/pkg/sftp".
//
package sftp

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// NewCli init sftp client.
func NewCli(username, passwd, host string, port int) (*sftp.Client, error) {

	sshConn, err := ssh.Dial("tcp", addr(host, port), config(username, passwd))
	if err != nil {
		return nil, fmt.Errorf("ssh connect failed: %s", err)
	}

	return sftp.NewClient(sshConn)
}

// Upload uploads file from local path (@fromPath) to home dir ~/foo.zip (@toPath = foo.zip).
func Upload(fromPath string, toPath string, cli *sftp.Client) error {

	srcFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := cli.Create(toPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	io.Pipe()
	_, err = io.Copy(dstFile, srcFile)
	return err
}

// generate whole addr string.
func addr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

// generate ssh client config instance point.
func config(username, passwd string) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(passwd)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil },
		Timeout:         30 * time.Second,
	}
}
