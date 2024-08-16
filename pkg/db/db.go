package db

import (
    "database/sql"
    "fmt"
    "log"

   "restApiSkladProject/pkg/mocks"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5431
    user     = "john"
    password = "ramzesik"
    dbname   = "goodsDB"
)

func Connect() *sql.DB {
    connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", connInfo)
    if err != nil {
        log.Fatal(err)
    }
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Successfully connected to db!")
    return db
}

func CloseConnection(db *sql.DB) {
    defer db.Close()
}

func CreateTable(db *sql.DB) {
    var exists bool
    if err := db.QueryRow("SELECT EXISTS (SELECT FROM pg_tables WHERE  schemaname = 'public' AND tablename = 'goodsDB' );").Scan(&exists); err != nil {
        fmt.Println("failed to execute query", err)
        return
    }
    if !exists {
        results, err := db.Query(
            "CREATE TABLE goodsDB (id VARCHAR(36) PRIMARY KEY, name VARCHAR(100) NOT NULL, provider integer NOT NULL, price integer NOT NULL, quantity integer NOT NULL);")
        if err != nil {
            fmt.Println("failed to execute query", err)
            return
        }
        fmt.Println("Table created successfully", results)

        for _, goods := range mocks.Goods {
            queryStmt := `INSERT INTO goodsDB (id,name,provider,price,quantity) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

            err := db.QueryRow(queryStmt, &goods.Id, &goods.Name, &goods.Provider, &goods.Price, &goods.Quantity).Scan(&goods.Id)
            if err != nil {
                log.Println("failed to execute query", err)
                return
            }
        }
        fmt.Println("Mock Articles included in Table", results)
    } else {
        fmt.Println("Table 'articles' already exists ")
    }

}