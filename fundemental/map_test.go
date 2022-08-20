package fundemental

import (
	"testing"
)

func Test_Map(t *testing.T) {
	var group = make(map[string]map[string][]string)
	group["Joe"] = map[string][]string{
		"address": {"joe_doe@163.com", "joe@qq.com"},
	}

	group["Jake"] = map[string][]string{
		"address": {"jake_01@163.com", "jake_02@163.com", "jake_01@qq.com"},
	}

	for name, info := range group {
		for _, address := range info["address"] {
			t.Logf("%s has address: %s \n", name, address)
		}

	}

	// Get not exists key from the map
	t.Logf("get Yangkun info: %s \n", group["Yangkun"])
}
