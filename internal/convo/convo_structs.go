package convo

import (
	"net"
)

// Config struct holds all the informations necessary to configure
// pluggabl apps (client, PrimaryNode and SecondaryNode)
type Config struct {
	Type                     string
	PrimaryNodeAddress       net.IP
	PrimaryNodePort          int
	SecondaryNodeAddress     net.IP
	SecondaryNodePort        int
	JobBinaryName            string
	GarbageCollectionTimeout int
	MaxThreads               int
	DatabaseAddress          net.IP
	DatabasePort             int
	DatabaseName             string
	DatabaseUsername         string
	DatabasePassword         string
}

type configJSON struct {
	Type       string `json:"type"`
	PAddr      string `json:"primary-node-address"`
	PPort      int    `json:"primary-node-port"`
	SAddr      string `json:"secondary-node-address"`
	SPort      int    `json:"secondary-node-port"`
	SCmd       string `json:"job-binary-name"`
	GcTimeout  int    `json:"garbage-collection-timeout"`
	MaxThreads int    `json:"max-threads"`
	DbAddress  string `json:"database-address"`
	DbPort     int    `json:"database-port"`
	DbName     string `json:"database-name"`
	DbUname    string `json:"database-username"`
	DbPasswd   string `json:"database-password"`
}

func jsonToConfig(cj configJSON) (cc Config) {
	cc.GarbageCollectionTimeout = cj.GcTimeout
	cc.MaxThreads = cj.MaxThreads
	cc.Type = cj.Type

	cc.PrimaryNodeAddress = net.ParseIP(cj.PAddr)
	cc.PrimaryNodePort = cj.PPort

	cc.SecondaryNodeAddress = net.ParseIP(cj.SAddr)
	cc.SecondaryNodePort = cj.SPort
	cc.JobBinaryName = cj.SCmd

	cc.DatabaseAddress = net.ParseIP(cj.DbAddress)
	cc.DatabasePort = cj.DbPort
	cc.DatabaseName = cj.DbName
	cc.DatabaseUsername = cj.DbUname
	cc.DatabasePassword = cj.DbPasswd

	return
}
