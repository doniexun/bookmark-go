// https://www.thepolyglotdeveloper.com/2017/04/load-json-configuration-file-golang-application/
package common

import "os"
import "fmt"
import "encoding/json"
import "io/ioutil"

type Config struct {
	Name     string `json:"name"`
	Database struct {
		Host     string `json:"host"`
		Password string `json:"password"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadConfig(filepath string) Config {
    // data 为 byte 类型
    data, err := ioutil.ReadFile(filepath)
    if err != nil {
        fmt.Printf("Program stopping with error %v", err)
        os.Exit(1)
    }

    var config Config
    json.Unmarshal(data, &config)
    fmt.Println(config)
    return config
}