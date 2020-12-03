package main

import(
	"github.com/kataras/iris/v12"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
  "fmt"
	"time"
  //"reflect"
)

type Posts struct{
	ID int
  Post string `json:"post"`
  Created string `json:"created" sql:"not null;type date"`
}

type any interface{}

func main(){

	app := iris.New()

	app.Get("/", func(ctx iris.Context){
		ctx.ServeFile("./index.html")
	})

	app.Get("/login", func(ctx iris.Context){
		ctx.ServeFile("./index.html")
	})

  app.Get("/logout", func(ctx iris.Context){
		ctx.ServeFile("./index.html")
	})

	app.Get("/signup", func(ctx iris.Context){
		ctx.ServeFile("./index.html")
	})

	app.Get("/data", func(ctx iris.Context){
    _, data := findData()
		ctx.JSON(iris.Map{"data": data})
	})

  app.Post("/",func(ctx iris.Context){
    //addData("問題なく追加できるようです")
 	})

	app.Post("/add", func(ctx iris.Context){
		todoitem := ctx.PostValue("textbox")

		if todoitem != "" {
			err := addData(todoitem)
      if err != nil {
				ctx.JSON(iris.Map{"message":"post error."})
      }else{
				ctx.JSON(iris.Map{"message":"success.", "post": todoitem})
      }
		}else{
			ctx.JSON(iris.Map{"message":"not a post."})
		}
	})

	app.Run(iris.Addr(":5000"))
	
  

/*
	addData("問題なく追加できるようです")
  findData()
*/
}

func findData() (error, []*Posts) {
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}else{
		fmt.Println("DB接続成功")
	}
	defer db.Close()

  result := []*Posts{}
	error := db.Limit(10).Find(&result).Error
	if error != nil || len(result) == 0 {
		
	}

 return error, result
}

func addData(post string) error{
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}else{
		fmt.Println("DB接続成功")
	}
	defer db.Close()

  error := db.Create(&Posts{
		Post: post,
    Created: getDate(),
	}).Error

	return error
}

func deleteData(){

}

func sqlConnect() (database *gorm.DB, err error){
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "go_sample"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

func getDate() string{
	const layout = "2006-01-02 15:04:05"
	now := time.Now()
	return now.Format(layout)
}