<template>
  <div id="app" class="app-container">
    <header class="navbar navbar-expand-lg navbar-light sticky-top" style="background: linear-gradient(135deg, #E8767B 0%, #FFB6C1 100%); box-shadow: 0 2px 12px rgba(232, 118, 123, 0.15);">
      <div class="container-lg">
        <router-link to="/" class="navbar-brand">
          <i class="bi bi-heart-fill me-2"></i>
          <span class="fw-bold">Catmi Craft</span>
        </router-link>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
          <ul class="navbar-nav ms-auto">
            <li class="nav-item">
              <router-link to="/" class="nav-link" :class="{ active: $route.path === '/' }">
                首頁
              </router-link>
            </li>
            <!-- <li class="nav-item">
              <router-link to="/items" class="nav-link" :class="{ active: $route.path === '/items' }">
                產品
              </router-link>
            </li> -->
          </ul>
        </div>
      </div>
    </header>

    <main class="flex-grow-1 py-4">
      <router-view />
    </main>

    <footer class="bg-light text-center py-4 border-top border-secondary border-opacity-25">
      <div class="container-lg">
        <p class="text-secondary mb-1">&copy; 2025 Catmi Craft. 所有權利保留。</p>
        <div class="d-flex justify-content-center gap-2 mt-2">
          <a href="#" class="text-decoration-none"><i class="bi bi-instagram"></i></a>
          <a href="#" class="text-decoration-none"><i class="bi bi-facebook"></i></a>
          <a href="#" class="text-decoration-none"><i class="bi bi-chat-dots"></i></a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
</script>

<style scoped>
.navbar-brand {
  color: white !important;
  font-size: 1.5rem;
}

.navbar-brand:hover {
  color: #FFF !important;
}

.nav-link {
  color: white !important;
  font-weight: 600;
  padding: 0.5rem 1rem !important;
  transition: all 0.3s;
  border-radius: 8px;
}

.nav-link:hover {
  background-color: rgba(255, 255, 255, 0.2);
  transform: translateY(-2px);
}

.nav-link.active {
  background-color: rgba(255, 255, 255, 0.3);
}

main {
  display: flex;
  flex-direction: column;
}
</style>
