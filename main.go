package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	// import from local package
	"github.com/paul-tong/buzzer_beater/controller"
	"github.com/paul-tong/buzzer_beater/model"
)

func main() {
	// setup  router
	controller.SetupRouter()

	// connect to database and setup db instance
	db := model.ConnectToDB()
	defer db.Close() // this will excute at end of main function
	model.SetDB(db)

	// create table and add data
	// model.CreatePostTable()
	/*post1 := model.PostTest{Title: "morning", Body: "Good morning!"}
	post2 := model.PostTest{Title: "noon", Body: "Good noon!"}
	post3 := model.PostTest{Title: "evening", Body: "Good evening!"}
	model.AddPost(post1)
	model.AddPost(post2)
	model.AddPost(post3)*/

	// start listening to the port
	http.ListenAndServe(":8888", nil)
}
