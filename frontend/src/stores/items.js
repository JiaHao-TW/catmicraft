import { defineStore } from 'pinia'
import { ref } from 'vue'
import api from '../api'

export const useItemStore = defineStore('items', () => {
  const items = ref([])
  const loading = ref(false)
  const error = ref(null)

  const fetchItems = async () => {
    loading.value = true
    error.value = null
    try {
      const response = await api.get('/items')
      items.value = response.data
    } catch (err) {
      error.value = '無法載入項目列表'
      console.error(err)
    } finally {
      loading.value = false
    }
  }

  const addItem = async (itemData) => {
    try {
      const response = await api.post('/items', itemData)
      items.value.push(response.data)
      return response.data
    } catch (err) {
      error.value = '無法新增項目'
      throw err
    }
  }

  const updateItem = async (id, itemData) => {
    try {
      const response = await api.put(`/items/${id}`, itemData)
      const index = items.value.findIndex(item => item._id === id)
      if (index !== -1) {
        items.value[index] = response.data
      }
      return response.data
    } catch (err) {
      error.value = '無法更新項目'
      throw err
    }
  }

  const deleteItem = async (id) => {
    try {
      await api.delete(`/items/${id}`)
      items.value = items.value.filter(item => item._id !== id)
    } catch (err) {
      error.value = '無法刪除項目'
      throw err
    }
  }

  return {
    items,
    loading,
    error,
    fetchItems,
    addItem,
    updateItem,
    deleteItem
  }
})
