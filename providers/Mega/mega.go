package MegaUpload

import (
	"log"

	"github.com/kilianp07/Prestashop-Backup-to-Cloud/config"
	"github.com/t3rm1n4l/go-mega"
)

type Client struct {
	megaClient *mega.Mega
	email      string
	password   string
}

func New(conf config.MegaConf) *Client {
	c := &Client{
		email:    conf.Username,
		password: conf.Password,
	}
	c.megaClient = mega.New()
	return c
}

func (c *Client) Login() error {
	log.Println("Login to Mega...")
	return c.megaClient.Login(c.email, c.password)
}

func (c *Client) Upload(filepath string) error {
	log.Println("Uploading file to Mega...")
	_, err := c.megaClient.UploadFile(filepath, c.megaClient.FS.GetRoot(), filepath, nil)
	return err
}
