package main
import (
  "encoding/json"
  "os"
  "fmt"
)

type Config struct {
    Database struct {
        Host     string `json:"host"`
        Port   string `json:"port"`
        User     string `json:"user"`
        Password string `json:"password"`
        DBName   string `json:"dbname"`
    } `json:"database"`
    Host string `json:"host"`
    Port string `json:"port"`
}

func LoadConfiguration() (Config,error) {
    var config Config
    configFile, err := os.Open("conf.json")
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config,nil
}