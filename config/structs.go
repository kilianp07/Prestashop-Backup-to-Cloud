package config

// Config is the struct that holds ssh and drive config.
type Config struct {
	SshConfig SshConf    `json:"ssh"`
	Folder    folderConf `json:"folder"`
	Mega      MegaConf   `json:"mega"`
}

// ssh is the struct that holds ssh config.
type SshConf struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// folder is the struct that holds folder config.
type folderConf struct {
	Location    string `json:"location"`
	Compression bool   `json:"compression"`
}

type MegaConf struct {
	Username string `json:"login"`
	Password string `json:"password"`
}
