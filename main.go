package main

import scpClient "github.com/kilianp07/Prestashop-Backup-to-Google-Drive/scp"

func main() {
	scpCli, err := scpClient.New()
	if err != nil {
		panic(err)
	}
	defer scpCli.Close()
	if err := scpCli.Copy(); err != nil {
		panic(err)
	}
}
