<template>
  <div class="task-form-overlay" @click="handleOverlayClick">
    <div class="task-form-modal" @click.stop>
      <div class="modal-header">
        <h3>{{ isEditing ? '编辑任务' : '新建任务' }}</h3>
        <button @click="$emit('close')" class="btn btn-ghost btn-sm">
          <X :size="20" />
        </button>
      </div>
      
      <form @submit.prevent="handleSubmit" class="task-form">
        <div class="form-group">
          <label for="content" class="form-label">任务内容</label>
          <textarea
            id="content"
            v-model="form.content"
            class="input task-textarea"
            placeholder="输入任务内容..."
            rows="3"
            required
          ></textarea>
        </div>
        
        <div class="form-row">
          <div class="form-group">
            <label for="priority" class="form-label">优先级</label>
            <select id="priority" v-model="form.priority" class="input">
              <option value="neither important or urgent">不重要不紧急</option>
              <option value="not important but urgent">不重要但紧急</option>
              <option value="important but not urgent">重要不紧急</option>
              <option value="important and urgent">重要且紧急</option>
            </select>
          </div>
          
          <div class="form-group">
            <label for="dueTime" class="form-label">截止时间</label>
            <input
              id="dueTime"
              v-model="form.dueTime"
              type="datetime-local"
              class="input"
            >
          </div>
        </div>
        
        <div class="form-actions">
          <button type="button" @click="$emit('close')" class="btn btn-secondary">
            取消
          </button>
          <button type="submit" class="btn btn-primary" :disabled="loading">
            <div v-if="loading" class="spinner"></div>
            {{ isEditing ? '更新' : '创建' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, watch, nextTick } from 'vue'
import { X } from 'lucide-vue-next'

const props = defineProps({
  task: {
    type: Object,
    default: null
  },
  loading: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['close', 'submit'])

const isEditing = ref(!!props.task)

const form = reactive({
  content: '',
  priority: 'neither important or urgent',
  dueTime: ''
})

// Initialize form data
function initForm() {
  if (props.task) {
    form.content = props.task.content || ''
    form.priority = props.task.priority || 'neither important or urgent'
    form.dueTime = props.task.dueTime ? formatDateTimeLocal(props.task.dueTime) : ''
  } else {
    form.content = ''
    form.priority = 'neither important or urgent'
    form.dueTime = ''
  }
}

function formatDateTimeLocal(timestamp) {
  if (!timestamp) return ''
  const date = new Date(timestamp)
  return date.toISOString().slice(0, 16)
}

function handleSubmit() {
  const taskData = {
    content: form.content.trim(),
    priority: form.priority,
    dueTime: form.dueTime ? new Date(form.dueTime).getTime() : null
  }
  
  if (isEditing.value) {
    taskData.id = props.task.id
    taskData.isCompleted = props.task.isCompleted
  }
  
  emit('submit', taskData)
}

function handleOverlayClick() {
  emit('close')
}

// Watch for task prop changes
watch(() => props.task, () => {
  isEditing.value = !!props.task
  initForm()
}, { immediate: true })

// Focus on content input when modal opens
nextTick(() => {
  const textarea = document.getElementById('content')
  if (textarea) {
    textarea.focus()
  }
})
</script>

<style scoped>
.task-form-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: var(--spacing-md);
}

.task-form-modal {
  background-color: var(--color-white);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-xl);
  width: 100%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: var(--spacing-lg);
  border-bottom: 1px solid var(--color-gray-200);
}

.modal-header h3 {
  margin: 0;
  font-size: var(--font-size-lg);
  font-weight: 600;
  color: var(--color-gray-900);
}

.task-form {
  padding: var(--spacing-lg);
}

.form-group {
  margin-bottom: var(--spacing-lg);
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--spacing-md);
}

.form-label {
  display: block;
  margin-bottom: var(--spacing-sm);
  font-weight: 500;
  color: var(--color-gray-700);
  font-size: var(--font-size-sm);
}

.task-textarea {
  resize: vertical;
  min-height: 80px;
}

.form-actions {
  display: flex;
  gap: var(--spacing-sm);
  justify-content: flex-end;
  margin-top: var(--spacing-xl);
  padding-top: var(--spacing-lg);
  border-top: 1px solid var(--color-gray-200);
}

@media (max-width: 768px) {
  .task-form-overlay {
    padding: var(--spacing-sm);
  }
  
  .form-row {
    grid-template-columns: 1fr;
  }
  
  .modal-header,
  .task-form {
    padding: var(--spacing-md);
  }
  
  .form-actions {
    flex-direction: column-reverse;
  }
  
  .form-actions .btn {
    width: 100%;
  }
}
</style>