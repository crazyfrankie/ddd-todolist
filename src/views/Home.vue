<template>
  <div class="home-page">
    <div class="container">
      <!-- Page Header -->
      <div class="page-header">
        <div class="header-content">
          <h1>我的任务</h1>
          <p>管理您的日常任务，提高工作效率</p>
        </div>
        <button @click="showTaskForm = true" class="btn btn-primary">
          <Plus :size="18" />
          新建任务
        </button>
      </div>

      <!-- Task Stats -->
      <div class="task-stats">
        <div class="stat-card">
          <div class="stat-icon pending">
            <Clock :size="24" />
          </div>
          <div class="stat-content">
            <div class="stat-number">{{ taskStore.pendingTasks.length }}</div>
            <div class="stat-label">待完成</div>
          </div>
        </div>
        
        <div class="stat-card">
          <div class="stat-icon completed">
            <CheckCircle :size="24" />
          </div>
          <div class="stat-content">
            <div class="stat-number">{{ taskStore.completedTasks.length }}</div>
            <div class="stat-label">已完成</div>
          </div>
        </div>
        
        <div class="stat-card">
          <div class="stat-icon total">
            <List :size="24" />
          </div>
          <div class="stat-content">
            <div class="stat-number">{{ taskStore.tasks.length }}</div>
            <div class="stat-label">总任务</div>
          </div>
        </div>
      </div>

      <!-- Task Filters -->
      <div class="task-filters">
        <button 
          v-for="filter in filters" 
          :key="filter.key"
          @click="activeFilter = filter.key"
          class="filter-btn"
          :class="{ active: activeFilter === filter.key }"
        >
          <component :is="filter.icon" :size="16" />
          {{ filter.label }}
          <span class="filter-count">{{ getFilterCount(filter.key) }}</span>
        </button>
      </div>

      <!-- Task Lists -->
      <div class="task-content">
        <!-- Priority Matrix View -->
        <div v-if="activeFilter === 'priority'" class="priority-matrix">
          <div 
            v-for="(tasks, priority) in taskStore.tasksByPriority" 
            :key="priority"
            class="priority-section"
            :data-priority="priority"
          >
            <div class="priority-header">
              <span 
                class="priority-indicator" 
                :class="getPriorityClass(priority)"
              ></span>
              <h3>{{ getPriorityLabel(priority) }}</h3>
              <span class="task-count">{{ tasks.length }}</span>
            </div>
            
            <div class="task-list">
              <TaskCard
                v-for="task in tasks"
                :key="task.id"
                :task="task"
                @toggle="handleToggleTask"
                @edit="handleEditTask"
                @delete="handleDeleteTask"
              />
              
              <div v-if="tasks.length === 0" class="empty-state">
                <FileText :size="48" />
                <p>暂无{{ getPriorityLabel(priority) }}任务</p>
              </div>
            </div>
          </div>
        </div>

        <!-- All Tasks View -->
        <div v-else class="task-list-view">
          <div class="task-list">
            <TaskCard
              v-for="task in filteredTasks"
              :key="task.id"
              :task="task"
              @toggle="handleToggleTask"
              @edit="handleEditTask"
              @delete="handleDeleteTask"
            />
            
            <div v-if="filteredTasks.length === 0" class="empty-state">
              <FileText :size="48" />
              <p v-if="taskStore.tasks.length === 0">
                还没有任务，点击"新建任务"开始吧！
              </p>
              <p v-else>
                没有找到符合条件的任务
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- Loading State -->
      <div v-if="taskStore.loading" class="loading-state">
        <div class="spinner"></div>
        <p>加载中...</p>
      </div>
    </div>

    <!-- Task Form Modal -->
    <TaskForm
      v-if="showTaskForm"
      :task="editingTask"
      :loading="taskStore.loading"
      @close="closeTaskForm"
      @submit="handleTaskSubmit"
    />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useTaskStore } from '../stores/tasks'
