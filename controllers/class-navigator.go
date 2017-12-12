package controllers

import (
	"encoding/json"
	"net/http"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"class-navigator/models"
	"strconv" 
	"io"
	"fmt"
)


func CheckErr(err error) {
  	if err != nil {
    	panic(err)
  	}
}

func GetAllClass(w http.ResponseWriter, r *http.Request, psqlInfo string ) {
  	w.Header().Set("Content-Type", "application/json")
  	db, err := sql.Open("postgres", psqlInfo)
  	CheckErr(err)
  	defer db.Close()


  	rows, err := db.Query(`SELECT id, code, name, building, floor FROM room `)
	CheckErr(err)
	defer db.Close()
  
  	var room_index []models.RoomIndex
  	
  	for rows.Next() {
        var code,name,building string
        var id, floor int

        rows.Scan(&id, &code, &name, &building, &floor)
		room_index = append(room_index, models.RoomIndex{id,code,name,building,floor} )
	}

	roomBytes, _ := json.Marshal(&room_index)
	w.Write(roomBytes)
	db.Close()
}

func GetClassByID(w http.ResponseWriter, r *http.Request, psqlInfo string, q string ) {
	w.Header().Set("Content-Type", "application/json")
	myid, _ := strconv.Atoi(q)
	db, err := sql.Open("postgres", psqlInfo)
	CheckErr(err)
	defer db.Close()

  	rows, err := db.Query(`SELECT code, name,  description, building, floor, long, lat FROM room  WHERE id = $1`,myid)
	CheckErr(err)
	defer db.Close()
  	
  	var room_navigator []models.RoomNavigator
  	for rows.Next() {
  		var code,name, description, building string
        var floor int
        var long, lat float64
        rows.Scan(&code, &name, &description, &building,&floor, &long, &lat)
		
		room_navigator = append(room_navigator, models.RoomNavigator{code,name,description,building,floor,long,lat})
	}
	roomBytes,_ := json.Marshal(&room_navigator)
	w.Write(roomBytes)
	db.Close()
}

func GetClass(w http.ResponseWriter, r *http.Request, psqlInfo string, q string ) {
	w.Header().Set("Content-Type", "application/json")
  	db, err := sql.Open("postgres", psqlInfo)
  	CheckErr(err)
  	defer db.Close()

  	rows, err := db.Query(`SELECT  code,  name,  description,  building, floor, long, lat FROM room WHERE name like $1`,q)
	CheckErr(err)
	defer db.Close()
  
	var room_navigator []models.RoomNavigator
  	for rows.Next() {
  		var code,name, description, building string
        var floor int
        var long, lat float64
        rows.Scan(&code, &name, &description, &building,&floor, &long, &lat)
		
		room_navigator = append(room_navigator, models.RoomNavigator{code,name,description,building,floor,long,lat})
	}
	roomBytes,_ := json.Marshal(&room_navigator)
	w.Write(roomBytes)
	db.Close()
}

//POST
func PostNewRoom(w http.ResponseWriter, r *http.Request, psqlInfo string) {
	out := make([] byte, 1024)

	bodyLen, err := r.Body.Read(out)

	if err != io.EOF {
		fmt.Printf(err.Error())
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	} 
	var k models.RoomNavigator
	err = json.Unmarshal(out[:bodyLen],&k)
	if err != nil {
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}

	idx := insertInDatabase(k,psqlInfo)
	log.Printf(" affect %d",idx)
	w.Write([]byte(`{"error":"success"}`))
}

func insertInDatabase(data models.RoomNavigator, psqlInfo string) int64 {
	db, err := sql.Open("postgres", psqlInfo)
  	CheckErr(err)
  	defer db.Close()

	result, err := db.Exec(`INSERT INTO room (code,name,description,building,floor, long, lat) VALUES ($1,$2,$3,$4,$5,$6,$7)`, 
		data.Code, data.Name, data.Description, data.Building, data.Floor, data.Long, data.Lat)
	CheckErr(err)
	affect,_ := result.RowsAffected(); 
	return affect
}

//PATCH
func PatchRoom(w http.ResponseWriter, r *http.Request, psqlInfo string, p string) {
	out := make([] byte, 1024)

	bodyLen, err := r.Body.Read(out)

	if err != io.EOF {
		fmt.Printf(err.Error())
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	} 
	var k models.RoomNavigator
	err = json.Unmarshal(out[:bodyLen],&k)
	if err != nil {
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}

	idx := patchInDatabase(k,psqlInfo,p)
	log.Printf(" affect %d",idx)
	w.Write([]byte(`{"error":"success"}`))
}

func patchInDatabase(data models.RoomNavigator, psqlInfo string, p string) int64 {
	myid, _ := strconv.Atoi(p)
	db, err := sql.Open("postgres", psqlInfo)
  	CheckErr(err)
  	defer db.Close()

	result, err := db.Exec(`UPDATE room SET code=$2, name=$3, description =$4, building = $5, floor=$6, long=$7, lat=$8 WHERE id=$1`, 
		myid, data.Code, data.Name, data.Description, data.Building, data.Floor, data.Long, data.Lat)
	
	CheckErr(err)
	affect,_ := result.RowsAffected(); 
	return affect
}

//DELETE
func DeleteRoom(w http.ResponseWriter, r *http.Request, psqlInfo string, p string) {
	myid, _ := strconv.Atoi(p)
	db, err := sql.Open("postgres", psqlInfo)
  	CheckErr(err)
  	defer db.Close()

	result, err := db.Exec(` DELETE FROM room WHERE id=$1`,myid)
	CheckErr(err)
	affect, err := result.RowsAffected()
    CheckErr(err)
    fmt.Println(affect)

}

