#Class Navigator - Web Service Project

##Deskripsi

Class Navigator adalah sebuah web service yang akan memerikan petunjuk ruangan dari sebuah ruang kelas di ITB. Penggunaan web service untuk memberikan informasi dan petunjuk jalan dari lokasi sekarang ke lokasi kelas tujuan

## Batasan

- Ruang kelas umum 

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
  "status": "OK",
  "geocoded_waypoints" : [
     {
        "geocoder_status" : "OK",
        "place_id" : "ChIJ7cv00DwsDogRAMDACa2m4K8",
        "types" : [ "locality", "political" ]
     },
     {
        "geocoder_status" : "OK",
        "place_id" : "ChIJ69Pk6jdlyIcRDqM1KDY3Fpg",
        "types" : [ "locality", "political" ]
     },
     {
        "geocoder_status" : "OK",
        "place_id" : "ChIJgdL4flSKrYcRnTpP0XQSojM",
        "types" : [ "locality", "political" ]
     },
     {
        "geocoder_status" : "OK",
        "place_id" : "ChIJE9on3F3HwoAR9AhGJW_fL-I",
        "types" : [ "locality", "political" ]
     }
  ],
  "routes": [ {
    "summary": "I-40 W",
    "legs": [ {
      "steps": [ {
        "travel_mode": "DRIVING",
        "start_location": {
          "lat": 41.8507300,
          "lng": -87.6512600
        },
        "end_location": {
          "lat": 41.8525800,
          "lng": -87.6514100
        },
        "polyline": {
          "points": "a~l~Fjk~uOwHJy@P"
        },
        "duration": {
          "value": 19,
          "text": "1 min"
        },
        "html_instructions": "Head \u003cb\u003enorth\u003c/b\u003e on \u003cb\u003eS Morgan St\u003c/b\u003e toward \u003cb\u003eW Cermak Rd\u003c/b\u003e",
        "distance": {
          "value": 207,
          "text": "0.1 mi"
        }
      },
      ...
      ... additional steps of this leg
    ...
    ... additional legs of this route
      "duration": {
        "value": 74384,
        "text": "20 hours 40 mins"
      },
      "distance": {
        "value": 2137146,
        "text": "1,328 mi"
      },
      "start_location": {
        "lat": 35.4675602,
        "lng": -97.5164276
      },
      "end_location": {
        "lat": 34.0522342,
        "lng": -118.2436849
      },
      "start_address": "Oklahoma City, OK, USA",
      "end_address": "Los Angeles, CA, USA"
    } ],
    "copyrights": "Map data ©2010 Google, Sanborn",
    "overview_polyline": {
      "points": "a~l~Fjk~uOnzh@vlbBtc~@tsE`vnApw{A`dw@~w\\|tNtqf@l{Yd_Fblh@rxo@b}@xxSfytAblk@xxaBeJxlcBb~t@zbh@jc|Bx}C`rv@rw|@rlhA~dVzeo@vrSnc}Axf]fjz@xfFbw~@dz{A~d{A|zOxbrBbdUvpo@`cFp~xBc`Hk@nurDznmFfwMbwz@bbl@lq~@loPpxq@bw_@v|{CbtY~jGqeMb{iF|n\\~mbDzeVh_Wr|Efc\\x`Ij{kE}mAb~uF{cNd}xBjp]fulBiwJpgg@|kHntyArpb@bijCk_Kv~eGyqTj_|@`uV`k|DcsNdwxAott@r}q@_gc@nu`CnvHx`k@dse@j|p@zpiAp|gEicy@`omFvaErfo@igQxnlApqGze~AsyRzrjAb__@ftyB}pIlo_BflmA~yQftNboWzoAlzp@mz`@|}_@fda@jakEitAn{fB_a]lexClshBtmqAdmY_hLxiZd~XtaBndgC"
    },
    "warnings": [ ],
    "waypoint_order": [ 0, 1 ],
    "bounds": {
      "southwest": {
        "lat": 34.0523600,
        "lng": -118.2435600
      },
      "northeast": {
        "lat": 41.8781100,
        "lng": -87.6297900
      }
    }
  } ]
}
	}
````

## Instalation and Ussage

Make sure you've installed Go Lang and MySQL
````
go build
````

You can access on [http://localhost:8181](http://localhost:8181)

## Use Case dan Modeling

In Progress

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
- Some Snack and Tea

## Deployment Environment

- Maybe VPS
- Localy ( Go Web Server package, mySQL )
- cURL

## Future Development

Some futute project you may concern
- an Android App to implements this service
- an Web-Base App

## Licence 

* Educational purpose only
* Google Maps, Go are copyrighted by Google

Created with ❤ for better grade 
Integrative Programming Class 2017