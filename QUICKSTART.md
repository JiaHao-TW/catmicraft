# Catmi Craft 快速開始指南

## 🎯 30秒快速開始

### 本地開發（無 MongoDB）

```bash
# 1. 進入項目目錄
cd catmicraft

# 2. 安裝所有依賴
npm install
cd frontend && npm install && cd ..
cd api && npm install && cd ..

# 3. 啟動開發伺服器
npm run dev
```

訪問 http://localhost:5173

### 完整部署（使用 MongoDB）

1. **創建 MongoDB 帳戶**
   - 訪問 [MongoDB Atlas](https://www.mongodb.com/cloud/atlas)
   - 創建免費帳戶和叢集
   - 獲取連接字符串

2. **設定環境變數**
   - 編輯 `api/.env.local`
   - 粘貼 MongoDB 連接字符串

3. **推送到 GitHub**
   ```bash
   git init
   git add .
   git commit -m "Initial commit"
   git remote add origin https://github.com/yourusername/catmicraft.git
   git push -u origin main
   ```

4. **部署後端到 Go 平台**
   - 選擇支援 Go 的部署服務，例如 Railway、Render、Heroku
   - 連接 GitHub 倉庫
   - 添加環境變數 `MONGODB_URI`
   - 自動部署

5. **啟用 GitHub Pages**
   - 在 GitHub 倉庫設定中找到 Pages
   - 選擇 "GitHub Actions" 作為源
   - 完成！

## 📂 主要文件

| 文件 | 用途 |
|------|------|
| `frontend/` | Vue 3 前端應用 |
| `api/` | Go 後端服務 |
| `DEPLOYMENT.md` | 詳細部署指南 |
| `DEVELOPMENT.md` | 開發文檔 |

## 🚀 命令參考

```bash
npm run dev           # 開發模式（同時運行前後端）
npm run build         # 構建項目
npm run build:frontend # 只構建前端
npm run build:api     # 只構建後端
npm run preview       # 預覽構建結果
```

## 🔧 技術棧

- **前端**: Vue 3 + Vite + Vue Router
- **後端**: Go + MongoDB
- **部署**: GitHub Pages + Go API Hosting

## 📝 API 端點

```
GET    /api/items           # 獲取所有項目
POST   /api/items           # 創建項目
GET    /api/items/:id       # 獲取單個項目
PUT    /api/items/:id       # 更新項目
DELETE /api/items/:id       # 刪除項目
GET    /api/health          # 健康檢查
```

## 🆘 幫助

- 遇到問題？查看 `DEPLOYMENT.md` 或 `DEVELOPMENT.md`
- 有任何疑問，檢查環境變數是否正確設定

## 📚 資源

- [Vue 3 文檔](https://vuejs.org/)
- [Go 文檔](https://go.dev/doc/)
- [MongoDB 文檔](https://docs.mongodb.com/)
- [Vite 文檔](https://vitejs.dev/)
