import axios from 'axios'

// Create axios instance
const api = axios.create({
  baseURL: 'http://localhost:8088/api',
  timeout: 10000,
  withCredentials: true
})

// Request interceptor
api.interceptors.request.use(
  (config) => {
    // Add access token to headers
    const token = localStorage.getItem('access_token')
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// Response interceptor
api.interceptors.response.use(
  (response) => {
    // Store access token from response headers
    const accessToken = response.headers['x-access-token']
    if (accessToken) {
      localStorage.setItem('access_token', accessToken)
    }
    return response
  },
  (error) => {
    // Handle 401 errors
    if (error.response?.status === 401) {
      localStorage.removeItem('access_token')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export default api