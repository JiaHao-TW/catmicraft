# 部署指南

本項目可以部署到 GitHub Pages（前端），後端採用 Go 提供 API。

## 📋 前置準備

1. **MongoDB Atlas 帳戶**
   - 訪問 https://www.mongodb.com/cloud/atlas
   - 創建免費帳戶
   - 創建一個叢集
   - 獲取連接字符串（格式：`mongodb+srv://username:password@cluster.mongodb.net/dbname`）

2. **GitHub 帳戶** 
   - 創建一個新的倉庫

3. **Go 部署平台帳戶**
   - 選擇任意支援 Go 的平台，例如 Railway、Render、Heroku
   - 使用 GitHub 帳戶連接或部署代碼

## 🚀 部署步驟

### 步驟 1: 準備本地環境

```bash
# 克隆倉庫
git clone https://github.com/yourusername/catmicraft.git
cd catmicraft

# 安裝根目錄依賴
npm install

# 安裝前端依賴
cd frontend && npm install && cd ..

# 安裝後端依賴
cd api && npm install && cd ..
```

### 步驟 2: 設定本地環境變數

**frontend/.env.local:**
```
VITE_API_URL=http://localhost:3001
```

**api/.env.local:**
```
MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/catmicraft
PORT=3001
```

### 步驟 3: 本地測試

```bash
# 運行開發伺服器（需要兩個終端）
npm run dev

# 或分別運行
# 終端 1: cd frontend && npm run dev
# 終端 2: cd api && go run main.go
```

訪問 http://localhost:5173

### 步驟 4: 部署後端（Go API）

1. 推送到 GitHub
```bash
git add .
git commit -m "Initial commit"
git push origin main
```

2. 使用任意支援 Go 的部署平台部署後端
   - 設定環境變數 `MONGODB_URI`
   - 選擇適合的部署 URL

3. 在前端配置中使用部署後端的 URL

### 步驟 5: 更新前端配置

編輯 `.github/workflows/deploy.yml`，將 API URL 更改為後端部署的 URL：

```yaml
env:
  VITE_API_URL: https://your-go-backend.example.com
```

編輯 `vite.config.js` 中的基礎路徑（如果需要）。

### 步驟 6: 啟用 GitHub Pages

1. 進入 GitHub 倉庫設定
2. 找到 "Pages" 部分
3. 在 "Build and deployment" 中選擇 "GitHub Actions"

### 步驟 7: 推送並部署

```bash
git add .
git commit -m "Deploy configuration"
git push origin main
```

GitHub Actions 將自動構建並部署前端到 GitHub Pages。

## 🔍 驗證部署

### 檢查前端
- 訪問 `https://yourusername.github.io/catmicraft`

### 檢查後端健康狀態
- 訪問 `https://your-go-backend.example.com/api/health`

### 檢查 API
- 嘗試獲取項目：`https://your-go-backend.example.com/api/items`

## 🆘 常見問題

### API 連接失敗
- 檢查部署平台環境變數是否設定
- 確保 MongoDB URI 正確
- 檢查 MongoDB Atlas IP 白名單

### GitHub Pages 未更新
- 檢查 GitHub Actions 是否成功運行
- 查看 Actions 選項卡中的運行日誌

### 跨域問題
- 確保後端 CORS 設定正確
- 檢查 API 端點是否可訪問

## 📝 連結

- [MongoDB Atlas](https://www.mongodb.com/cloud/atlas)
- [Go 部署文檔](https://go.dev/doc/deploy/)
- [GitHub Pages 文檔](https://pages.github.com/)
- [Vue 3 文檔](https://vuejs.org/)
