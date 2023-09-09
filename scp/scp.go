package scpClient

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/kilianp07/Prestashop-Backup-to-Cloud/config"
	sshClientLib "github.com/kilianp07/Prestashop-Backup-to-Cloud/ssh"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	sshClient *ssh.Client
	scpClient scp.Client
	config    *config.Config
	file      *os.File
	filePath  string
}

func New(conf *config.Config) (*Client, error) {
	c := &Client{
		config: conf,
	}

	err := c.createSshClient()
	if err != nil {
		return nil, err
	}

	err = c.createScpClient()
	if err != nil {
		return nil, err
	}

	// We want to copy www folder so if compression is enabled to create a tar file else create a folder
	if c.config.Folder.Compression {

		if err = c.createDestinationFile(); err != nil {
			return nil, err
		}

		if err = c.compress(); err != nil {
			return nil, err
		}
	} else {
		err = c.createDestinationFolder()
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) compress() error {
	session, err := c.sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// configure terminal mode
	modes := ssh.TerminalModes{
		ssh.ECHO: 0, // supress echo

	}
	// run terminal session
	if err := session.RequestPty("xterm", 50, 80, modes); err != nil {
		return err
	}

	// Create a pipe to hold the output
	stdoutPipe, stdoutWriter := io.Pipe()
	session.Stdout = stdoutWriter

	commands := []string{
		fmt.Sprintf("cd %s && tar -czf backup.tar %s", c.config.Folder.Location, c.config.Folder.Location),
	}

	log.Printf("Compressing %s...", c.config.Folder.Location)
	for _, cmd := range commands {
		if err := session.Start(cmd); err != nil {
			return err
		}
		go func() {
			_, err := io.Copy(os.Stdout, stdoutPipe)
			if err != nil {
				log.Println("Error during output acquisition :", err)
			}
		}()
		_ = session.Wait()
	}
	log.Println("Compressing files... Done")
	return nil
}

func (c *Client) createScpClient() error {
	client, err := scp.NewClientBySSH(c.sshClient)
	if err != nil {
		return err
	}
	c.scpClient = client
	return nil
}

func (c *Client) createSshClient() error {
	client, err := sshClientLib.New(*c.config.GetSSHConfig())
	if err != nil {
		return err
	}
	err = client.Connect()
	if err != nil {
		return err
	}

	c.sshClient = client.GetConnection()
	return nil
}

func (c *Client) createDestinationFile() (err error) {
	c.filePath = fmt.Sprintf("backup_%s.tar", time.Now().Format("2006-01-02_15-04-05"))
	c.file, err = os.Create(c.filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (c *Client) createDestinationFolder() error {
	return os.Mkdir("backup", 0777)
}

func (c *Client) Close() error {
	return c.sshClient.Close()
}

func (c *Client) Copy() error {
	var filename string
	if c.config.Folder.Compression {
		filename = fmt.Sprintf("%s/backup.tar", c.config.Folder.Location)
	} else {
		filename = c.config.Folder.Location
	}
	log.Printf("Copying %s...", filename)
	return c.scpClient.CopyFromRemote(context.Background(), c.file, filename)
}

func (c *Client) GetFilePath() string {
	return c.filePath
}

func (c *Client) SupressRemoteFile() error {
	session, err := c.sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// configure terminal mode
	modes := ssh.TerminalModes{
		ssh.ECHO: 0, // supress echo

	}
	// run terminal session
	if err := session.RequestPty("xterm", 50, 80, modes); err != nil {
		return err
	}

	// Create a pipe to hold the output
	stdoutPipe, stdoutWriter := io.Pipe()
	session.Stdout = stdoutWriter

	commands := []string{
		fmt.Sprintf("cd %s && rm backup.tar", c.config.Folder.Location),
	}
	log.Printf("Deleting compressed file %s...", c.config.Folder.Location)
	for _, cmd := range commands {
		if err := session.Start(cmd); err != nil {
			return err
		}
		go func() {
			_, err := io.Copy(os.Stdout, stdoutPipe)
			if err != nil {
				log.Println("Error during output acquisition :", err)
			}
		}()
		_ = session.Wait()
	}
	log.Println("Deleting file... Done")
	return nil
}
