package main

import (
	"log"
	"os"

	"github.com/kilianp07/Prestashop-Backup-to-Cloud/config"
	scpClient "github.com/kilianp07/Prestashop-Backup-to-Cloud/scp"

	MegaUpload "github.com/kilianp07/Prestashop-Backup-to-Cloud/providers/Mega"
)

func main() {

	var (
		conf   *config.Config
		err    error
		mega   *MegaUpload.Client
		scpCli *scpClient.Client
	)

	if conf, err = config.New(); err != nil {
		log.Fatalf("Error while loading config: %s", err.Error())
		return
	}

	scpCli, err = scpClient.New(conf)
	if err != nil {
		panic(err)
	}

	if err := scpCli.Copy(); err != nil {
		panic(err)
	}

	if err := scpCli.SupressRemoteFile(); err != nil {
		panic(err)
	}

	scpCli.Close()

	mega = MegaUpload.New(conf.Mega)
	if err := mega.Login(); err != nil {
		log.Fatalf("Error while login to Mega: %s", err.Error())
		return
	}

	if err := mega.Upload(scpCli.GetFilePath()); err != nil {
		log.Fatalf("Error while uploading file to Mega: %s", err.Error())
		return
	}

	if err = os.Remove(scpCli.GetFilePath()); err != nil {
		log.Fatalf("Error while removing file: %s", err.Error())
		return
	}
}
