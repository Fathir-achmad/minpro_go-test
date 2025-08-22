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
       os.Getenv("PGHOST"),
       os.Getenv("PGPORT"),
       os.Getenv("PGUSER"),
       os.Getenv("PGPASSWORD"),
       os.Getenv("PGDATABASE"),
    )

    // Open DB
    DB, err = sql.Open(driverName:"postgres", psqlInfo)
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

    router.Run(":" + os.Getenv(key: "PORT"))
}
