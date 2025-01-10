package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Huselt struct {
	Neg   int `json:"neg"`
	Hoyor int `json:"hoyor"`
}

func calcNiilber(c *gin.Context) {
	huselt := Huselt{}
	c.ShouldBindJSON(&huselt)
	c.JSON(http.StatusOK, gin.H{"niilber": huselt.Neg + huselt.Hoyor})
}

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "1234"
	dbname   = "db_tamir"
)

type User struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func createUser(c *gin.Context) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password= %s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error")
		panic(err)
	}
	defer db.Close()

	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	query := fmt.Sprintf("INSERT INTO t_user (\"Age\", \"Name\", \"Pass\") VALUES (%d,'%s','%s')", newUser.Age, newUser.Name, newUser.Pass)
	_, err = db.Exec(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func main() {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password= %s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	fmt.Println("error")
	// 	panic(err)
	// }
	// defer db.Close()

	// rows, err := db.Query("SELECT * FROM t_user")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var users []User
	// for rows.Next() {
	// 	var singleUser User
	// 	err := rows.Scan(&singleUser.Id, &singleUser.Age, &singleUser.Name, &singleUser.Pass)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	users = append(users, singleUser)

	// }

	// rows.Close()
	// db.Close()
	// fmt.Println(users)
	router := gin.Default()
	router.GET("/niilber", calcNiilber)
	router.POST("/user", createUser)
	log.Fatal(router.Run(":3456"))

}
