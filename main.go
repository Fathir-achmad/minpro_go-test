package main

import (
    "database/sql"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "minpro-practice/controllers"
    "minpro-practice/database"
    "os"

    _ "github.com/lib/pq"
)

var (
    DB  *sql.DB
    err error
)

func main() {
    // Load .env
    err = godotenv.Load("config/.env")
    if err != nil {
       panic("Error loading .env file")
    }

    // Build connection string
    psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
       os.Getenv("DB_HOST"),
       os.Getenv("DB_PORT"),
       os.Getenv("DB_USER"),
       os.Getenv("DB_PASSWORD"),
       os.Getenv("DB_NAME"),
    )

    // Open DB
    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
       panic(err)
    }

    // Test connection
    err = DB.Ping()
    if err != nil {
       panic(err)
    }

    // Run migration
    database.DBMigrate(DB)

    // Router
    router := gin.Default()
    router.GET("/persons", controllers.GetAllPerson)
    router.POST("/persons", controllers.InsertPerson)
    router.PUT("/persons/:id", controllers.UpdatePerson)
    router.DELETE("/persons/:id", controllers.DeletePerson)

    router.Run(":8080")
}
