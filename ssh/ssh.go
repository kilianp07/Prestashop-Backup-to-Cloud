package sshClient

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kilianp07/Prestashop-Backup-to-Google-Drive/config"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

type Client struct {
	sshConfig ssh.ClientConfig
	conn      *ssh.Client
	host      string
}

func New(sshConfig config.SshConf) (*Client, error) {
	hostkey, error := knownhosts.New(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if error != nil {
		return nil, error
	}

	host := sshConfig.Host + ":" + fmt.Sprint(sshConfig.Port)
	return &Client{
		sshConfig: ssh.ClientConfig{
			User: sshConfig.User,
			Auth: []ssh.AuthMethod{
				ssh.Password(sshConfig.Password),
			},
			HostKeyCallback:   hostkey,
			HostKeyAlgorithms: []string{ssh.KeyAlgoED25519},
		},
		host: host,
	}, nil
}

// Connect opens a ssh connection to the host.
func (c *Client) Connect() error {
	var err error
	c.conn, err = ssh.Dial("tcp", c.host, &c.sshConfig)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the ssh connection.
func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) GetConnection() *ssh.Client {
	return c.conn
}
