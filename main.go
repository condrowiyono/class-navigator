package main
import (
  "net/http"
  "fmt"
  "log"
  "strconv"
  "class-navigator/views"
  "class-navigator/controllers"
)


var index *views.View
var psqlInfo string


func main() {
	//LOADING CONFIGURATION
	config,err := LoadConfiguration()
	if err != nil { log.Fatal(err) }
  	intport,err :=strconv.Atoi(config.Database.Port)
	if err != nil { log.Fatal(err) }
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			    "password=%s dbname=%s sslmode=disable",
			    config.Database.Host, intport, 
			    config.Database.User, 
			    config.Database.Password, 
			    config.Database.DBName)
	
	//CONFIGURE THE PORT
	addr := config.Port

	//ROUTER
	http.HandleFunc("/", indexHandler)
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
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",addr),nil))
}


func indexHandler(w http.ResponseWriter, r *http.Request) {  
	index = views.NewView("bootstrap", "views/index.gohtml" )
  	index.Render(w, nil)
}

