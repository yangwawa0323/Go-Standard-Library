package yamltest

import (
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func Test_YAML_Read(t *testing.T) {
	yfile, err := ioutil.ReadFile("./items.yaml")
	if err != nil {
		t.Fatal(err)
	}

	data := make(map[interface{}]interface{})

	err2 := yaml.Unmarshal(yfile, &data)

	if err2 != nil {
		t.Fatal(err2)
	}

	for k, v := range data {
		t.Logf("%s -> %d \n", k, v)
	}
}

type User struct {
	Name       string
	Occupation string
}

func Test_YAML_Read_Into_Struct(t *testing.T) {
	yfile, err := ioutil.ReadFile("./users.yaml")
	if err != nil {
		t.Fatal(err)
	}

	data := make(map[string]User)

	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range data {
		t.Logf("%s: %#v\n", k, v)
	}
}

func Test_YAML_Write(t *testing.T) {
	user1 := &User{
		"Yangwawa", "instructor",
	}
	user2 := &User{
		"John doe", "singer",
	}

	const outfile = "./output_userlist.yaml"

	userlist := []*User{user1, user2}

	data, err := yaml.Marshal(userlist)

	if err != nil {
		t.Fatal(err)
	}

	if err := ioutil.WriteFile(outfile, data, 0); err != nil {
		t.Fatal(err)
	}

	t.Logf("YAML %s has been write successfully", outfile)
}
