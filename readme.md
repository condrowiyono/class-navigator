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

````
/class-navigator?q=7601
```` 

**Proses**

Data masukan akan diolah sedemikian sehingga menghasilkan hasil berupa informasi ruangan ( di gedung apa, lantai berapa, di jalan apa) dan juga hasil navigasi dengan Google Navigator API ( jika ada )

**Output**

Sample output
````json
{
   "Code": "7601",
   "Name": "Ruang Kuliah 7601",
   "Description": "Ruang kelas reguler untuk prodi IF",
   "Building": "Labtek V / Gd Benny Subianto",
   "Floor": 1,
   "Long": -6.890581130981445,
   "Lat": 107.60758972167969
}
````

## Instalation and Ussage

- Make sure you've installed Go Lang and PostgreSQL
- Clone the git
- Create new database on PostgreSQL
- Create this table

````
CREATE TABLE room ( id SERIAL PRIMARY KEY, code VARCHAR(10), name VARCHAR(50), description VARCHAR(100), building VARCHAR(50), floor INTEGER, long REAL, lat REAL);
````

- Edit the configure file (conf.json)

````
{
    "database": {
        "host": "localhost",
        "port": "5432",
        "user": "postgres",
        "password" : "root",
        "dbname" : "room_navigator"
    },
    "host": "localhost",
    "port": "8080"
}
````

- Run this

````
go build class-navigator
````

You can access on [http://localhost:8080](http://localhost:8080)

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
- Database : PostgreSQL / MariaDB
- Other Tools : StarUML, Terminal, some refreshement stuf (Spotify, TuxRacer,Frozen Bubble)
- Some Snack and Tea (hot Chocho is the best)

## Deployment Environment

- Maybe VPS
- Localy ( Go Web Server package, PostgreSQL )
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