# 項目概述

## 📖 Catmi Craft 網站

這是一個完整的全棧網站項目，前端使用 Vue 3，後端使用 Go + MongoDB。該項目展示了現代 Web 開發的最佳實踐，包括：

- ✅ 響應式設計
- ✅ RESTful API
- ✅ 數據庫集成
- ✅ 自動化部署
- ✅ 完整文檔

## 🎯 項目目標

創建一個可以在 GitHub Pages 上部署前端、並搭配 Go 後端的生產級別網站，展示完整的 Web 開發能力。

## 📊 項目統計

| 項目 | 數量 |
|------|------|
| 前端文件 | 15+ |
| 後端文件 | 10+ |
| 文檔頁面 | 7 |
| API 端點 | 6 |
| 部署提供商 | 2 |

## 🏗️ 架構圖

```
┌─────────────────────────────────────┐
│      GitHub Pages (前端)            │
│      https://user.github.io/        │
│                                     │
│   Vue 3 + Vite + Vue Router         │
│   ├── Home Page                     │
│   └── Items Page                    │
└──────────────┬──────────────────────┘
               │ API Calls
               │
┌──────────────▼──────────────────────┐
│      Go API Hosting 平台            │
│      https://your-go-backend.example.com/ │
│                                     │
│   Go HTTP API                       │
│   ├── GET /api/items                │
│   ├── POST /api/items               │
│   ├── PUT /api/items/:id            │
│   ├── DELETE /api/items/:id         │
│   └── GET /api/health               │
└──────────────┬──────────────────────┘
               │ Database Calls
               │
┌──────────────▼──────────────────────┐
│     MongoDB Atlas (數據庫)          │
│   mongodb+srv://cluster...          │
│                                     │
│   ├── Items Collection              │
│   │   ├── _id (Object ID)           │
│   │   ├── name (String)             │
│   │   ├── description (String)      │
│   │   ├── price (Number)            │
│   │   └── createdAt (Date)          │
│   └── ...                           │
└─────────────────────────────────────┘
```

## 📚 文檔目錄

### 快速參考
- [快速開始指南](./QUICKSTART.md) - 30秒快速開始

### 詳細指南
- [部署指南](./DEPLOYMENT.md) - 完整部署流程
- [開發指南](./DEVELOPMENT.md) - 開發和維護
- [API 文檔](./API.md) - API 參考
- [貢獻指南](./CONTRIBUTING.md) - 如何貢獻
- [測試指南](./TESTING.md) - 測試和 CI/CD
- [變更日誌](./CHANGELOG.md) - 版本歷史

### 項目文件
- [README.md](./README.md) - 項目簡介
- [LICENSE](./LICENSE) - MIT 許可證

## 🔄 開發流程

```
1. 本地開發
   └─> npm run dev

2. 測試
   └─> npm run test

3. 推送到 GitHub
   └─> git push origin main

4. 自動部署
   ├─> GitHub Actions (前端 → GitHub Pages)
   └─> Go API 平台 (後端自動部署)

5. 訪問網站
   ├─> 前端: https://user.github.io/catmicraft
   └─> 後端: https://your-go-backend.example.com
```

## 🚀 核心特性

### 前端特性
- 🎨 Vue 3 Composition API
- 🛣️ Vue Router 客戶端路由
- 📦 Pinia 狀態管理
- 🎯 Axios HTTP 客戶端
- 📱 響應式移動設計
- ⚡ Vite 快速構建

### 後端特性
- 🔌 Go REST API
- 🗄️ MongoDB 數據庫
- 🔐 CORS 安全配置
- 📊 健康檢查端點
- 🚀 可部署到任何 Go 支持平台
- ✔️ 錯誤處理

### 部署特性
- 🔄 GitHub Actions CI/CD
- 📲 GitHub Pages 靜態托管
- ☁️ Go 後端 API
- 🗃️ MongoDB Atlas 雲數據庫
- 🔐 環境變數管理

## 💡 學習資源

此項目展示了以下技術：

### 前端
- [Vue 3 官方文檔](https://vuejs.org/)
- [Vite 文檔](https://vitejs.dev/)
- [Vue Router 文檔](https://router.vuejs.org/)
- [Pinia 文檔](https://pinia.vuejs.org/)

### 後端
- [Go 文檔](https://go.dev/doc/)
- [MongoDB Go Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)
- [MongoDB 文檔](https://docs.mongodb.com/)

### 部署
- [GitHub Pages 文檔](https://pages.github.com/)
- [Railway 文檔](https://railway.app/docs)
- [MongoDB Atlas 文檔](https://docs.atlas.mongodb.com/)

## 🎓 適合學習者

此項目適合以下學習者：

- 🔰 初學者 - 了解全棧開發結構
- 📈 中級開發者 - 學習最佳實踐和部署
- 🚀 進階開發者 - 參考架構和優化

## ✨ 下一步

現在您可以：

1. **本地運行**: 按照 [QUICKSTART.md](./QUICKSTART.md)
2. **部署上線**: 按照 [DEPLOYMENT.md](./DEPLOYMENT.md)
3. **開發功能**: 按照 [DEVELOPMENT.md](./DEVELOPMENT.md)
4. **貢獻代碼**: 按照 [CONTRIBUTING.md](./CONTRIBUTING.md)

## 📞 支持

遇到問題？

- 📖 查看相關文檔
- 🔍 搜索已有的 Issues
- 💬 創建新的 Issue
- 📧 聯繫項目維護者

## 📄 許可證

此項目根據 MIT 許可證授權 - 查看 [LICENSE](./LICENSE) 文件了解詳情。

---

**開始探索吧！** 祝您開發愉快！🎉
