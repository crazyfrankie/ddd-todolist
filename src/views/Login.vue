<template>
  <div class="auth-page">
    <div class="auth-container">
      <div class="auth-card">
        <div class="auth-header">
          <h1>登录</h1>
          <p>欢迎回来，请登录您的账户</p>
        </div>
        
        <form @submit.prevent="handleLogin" class="auth-form">
          <div class="form-group">
            <label for="email" class="form-label">邮箱</label>
            <input
              id="email"
              v-model="form.email"
              type="email"
              class="input"
              placeholder="请输入邮箱"
              required
            >
          </div>
          
          <div class="form-group">
            <label for="password" class="form-label">密码</label>
            <div class="password-input">
              <input
                id="password"
                v-model="form.password"
                :type="showPassword ? 'text' : 'password'"
                class="input"
                placeholder="请输入密码"
                required
              >
              <button
                type="button"
                @click="showPassword = !showPassword"
                class="password-toggle"
              >
                <Eye v-if="!showPassword" :size="18" />
                <EyeOff v-else :size="18" />
              </button>
            </div>
          </div>
          
          <div v-if="authStore.error" class="error-message">
            {{ authStore.error }}
          </div>
          
          <button type="submit" class="btn btn-primary btn-lg auth-submit" :disabled="authStore.loading">
            <div v-if="authStore.loading" class="spinner"></div>
            登录
          </button>
        </form>
        
        <div class="auth-footer">
          <p>
            还没有账户？
            <router-link to="/register" class="auth-link">立即注册</router-link>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { Eye, EyeOff } from 'lucide-vue-next'

const router = useRouter()
const authStore = useAuthStore()

const showPassword = ref(false)

const form = reactive({
  email: '',
  password: ''
})

async function handleLogin() {
  try {
    await authStore.login(form.email, form.password)
    router.push('/')
  } catch (error) {
    // Error is handled by the store
  }
}

onMounted(() => {
  authStore.clearError()
})
</script>

<style scoped>
.auth-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--color-white);
  padding: var(--spacing-md);
}

.auth-container {
  width: 100%;
  max-width: 400px;
}

.auth-card {
  background-color: var(--color-white);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-xl);
  border: 1px solid var(--color-gray-200);
  padding: var(--spacing-2xl);
}

.auth-header {
  text-align: center;
  margin-bottom: var(--spacing-2xl);
}

.auth-header h1 {
  font-size: var(--font-size-3xl);
  font-weight: 700;
  color: var(--color-gray-900);
  margin-bottom: var(--spacing-sm);
}

.auth-header p {
  color: var(--color-gray-600);
  font-size: var(--font-size-base);
}

.auth-form {
  margin-bottom: var(--spacing-xl);
}

.form-group {
  margin-bottom: var(--spacing-lg);
}

.form-label {
  display: block;
  margin-bottom: var(--spacing-sm);
  font-weight: 500;
  color: var(--color-gray-700);
  font-size: var(--font-size-sm);
}

.password-input {
  position: relative;
}

.password-toggle {
  position: absolute;
  right: var(--spacing-md);
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  color: var(--color-gray-400);
  cursor: pointer;
  padding: var(--spacing-xs);
  border-radius: var(--radius-sm);
  transition: color var(--transition-fast);
}

.password-toggle:hover {
  color: var(--color-gray-600);
}

.error-message {
  background-color: rgb(239 68 68 / 0.1);
  color: var(--color-error);
  padding: var(--spacing-md);
  border-radius: var(--radius-md);
  font-size: var(--font-size-sm);
  margin-bottom: var(--spacing-lg);
  border: 1px solid rgb(239 68 68 / 0.2);
}

.auth-submit {
  width: 100%;
  margin-bottom: var(--spacing-lg);
}

.auth-footer {
  text-align: center;
  padding-top: var(--spacing-lg);
  border-top: 1px solid var(--color-gray-200);
}

.auth-footer p {
  color: var(--color-gray-600);
  font-size: var(--font-size-sm);
}

.auth-link {
  color: var(--color-primary);
  text-decoration: none;
  font-weight: 500;
  transition: color var(--transition-fast);
}

.auth-link:hover {
  color: var(--color-primary-dark);
  text-decoration: underline;
}

@media (max-width: 768px) {
  .auth-card {
    padding: var(--spacing-xl);
  }
  
  .auth-page {
    padding: var(--spacing-sm);
  }
}
</style>