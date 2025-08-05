import api from './api'

export const taskService = {
  // Create a new task
  async createTask(taskData) {
    const response = await api.post('/tasks', {
      content: taskData.content,
      date: taskData.dueTime,
      priority: taskData.priority
    })
    return response.data
  },

  // Get task list
  async getTasks() {
    const response = await api.get('/tasks')
    return response.data
  },

  // Get task detail
  async getTask(taskId) {
    const response = await api.get(`/tasks/${taskId}`)
    return response.data
  },

  // Update task
  async updateTask(taskData) {
    const response = await api.put('/tasks', {
      task_id: taskData.id,
      content: taskData.content,
      priority: taskData.priority,
      date: taskData.dueTime,
      isCompleted: taskData.isCompleted
    })
    return response.data
  },

  // Delete task
  async deleteTask(taskId) {
    const response = await api.delete(`/tasks/${taskId}`)
    return response.data
  },

  // Priority mapping
  getPriorityLabel(priority) {
    const priorityMap = {
      'important and urgent': '重要且紧急',
      'important but not urgent': '重要不紧急',
      'not important but urgent': '不重要但紧急',
      'neither important or urgent': '不重要不紧急'
    }
    return priorityMap[priority] || '普通'
  },

  getPriorityClass(priority) {
    const classMap = {
      'important and urgent': 'priority-urgent-important',
      'important but not urgent': 'priority-important',
      'not important but urgent': 'priority-urgent',
      'neither important or urgent': 'priority-normal'
    }
    return classMap[priority] || 'priority-normal'
  }
}