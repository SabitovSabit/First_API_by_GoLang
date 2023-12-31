package main

import (
	"errors"
	"firstApi/database"
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

func carById(c *gin.Context){
	id:=c.Param("id")

	scar,err:=getCarById(id)

	if err!=nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"not found!"})
		return 
	}

	c.IndentedJSON(http.StatusOK,scar)
}
func getCarById(id string) (*car,error){
	
	for i,item :=range cars{
		 if item.ID==id{
			return  &cars[i],nil
		 }
	}

	return nil,errors.New("car not found")
}

func createCar(c *gin.Context){

	var newCar car

	if err:=c.BindJSON(&newCar);err!=nil{
		return
	}

	cars=append(cars,newCar)
	c.IndentedJSON(http.StatusCreated,newCar)
}
func deleteById(c *gin.Context){
	id,ok:=c.GetQuery("id")
    var scar []car
	
	if ok{
		scar=DeleteCarById(id)
	}else{
       c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"missing parametr"})
	   return
	}

	c.IndentedJSON(http.StatusOK,scar)
}
func DeleteCarById(id string) []car{
	var newcars []car

	for _,item:=range cars{
       if item.ID !=id{
		   newcars=append(newcars, item)
	   }
	}
    cars=newcars
	return cars
}

func main(){
	// router:=gin.Default()
	// router.DELETE("/delete",deleteById)
	// router.POST("/cars",createCar)
	// router.GET("/cars/:id",carById)
	// router.GET("/cars",getCars)
	// router.Run("localhost:8080")

	database.Main()
}

