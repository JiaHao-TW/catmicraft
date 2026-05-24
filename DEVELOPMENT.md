# 開發指南

## 項目結構

```
catmicraft/
├── frontend/                 # Vue 3 前端應用
│   ├── src/
│   │   ├── pages/           # 頁面組件
│   │   ├── router/          # 路由配置
│   │   ├── api.js           # API 客戶端
│   │   ├── App.vue          # 根組件
│   │   ├── main.js          # 入口文件
│   │   └── style.css        # 全局樣式
│   ├── index.html           # HTML 模板
│   ├── vite.config.js       # Vite 配置
│   └── package.json
│
├── api/                      # Go 後端
│   ├── main.go              # Go API 伺服器
│   ├── go.mod               # Go 模組定義
│   ├── go.sum               # Go 依賴鎖定
│   ├── .env.example         # 環境變數模板
│   └── .env.local           # 本地環境變數
│
├── .github/
│   └── workflows/
│       └── deploy.yml       # GitHub Actions 工作流
│
├── DEPLOYMENT.md            # 部署指南
├── DEVELOPMENT.md           # 開發指南（本文件）
├── README.md                # 項目說明
├── package.json             # 根目錄依賴
└── .gitignore              # Git 忽略文件
```

## 開發命令

### 安裝依賴
```bash
npm install              # 安裝根目錄依賴
cd frontend && npm install && cd ..  # 安裝前端依賴
cd api && go mod tidy && cd ..       # 安裝後端依賴
```

### 開發模式
```bash
npm run dev              # 同時運行前端和後端
npm run dev:frontend     # 只運行前端
npm run dev:api         # 只運行後端
```

### 構建
```bash
npm run build            # 構建前端和後端
npm run build:frontend   # 只構建前端
npm run build:api       # 只構建後端
```

### 預覽
```bash
npm run preview          # 預覽構建結果
```

## 前端開發

### 技術棧
- **Vue 3** - 漸進式 JavaScript 框架
- **Vite** - 下一代前端構建工具
- **Vue Router** - 官方路由庫
- **Pinia** - Vue 狀態管理
- **Axios** - HTTP 客戶端

### 項目結構說明

#### `src/main.js`
應用入口文件，初始化 Vue 應用和所有插件。

#### `src/api.js`
Axios 實例配置，所有 API 請求都通過此文件。

```javascript
import api from './api'
api.get('/items')
api.post('/items', data)
```

#### `src/App.vue`
根組件，定義了頁面布局（頭部、導航、主體、頁腳）。

#### `src/router/index.js`
路由配置，定義所有可用的頁面和路由。

#### `src/pages/`
頁面組件：
- `Home.vue` - 首頁
- `Items.vue` - 產品列表

### 添加新頁面

1. 在 `src/pages/` 創建新的 `.vue` 文件
2. 在 `src/router/index.js` 中添加路由
3. 在 `src/App.vue` 中添加導航鏈接

### 環境變數

在 `frontend/.env.local` 中設定：

```
VITE_API_URL=http://localhost:3001
```

在代碼中使用：
```javascript
const apiUrl = import.meta.env.VITE_API_URL
```

## 後端開發

### 技術棧
- **Go** - 後端開發語言
- **Chi** - Go 路由框架
- **MongoDB** - 數據庫
- **MongoDB Go Driver** - 官方驅動

### API 端點

#### 項目相關
- `GET /api/items` - 獲取所有項目
- `POST /api/items` - 創建新項目
- `GET /api/items/:id` - 獲取單個項目
- `PUT /api/items/:id` - 更新項目
- `DELETE /api/items/:id` - 刪除項目

#### 系統
- `GET /api/health` - 健康檢查

### 添加新的 API 端點

1. 在 `api/main.go` 中定義路由
2. 使用 Chi 處理器：

```go
r.Get("/api/new-endpoint", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "ok"})
})
```

### 環境變數

在 `api/.env.local` 中設定：

```
MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/dbname
PORT=3001
```

### MongoDB 連接

```go
client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
if err != nil {
    log.Fatalf("MongoDB connection failed: %v", err)
}
collection := client.Database("catmicraft").Collection("items")
```

## 常見開發任務

### 添加新的數據模型

1. 在 `api/main.go` 定義 Schema 和 Model
2. 創建對應的 API 路由

### 修改 API 端點

1. 編輯 `api/main.go` 中的相應處理器
2. 測試本地運行

### 修改前端界面

1. 編輯 `src/pages/` 中的 Vue 組件
2. 前端會自動熱重載
3. 檢查瀏覽器開發工具查看改變

## 調試

### 前端調試
- 使用 Vue DevTools 瀏覽器擴展
- 打開瀏覽器開發工具（F12）
- 在 Console 選項卡查看日誌

### 後端調試
- 檢查終端中的日誌輸出
- 使用 `console.log()` 或 `console.error()`
- 訪問 `/api/health` 檢查伺服器狀態

### API 測試
使用 Postman 或類似工具測試 API 端點。

## 提交和推送

```bash
git add .
git commit -m "描述你的改變"
git push origin main
```

## 獲取幫助

- [Vue 3 文檔](https://vuejs.org/)
- [Go 文檔](https://go.dev/doc/)
- [MongoDB 文檔](https://docs.mongodb.com/)
- [Vite 文檔](https://vitejs.dev/)
