package files

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func WriteFile(path string, name string, data interface{}) {
	file, _ := json.MarshalIndent(data, "", "")
	e := ioutil.WriteFile("./files/"+path+"/"+name, file, 0700)

	if e != nil {
		fmt.Println(e)
	}
}

func ReadFile(structure interface{}, path string, name string) {
	bytes, e := ioutil.ReadFile("./files/" + path + "/" + name)

	if e != nil {
		fmt.Println(e)
	}

	json.Unmarshal(bytes, structure)
}

func ReadAllFiles(path string, getStructure func(file os.FileInfo) interface{}) {
	files, _ := ioutil.ReadDir("./files/" + path)
	for _, f := range files {
		structure := getStructure(f)
		ReadFile(structure, path, f.Name())
	}
}
