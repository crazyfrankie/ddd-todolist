<template>
  <div class="task-card" :class="{ 'task-completed': task.isCompleted }">
    <div class="task-header">
      <div class="task-checkbox">
        <input 
          type="checkbox" 
          :checked="task.isCompleted"
          @change="$emit('toggle', task.id)"
          class="checkbox"
        >
      </div>
      
      <div class="task-priority">
        <span 
          class="priority-indicator" 
          :class="getPriorityClass(task.priority)"
          :title="getPriorityLabel(task.priority)"
        ></span>
      </div>
      
      <div class="task-actions">
        <button 
          @click="$emit('edit', task)"
          class="btn btn-ghost btn-sm"
          title="编辑"
        >
          <Edit2 :size="16" />
        </button>
        <button 
          @click="$emit('delete', task.id)"
          class="btn btn-ghost btn-sm text-error"
          title="删除"
        >
          <Trash2 :size="16" />
        </button>
      </div>
    </div>
    
    <div class="task-content">
      <p class="task-text" :class="{ 'completed': task.isCompleted }">
        {{ task.content }}
      </p>
      
      <div class="task-meta" v-if="task.dueTime">
        <div class="task-due-time" :class="getDueTimeClass(task.dueTime)">
          <Calendar :size="14" />
          <span>{{ formatDueTime(task.dueTime) }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Edit2, Trash2, Calendar } from 'lucide-vue-next'
import { taskService } from '../services/tasks'

defineProps({
  task: {
    type: Object,
    required: true
  }
})

defineEmits(['toggle', 'edit', 'delete'])

function getPriorityLabel(priority) {
  return taskService.getPriorityLabel(priority)
}

function getPriorityClass(priority) {
  return taskService.getPriorityClass(priority)
}

function formatDueTime(timestamp) {
  if (!timestamp) return ''
  
  const date = new Date(timestamp)
  const now = new Date()
  const today = new Date(now.getFullYear(), now.getMonth(), now.getDate())
  const taskDate = new Date(date.getFullYear(), date.getMonth(), date.getDate())
  
  const diffTime = taskDate.getTime() - today.getTime()
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  
  if (diffDays === 0) {
    return '今天'
  } else if (diffDays === 1) {
    return '明天'
  } else if (diffDays === -1) {
    return '昨天'
  } else if (diffDays > 1 && diffDays <= 7) {
    return `${diffDays}天后`
  } else if (diffDays < -1 && diffDays >= -7) {
    return `${Math.abs(diffDays)}天前`
  } else {
    return date.toLocaleDateString('zh-CN', {
      month: 'short',
      day: 'numeric'
    })
  }
}

function getDueTimeClass(timestamp) {
  if (!timestamp) return ''
  
  const date = new Date(timestamp)
  const now = new Date()
  const diffTime = date.getTime() - now.getTime()
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24))
  
  if (diffDays < 0) {
    return 'overdue'
  } else if (diffDays === 0) {
    return 'due-today'
  } else if (diffDays <= 3) {
    return 'due-soon'
  }
  
  return ''
}
</script>

<style scoped>
.task-card {
  background-color: var(--color-white);
  border: 1px solid var(--color-gray-200);
  border-radius: var(--radius-lg);
  padding: var(--spacing-md);
  transition: all var(--transition-fast);
  cursor: pointer;
}

.task-card:hover {
  border-color: var(--color-primary-light);
  box-shadow: var(--shadow-lg);
  transform: translateY(-2px);
}

.task-completed {
  opacity: 0.7;
  background-color: var(--color-gray-50);
}

.task-header {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  margin-bottom: var(--spacing-sm);
}

.task-checkbox {
  flex-shrink: 0;
}

.checkbox {
  width: 18px;
  height: 18px;
  border-radius: var(--radius-sm);
  border: 2px solid var(--color-gray-300);
  cursor: pointer;
  transition: all var(--transition-fast);
}

.checkbox:checked {
  background: linear-gradient(135deg, var(--color-success) 0%, var(--color-success-dark) 100%);
  border-color: var(--color-success);
}

.task-priority {
  flex-shrink: 0;
}

.task-actions {
  margin-left: auto;
  display: flex;
  gap: var(--spacing-xs);
  opacity: 0;
  transition: opacity var(--transition-fast);
}

.task-card:hover .task-actions {
  opacity: 1;
}

.text-error {
  color: var(--color-error);
}

.text-error:hover {
  background-color: rgb(239 68 68 / 0.1);
}

.task-content {
  margin-left: calc(18px + var(--spacing-sm) + 8px + var(--spacing-sm));
}

.task-text {
  margin: 0 0 var(--spacing-sm) 0;
  line-height: 1.5;
  color: var(--color-gray-900);
  word-wrap: break-word;
}

.task-text.completed {
  text-decoration: line-through;
  color: var(--color-gray-500);
}

.task-meta {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.task-due-time {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  font-size: var(--font-size-xs);
  color: var(--color-gray-500);
  padding: var(--spacing-xs) var(--spacing-sm);
  border-radius: var(--radius-sm);
  background-color: var(--color-gray-100);
}

.task-due-time.due-today {
  background-color: rgb(59 130 246 / 0.1);
  color: var(--color-info);
}

.task-due-time.due-soon {
  background-color: rgb(245 158 11 / 0.1);
  color: var(--color-warning);
}

.task-due-time.overdue {
  background-color: rgb(239 68 68 / 0.1);
  color: var(--color-error);
}

@media (max-width: 768px) {
  .task-actions {
    opacity: 1;
  }
  
  .task-card {
    padding: var(--spacing-sm);
  }
}
</style>