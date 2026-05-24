# 測試

此項目使用以下工具進行測試：

## 前端測試

### 單元測試
```bash
cd frontend
npm run test
```

### E2E 測試
```bash
cd frontend
npm run test:e2e
```

## 後端測試

### API 測試
```bash
cd api
npm run test
```

## 集成測試

運行完整的集成測試：
```bash
npm run test:integration
```

## 覆蓋率報告

```bash
npm run coverage
```

## CI/CD

所有推送到 `main` 分支的代碼都會自動通過 GitHub Actions 進行測試。

查看 `.github/workflows/` 了解更多詳情。
