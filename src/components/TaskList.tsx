'use client';

import { useState } from 'react';
import { apiClient } from '@/lib/api';
import { Task, TaskPriority } from '@/types';

interface TaskListProps {
  tasks: Task[];
  onTaskUpdated: (task: Task) => void;
  onTaskDeleted: (taskId: number) => void;
}

const priorities: { value: TaskPriority; label: string; color: string }[] = [
  { value: 'important and urgent', label: '重要且紧急', color: 'bg-red-100 text-red-800' },
  { value: 'important but not urgent', label: '重要但不紧急', color: 'bg-yellow-100 text-yellow-800' },
  { value: 'not important but urgent', label: '不重要但紧急', color: 'bg-blue-100 text-blue-800' },
  { value: 'neither important or urgent', label: '既不重要也不紧急', color: 'bg-gray-100 text-gray-800' },
];

export default function TaskList({ tasks, onTaskUpdated, onTaskDeleted }: TaskListProps) {
  const [editingTask, setEditingTask] = useState<number | null>(null);
  const [editContent, setEditContent] = useState('');
  const [editPriority, setEditPriority] = useState<TaskPriority>('important and urgent');

  // 确保 tasks 是数组
  const taskList = tasks || [];

  const getPriorityInfo = (priority?: TaskPriority) => {
    return priorities.find(p => p.value === priority) || priorities[0];
  };

  const handleToggleComplete = async (task: Task) => {
    try {
      if (task.isCompleted) {
        await apiClient.uncompleteTask(task.id);
      } else {
        await apiClient.completeTask(task.id);
      }
      onTaskUpdated({ ...task, isCompleted: !task.isCompleted });
    } catch (error) {
      console.error('更新任务状态失败:', error);
    }
  };

  const handleDelete = async (taskId: number) => {
    if (!confirm('确定要删除这个任务吗？')) return;
    
    try {
      await apiClient.deleteTask(taskId);
      onTaskDeleted(taskId);
    } catch (error) {
      console.error('删除任务失败:', error);
    }
  };

  const handleEdit = (task: Task) => {
    setEditingTask(task.id);
    setEditContent(task.content);
    setEditPriority(task.priority || 'important and urgent');
  };

  const handleSaveEdit = async (taskId: number) => {
    try {
      await apiClient.updateTask({
        task_id: taskId,
        content: editContent,
        priority: editPriority,
      });
      
      const updatedTask = tasks.find(t => t.id === taskId);
      if (updatedTask) {
        onTaskUpdated({
          ...updatedTask,
          content: editContent,
          priority: editPriority,
        });
      }
      
      setEditingTask(null);
    } catch (error) {
      console.error('更新任务失败:', error);
    }
  };

  const handleCancelEdit = () => {
    setEditingTask(null);
    setEditContent('');
    setEditPriority('important and urgent');
  };

  if (tasks.length === 0) {
    return (
      <div className="bg-white shadow rounded-lg p-6 text-center">
        <p className="text-gray-500">暂无任务</p>
      </div>
    );
  }

  return (
    <div className="bg-white shadow rounded-lg">
      <div className="px-6 py-4 border-b border-gray-200">
        <h2 className="text-lg font-medium text-gray-900">任务列表</h2>
      </div>
      <ul className="divide-y divide-gray-200">
        {tasks.map((task) => {
          const priorityInfo = getPriorityInfo(task.priority);
          const isEditing = editingTask === task.id;

          return (
            <li key={task.id} className="px-6 py-4">
              <div className="flex items-start space-x-3">
                <input
                  type="checkbox"
                  checked={task.isCompleted || false}
                  onChange={() => handleToggleComplete(task)}
                  className="mt-1 h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded"
                />
                
                <div className="flex-1 min-w-0">
                  {isEditing ? (
                    <div className="space-y-4">
                      <textarea
                        value={editContent}
                        onChange={(e) => setEditContent(e.target.value)}
                        className="w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm placeholder-gray-400 text-gray-900 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition-all duration-200 hover:border-gray-400 resize-none"
                        rows={3}
                        placeholder="编辑任务内容..."
                      />
                      <select
                        value={editPriority}
                        onChange={(e) => setEditPriority(e.target.value as TaskPriority)}
                        className="w-full px-3 py-2 border border-gray-300 rounded-lg shadow-sm text-gray-900 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition-all duration-200 hover:border-gray-400 bg-white"
                      >
                        {priorities.map((p) => (
                          <option key={p.value} value={p.value}>
                            {p.label}
                          </option>
                        ))}
                      </select>
                      <div className="flex space-x-3">
                        <button
                          onClick={() => handleSaveEdit(task.id)}
                          className="flex items-center px-4 py-2 bg-green-600 text-white text-sm rounded-lg hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 transition-all duration-200 transform hover:scale-105"
                        >
                          <svg className="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
                          </svg>
                          保存
                        </button>
                        <button
                          onClick={handleCancelEdit}
                          className="flex items-center px-4 py-2 bg-gray-500 text-white text-sm rounded-lg hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-gray-500 transition-all duration-200 transform hover:scale-105"
                        >
                          <svg className="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                          </svg>
                          取消
                        </button>
                      </div>
                    </div>
                  ) : (
                    <>
                      <p className={`text-sm text-gray-900 ${task.isCompleted ? 'line-through text-gray-500' : ''}`}>
                        {task.content}
                      </p>
                      <div className="mt-2 flex items-center space-x-2">
                        <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${priorityInfo.color}`}>
                          {priorityInfo.label}
                        </span>
                        {task.due_time && (
                          <span className="text-xs text-gray-500">
                            截止: {new Date(task.due_time * 1000).toLocaleDateString()}
                          </span>
                        )}
                      </div>
                    </>
                  )}
                </div>

                {!isEditing && (
                  <div className="flex space-x-2">
                    <button
                      onClick={() => handleEdit(task)}
                      className="text-indigo-600 hover:text-indigo-900 text-sm"
                    >
                      编辑
                    </button>
                    <button
                      onClick={() => handleDelete(task.id)}
                      className="text-red-600 hover:text-red-900 text-sm"
                    >
                      删除
                    </button>
                  </div>
                )}
              </div>
            </li>
          );
        })}
      </ul>
    </div>
  );
}