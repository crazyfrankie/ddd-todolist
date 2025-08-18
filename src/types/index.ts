// 用户相关类型
export interface User {
  userID: string;
  name: string;
  user_unique_name: string;
  email: string;
  avatarURL: string;
  screen_name?: string;
  userCreateTime: number;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
}

// 任务相关类型
export interface Task {
  id: number;
  content: string;
  date?: string;
  taskTyp: string;
  priority?: TaskPriority;
  isCompleted?: boolean;
  due_time?: number;
}

export interface CreateTaskRequest {
  content: string;
  date?: number;
  priority?: TaskPriority;
}

export interface UpdateTaskRequest {
  task_id: number;
  content?: string;
  priority?: TaskPriority;
  date?: number;
  isCompleted?: boolean;
}

export type TaskPriority = 
  | 'important and urgent'
  | 'important but not urgent' 
  | 'not important but urgent'
  | 'neither important or urgent';

// API响应类型
export interface ApiResponse<T = any> {
  code: number;
  message: string;
  data: T;
}