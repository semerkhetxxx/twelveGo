package database

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "log"
)

func InitDB() {
    connStr := "user=yourusername dbname=yourdbname sslmode=disable password=yourpassword" // 請替換為實際的連接信息
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("無法連接到數據庫: ", err)
    }

    // 設置數據庫連接池參數
    db.SetMaxOpenConns(20)
    db.SetMaxIdleConns(20)
    db.SetConnMaxLifetime(0)

    // 測試數據庫連接
    if err := db.Ping(); err != nil {
        log.Fatal("無法 ping 通數據庫: ", err)
    }

    fmt.Println("數據庫連接成功！")
    // 確保 db 對象全局可用或按需傳遞
}
