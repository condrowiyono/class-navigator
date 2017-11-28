package main
import (
  "encoding/json"
  "os"
  "net/http"
  "fmt"
  "log"
)

type Config struct {
    Database struct {
        Host     string `json:"host"`
        Password string `json:"password"`
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

func main() {
	config,err := LoadConfiguration()
	if err != nil {
    	log.Fatal(err)
  	}
	addr := config.Port
	http.HandleFunc("/", hello)
	log.Printf("Server starting on port %v\n", addr)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",addr),nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello World")
}
