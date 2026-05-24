<template>
  <div class="home-page">
    <!-- 英雄區域 -->
    <section class="hero-section mb-5">
      <div class="container-lg">
        <div class="row align-items-center">
          <div class="col-lg-6 mb-4 mb-lg-0">
            <h1 class="display-4 fw-bold mb-3">
              <span style="color: #E8767B;">歡迎來到</span><br>
              <span style="color: #6B4A47;">Catmi Craft</span>
            </h1>
            <p class="lead text-secondary mb-4">
              用創意把日常過成喜歡的樣子<br>
              <span class="text-muted">一起慢慢生活，找到喜歡的自己</span>
            </p>
            <button  class="btn btn-primary btn-lg">
              一起慢慢生活<i class="ms-2 bi bi-arrow-right me-2"></i>
            </button>
          </div>
          <div class="col-lg-6">
            <div class="hero-image text-center">
              <div class="p-5" style="background: linear-gradient(135deg, #FFE4E8 0%, #F5F1E8 100%); border-radius: 20px;">
                <i class="bi bi-heart" style="font-size: 8rem; color: #E8767B; opacity: 0.3;"></i>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- 特色區域 -->
    <section class="features-section mb-5">
      <div class="container-lg">
        <h2 class="text-center mb-5 fw-bold">
          <span style="color: #E8767B;">Catmi</span> 的故事
        </h2>
        <div class="row g-4">
          <div class="col-md-6 col-lg-3">
            <div class="feature-card">
              <div class="feature-icon mb-3">
                <img src="../images/catClaws.png" alt="品牌故事" class="feature-icon-image" />
              </div>
              <h5 class="fw-bold"><p>
              走進<br>
              <span class="text-muted">Catmi 的世界</span>
            </p> </h5>
              <p class="text-secondary">認識 Catmi 的初心與品牌故事。</p>
              <div class="feature-image mb-3">
                <img src="../images/catClaws.png" alt="品牌故事" class="feature-rect" />
              </div>
              <div><button  class="featurebtn btn btn-primary btn-lg">
              <i class="ms-2 bi bi-arrow-right me-2"></i>
            </button>
              </div>
            </div>
          </div>
          <div class="col-md-6 col-lg-3">
            <div class="feature-card">
              <div class="feature-icon mb-3">
                <img src="../images/catClaws.png" alt="手工產品" class="feature-icon-image" />
              </div>
              <h5 class="fw-bold"><p>
              留下<br>
              <span class="text-muted">一點喜歡的日常</span>
            </p> </h5>
              <p class="text-secondary">欣賞作品與手作商品，收藏生活的美好。</p>
              <div class="feature-image mb-3">
                <img src="../images/catClaws.png" alt="品牌故事" class="feature-rect" />
              </div>
              <div><button  class="featurebtn btn btn-primary btn-lg">
              <i class="ms-2 bi bi-arrow-right me-2"></i>
            </button>
              </div>
            </div>
          </div>
          <div class="col-md-6 col-lg-3">
            <div class="feature-card">
              <div class="feature-icon mb-3">
                <img src="../images/catClaws.png" alt="溫柔陪伴" class="feature-icon-image" />
              </div>
              <h5 class="fw-bold"><p>
              開始<br>
              <span class="text-muted">一段療癒時光</span>
            </p> </h5>
              <p class="text-secondary">探索手作課程，體驗慢下來的幸福。</p>
              <div class="feature-image mb-3">
                <img src="../images/catClaws.png" alt="品牌故事" class="feature-rect" />
              </div>
              <div><button  class="featurebtn btn btn-primary btn-lg">
              <i class="ms-2 bi bi-arrow-right me-2"></i>
            </button>
              </div>
            </div>
          </div>
          <div class="col-md-6 col-lg-3">
            <div class="feature-card">
              <div class="feature-icon mb-3">
                <img src="../images/catClaws.png" alt="生活美學" class="feature-icon-image" />
              </div>
              <h5 class="fw-bold"><p>
              聆聽<br>
              <span class="text-muted">溫柔的陪伴</span>
            </p> </h5>
              <p class="text-secondary">最新消息與日常分享，與你一起慢慢生活。</p>
              <div class="feature-image mb-3">
                <img src="../images/catClaws.png" alt="品牌故事" class="feature-rect" />
              </div>
              <div><button  class="featurebtn btn btn-primary btn-lg">
              <i class="ms-2 bi bi-arrow-right me-2"></i>
            </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- CTA 區域 -->
    <!-- <section class="cta-section mb-5">
      <div class="container-lg">
        <div class="row g-4">
          <div class="col-md-6">
            <div class="cta-card text-center">
              <div class="cta-icon mb-3">
                <img src="../images/catClaws.png" alt="走進 Catmi 的世界" class="feature-icon-image" />
              </div>
              <h4 class="fw-bold mb-2">走進 Catmi 的世界</h4>
              <p class="text-secondary">探索獨特的手工創意產品</p>
            </div>
          </div>
          <div class="col-md-6">
            <div class="cta-card text-center">
              <div class="cta-icon mb-3">
                <img src="../images/catClaws.png" alt="加入我們的社群" class="feature-icon-image" />
              </div>
              <h4 class="fw-bold mb-2">加入我們的社群</h4>
              <p class="text-secondary">分享生活故事，互動交流</p>
            </div>
          </div>
        </div>
      </div>
    </section> -->

    
  </div>
</template>

<script setup>
</script>

<style scoped>
.hero-section {
  padding: 3rem 0;
}

.hero-image {
  animation: float 3s ease-in-out infinite;
}

@keyframes float {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-20px);
  }
}

.feature-card {
  /* background: white; */
  padding: 2rem 1.5rem;
  border-radius: 16px;
  text-align: center;
  transition: all 0.3s ease;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.feature-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 8px 24px rgba(232, 118, 123, 0.15);
}

.feature-icon {
  width: 60px;
  height: 60px;
  margin: 0 auto;
  background: linear-gradient(135deg, #FFE4E8 0%, #F5F1E8 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2rem;
  color: #E8767B;
}

.feature-icon-image {
  height: 150%;
  object-fit: contain;
  border-radius: 8px;
}

.feature-image {
  width: 100%;
  height: auto;
  overflow: hidden;
  border-radius: 0; /* rectangle */
  background: transparent; /* remove background color */
  aspect-ratio: 16 / 9;
}

.feature-image img,
.feature-rect {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}

.featurebtn {
  width: 45px;
  height: 45px;
  padding: 0;
  border-radius: 50%;
  margin: 1.5rem auto 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.featurebtn i {
  margin: 0 !important;
}

.feature-card > div {
  text-align: center;
}

.cta-card {
  /* background: white; */
  padding: 2.5rem 2rem;
  border-radius: 16px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease;
}

.cta-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 20px rgba(232, 118, 123, 0.15);
}

.cta-icon {
  width: 80px;
  height: 80px;
  margin: 0 auto;
  background: linear-gradient(135deg, #FFE4E8 0%, #F5F1E8 100%);
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2.5rem;
  color: #E8767B;
}

.bottom-section {
  background: linear-gradient(135deg, #FFE4E8 0%, #F5F1E8 100%);
  border-radius: 16px;
  margin: 0 1rem;
}

@media (max-width: 768px) {
  .hero-section {
    padding: 2rem 0;
  }

  h1 {
    font-size: 2rem;
  }

  .feature-card {
    padding: 1.5rem 1rem;
  }
}
</style>
