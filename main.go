package main

import (
	_ "fmt"
	_ "log"
	"net/http"
	_ "os"

	"github.com/LINK/router"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, KyotoHack Team F! :)")
	router := router.GetRouter()
	router.Run(":8000")
}

func main() {
	//http.HandleFunc("/", handler)
	//http.HandleFunc("/alive", handler)
	////http.HandleFunc("/create", create_handler)
	//if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
	//	log.Fatal(err)
	//}

	r := router.GetRouter()
	r.Run(":3000")
	/*

	   var db = NewDBConn()

	   func NewDBConn() *gorm.DB {
	   	db, err := gorm.Open("mysql", "root:password@tcp(kyotohack2018-f-1.cm68cjyvpekc.ap-northeast-1.rds.amazonaws.com)/kyotohack2018_f?charset=utf8&parseTime=True&loc=Local")
	   	if err != nil {
	   		panic(err)
	   	}

	   	return db
	   }

	   func GetDBConn() *gorm.DB {
	   	return db
	   }

	   type Category struct {
	   	Id   uint
	   	Name string
	   }

	   func create_handler(w http.ResponseWriter, r *http.Request) {
	   	category := Category{Id: 1, Name: "go"}
	   	db.Create(&category)
	   }
	*/

}
