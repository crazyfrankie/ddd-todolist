import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { taskService } from '../services/tasks'

export const useTaskStore = defineStore('tasks', () => {
  const tasks = ref([])
  const loading = ref(false)
  const error = ref(null)

  // Computed properties
  const completedTasks = computed(() => 
    tasks.value.filter(task => task.isCompleted)
  )
  
  const pendingTasks = computed(() => 
    tasks.value.filter(task => !task.isCompleted)
  )

  const tasksByPriority = computed(() => {
    const groups = {
      'important and urgent': [],
      'important but not urgent': [],
      'not important but urgent': [],
      'neither important or urgent': []
    }
    
    pendingTasks.value.forEach(task => {
      const priority = task.priority || 'neither important or urgent'
      if (groups[priority]) {
        groups[priority].push(task)
      }
    })
    
    return groups
  })

  // Fetch all tasks
  async function fetchTasks() {
    loading.value = true
    error.value = null
    
    try {
      const response = await taskService.getTasks()
      tasks.value = response.data || []
    } catch (err) {
      error.value = err.response?.data?.message || '获取任务失败'
      console.error('Fetch tasks error:', err)
    } finally {
      loading.value = false
    }
  }

  // Create new task
  async function createTask(taskData) {
    loading.value = true
    error.value = null
    
    try {
      const response = await taskService.createTask(taskData)
      const newTask = response.data
      tasks.value.unshift(newTask)
      return newTask
    } catch (err) {
      error.value = err.response?.data?.message || '创建任务失败'
      throw err
    } finally {
      loading.value = false
    }
  }

  // Update task
  async function updateTask(taskData) {
    loading.value = true
    error.value = null
    
    try {
      await taskService.updateTask(taskData)
      
      // Update local task
      const index = tasks.value.findIndex(task => task.id === taskData.id)
      if (index !== -1) {
        tasks.value[index] = { ...tasks.value[index], ...taskData }
      }
    } catch (err) {
      error.value = err.response?.data?.message || '更新任务失败'
      throw err
    } finally {
      loading.value = false
    }
  }

  // Delete task
  async function deleteTask(taskId) {
    loading.value = true
    error.value = null
    
    try {
      await taskService.deleteTask(taskId)
      
      // Remove from local tasks
      const index = tasks.value.findIndex(task => task.id === taskId)
      if (index !== -1) {
        tasks.value.splice(index, 1)
      }
    } catch (err) {
      error.value = err.response?.data?.message || '删除任务失败'
      throw err
    } finally {
      loading.value = false
    }
  }

  // Toggle task completion
  async function toggleTaskCompletion(taskId) {
    const task = tasks.value.find(t => t.id === taskId)
    if (!task) return
    
    const updatedTask = {
      ...task,
      isCompleted: !task.isCompleted
    }
    
    await updateTask(updatedTask)
  }

  // Clear error
  function clearError() {
    error.value = null
  }

  // Clear all tasks (for logout)
  function clearTasks() {
    tasks.value = []
  }

  return {
    tasks,
    loading,
    error,
    completedTasks,
    pendingTasks,
    tasksByPriority,
    fetchTasks,
    createTask,
    updateTask,
    deleteTask,
    toggleTaskCompletion,
    clearError,
    clearTasks
  }
})