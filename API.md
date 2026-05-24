# API 文檔

## 基本信息

**基礎 URL**: 
- 本地: `http://localhost:3001`
- 部署: `https://your-go-backend.example.com`

**認證**: 當前無需認證（公開 API）

## 端點列表

### 項目管理

#### 獲取所有項目
```
GET /api/items
```

**回應示例**:
```json
[
  {
    "_id": "507f1f77bcf86cd799439011",
    "name": "毛線球",
    "description": "柔軟的毛線球",
    "price": 299,
    "createdAt": "2025-05-24T00:00:00Z"
  }
]
```

---

#### 創建項目
```
POST /api/items
Content-Type: application/json
```

**請求體**:
```json
{
  "name": "產品名稱",
  "description": "產品描述",
  "price": 299
}
```

**回應** (201 Created):
```json
{
  "_id": "507f1f77bcf86cd799439011",
  "name": "產品名稱",
  "description": "產品描述",
  "price": 299,
  "createdAt": "2025-05-24T00:00:00Z"
}
```

**錯誤** (400 Bad Request):
```json
{
  "error": "名稱和描述為必填項"
}
```

---

#### 獲取單個項目
```
GET /api/items/:id
```

**參數**:
- `id` (string) - 項目 ID

**回應** (200 OK):
```json
{
  "_id": "507f1f77bcf86cd799439011",
  "name": "毛線球",
  "description": "柔軟的毛線球",
  "price": 299,
  "createdAt": "2025-05-24T00:00:00Z"
}
```

**錯誤** (404 Not Found):
```json
{
  "error": "項目未找到"
}
```

---

#### 更新項目
```
PUT /api/items/:id
Content-Type: application/json
```

**參數**:
- `id` (string) - 項目 ID

**請求體**:
```json
{
  "name": "新名稱",
  "description": "新描述",
  "price": 399
}
```

**回應** (200 OK):
```json
{
  "_id": "507f1f77bcf86cd799439011",
  "name": "新名稱",
  "description": "新描述",
  "price": 399,
  "createdAt": "2025-05-24T00:00:00Z"
}
```

---

#### 刪除項目
```
DELETE /api/items/:id
```

**參數**:
- `id` (string) - 項目 ID

**回應** (200 OK):
```json
{
  "message": "項目已刪除",
  "item": {
    "_id": "507f1f77bcf86cd799439011",
    "name": "毛線球",
    "description": "柔軟的毛線球",
    "price": 299,
    "createdAt": "2025-05-24T00:00:00Z"
  }
}
```

---

### 系統

#### 健康檢查
```
GET /api/health
```

**回應** (200 OK):
```json
{
  "status": "OK",
  "timestamp": "2025-05-24T12:34:56.789Z",
  "mongodb": "連接成功"
}
```

## 使用示例

### JavaScript / Fetch API

```javascript
// 獲取所有項目
const response = await fetch('http://localhost:3001/api/items')
const items = await response.json()

// 創建項目
const newItem = await fetch('http://localhost:3001/api/items', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    name: '新產品',
    description: '描述',
    price: 299
  })
})
const createdItem = await newItem.json()
```

### Axios

```javascript
import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:3001'
})

// 獲取所有項目
const items = await api.get('/api/items')

// 創建項目
const newItem = await api.post('/api/items', {
  name: '新產品',
  description: '描述',
  price: 299
})

// 刪除項目
await api.delete(`/api/items/${itemId}`)
```

### cURL

```bash
# 獲取所有項目
curl http://localhost:3001/api/items

# 創建項目
curl -X POST http://localhost:3001/api/items \
  -H "Content-Type: application/json" \
  -d '{"name":"新產品","description":"描述","price":299}'

# 刪除項目
curl -X DELETE http://localhost:3001/api/items/507f1f77bcf86cd799439011
```

## 錯誤處理

所有錯誤響應都遵循此格式：

```json
{
  "error": "錯誤消息說明"
}
```

### 常見錯誤代碼

| 代碼 | 含義 |
|------|------|
| 200 | 成功 |
| 201 | 已創建 |
| 400 | 請求錯誤 |
| 404 | 未找到 |
| 500 | 服務器錯誤 |

## CORS 政策

API 允許來自以下來源的跨域請求：
- `http://localhost:5173` (本地開發)
- `https://yourusername.github.io` (GitHub Pages)

## 速率限制

當前無速率限制。未來可能會添加。

## 版本控制

API 當前版本：v1

未來的更新可能會在 `/api/v2/` 下提供。
