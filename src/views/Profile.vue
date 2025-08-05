<template>
  <div class="profile-page">
    <div class="container">
      <div class="page-header">
        <h1>个人资料</h1>
        <p>管理您的账户信息和偏好设置</p>
      </div>

      <div class="profile-content">
        <!-- Avatar Section -->
        <div class="profile-section">
          <div class="section-header">
            <h2>头像</h2>
            <p>上传您的个人头像</p>
          </div>
          
          <div class="avatar-section">
            <div class="current-avatar">
              <img 
                v-if="authStore.user?.avatarURL" 
                :src="authStore.user.avatarURL" 
                :alt="authStore.user.name"
                class="avatar-image"
              >
              <div v-else class="avatar-placeholder">
                {{ authStore.user?.name?.charAt(0)?.toUpperCase() || 'U' }}
              </div>
            </div>
            
            <div class="avatar-actions">
              <input
                ref="avatarInput"
                type="file"
                accept="image/*"
                @change="handleAvatarChange"
                style="display: none"
              >
              <button @click="$refs.avatarInput.click()" class="btn btn-primary">
                <Upload :size="16" />
                上传头像
              </button>
              <p class="avatar-hint">支持 JPG、PNG 格式，文件大小不超过 5MB</p>
            </div>
          </div>
        </div>

        <!-- Profile Information -->
        <div class="profile-section">
          <div class="section-header">
            <h2>基本信息</h2>
            <p>更新您的个人信息</p>
          </div>
          
          <form @submit.prevent="handleUpdateProfile" class="profile-form">
            <div class="form-row">
              <div class="form-group">
                <label for="name" class="form-label">姓名</label>
                <input
                  id="name"
                  v-model="profileForm.name"
                  type="text"
                  class="input"
                  placeholder="请输入姓名"
                >
              </div>
              
              <div class="form-group">
                <label for="userUniqueName" class="form-label">用户名</label>
                <input
                  id="userUniqueName"
                  v-model="profileForm.userUniqueName"
                  type="text"
                  class="input"
                  placeholder="请输入用户名"
                >
              </div>
            </div>
            
            <div class="form-group">
              <label for="email" class="form-label">邮箱</label>
              <input
                id="email"
                :value="authStore.user?.email"
                type="email"
                class="input"
                disabled
              >
              <p class="form-hint">邮箱地址不可修改</p>
            </div>
            
            <div v-if="authStore.error" class="error-message">
              {{ authStore.error }}
            </div>
            
            <div class="form-actions">
              <button type="submit" class="btn btn-primary" :disabled="authStore.loading">
                <div v-if="authStore.loading" class="spinner"></div>
                保存更改
              </button>
            </div>
          </form>
        </div>

        <!-- Password Section -->
        <div class="profile-section">
          <div class="section-header">
            <h2>密码</h2>
            <p>更改您的登录密码</p>
          </div>
          
          <form @submit.prevent="handleResetPassword" class="profile-form">
            <div class="form-group">
              <label for="newPassword" class="form-label">新密码</label>
              <div class="password-input">
                <input
                  id="newPassword"
                  v-model="passwordForm.newPassword"
                  :type="showNewPassword ? 'text' : 'password'"
                  class="input"
                  placeholder="请输入新密码（至少6位）"
                  minlength="6"
                >
                <button
                  type="button"
                  @click="showNewPassword = !showNewPassword"
                  class="password-toggle"
                >
                  <Eye v-if="!showNewPassword" :size="18" />
                  <EyeOff v-else :size="18" />
                </button>
              </div>
            </div>
            
            <div class="form-group">
              <label for="confirmNewPassword" class="form-label">确认新密码</label>
              <div class="password-input">
                <input
                  id="confirmNewPassword"
                  v-model="passwordForm.confirmNewPassword"
                  :type="showConfirmPassword ? 'text' : 'password'"
                  class="input"
                  placeholder="请再次输入新密码"
                >
                <button
                  type="button"
                  @click="showConfirmPassword = !showConfirmPassword"
                  class="password-toggle"
                >
                  <Eye v-if="!showConfirmPassword" :size="18" />
                  <EyeOff v-else :size="18" />
                </button>
              </div>
            </div>
            
            <div v-if="passwordValidationError" class="error-message">
              {{ passwordValidationError }}
            </div>
            
            <div class="form-actions">
              <button 
                type="submit" 
                class="btn btn-primary" 
                :disabled="authStore.loading || !passwordForm.newPassword"
              >
                <div v-if="authStore.loading" class="spinner"></div>
                更新密码
              </button>
            </div>
          </form>
        </div>

        <!-- Account Stats -->
        <div class="profile-section">
          <div class="section-header">
            <h2>账户统计</h2>
            <p>您的使用情况概览</p>
          </div>
          
          <div class="stats-grid">
            <div class="stat-item">
              <div class="stat-icon">
                <CheckCircle :size="24" />
              </div>
              <div class="stat-content">
                <div class="stat-number">{{ taskStore.completedTasks.length }}</div>
                <div class="stat-label">已完成任务</div>
              </div>
            </div>
            
            <div class="stat-item">
              <div class="stat-icon">
                <Clock :size="24" />
              </div>
              <div class="stat-content">
                <div class="stat-number">{{ taskStore.pendingTasks.length }}</div>
                <div class="stat-label">待完成任务</div>
              </div>
            </div>
            
            <div class="stat-item">
              <div class="stat-icon">
                <Calendar :size="24" />
              </div>
              <div class="stat-content">
                <div class="stat-number">{{ formatJoinDate(authStore.user?.userCreateTime) }}</div>
                <div class="stat-label">加入时间</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useAuthStore } from '../stores/auth'
