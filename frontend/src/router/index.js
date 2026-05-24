import { createRouter, createWebHistory } from 'vue-router'
import Home from '../pages/Home.vue'
import Items from '../pages/Items.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/items',
    name: 'Items',
    component: Items
  }
]

const router = createRouter({
  history: createWebHistory('/catmicraft/'),
  routes
})

export default router
