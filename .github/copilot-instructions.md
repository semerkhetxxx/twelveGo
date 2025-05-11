# Copilot 使用說明



此專案主要使用 **Go 語言**，並搭配 **Gin 框架** 開發 API 服務（參考 [cmd/api/main.go](cmd/api/main.go)），並使用 **GORM** 進行資料庫操作（參考 [internal/database/db.go](internal/database/db.go)）。  
此外，使用 **Viper** 來管理設定檔，並採用 **Zap 日誌庫** 進行日誌記錄。

### Guidelines for GO

#### GIN

- Use middleware for cross-cutting concerns like authentication, logging, and request validation
- Implement structured logging with context for better debugging of {{error_scenarios}}
- Use binding validation for request payloads with custom validators for complex business rules
- Apply the context package properly to manage request-scoped values and cancellation signals
- Implement proper error handling with custom error types and consistent HTTP status codes
- Use the gin.H map for JSON responses consistently across handlers for {{api_endpoints}}


## 📌 **程式碼風格與最佳實踐**
請遵循 Google 的 **[Go Code Style Guide](https://google.github.io/styleguide/go/)**，並遵守以下規範：

### **1️⃣ 命名規範**
- **變數與函式**：使用 **駝峰式命名法 (camelCase)**，如 `userID`, `getUserInfo()`
- **常數**：使用 **全大寫與底線 (SNAKE_CASE)**，如 `MAX_RETRY_COUNT`
- **結構體**：使用 **駝峰式 (CamelCase)**，如：
  ```go
  type User struct {
      ID   uuid.UUID
      Name string
  }
  ```
- **介面 (interface)**：命名應以 `er` 結尾，如 `Logger`, `UserFetcher`
- **資料庫模型欄位**：
  - SQL 欄位請使用 **snake_case**（如 `created_at`）
  - Go 結構體對應欄位請使用 **駝峰命名**（如 `CreatedAt`）

### **2️⃣ 代碼格式**
- **每行程式碼不超過 80 字元**
- **使用 `gofmt` 自動格式化**
- **避免未使用的變數**，可以使用 `_` 來忽略返回值：
  ```go
  _, err := someFunc()
  ```

---

## 📌 **日誌記錄**
- **所有日誌請使用 [Zap](https://github.com/uber-go/zap)**
- **初始化日誌記錄器 (`zap.Logger`)**
- **區分日誌等級**：`Info`、`Warn`、`Error`
- **務必在錯誤處理時記錄錯誤**
- **範例：**
  ```go
  import "go.uber.org/zap"

  var logger, _ = zap.NewProduction()

  func GetUser(id string) {
      logger.Info("Fetching user", zap.String("user_id", id))
      user, err := fetchUserFromDB(id)
      if err != nil {
          logger.Error("Failed to fetch user", zap.Error(err))
      }
  }
  ```

---

## 📌 **框架與套件**
請優先使用以下 **框架與套件**：
| **功能**       | **推薦套件**                | **說明**                  |
|--------------|--------------------------|---------------------------|
| **Web 框架**   | [Gin](https://gin-gonic.com/) | 高效能 HTTP 框架 |
| **資料庫 ORM** | [GORM](https://gorm.io/) | 提供 Golang ORM 支援 |
| **設定管理**   | [Viper](https://github.com/spf13/viper) | 讀取 `config.yaml` |
| **日誌記錄**   | [Zap](https://go.uber.org/zap) | 高效能結構化日誌 |

---

## 📌 **錯誤處理**
- **所有函式皆應檢查 `error`**
- **避免忽略錯誤**
- **建議封裝錯誤並提供清晰訊息**
- **範例：**
  ```go
  func GetUser(id string) (*User, error) {
      user, err := fetchUserFromDB(id)
      if err != nil {
          return nil, fmt.Errorf("GetUser failed: %w", err)
      }
      return user, nil
  }
  ```

---

## 📌 **測試**
- **所有函式應提供單元測試**
- **使用 Go 內建 `testing` 套件**
- **測試命名規則**
  - `Test_函式名稱`
  - `Benchmark_函式名稱`
- **範例：**
  ```go
  import "testing"

  func TestGetUser(t *testing.T) {
      user, err := GetUser("123")
      if err != nil {
          t.Errorf("GetUser failed: %v", err)
      }
      if user == nil {
          t.Error("Expected user, got nil")
      }
  }
  ```
- **執行測試**
  ```sh
  go test ./...
  ```

---

## 📌 **REST API 回應規範**
- **所有 API 回應應符合以下格式**
- **回應格式範例**
  ```json
  {
      "status": "success",
      "message": "用戶資料獲取成功",
      "data": {
          "user_id": "123",
          "name": "Alice"
      }
  }
  ```
- **錯誤回應應包含 `error_code`**
  ```json
  {
      "status": "error",
      "error_code": "USER_NOT_FOUND",
      "message": "找不到該用戶"
  }
  ```
- **在 Gin 控制器中返回 JSON**
  ```go
  c.JSON(http.StatusOK, gin.H{
      "status": "success",
      "message": "用戶資料獲取成功",
      "data": user,
  })
  ```

---

## 📌 **專案目錄結構**
| **目錄**        | **描述** |
|----------------|-------------------------------|
| `cmd/api/`     | 服務啟動點 (`main.go`) |
| `pkg/api/`     | API 控制器 |
| `pkg/model/`   | 資料庫模型 |
| `internal/database/` | DB 連線池與 GORM 操作 |
| `pkg/middleware/` | Gin 中間件 |
| `config/`      | 設定檔 (`config.yaml`) |
| `test/`        | 測試程式 |

---

## 📌 **請求回應**
- **請確保所有函式保持簡潔且單一職責**
- **避免產生與專案無關的輸出**
- **API 回應格式應符合上述標準**

---

## 📌 **參考檔案**
- **[cmd/api/main.go](../cmd/api/main.go)** - 主要 API 入口
- **[internal/database/db.go](../internal/database/db.go)** - 資料庫初始化與連線管理
- **[pkg/middleware/logger.go](../pkg/middleware/logger.go)** - Zap 日誌記錄
- **[pkg/models/user.go](../pkg/models/user.go)** - `User` 資料表對應的 Go Struct
- **[test/auth_test.go](../test/auth_test.go)** - `auth` API 測試

---

## 📌 **如何使用**
### 1️⃣ **安裝依賴**
```sh
go mod tidy
```
### 2️⃣ **執行應用**
```sh
go run cmd/api/main.go
```
### 3️⃣ **執行測試**
```sh
go test ./...
```

---

## 📌 **結論**
✅ **遵守 Google Go Coding Style**  
✅ **使用 Gin、GORM、Viper、Zap 作為標準框架**  
✅ **所有 API 回應標準化**  
✅ **確保單元測試覆蓋率**  
✅ **解釋以及說明永遠使用zh-tw**  

🚀 **請務必遵守本文件，以確保程式碼品質！**


## DATABASE

### Guidelines for NOSQL

#### DYNAMODB

- Design access patterns first, then create tables and indexes to support {{query_requirements}}
- Implement single-table design for related entities to minimize RCU/WCU costs
- Use sparse indexes and composite keys for efficient querying

#### POSTGRES

- Use connection pooling to manage database connections efficiently
- Implement JSONB columns for semi-structured data instead of creating many tables for {{flexible_data}}
- Use materialized views for complex, frequently accessed read-only data

