# Catmi Craft 網站

一個現代化的 Vue 3 + MongoDB 全棧應用，前端部署到 GitHub Pages，後端採用 Go 提供 API。

## 📋 項目結構

```
catmicraft/
├── frontend/          # Vue 3 前端應用
├── api/              # Go 後端服務
├── .github/
│   └── workflows/    # GitHub Actions 工作流
└── package.json      # 根目錄依賴
```

## 🚀 快速開始

### 前提條件
- Node.js 16+
- Go 1.20+
- MongoDB Atlas 帳戶
- GitHub 帳戶

### 本地開發

#### 1. 安裝依賴
```bash
npm install
cd frontend && npm install && cd ..
cd api && go mod tidy && cd ..
```

#### 2. 設定環境變數
在 `frontend/.env.local` 創建：
```
VITE_API_URL=http://localhost:3001
```

在 `api/.env.local` 創建：
```
MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/catmicraft
```

#### 3. 運行開發伺服器
```bash
# 終端 1 - 前端
cd frontend
npm run dev

# 終端 2 - 後端
cd api
npm run dev
```

訪問 http://localhost:5173

## 📦 部署

### 部署架構

1. **推送到 GitHub**
```bash
git push origin main
```

2. **後端部署**
   - 後端已改為 Go，請部署到任何支援 Go 的平台（例如 Railway、Heroku、Render 等）
   - 設定環境變數 `MONGODB_URI`

3. **前端部署**
   - GitHub Actions 自動構建並部署前端到 GitHub Pages
   - 訪問 `https://yourusername.github.io/catmicraft`

## 🔧 技術棧

- **前端**: Vue 3, Vite, Axios
- **後端**: Go, MongoDB
- **部署**: GitHub Pages + Go API Hosting
- **數據庫**: MongoDB Atlas

## 📝 API 端點

- `GET /api/items` - 獲取所有項目
- `POST /api/items` - 創建項目
- `GET /api/items/:id` - 獲取單個項目
- `PUT /api/items/:id` - 更新項目
- `DELETE /api/items/:id` - 刪除項目

## 🤝 貢獻

歡迎提交 Pull Request！

## 📄 許可

MIT
