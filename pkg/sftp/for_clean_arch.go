package sftp

import (
	"errors"

	"github.com/pkg/sftp"
)

// Point is one available sft client handler.
type Point struct {
	client *sftp.Client
}

// NewPoint inits one sftp client point.
func NewPoint(username, passwd, host string, port int) (*Point, error) {
	cli, err := NewCli(username, passwd, host, port)
	if err != nil {
		return nil, err
	}

	return &Point{
		client: cli,
	}, nil
}

// Upload is used for top level call direct for upload file.
func (p *Point) Upload(fromPath, toPath string) error {
	if p.client == nil {
		return errors.New("component not be init")
	}

	return Upload(fromPath, toPath, p.client)
}
