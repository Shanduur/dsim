package convo

import (
	"errors"
	"fmt"
	"net"
	"testing"
)

func Equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestLoadConfiguration(t *testing.T) {
	matrix := [3]string{"client", "manager", "worker"}

	err := errors.New("")

	for _, m := range matrix {
		path := fmt.Sprintf("test/config_%v.json", m)

		conf, err := LoadConfiguration(path)
		if err != nil {
			t.Errorf("LoadConfiguration: while processing %v got error: %v", m, err)
		}

		managerIP := net.ParseIP("192.168.0.105")
		workerIP := net.ParseIP("192.168.0.106")

		if m != "client" {
			if conf.GarbageCollectionTimeout != 1000 {
				t.Errorf("GCT: got %v wanted %v", conf.GarbageCollectionTimeout, 1000)
			}

			if !Equal(conf.WorkerAddress, workerIP) {
				t.Errorf("%v WA: got %v wanted %v", m, conf.WorkerAddress, workerIP)
			}

			if conf.WorkerPort != 4004 {
				t.Errorf("%v WP: got %v wanted %v", m, conf.WorkerPort, 4004)
			}

			if conf.DatabaseName != "database" {
				t.Errorf("%v DN: got %v wanted %v", m, conf.DatabaseName, 4004)
			}

			if conf.DatabaseUsername != "admin" {
				t.Errorf("%v DU: got %v wanted %v", m, conf.DatabaseUsername, 4004)
			}

			if conf.DatabasePassword != "password" {
				t.Errorf("%v DP: got %v wanted %v", m, conf.DatabasePassword, 4004)
			}
		}

		if m == "worker" {
			if conf.MaxThreads != 4 {
				t.Errorf("%v MT: got %v wanted %v", m, conf.MaxThreads, 4)
			}
		}

		if !Equal(conf.ManagerAddress, managerIP) {
			t.Errorf("%v MA: got %v wanted %v", m, conf.ManagerAddress, managerIP)
		}

		if conf.ManagerPort != 4004 {
			t.Errorf("%v MP: got %v wanted %v", m, conf.WorkerPort, 4004)
		}
	}

	_, err = LoadConfiguration("test/unmarshalable_json.json")
	if err == nil {
		t.Errorf("Wanted error")
	}

	_, err = LoadConfiguration("test/no_such_file.json")
	if err == nil {
		t.Errorf("Wanted error")
	}
}
