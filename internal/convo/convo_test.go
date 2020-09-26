package convo

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	matrix := [3]string{"client", "manager", "worker"}

	root, err := os.Getwd()
	if err != nil {
		t.Errorf("error while geting path: %v", err)
	}

	for _, m := range matrix {
		path := fmt.Sprintf("%v/../../config/config_%v.json", root, m)

		err = LoadConfiguration(path)
		if err != nil {
			t.Errorf("LoadConfiguration: while processing %v got error: %v", m, err)
		}
	}

	err = LoadConfiguration("test/unmarshalable_json.json")
	if err == nil {
		t.Errorf("Wanted error")
	}

	err = LoadConfiguration("test/no_such_file.json")
	if err == nil {
		t.Errorf("Wanted error")
	}
}
