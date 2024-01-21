package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	mongodb "main.go/mongoDB"
	"main.go/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Some error occured. Err: %s \n", err)
	}

	m, err := mongodb.InitDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r := gin.Default()
	r.Use(cors.Default())
	routes.PostProductInfo(r, m)
	routes.GetSamePhones(r, m)
	routes.SearchProductsFromSearchBar(r, m)
	routes.GetSimilarProducts(r, m)
	routes.SearchProducts(r, m)
	routes.IncrementField(r, m)
	routes.GetTopProducts(r, m)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error running server")
	}

	fmt.Println(m.Client)
}
