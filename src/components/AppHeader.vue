<template>
  <header class="app-header">
    <div class="container">
      <div class="header-content">
        <!-- Logo -->
        <div class="logo">
          <h1>TodoList</h1>
        </div>

        <!-- Navigation -->
        <nav class="nav" v-if="isAuthenticated">
          <router-link to="/" class="nav-link" active-class="nav-link-active">
            <CheckSquare :size="18" />
            任务
          </router-link>
          <router-link to="/profile" class="nav-link" active-class="nav-link-active">
            <User :size="18" />
            个人资料
          </router-link>
        </nav>

        <!-- User menu -->
        <div class="user-menu" v-if="isAuthenticated">
          <div class="user-info" @click="toggleUserMenu">
            <img 
              v-if="user?.avatarURL" 
              :src="user.avatarURL" 
              :alt="user.name"
              class="user-avatar"
            >
            <div v-else class="user-avatar-placeholder">
              {{ user?.name?.charAt(0)?.toUpperCase() || 'U' }}
            </div>
            <span class="user-name">{{ user?.name || 'User' }}</span>
            <ChevronDown :size="16" />
          </div>
          
          <div v-if="showUserMenu" class="user-dropdown" @click="showUserMenu = false">
            <router-link to="/profile" class="dropdown-item">
              <Settings :size="16" />
              设置
            </router-link>
            <button @click="handleLogout" class="dropdown-item">
              <LogOut :size="16" />
              退出登录
            </button>
          </div>
        </div>

        <!-- Auth buttons for guests -->
        <div class="auth-buttons" v-else>
          <router-link to="/login" class="btn btn-ghost">登录</router-link>
          <router-link to="/register" class="btn btn-primary">注册</router-link>
        </div>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { CheckSquare, User, ChevronDown, Settings, LogOut } from 'lucide-vue-next'

const router = useRouter()
const authStore = useAuthStore()

const showUserMenu = ref(false)

const isAuthenticated = computed(() => authStore.isAuthenticated)
const user = computed(() => authStore.user)

function toggleUserMenu() {
  showUserMenu.value = !showUserMenu.value
}

async function handleLogout() {
  try {
    await authStore.logout()
    router.push('/login')
  } catch (error) {
    console.error('Logout error:', error)
  }
}

// Close user menu when clicking outside
function handleClickOutside(event) {
  if (!event.target.closest('.user-menu')) {
    showUserMenu.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.app-header {
  background-color: var(--color-white);
  border-bottom: 1px solid var(--color-gray-200);
  box-shadow: var(--shadow-sm);
  position: sticky;
  top: 0;
  z-index: 100;
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 64px;
}

.logo h1 {
  font-size: var(--font-size-xl);
  font-weight: 700;
  color: var(--color-primary);
  margin: 0;
}

.nav {
  display: flex;
  gap: var(--spacing-lg);
}

.nav-link {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-sm) var(--spacing-md);
  border-radius: var(--radius-md);
  text-decoration: none;
  color: var(--color-gray-600);
  font-weight: 500;
  transition: all var(--transition-fast);
}

.nav-link:hover {
  color: var(--color-gray-900);
  background-color: var(--color-gray-100);
}

.nav-link-active {
  color: var(--color-primary);
  background-color: var(--color-primary-light);
  background-color: rgb(99 102 241 / 0.1);
}

.user-menu {
  position: relative;
}

.user-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-sm);
  border-radius: var(--radius-md);
  cursor: pointer;
  transition: all var(--transition-fast);
}

.user-info:hover {
  background-color: var(--color-gray-100);
}

.user-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  object-fit: cover;
}

.user-avatar-placeholder {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: var(--color-primary);
  color: var(--color-white);
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: var(--font-size-sm);
}

.user-name {
  font-weight: 500;
  color: var(--color-gray-900);
}

.user-dropdown {
  position: absolute;
  top: 100%;
  right: 0;
  margin-top: var(--spacing-sm);
  background-color: var(--color-white);
  border: 1px solid var(--color-gray-200);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-lg);
  min-width: 180px;
  overflow: hidden;
  z-index: 50;
}

.dropdown-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  width: 100%;
  padding: var(--spacing-md);
  border: none;
  background: none;
  text-align: left;
  text-decoration: none;
  color: var(--color-gray-700);
  font-size: var(--font-size-sm);
  cursor: pointer;
  transition: all var(--transition-fast);
}

.dropdown-item:hover {
  background-color: var(--color-gray-50);
  color: var(--color-gray-900);
}

.auth-buttons {
  display: flex;
  gap: var(--spacing-sm);
}

@media (max-width: 768px) {
  .nav {
    display: none;
  }
  
  .user-name {
    display: none;
  }
  
  .logo h1 {
    font-size: var(--font-size-lg);
  }
}
</style>