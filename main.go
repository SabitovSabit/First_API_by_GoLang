package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

type car struct{
	ID    string `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var cars= []car{
	{ID:"1",Make:"Toyota",Model:"Corolla",Price:40000},
	{ID:"2",Make:"Honda",Model:"Civic",Price:50000},
}

func getCars(c *gin.Context){
	c.IndentedJSON(http.StatusOK,cars)
}

func main(){
	router:=gin.Default()
	router.GET("/cars",getCars)
	router.Run("localhost:8080")
}

