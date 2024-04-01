package model

import (
    "github.com/google/uuid"
    "time"
)

// User 定義了用戶的結構體
type User struct {
    ID        uuid.UUID `json:"id" db:"id"` // 使用 UUID 作為主鍵
    Email     string    `json:"email" db:"email"`
    Name      string    `json:"name" db:"name"` // 添加名稱字段
    GoogleID  string    `json:"google_id,omitempty" db:"google_id"`
    LineID    string    `json:"line_id,omitempty" db:"line_id"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
