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
	matrix := []string{"client", "primary", "secondary", "db"}

	err := errors.New("")

	for _, m := range matrix {
		path := fmt.Sprintf("test/config_%v.json", m)

		conf, err := LoadConfiguration(path)
		if err != nil {
			t.Errorf("LoadConfiguration: while processing %v got error: %v", m, err)
		}

		fmt.Printf("%v", conf.String())

		if m == "db" {
			if conf.Port != 5432 {
				t.Errorf("%v got %v wanted %v", m, conf.Port, 5432)
			}

			if conf.DatabaseName != "database-name" {
				t.Errorf("%v Dname: got %v wanted %v", m, conf.DatabaseName, 4004)
			}

			if conf.DatabaseUsername != "admin" {
				t.Errorf("%v Uname: got %v wanted %v", m, conf.DatabaseUsername, 4004)
			}

			if conf.DatabasePassword != "database-password" {
				t.Errorf("%v Pass: got %v wanted %v", m, conf.DatabasePassword, 4004)
			}
		} else if m != "client" {
			if conf.GarbageCollectionTimeout != 1000 {
				t.Errorf("GCT: got %v wanted %v", conf.GarbageCollectionTimeout, 1000)
			}
		} else {
			testIP := net.ParseIP("192.168.0.105")

			if !Equal(conf.Address, testIP) {
				t.Errorf("%v got %v wanted %v", m, conf.Address, testIP)
			}

			if conf.Port != 4004 {
				t.Errorf("%v got %v wanted %v", m, conf.Port, 4004)
			}

			if m == "secondary" {
				if conf.MaxThreads != 4 {
					t.Errorf("%v MT: got %v wanted %v", m, conf.MaxThreads, 4)
				}

				if conf.ExternalPort != 4010 {
					t.Errorf("%v got %v wanted %v", m, conf.ExternalPort, 4010)
				}
			} else {
				if conf.ExternalPort != conf.Port {
					t.Errorf("%v got %v wanted %v", m, conf.ExternalPort, conf.Port)
				}
			}
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
