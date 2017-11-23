# Class Navigator - Web Service Project

## Deskripsi

Class Navigator adalah sebuah web service yang akan memerikan petunjuk ruangan dari sebuah ruang kelas di ITB. Penggunaan web service untuk memberikan informasi dan petunjuk jalan dari lokasi sekarang ke lokasi kelas tujuan

## Batasan

- Ruang kelas umum 
- Ruang umum

## Human

Condro Wiyono - 18215042

## Sample Input and Output 

**Input**

Sample input : 

````json
[
	{
		"curr_long": "107.60945",
		"curr_lat": "-6.89065",
		"classroom": "7601",
	}
]
```` 

**Proses**

Data masukan akan diolah sedemikian sehingga menghasilkan hasil berupa informasi ruangan ( di gedung apa, lantai berapa, di jalan apa) dan juga hasil navigasi dengan Google Navigator API ( jika ada )

**Output**

Sample output
````json
	{
		"curr_long": "107.60945",
		"curr_lat": "-6.89065",
		"classroom": { 
			"namaruang" : "7601", 
			"namagedung" : "Labtek V",
			"namalain" : "Gedung Benny Subianto",
			"lantai": 1,
			"class_lat" : "107.60945",
			"class_long" :  "-6.89065", 	
				   },
		"navigation": 
		{
   "geocoded_waypoints" : [
      {
         "geocoder_status" : "OK",
         "place_id" : "ChIJgYxZXlfmaC4R3mmubrgd0KQ",
         "types" : [ "route" ]
      },
      {
         "geocoder_status" : "OK",
         "place_id" : "ChIJO08GFljmaC4RumV5itwJCgo",
         "types" : [ "route" ]
      }
   ],
   "routes" : [
      {
         "bounds" : {
            "northeast" : {
               "lat" : -6.8881308,
               "lng" : 107.6099633
            },
            "southwest" : {
               "lat" : -6.8910557,
               "lng" : 107.6082303
            }
         },
         "copyrights" : "Data peta ©2017 Google",
         "legs" : [
            {
               "distance" : {
                  "text" : "0,6 km",
                  "value" : 612
               },
               "duration" : {
                  "text" : "8 menit",
                  "value" : 458
               },
               "end_address" : "\"Jalan D ITB, Lb. Siliwangi, Coblong, Kota Bandung, Jawa Barat 40132, Indonesia",
               "end_location" : {
                  "lat" : -6.8881308,
                  "lng" : 107.6083307
               },
               "start_address" : "Jl. II, Lb. Siliwangi, Coblong, Kota Bandung, Jawa Barat 40132, Indonesia",
               "start_location" : {
                  "lat" : -6.8908204,
                  "lng" : 107.6099633
               },
               "steps" : [
                  {
                     "distance" : {
                        "text" : "75 m",
                        "value" : 75
                     },
                     "duration" : {
                        "text" : "1 menit",
                        "value" : 53
                     },
                     "end_location" : {
                        "lat" : -6.891032,
                        "lng" : 107.609463
                     },
                     "html_instructions" : "Pergilah ke \u003cb\u003ebarat\u003c/b\u003e menuju ke \u003cb\u003eJl. II\u003c/b\u003e",
                     "polyline" : {
                        "points" : "rz`i@gqxoS@xA@B@@@@@@@?@?Z?"
                     },
                     "start_location" : {
                        "lat" : -6.8908204,
                        "lng" : 107.6099633
                     },
                     "travel_mode" : "WALKING"
                  },
                  {
                     "distance" : {
                        "text" : "84 m",
                        "value" : 84
                     },
                     "duration" : {
                        "text" : "1 menit",
                        "value" : 64
                     },
                     "end_location" : {
                        "lat" : -6.8910557,
                        "lng" : 107.6087048
                     },
                     "html_instructions" : "Belok \u003cb\u003ekanan\u003c/b\u003e ke \u003cb\u003eJl. II\u003c/b\u003e",
                     "maneuver" : "turn-right",
                     "polyline" : {
                        "points" : "|{`i@cnxoSBtB?\\@B"
                     },
                     "start_location" : {
                        "lat" : -6.891032,
                        "lng" : 107.609463
                     },
                     "travel_mode" : "WALKING"
                  },
                  {
                     "distance" : {
                        "text" : "0,1 km",
                        "value" : 137
                     },
                     "duration" : {
                        "text" : "2 menit",
                        "value" : 104
                     },
                     "end_location" : {
                        "lat" : -6.890138299999999,
                        "lng" : 107.6090622
                     },
                     "html_instructions" : "Belok \u003cb\u003ekanan\u003c/b\u003e ke \u003cb\u003eJl. C ITB\u003c/b\u003e",
                     "maneuver" : "turn-right",
                     "polyline" : {
                        "points" : "b|`i@kixoSmA??KCo@ACACCACAEAE?Y?[BE?QA"
                     },
                     "start_location" : {
                        "lat" : -6.8910557,
                        "lng" : 107.6087048
                     },
                     "travel_mode" : "WALKING"
                  },
                  {
                     "distance" : {
                        "text" : "92 m",
                        "value" : 92
                     },
                     "duration" : {
                        "text" : "1 menit",
                        "value" : 65
                     },
                     "end_location" : {
                        "lat" : -6.8901425,
                        "lng" : 107.6082303
                     },
                     "html_instructions" : "Belok \u003cb\u003ekiri\u003c/b\u003e menuju \u003cb\u003e\"Jalan D ITB\u003c/b\u003e",
                     "maneuver" : "turn-left",
                     "polyline" : {
                        "points" : "jv`i@skxoS?@AB@xA?dA"
                     },
                     "start_location" : {
                        "lat" : -6.890138299999999,
                        "lng" : 107.6090622
                     },
                     "travel_mode" : "WALKING"
                  },
                  {
                     "distance" : {
                        "text" : "0,2 km",
                        "value" : 224
                     },
                     "duration" : {
                        "text" : "3 menit",
                        "value" : 172
                     },
                     "end_location" : {
                        "lat" : -6.8881308,
                        "lng" : 107.6083307
                     },
                     "html_instructions" : "Belok \u003cb\u003ekanan\u003c/b\u003e ke \u003cb\u003e\"Jalan D ITB\u003c/b\u003e\u003cdiv style=\"font-size:0.9em\"\u003eTujuan ada di sebelah kanan.\u003c/div\u003e",
                     "maneuver" : "turn-right",
                     "polyline" : {
                        "points" : "jv`i@mfxoS}AA{AAeBAuAEWECA"
                     },
                     "start_location" : {
                        "lat" : -6.8901425,
                        "lng" : 107.6082303
                     },
                     "travel_mode" : "WALKING"
                  }
               ],
               "traffic_speed_entry" : [],
               "via_waypoint" : []
            }
         ],
         "overview_polyline" : {
            "points" : "rz`i@gqxoSB|ABB`@@BrC@BmA?C{@CGSEmA@?dDyDC{DG[G"
         },
         "summary" : "\"Jalan D ITB",
         "warnings" : [
            "Petunjuk arah berjalan kaki masih versi beta. Waspadalah – Mungkin ada trotoar atau jalan setapak yang tidak tampak dalam rute ini."
         ],
         "waypoint_order" : []
      }
   ],
   "status" : "OK"
}
````

## Instalation and Ussage

- Make sure you've installed Go Lang and MySQL
- Clone the git
- Run this
````
go build 
````

You can access on [http://localhost:8181](http://localhost:8181)

## Use Case dan Modeling

See here :
(https://github.com/condrowiyono/class-navigator/tree/master/modelling)

## Requirements

**Functional Requirements**

- Sistem dapat memberikan informasi tentang sebuah ruang kelas
- Sistem dapat memberikan petunjuk arah dari lokasi saat ini dengan ruang kelas yang dituju
- Sistem dapat menerima masukan berupa MIME application/json
- Sistem dapat menyediakan luaran bkerupa JSON yang dapat digunakan oleh sebuah aplikasi untuk menampilkan informasi dan membuat petunjuk arah melalui Google Maps
- Sistem dapat menyimpan data log berupa lokasi sekarang dan kelas 
- Sistem dapat menerima masukan baru data ruangan (*permission needed*)

**Non-Functional Requirements**

- Sistem dapat menerima penggunaan secara bersamaan untuk 5 - 10 pengguna
- Sistem memiliki penyimpanan yang cukup untuk data 1 bulan
- Respon time untuk pencarian maksimal 20ms, dan 1s untuk memperoleh data dari Google Maps
- Up-time web service adalah 24/7, dengan toleransi sistem mati sebesar 5%
- Kesalahan sistem yang ditoleransi adalah 3 %

## Technology Used for Implementation

- Linux-based Operating System
- Server (VPS, or local server, TBA)
- REST Client ( cURL, Browser, dll)
- Go Lang
- Google Maps Direction API

## Development Environment

- Dell Inspiron 3458 with Core i5 5200U
- 4 GB of RAM
- 500 GB Storage
- OS : Lubuntu 16.04 Long Term Service
- Editor : Sublime
- Compiler : GoLang
- Web Server : Golang based net package
- Database : MySQL / MariaDB
- Other Tools : StarUML, Terminal, some refreshement stuf (Spotify, TuxRacer,Frozen Bubble)
- Some Snack and Tea (hot Chocho is the best)

## Deployment Environment

- Maybe VPS
- Localy ( Go Web Server package, mySQL )
- cURL

## Future Development

Some futute project you may concern
- an Android App to implements this service
- a Web-Base App

## Licence 

* Educational purpose only
* Google Maps, Go are copyrighted by Google

Created with ❤ for better grade 
Integrative Programming Class 2017