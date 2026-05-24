<template>
  <div class="items-page">
    <div class="container-lg">
      <!-- 頁面標題 -->
      <div class="page-header mb-5">
        <h1 class="fw-bold mb-2">
          <i class="bi bi-bag-heart me-2"></i>
          <span style="color: #E8767B;">Catmi</span> 的產品
        </h1>
        <p class="lead text-secondary">探索我們精心準備的手工創意產品</p>
      </div>

      <!-- 新增產品表單 -->
      <div class="mb-5">
        <div class="card border-0 shadow-sm">
          <div class="card-header bg-light border-0 py-3">
            <h5 class="mb-0 fw-bold">
              <i class="bi bi-plus-circle me-2" style="color: #E8767B;"></i>
              新增產品
            </h5>
          </div>
          <div class="card-body p-4">
            <form @submit.prevent="addItem" class="row g-3">
              <div class="col-md-6">
                <label class="form-label fw-bold">產品名稱</label>
                <input
                  v-model="newItem.name"
                  type="text"
                  class="form-control"
                  placeholder="輸入產品名稱"
                  required
                />
              </div>
              <div class="col-md-6">
                <label class="form-label fw-bold">價格</label>
                <div class="input-group">
                  <span class="input-group-text bg-light border-end-0">NT$</span>
                  <input
                    v-model.number="newItem.price"
                    type="number"
                    class="form-control border-start-0"
                    placeholder="0"
                    min="0"
                  />
                </div>
              </div>
              <div class="col-12">
                <label class="form-label fw-bold">產品描述</label>
                <textarea
                  v-model="newItem.description"
                  class="form-control"
                  rows="3"
                  placeholder="輸入產品描述"
                  required
                ></textarea>
              </div>
              <div class="col-12">
                <button type="submit" class="btn btn-primary btn-lg w-md">
                  <i class="bi bi-plus me-2"></i>新增產品
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>

      <!-- 載入狀態 -->
      <div v-if="loading" class="text-center py-5">
        <div class="spinner-border" style="color: #E8767B;" role="status">
          <span class="visually-hidden">載入中...</span>
        </div>
        <p class="mt-3 text-secondary">正在載入產品列表...</p>
      </div>

      <!-- 錯誤訊息 -->
      <div v-if="error" class="alert alert-danger alert-dismissible fade show" role="alert">
        <i class="bi bi-exclamation-circle me-2"></i>
        {{ error }}
        <button type="button" class="btn-close" @click="error = null"></button>
      </div>

      <!-- 產品網格 -->
      <div v-if="!loading && items.length > 0" class="items-grid">
        <div class="row g-4">
          <div v-for="item in items" :key="item._id" class="col-md-6 col-lg-4">
            <div class="product-card card border-0 h-100">
              <div class="card-img-top bg-light d-flex align-items-center justify-content-center" style="height: 200px;">
                <i class="bi bi-bag-heart" style="font-size: 4rem; color: #FFB6C1; opacity: 0.5;"></i>
              </div>
              <div class="card-body d-flex flex-column">
                <h5 class="card-title fw-bold text-truncate">{{ item.name }}</h5>
                <p class="card-text text-secondary flex-grow-1">{{ item.description }}</p>
                <div class="d-flex justify-content-between align-items-center mt-3 pt-3 border-top">
                  <div class="price-tag">
                    <span class="fw-bold" style="font-size: 1.5rem; color: #E8767B;">NT$ {{ item.price }}</span>
                  </div>
                  <button 
                    @click="deleteItem(item._id)"
                    class="btn btn-sm btn-outline-danger"
                    title="刪除產品"
                  >
                    <i class="bi bi-trash"></i>
                  </button>
                </div>
                <small class="text-muted d-block mt-2">
                  <i class="bi bi-calendar-event me-1"></i>
                  {{ formatDate(item.createdAt) }}
                </small>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 空狀態 -->
      <div v-if="!loading && items.length === 0" class="text-center py-5">
        <div class="mb-4">
          <i class="bi bi-inbox" style="font-size: 4rem; color: #FFB6C1; opacity: 0.5;"></i>
        </div>
        <h4 class="fw-bold text-secondary mb-2">還沒有產品</h4>
        <p class="text-secondary mb-4">試著新增一個產品吧！</p>
        <button class="btn btn-primary" @click="$router.push('/items')">
          <i class="bi bi-arrow-clockwise me-2"></i>重新載入
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../api'

const items = ref([])
const loading = ref(false)
const error = ref(null)
const newItem = ref({
  name: '',
  description: '',
  price: 0
})

const fetchItems = async () => {
  loading.value = true
  error.value = null
  try {
    const response = await api.get('/items')
    items.value = response.data
  } catch (err) {
    error.value = '無法載入產品列表，請稍後重試'
    console.error('Error fetching items:', err)
  } finally {
    loading.value = false
  }
}

const addItem = async () => {
  if (!newItem.value.name || !newItem.value.description) {
    error.value = '請填寫所有欄位'
    return
  }

  try {
    const response = await api.post('/items', newItem.value)
    items.value.unshift(response.data)
    newItem.value = { name: '', description: '', price: 0 }
    error.value = null
  } catch (err) {
    error.value = '新增產品失敗：' + (err.response?.data?.error || err.message)
    console.error('Error adding item:', err)
  }
}

const deleteItem = async (id) => {
  if (!confirm('確認刪除此產品？')) return

  try {
    await api.delete(`/items/${id}`)
    items.value = items.value.filter(item => item._id !== id)
    error.value = null
  } catch (err) {
    error.value = '刪除產品失敗：' + (err.response?.data?.error || err.message)
    console.error('Error deleting item:', err)
  }
}

const formatDate = (date) => {
  if (!date) return ''
  return new Date(date).toLocaleDateString('zh-TW', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })
}

onMounted(() => {
  fetchItems()
})
</script>

<style scoped>
.items-page {
  padding: 2rem 1rem;
  background: linear-gradient(135deg, #FFF5F7 0%, #F5F1E8 100%);
  min-height: calc(100vh - 200px);
}

.page-header {
  text-align: center;
}

.page-header h1 {
  font-size: 2.5rem;
}

.card {
  transition: all 0.3s ease;
}

.product-card {
  overflow: hidden;
  transition: all 0.3s ease;
  border-radius: 16px !important;
}

.product-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 12px 24px rgba(232, 118, 123, 0.15) !important;
}

.card-title {
  color: #E8767B;
  font-size: 1.1rem;
}

.form-label {
  color: #6B4A47;
}

.btn-primary {
  background-color: #E8767B;
  border-color: #E8767B;
}

.btn-primary:hover {
  background-color: #D4627B;
  border-color: #D4627B;
}

.btn-outline-danger {
  color: #E8767B;
  border-color: #FFE4E8;
}

.btn-outline-danger:hover {
  background-color: #FFE4E8;
  border-color: #E8767B;
}

.alert-danger {
  background-color: #FFE4E8;
  border-color: #E8767B;
  color: #D4627B;
}

.spinner-border {
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

@media (max-width: 768px) {
  .page-header h1 {
    font-size: 1.8rem;
  }

  .items-grid {
    grid-template-columns: 1fr;
  }
}
</style>
