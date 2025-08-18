'use client';

import { useState } from 'react';
import { apiClient } from '@/lib/api';
import { Task, TaskPriority } from '@/types';

interface TaskFormProps {
  onTaskCreated: (task: Task) => void;
}

const priorities: { value: TaskPriority; label: string }[] = [
  { value: 'important and urgent', label: '重要且紧急' },
  { value: 'important but not urgent', label: '重要但不紧急' },
  { value: 'not important but urgent', label: '不重要但紧急' },
  { value: 'neither important or urgent', label: '既不重要也不紧急' },
];

export default function TaskForm({ onTaskCreated }: TaskFormProps) {
  const [content, setContent] = useState('');
  const [priority, setPriority] = useState<TaskPriority>('important and urgent');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!content.trim()) return;

    setLoading(true);
    try {
      const response = await apiClient.createTask({
        content: content.trim(),
        priority,
      });
      onTaskCreated(response.data);
      setContent('');
      setPriority('important and urgent');
    } catch (error) {
      console.error('创建任务失败:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="bg-white shadow-lg rounded-xl p-8 border border-gray-100">
      <h2 className="text-xl font-semibold text-gray-900 mb-6 flex items-center">
        <span className="w-2 h-2 bg-indigo-500 rounded-full mr-3"></span>
        创建新任务
      </h2>
      <form onSubmit={handleSubmit} className="space-y-6">
        <div>
          <label htmlFor="content" className="block text-sm font-medium text-gray-700 mb-2">
            任务内容
          </label>
          <textarea
            id="content"
            rows={4}
            className="w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm placeholder-gray-400 text-gray-900 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition-all duration-200 hover:border-gray-400 resize-none"
            placeholder="描述您要完成的任务..."
            value={content}
            onChange={(e) => setContent(e.target.value)}
            required
          />
        </div>

        <div>
          <label htmlFor="priority" className="block text-sm font-medium text-gray-700 mb-2">
            优先级
          </label>
          <select
            id="priority"
            className="w-full px-4 py-3 border border-gray-300 rounded-lg shadow-sm text-gray-900 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition-all duration-200 hover:border-gray-400 bg-white"
            value={priority}
            onChange={(e) => setPriority(e.target.value as TaskPriority)}
          >
            {priorities.map((p) => (
              <option key={p.value} value={p.value}>
                {p.label}
              </option>
            ))}
          </select>
        </div>

        <div>
          <button
            type="submit"
            disabled={loading || !content.trim()}
            className="w-full flex justify-center items-center py-3 px-6 border border-transparent rounded-lg shadow-sm text-base font-medium text-white bg-gradient-to-r from-indigo-600 to-purple-600 hover:from-indigo-700 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-200 transform hover:scale-[1.02]"
          >
            {loading ? (
              <>
                <svg className="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                  <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                  <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                创建中...
              </>
            ) : (
              <>
                <svg className="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
                </svg>
                创建任务
              </>
            )}
          </button>
        </div>
      </form>
    </div>
  );
}