# Bootstrap 集成更新

## ✨ 新增功能

本次更新已完成以下改進：

### 🎨 視覺設計
- ✅ 添加 Bootstrap 5.3 框架
- ✅ 集成 Bootstrap Icons 圖標庫
- ✅ 根據 Catmi Craft 官方品牌色系調整配色
- ✅ 重新設計所有頁面排版
- ✅ 實現響應式移動設計

### 🎯 品牌色系
- **主色**: #E8767B (粉紅色)
- **次色**: #FFB6C1 (淡粉紅)
- **亮色**: #FFF5F7 (淡米色)
- **米色**: #F5F1E8 (米色背景)
- **深色**: #6B4A47 (棕色文字)
- **強調色**: #D4A574 (淡棕色)

### 📱 頁面更新

#### 首頁 (Home.vue)
- 英雄區域 - 動畫浮動效果
- 特色卡片 - 4 大特色展示
- CTA 區域 - 行動呼籲
- 底部區域 - 品牌寄語

#### 產品頁 (Items.vue)
- 產品表單 - 優雅的輸入區域
- 產品網格 - 3 欄響應式布局
- 產品卡片 - 懸停效果和動畫
- 空狀態 - 友好的提示界面

#### 導航欄 (App.vue)
- 漸變背景導航欄
- 響應式漢堡菜單
- 醒目的品牌標識
- 社交媒體鏈接

### 🔧 技術改進
- 使用 Bootstrap 的柵欄系統
- 統一的間距和邊距
- 專業的陰影效果
- 平滑的過渡動畫
- 完全響應式設計

## 📦 新增依賴

```json
{
  "bootstrap": "^5.3.0",
  "bootstrap-icons": "^1.11.0"
}
```

## 🚀 安裝步驟

### 1. 安裝新依賴
```bash
cd frontend
npm install
```

### 2. 本地運行
```bash
npm run dev
```

訪問 http://localhost:5173

### 3. 生產構建
```bash
npm run build
```

## 📋 文件變更

已修改文件：
- `package.json` - 添加 Bootstrap 和 Icons
- `src/main.js` - 導入 Bootstrap CSS 和 JavaScript
- `src/style.css` - 自定義品牌色系和組件樣式
- `src/App.vue` - 使用 Bootstrap 導航欄和佈局
- `src/pages/Home.vue` - 重新設計首頁
- `src/pages/Items.vue` - 重新設計產品頁面
- `vite.config.js` - 優化 Bootstrap 構建
- `index.html` - 更新 meta 標籤和標題

新增文件：
- `.env.production` - 生產環境配置

## 🎨 自定義樣式

所有自定義樣式都在 `src/style.css` 中定義：

```css
:root {
  /* Catmi Craft 品牌色系 */
  --catmi-primary: #E8767B;
  --catmi-secondary: #FFB6C1;
  --catmi-light: #FFF5F7;
  --catmi-cream: #F5F1E8;
  --catmi-dark: #6B4A47;
  --catmi-accent: #D4A574;
}
```

可在任何組件中使用：
```vue
<div style="color: var(--catmi-primary);">文本</div>
```

## 🔄 更新 Bootstrap 版本

如需更新 Bootstrap：

```bash
npm update bootstrap bootstrap-icons
```

## 📚 文檔連結

- [Bootstrap 官方文檔](https://getbootstrap.com/)
- [Bootstrap Icons](https://icons.getbootstrap.com/)
- [Bootstrap Vue 集成](https://bootstrap-vue.org/)

## ✅ 下一步

1. 運行本地開發伺服器測試新設計
2. 根據需要自定義顏色和樣式
3. 測試所有設備的響應式設計
4. 部署前端到 GitHub Pages，後端到 Go 支援平台

## 💬 常見問題

### Q: 如何修改品牌顏色？
A: 編輯 `src/style.css` 中的 `:root` 變數即可。

### Q: Bootstrap 組件未響應？
A: 確保在 `src/main.js` 中已導入 `import * as bootstrap from 'bootstrap'`

### Q: 如何添加新的 Bootstrap 組件？
A: 直接在 Vue 模板中使用 Bootstrap 的 class 和結構即可。

---

**現在您可以享受現代化、美觀的 Catmi Craft 網站了！** 🎉