import { taskService } from '../services/tasks'
import TaskCard from '../components/TaskCard.vue'
import TaskForm from '../components/TaskForm.vue'
import { 
  Plus, 
  Clock, 
  CheckCircle, 
  List, 
  FileText,
  Target
} from 'lucide-vue-next'

const taskStore = useTaskStore()

const showTaskForm = ref(false)
const editingTask = ref(null)
const activeFilter = ref('all')

const filters = [
  { key: 'all', label: '全部', icon: List },
  { key: 'pending', label: '待完成', icon: Clock },
  { key: 'completed', label: '已完成', icon: CheckCircle },
  { key: 'priority', label: '优先级', icon: Target }
]

const filteredTasks = computed(() => {
  switch (activeFilter.value) {
    case 'pending':
      return taskStore.pendingTasks
    case 'completed':
      return taskStore.completedTasks
    case 'all':
    default:
      return taskStore.tasks
  }
})

function getFilterCount(filterKey) {
  switch (filterKey) {
    case 'pending':
      return taskStore.pendingTasks.length
    case 'completed':
      return taskStore.completedTasks.length
    case 'priority':
      return taskStore.pendingTasks.length
    case 'all':
    default:
      return taskStore.tasks.length
  }
}

function getPriorityLabel(priority) {
  return taskService.getPriorityLabel(priority)
}

function getPriorityClass(priority) {
  return taskService.getPriorityClass(priority)
}

async function handleToggleTask(taskId) {
  try {
    await taskStore.toggleTaskCompletion(taskId)
  } catch (error) {
    console.error('Toggle task error:', error)
  }
}

function handleEditTask(task) {
  editingTask.value = task
  showTaskForm.value = true
}

async function handleDeleteTask(taskId) {
  if (confirm('确定要删除这个任务吗？')) {
    try {
      await taskStore.deleteTask(taskId)
    } catch (error) {
      console.error('Delete task error:', error)
    }
  }
}

async function handleTaskSubmit(taskData) {
  try {
    if (editingTask.value) {
      await taskStore.updateTask(taskData)
    } else {
      await taskStore.createTask(taskData)
    }
    closeTaskForm()
  } catch (error) {
    console.error('Task submit error:', error)
  }
}

function closeTaskForm() {
  showTaskForm.value = false
  editingTask.value = null
}

onMounted(async () => {
  await taskStore.fetchTasks()
})
</script>

<style scoped>
.home-page {
  min-height: calc(100vh - 64px);
  padding: var(--spacing-xl) 0;
}

.page-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: var(--spacing-2xl);
}

.header-content h1 {
  font-size: var(--font-size-3xl);
  font-weight: 700;
  color: var(--color-gray-900);
  margin: 0 0 var(--spacing-sm) 0;
}

.header-content p {
  color: var(--color-gray-600);
  margin: 0;
}

.task-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: var(--spacing-lg);
  margin-bottom: var(--spacing-2xl);
}

.stat-card {
  background-color: var(--color-white);
  border-radius: var(--radius-lg);
  padding: var(--spacing-lg);
  box-shadow: var(--shadow-sm);
  border: 1px solid var(--color-gray-200);
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
  transition: all var(--transition-fast);
}

.stat-card:hover {
  box-shadow: var(--shadow-md);
  transform: translateY(-2px);
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: var(--radius-lg);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.stat-icon.pending {
  background: linear-gradient(135deg, var(--color-warning) 0%, var(--color-warning-dark) 100%);
  color: var(--color-white);
}

.stat-icon.completed {
  background: linear-gradient(135deg, var(--color-success) 0%, var(--color-success-dark) 100%);
  color: var(--color-white);
}

.stat-icon.total {
  background: linear-gradient(135deg, var(--color-primary) 0%, var(--color-info) 100%);
  color: var(--color-white);
}

.stat-number {
  font-size: var(--font-size-2xl);
  font-weight: 700;
  color: var(--color-gray-900);
  line-height: 1;
}

.stat-label {
  font-size: var(--font-size-sm);
  color: var(--color-gray-600);
  margin-top: var(--spacing-xs);
}

.task-filters {
  display: flex;
  gap: var(--spacing-sm);
  margin-bottom: var(--spacing-xl);
  overflow-x: auto;
  padding-bottom: var(--spacing-xs);
}

.filter-btn {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-sm) var(--spacing-md);
  border: 1px solid var(--color-gray-300);
  background-color: var(--color-white);
  border-radius: var(--radius-lg);
  font-size: var(--font-size-sm);
  font-weight: 500;
  color: var(--color-gray-600);
  cursor: pointer;
  transition: all var(--transition-fast);
  white-space: nowrap;
}

