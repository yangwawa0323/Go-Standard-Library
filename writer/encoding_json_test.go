package writer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Classmate struct {
	Students []Student `json:"students"`
}

func Test_Json_Write(t *testing.T) {
	var class_301 = Classmate{
		[]Student{
			{"Yangwawa", 47},
			{"WangSheng", 42},
		},
	}

	var buf = new(bytes.Buffer)
	json_encoder := json.NewEncoder(buf)
	if err := json_encoder.Encode(class_301); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", buf.String())
	buf.ReadString('\n')
	fmt.Printf("%s\n", buf.String())

}