import { useTaskStore } from '../stores/tasks'
import { authService } from '../services/auth'
import { 
  Upload, 
  Eye, 
  EyeOff, 
  CheckCircle, 
  Clock, 
  Calendar 
} from 'lucide-vue-next'

const authStore = useAuthStore()
const taskStore = useTaskStore()

const showNewPassword = ref(false)
const showConfirmPassword = ref(false)

const profileForm = reactive({
  name: '',
  userUniqueName: ''
})

const passwordForm = reactive({
  newPassword: '',
  confirmNewPassword: ''
})

const passwordValidationError = computed(() => {
  if (passwordForm.newPassword && passwordForm.confirmNewPassword && 
      passwordForm.newPassword !== passwordForm.confirmNewPassword) {
    return '两次输入的密码不一致'
  }
  return null
})

function initProfileForm() {
  if (authStore.user) {
    profileForm.name = authStore.user.name || ''
    profileForm.userUniqueName = authStore.user.user_unique_name || ''
  }
}

async function handleAvatarChange(event) {
  const file = event.target.files[0]
  if (!file) return
  
  // Validate file size (5MB)
  if (file.size > 5 * 1024 * 1024) {
    alert('文件大小不能超过 5MB')
    return
  }
  
  // Validate file type
  if (!file.type.startsWith('image/')) {
    alert('请选择图片文件')
    return
  }
  
  try {
    await authStore.updateAvatar(file)
    alert('头像更新成功')
  } catch (error) {
    alert('头像更新失败：' + (error.message || '未知错误'))
  }
}

async function handleUpdateProfile() {
  try {
    const updateData = {}
    
    if (profileForm.name !== authStore.user?.name) {
      updateData.name = profileForm.name
    }
    
    if (profileForm.userUniqueName !== authStore.user?.user_unique_name) {
      updateData.userUniqueName = profileForm.userUniqueName
    }
    
    if (Object.keys(updateData).length > 0) {
      await authStore.updateProfile(updateData)
      alert('个人信息更新成功')
    }
  } catch (error) {
    alert('更新失败：' + (error.message || '未知错误'))
  }
}

