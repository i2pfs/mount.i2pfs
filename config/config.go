package config

import (
	"github.com/i2pfs/mount.i2pfs/log"
	"gopkg.in/ini.v1"
)

type Config struct {
	SamUrl        string
	PeersFilePath string
}

func GetDefaults() Config {
	return Config{
		SamUrl:        "127.0.0.1:7656",
		PeersFilePath: "/var/lib/i2pfsd/peers.txt",
	}
}

func getParameter(section *ini.Section, result *string, parameter_name string) {
	value := section.Key(parameter_name).String()
	if value == "" {
		return
	}
	*result = value
}

func (cfg *Config) ReloadConfig() {
	cfgFile, err := ini.Load("~/.i2pfs.ini")
	log.Debugf(`Cannot open file "~/.i2pfs.ini": %v`, err)
	if err != nil {
		cfgFile, err = ini.Load("/etc/i2pfs.ini")
		if err != nil {
			log.Debugf(`Cannot open file "/etc/i2pfs.ini": %v`, err)
		}
	}
	if err != nil {
		return
	}
	section := cfgFile.Section("i2pfs")

	getParameter(section, &cfg.SamUrl, "sam_url")
	getParameter(section, &cfg.PeersFilePath, "peers_file_path")

	return
}
