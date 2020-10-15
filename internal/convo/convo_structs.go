package convo

import (
	"net"
)

// Config struct holds all the informations necessary to configure
// pluggabl apps (client, manager and worker)
type Config struct {
	Type                     string
	ManagerAddress           net.IP
	ManagerPort              int
	WorkerAddress            net.IP
	WorkerPort               int
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
	MAddr      string `json:"manager-address"`
	MPort      int    `json:"manager-port"`
	WAddr      string `json:"worker-address"`
	WPort      int    `json:"worker-port"`
	GcTimeout  int    `json:"garbage-collection-timeout"`
	MaxThreads int    `json:"max-threads"`
	DbAddress  string `json:"address"`
	DbPort     int    `json:"port"`
	DbName     string `json:"database"`
	DbUname    string `json:"username"`
	DbPasswd   string `json:"password"`
}

func jsonToConfig(cj configJSON) (cc Config) {
	cc.GarbageCollectionTimeout = cj.GcTimeout
	cc.MaxThreads = cj.MaxThreads
	cc.Type = cj.Type

	cc.ManagerAddress = net.ParseIP(cj.MAddr)
	cc.ManagerPort = cj.MPort

	cc.WorkerAddress = net.ParseIP(cj.WAddr)
	cc.WorkerPort = cj.WPort

	cc.DatabaseAddress = net.ParseIP(cj.DbAddress)
	cc.DatabasePort = cj.DbPort
	cc.DatabaseName = cj.DbName
	cc.DatabaseUsername = cj.DbUname
	cc.DatabasePassword = cj.DbPasswd

	return
}