package convo

import (
	"fmt"
	"net"
)

// Config struct holds all the informations necessary to configure
// pluggabl apps (client, PrimaryNode and SecondaryNode)
type Config struct {
	Type                     string
	Address                  net.IP
	Port                     int
	JobBinaryName            string
	GarbageCollectionTimeout int
	MaxThreads               int
	DatabaseName             string
	DatabaseUsername         string
	DatabasePassword         string
}

type configJSON struct {
	Type       string `json:"type"`
	Addr       string `json:"address"`
	Port       int    `json:"port"`
	Cmd        string `json:"job-binary-name"`
	GcTimeout  int    `json:"garbage-collection-timeout"`
	MaxThreads int    `json:"max-threads"`
	DbName     string `json:"database-name"`
	DbUname    string `json:"database-username"`
	DbPasswd   string `json:"database-password"`
}

func (cc *Config) jsonToConfig(cj configJSON) {
	cc.GarbageCollectionTimeout = cj.GcTimeout
	cc.MaxThreads = cj.MaxThreads
	cc.Type = cj.Type

	cc.Address = net.ParseIP(cj.Addr)
	cc.Port = cj.Port
	cc.JobBinaryName = cj.Cmd

	cc.DatabaseName = cj.DbName
	cc.DatabaseUsername = cj.DbUname
	cc.DatabasePassword = cj.DbPasswd

	return
}

func (cc Config) String() (s string) {
	s = fmt.Sprintf("\tWelcome to pluggabl!\n"+
		"\tThis server runs as %v node.\n"+
		"\tYou can acces it at %v:%v.\n",
		cc.Type, cc.Address, cc.Port)

	if cc.Type == "secondary" {
		s = fmt.Sprintf("%v"+
			"\tNumber of available concurrent processing threads is set to %v.\n"+
			"\tYou will be processing jobs using %v.\n",
			s, cc.MaxThreads, cc.JobBinaryName)
	}

	s = fmt.Sprintf("%v"+
		"\tGarbage collection target percentage is set to %v.\n",
		s, cc.GarbageCollectionTimeout)

	return
}
