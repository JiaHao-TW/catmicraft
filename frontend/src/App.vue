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
              <router-link to="/" class="nav-link d-flex align-items-center gap-2" :class="{ active: $route.path === '/' }">
                <img src="./images/catClaws.png" alt="生活美學" class="header-btn-image" />
                <span>首頁</span>
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

    <footer class="text-center ">
      <!-- 底部區域 -->
    <section class="bottom-section py-1">
        <div class="container-lg">
          <div class="row align-items-center">
            <div class="col-md-6 mb-4 mb-md-0">
              <img src="./images/footer-decor.png" alt="Catmi 裝飾圖" class="footer-image footer-decor" />
            </div>
            <div class="col-md-6">
              <div class="footer-social-list">
                 <div class="social-item d-flex align-items-start mb-3">
                  <h5 class="fw-bold">一起在日常裡，<br>找到喜歡的自己。</h5>
                </div>

                <div class="social-item d-flex align-items-center mb-3">
                  <span class="social-icon me-3">
                    <img src="./images/Instagram_Glyph_White.png" alt="Instagram" class="social-icon-img" />
                  </span>
                  <div class="social-content d-flex align-items-center gap-3 flex-wrap">
                    <span class="social-title fw-bold">Instagram</span>
                    <span class="text-muted fw-bold">追蹤最新作品與日常</span>
                  </div>
                </div>

                <div class="social-item d-flex align-items-center mb-3">
                  <span class="social-icon me-3">
                     <img src="./images/threads-logo-white.png" alt="Instagram" class="social-icon-img" />
                  </span>
                  <div class="social-content d-flex align-items-center gap-3 flex-wrap">
                    <span class="social-title fw-bold">Threads</span>
                    <span class="text-muted fw-bold">一起分享生活小事</span>
                  </div>
                </div>

                <div class="social-item d-flex align-items-center mb-3">
                  <span class="social-icon me-3">
                    <img src="./images/Facebook_Logo_Secondary.png" alt="Facebook" class="social-icon-img" />
                  </span>
                  <div class="social-content d-flex align-items-center gap-3 flex-wrap">
                    <span class="social-title fw-bold">Facebook</span>
                    <span class="text-muted fw-bold">加入 Catmi 粉絲專頁</span>
                  </div>
                </div>

               </div>
            </div>
          </div>
        </div>
      </section>
      <div class="container-lg">
        <p class="text-secondary mb-1">&copy; 2026 Catmi Craft. 所有權利保留。</p>
        <!-- <div class="d-flex justify-content-center gap-2 mt-2">
          <a href="#" class="text-decoration-none"><i class="bi bi-instagram"></i></a>
          <a href="#" class="text-decoration-none"><i class="bi bi-facebook"></i></a>
          <a href="#" class="text-decoration-none"><i class="bi bi-chat-dots"></i></a>
        </div> -->
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
  /* background-color: rgba(255, 255, 255, 0.2); */
  transform: translateY(-2px);
}

.nav-link.active {
  /* background-color: rgba(255, 255, 255, 0.3); */
}

.header-btn-image {
    width: 22px;
    height: 36px;
    object-fit: cover;
}

main {
  display: flex;
  flex-direction: column;
}

.social-icon {
  width: 40px;
  height: 40px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg,#000000 0%,#2e2e2e 100%);
  color: #E8767B;
  border-radius: 12px;
}

.social-icon-img {
  width: 24px;
  height: 24px;
  object-fit: contain;
}

.social-content {
  flex-wrap: wrap;
}

.social-title {
  white-space: nowrap;
}

.social-item .text-muted.small {
  white-space: nowrap;
}
.footer-image{
  
  height: 230px;
  object-fit: cover;
}
</style>
