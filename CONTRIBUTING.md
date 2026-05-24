# 貢獻指南

感謝您對 Catmi Craft 的興趣！

## 如何貢獻

### 報告 Bug

1. 檢查 [Issues](../../issues) 以確保 Bug 尚未報告
2. 如果您找不到現有的 Issue，請 [創建新的 Issue](../../issues/new)
3. 提供詳細的錯誤描述和復現步驟

### 建議功能

1. 檢查 [Issues](../../issues) 了解現有的建議
2. 創建新的 Issue 並標記為 "enhancement"
3. 詳細描述功能及其好處

### 代碼貢獻

1. Fork 此倉庫
2. 創建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 開啟 Pull Request

## 代碼風格

我們使用以下工具維護代碼風格：

- **Prettier** - 代碼格式化
- **ESLint** - 代碼質量
- **EditorConfig** - 編輯器配置

在提交之前，請確保代碼符合這些標準：

```bash
npm run lint
npm run format
```

## 提交消息

請遵循以下格式提交消息：

```
type(scope): subject

body

footer
```

示例：
```
feat(items): add delete functionality

Implement item deletion feature with confirmation dialog.

Closes #123
```

### Types:
- `feat`: 新功能
- `fix`: 修復錯誤
- `docs`: 文檔更改
- `style`: 代碼風格（不影響功能）
- `refactor`: 代碼重構
- `test`: 添加或修改測試
- `chore`: 維護任務

## 開發設置

```bash
git clone https://github.com/yourusername/catmicraft.git
cd catmicraft
npm install
cd frontend && npm install && cd ..
cd api && npm install && cd ..
npm run dev
```

## 測試

在提交 Pull Request 前，請確保所有測試都通過：

```bash
npm run test
npm run test:frontend
npm run test:api
```

## 許可

通過提交代碼，您同意您的貢獻根據 MIT 許可證發布。

## 行為準則

### 我們的承諾

我們致力於提供一個開放且歡迎的環境。

### 標准

- 使用熱情的語言
- 尊重不同的觀點
- 接受建設性的批評
- 專注於對社區最有益的事情

### 執行

不適當行為的例子包括：
- 使用帶有性暗示的語言或意象
- 人身攻擊
- 騷擾或辱罵評論
- 未經許可發布他人的私人信息

## 問題？

如有任何問題，請在 [Discussions](../../discussions) 中提出。

---

謝謝您的貢獻！🙏