async function handleResetPassword() {
  if (passwordValidationError.value) {
    return
  }
  
  try {
    await authService.resetPassword(authStore.user.email, passwordForm.newPassword)
    passwordForm.newPassword = ''
    passwordForm.confirmNewPassword = ''
    alert('密码更新成功')
  } catch (error) {
    alert('密码更新失败：' + (error.response?.data?.message || '未知错误'))
  }
}

function formatJoinDate(timestamp) {
  if (!timestamp) return '-'
  
  const date = new Date(timestamp)
  return date.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'short'
  })
}

onMounted(() => {
  initProfileForm()
  taskStore.fetchTasks()
})
</script>

<style scoped>
.profile-page {
  min-height: calc(100vh - 64px);
  padding: var(--spacing-xl) 0;
}

.page-header {
  margin-bottom: var(--spacing-2xl);
}

.page-header h1 {
  font-size: var(--font-size-3xl);
  font-weight: 700;
  color: var(--color-gray-900);
  margin: 0 0 var(--spacing-sm) 0;
}

.page-header p {
  color: var(--color-gray-600);
  margin: 0;
}

.profile-content {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-2xl);
}

.profile-section {
  background-color: var(--color-white);
  border-radius: var(--radius-xl);
  border: 1px solid var(--color-gray-200);
  overflow: hidden;
}

.section-header {
  padding: var(--spacing-xl);
  border-bottom: 1px solid var(--color-gray-200);
  background-color: var(--color-gray-50);
}

.section-header h2 {
  font-size: var(--font-size-xl);
  font-weight: 600;
  color: var(--color-gray-900);
  margin: 0 0 var(--spacing-xs) 0;
}

.section-header p {
  color: var(--color-gray-600);
  margin: 0;
  font-size: var(--font-size-sm);
}

.avatar-section {
  padding: var(--spacing-xl);
  display: flex;
  align-items: center;
  gap: var(--spacing-xl);
}

.current-avatar {
  flex-shrink: 0;
}

.avatar-image {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  object-fit: cover;
  border: 3px solid var(--color-gray-200);
}

.avatar-placeholder {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background-color: var(--color-primary);
  color: var(--color-white);
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: var(--font-size-xl);
  border: 3px solid var(--color-gray-200);
}

.avatar-actions {
  flex: 1;
}

.avatar-hint {
  margin-top: var(--spacing-sm);
  font-size: var(--font-size-xs);
  color: var(--color-gray-500);
}

.profile-form {
  padding: var(--spacing-xl);
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--spacing-lg);
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

.form-hint {
  margin-top: var(--spacing-sm);
  font-size: var(--font-size-xs);
  color: var(--color-gray-500);
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

.form-actions {
  display: flex;
  justify-content: flex-end;
  margin-top: var(--spacing-xl);
  padding-top: var(--spacing-lg);
  border-top: 1px solid var(--color-gray-200);
}

.stats-grid {
  padding: var(--spacing-xl);
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-lg);
}

.stat-item {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  padding: var(--spacing-lg);
  background-color: var(--color-gray-50);
  border-radius: var(--radius-lg);
  border: 1px solid var(--color-gray-200);
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: var(--radius-lg);
  background: linear-gradient(135deg, var(--color-primary) 0%, var(--color-info) 100%);
  color: var(--color-white);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.stat-number {
  font-size: var(--font-size-xl);
  font-weight: 700;
  color: var(--color-gray-900);
  line-height: 1;
}

.stat-label {
  font-size: var(--font-size-sm);
  color: var(--color-gray-600);
  margin-top: var(--spacing-xs);
}

@media (max-width: 768px) {
  .avatar-section {
    flex-direction: column;
    text-align: center;
  }
  
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .stats-grid {
    grid-template-columns: 1fr;
  }
  
  .form-actions {
    justify-content: stretch;
  }
  
  .form-actions .btn {
    width: 100%;
  }
}
</style>