.filter-btn:hover {
  border-color: var(--color-primary);
  color: var(--color-primary);
}

.filter-btn.active {
  background: linear-gradient(135deg, var(--color-primary) 0%, var(--color-info) 100%);
  border-color: var(--color-primary);
  color: var(--color-white);
  box-shadow: var(--shadow-md);
}

.filter-count {
  background-color: var(--color-gray-200);
  color: var(--color-gray-700);
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-sm);
  font-size: var(--font-size-xs);
  font-weight: 600;
  min-width: 20px;
  text-align: center;
}

.filter-btn.active .filter-count {
  background-color: rgba(255, 255, 255, 0.2);
  color: var(--color-white);
}

.task-content {
  margin-bottom: var(--spacing-2xl);
}

.priority-matrix {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: var(--spacing-xl);
}

.priority-section {
  background-color: var(--color-white);
  border-radius: var(--radius-lg);
  border: 1px solid var(--color-gray-200);
  overflow: hidden;
  transition: all var(--transition-fast);
}

.priority-section:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-lg);
}

.priority-section[data-priority="important and urgent"] {
  background: linear-gradient(135deg, var(--color-priority-urgent-important-bg) 0%, #ffffff 100%);
  border-left: 4px solid var(--color-priority-urgent-important);
}

.priority-section[data-priority="important but not urgent"] {
  background: linear-gradient(135deg, var(--color-priority-important-bg) 0%, #ffffff 100%);
  border-left: 4px solid var(--color-priority-important);
}

.priority-section[data-priority="not important but urgent"] {
  background: linear-gradient(135deg, var(--color-priority-urgent-bg) 0%, #ffffff 100%);
  border-left: 4px solid var(--color-priority-urgent);
}

.priority-section[data-priority="neither important or urgent"] {
  background: linear-gradient(135deg, var(--color-priority-normal-bg) 0%, #ffffff 100%);
  border-left: 4px solid var(--color-priority-normal);
}

.priority-header {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-lg);
  background-color: var(--color-gray-50);
  border-bottom: 1px solid var(--color-gray-200);
}

.priority-header h3 {
  font-size: var(--font-size-lg);
  font-weight: 600;
  color: var(--color-gray-900);
  margin: 0;
  flex: 1;
}

.task-count {
  background-color: var(--color-gray-200);
  color: var(--color-gray-700);
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-sm);
  font-size: var(--font-size-xs);
  font-weight: 600;
}

.task-list {
  padding: var(--spacing-lg);
  display: flex;
  flex-direction: column;
  gap: var(--spacing-md);
  min-height: 200px;
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: var(--spacing-2xl);
  color: var(--color-gray-400);
  text-align: center;
  flex: 1;
}

.empty-state p {
  margin-top: var(--spacing-md);
  font-size: var(--font-size-base);
}

.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: var(--spacing-2xl);
  color: var(--color-gray-500);
}

.loading-state p {
  margin-top: var(--spacing-md);
}

@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--spacing-md);
  }
  
  .task-stats {
    grid-template-columns: 1fr;
  }
  
  .priority-matrix {
    grid-template-columns: 1fr;
  }
  
  .task-filters {
    margin-left: calc(-1 * var(--spacing-md));
    margin-right: calc(-1 * var(--spacing-md));
    padding-left: var(--spacing-md);
    padding-right: var(--spacing-md);
  }
}
</style>