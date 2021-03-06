package parser

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

const _dataFolder = "../../data"

func Test_can_parse_all_examples(t *testing.T) {
	for _, name := range []string{"a", "b", "c", "d", "e", "f"} {
		i := LoadInput(_dataFolder, "a")
		bytes, err := json.MarshalIndent(i, "", " ")
		if err != nil {
			t.FailNow()
		}
		err = ioutil.WriteFile(filepath.Join("../../data", name+".json"), bytes, 0775)
		if err != nil {
			t.FailNow()
		}
	}
}

