import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authService } from '../services/auth'

export const useAuthStore = defineStore('auth', () => {
  const user = ref(null)
  const loading = ref(false)
  const error = ref(null)

  const isAuthenticated = computed(() => !!user.value)

  // Login
  async function login(email, password) {
    loading.value = true
    error.value = null
    
    try {
      const response = await authService.login(email, password)
      user.value = response.data
      return response
    } catch (err) {
      error.value = err.response?.data?.message || '登录失败'
      throw err
    } finally {
      loading.value = false
    }
  }

  // Register
  async function register(email, password) {
    loading.value = true
    error.value = null
    
    try {
      const response = await authService.register(email, password)
      user.value = response.data
      return response
    } catch (err) {
      error.value = err.response?.data?.message || '注册失败'
      throw err
    } finally {
      loading.value = false
    }
  }

  // Logout
  async function logout() {
    loading.value = true
    
    try {
      await authService.logout()
    } catch (err) {
      console.error('Logout error:', err)
    } finally {
      user.value = null
      loading.value = false
    }
  }

  // Get user profile
  async function fetchProfile() {
    if (!authService.isAuthenticated()) return
    
    loading.value = true
    
    try {
      const response = await authService.getProfile()
      user.value = response.data
    } catch (err) {
      console.error('Fetch profile error:', err)
      if (err.response?.status === 401) {
        user.value = null
      }
    } finally {
      loading.value = false
    }
  }

  // Update profile
  async function updateProfile(data) {
    loading.value = true
    error.value = null
    
    try {
      await authService.updateProfile(data)
      await fetchProfile() // Refresh user data
    } catch (err) {
      error.value = err.response?.data?.message || '更新失败'
      throw err
    } finally {
      loading.value = false
    }
  }

  // Update avatar
  async function updateAvatar(file) {
    loading.value = true
    error.value = null
    
    try {
      const response = await authService.updateAvatar(file)
      await fetchProfile() // Refresh user data
      return response
    } catch (err) {
      error.value = err.response?.data?.message || '头像更新失败'
      throw err
    } finally {
      loading.value = false
    }
  }

  // Clear error
  function clearError() {
    error.value = null
  }

  return {
    user,
    loading,
    error,
    isAuthenticated,
    login,
    register,
    logout,
    fetchProfile,
    updateProfile,
    updateAvatar,
    clearError
  }
})