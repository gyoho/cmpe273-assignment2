package main

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "fmt"
    "log"
    "gopkg.in/mgo.v2"
    "./controllers"
)

const (
     dbUser string = "admin"
     dbPassword string = "admin"
     dbServer string = "ds043694.mongolab.com"
     dbPort string = "43694"
     dbName string = "cmpe273_asgmt2"
)

func main() {
    // Instantiate a new router
    router := httprouter.New()

    // Get a UserController instance
    userController := controllers.NewUserController(getMongoSession())

    // Get a user resource
    router.GET("/locations/:id", userController.GetUser)
    router.POST("/locations", userController.CreateUser)
    router.PUT("/locations/:id", userController.UpdateUser)
    router.DELETE("/locations/:id", userController.RemoveUser)

    // Fire up the server
    fmt.Println("Server listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getMongoSession() *mgo.Session {
    // Test
    // session, err := mgo.Dial("mongodb://localhost")
    
    // Production
    url := "mongodb://" + dbUser + ":" + dbPassword + "@" + dbServer + ":" + dbPort + "/" + dbName
    session, err := mgo.Dial(url)

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }

    return session
}
