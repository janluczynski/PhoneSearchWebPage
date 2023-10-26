package main

import mongodb "main.go/mongoDB"

func main() {
	m, err := mongodb.InitDB()
	if err != nil {
		panic(err)
	}
	//

	err = m.AddProducts([]interface{}{"test"})
	if err != nil {
		panic(err)
	}
}
