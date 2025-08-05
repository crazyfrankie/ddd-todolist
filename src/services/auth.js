import api from './api'

export const authService = {
  // User registration
  async register(email, password) {
    const response = await api.post('/user/register', {
      email,
      password
    })
    return response.data
  },

  // User login
  async login(email, password) {
    const response = await api.post('/user/login', {
      email,
      password
    })
    return response.data
  },

  // User logout
  async logout() {
    const response = await api.get('/user/logout')
    localStorage.removeItem('access_token')
    return response.data
  },

  // Get user profile
  async getProfile() {
    const response = await api.get('/user/profile')
    return response.data
  },

  // Update user profile
  async updateProfile(data) {
    const response = await api.put('/user/profile', data)
    return response.data
  },

  // Update user avatar
  async updateAvatar(file) {
    const formData = new FormData()
    formData.append('avatar', file)
    
    const response = await api.put('/user/avatar', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
    return response.data
  },

  // Reset password
  async resetPassword(email, password) {
    const response = await api.post('/user/reset-password', {
      email,
      password
    })
    return response.data
  },

  // Check if user is authenticated
  isAuthenticated() {
    return !!localStorage.getItem('access_token')
  }
}