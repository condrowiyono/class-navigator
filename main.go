package main
import (
  "net/http"
  "fmt"
  "log"
  "strconv"
  "github.com/lib/pq"
  "class-navigator/views"
  "class-navigator/controllers"
  "os"
)


var index *views.View
var documentation *views.View
var about *views.View
var psqlInfo string
var config Config

func Init() {
	config,_ = LoadConfiguration()
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Port
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}

func GetDBURI() string {
	var connection string;
	intport,_:=strconv.Atoi(config.Database.Port)
	url := os.Getenv("DATABASE_URL")
	if url !="" {
		connection, _ = pq.ParseURL(url)
    	connection += " sslmode=require"
	} else {
		connection =  fmt.Sprintf("host=%s port=%d user=%s "+
			    "password=%s dbname=%s sslmode=disable",
			    config.Database.Host, intport, 
			    config.Database.User, 
			    config.Database.Password, 
			    config.Database.DBName)
		log.Println("[-] No DATABASE_URL environment variable detected. Setting to local ")
	}
	return connection;


}

func main() {
	//LOADING CONFIGURATION
	Init()
  	
	//CONFIGURE DB
	psqlInfo = GetDBURI()
	
	//CONFIGURE THE PORT
	addr := GetPort() 

	//ROUTER
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/documentation", docHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/class-navigator/", func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case "GET" :
			s := r.URL.Path[len("/class-navigator/"):]
			if s != "" {
				controllers.GetClassByID(w,r,psqlInfo,s)
			} else {
				controllers.GetAllClass(w,r,psqlInfo)
			}
		case "POST" : controllers.PostNewRoom(w,r,psqlInfo)
		case "PATCH" :
			p := r.URL.Path[len("/class-navigator/"):]
			if p != "" {
				controllers.PatchRoom(w,r,psqlInfo,p)
			}				
		case "DELETE" :
			d := r.URL.Path[len("/class-navigator/"):]
			if d != "" {
				controllers.DeleteRoom(w,r,psqlInfo,d)
			} 				
		}
	})
	http.HandleFunc("/class-navigator", func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case "GET" :
			var q string
			q = fmt.Sprint("%"+ r.FormValue("q") + "%")
			controllers.GetClass(w,r,psqlInfo,q)		
		}
	})

	log.Printf("Server starting on port %v\n", addr)
	log.Fatal(http.ListenAndServe(addr,nil))
}


func indexHandler(w http.ResponseWriter, r *http.Request) {  
	index = views.NewView("bootstrap", "views/index.gohtml" )
  	index.Render(w, nil)
}
func docHandler(w http.ResponseWriter, r *http.Request) {  
	documentation = views.NewView("bootstrap", "views/documentation.gohtml" )
  	documentation.Render(w, nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {  
	about = views.NewView("bootstrap", "views/about.gohtml" )
  	about.Render(w, nil)
}
