import {
  ApiResponse,
  LoginRequest,
  RegisterRequest,
  User,
  Task,
  CreateTaskRequest,
  UpdateTaskRequest,
} from "@/types";

const API_BASE_URL =
  process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api";

class ApiClient {
  private baseURL: string;
  private token: string | null = null;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
    // 从localStorage获取token
    if (typeof window !== "undefined") {
      this.token = localStorage.getItem("access_token");
    }
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<ApiResponse<T>> {
    const url = `${this.baseURL}${endpoint}`;

    const headers: Record<string, string> = {
      "Content-Type": "application/json",
      ...(options.headers as Record<string, string>),
    };

    if (this.token) {
      headers["Authorization"] = `Bearer ${this.token}`;
    }

    const response = await fetch(url, {
      ...options,
      headers,
      credentials: "include", // 包含cookies
    });

    // 检查响应头中的access token
    const accessToken = response.headers.get("x-access-token");
    if (accessToken) {
      this.token = accessToken;
      if (typeof window !== "undefined") {
        localStorage.setItem("access_token", accessToken);
      }
    }

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(
        errorData.message || `HTTP error! status: ${response.status}`
      );
    }

    return response.json();
  }

  // 用户相关API
  async login(data: LoginRequest): Promise<ApiResponse<User>> {
    const response = await this.request<User>("/user/login", {
      method: "POST",
      body: JSON.stringify(data),
    });
    return response;
  }

  async register(data: RegisterRequest): Promise<ApiResponse<User>> {
    const response = await this.request<User>("/user/register", {
      method: "POST",
      body: JSON.stringify(data),
    });
    return response;
  }

  async logout(): Promise<ApiResponse<null>> {
    const response = await this.request<null>("/user/logout", {
      method: "GET",
    });

    // 清除本地token
    this.token = null;
    if (typeof window !== "undefined") {
      localStorage.removeItem("access_token");
    }

    return response;
  }

  async getUserProfile(): Promise<ApiResponse<User>> {
    return this.request<User>("/user/profile");
  }

  // 任务相关API
  async getTasks(): Promise<ApiResponse<Task[]>> {
    return this.request<Task[]>("/tasks");
  }

  async getTask(taskId: number): Promise<ApiResponse<Task>> {
    return this.request<Task>(`/tasks/${taskId}`);
  }

  async createTask(data: CreateTaskRequest): Promise<ApiResponse<Task>> {
    return this.request<Task>("/tasks", {
      method: "POST",
      body: JSON.stringify(data),
    });
  }

  async updateTask(data: UpdateTaskRequest): Promise<ApiResponse<null>> {
    return this.request<null>("/tasks", {
      method: "PUT",
      body: JSON.stringify(data),
    });
  }

  async deleteTask(taskId: number): Promise<ApiResponse<null>> {
    return this.request<null>(`/tasks/${taskId}`, {
      method: "DELETE",
    });
  }

  setToken(token: string) {
    this.token = token;
    if (typeof window !== "undefined") {
      localStorage.setItem("access_token", token);
    }
  }

  clearToken() {
    this.token = null;
    if (typeof window !== "undefined") {
      localStorage.removeItem("access_token");
    }
  }

  // 检查是否已登录
  isAuthenticated(): boolean {
    return !!this.token;
  }

  // 获取当前token
  getToken(): string | null {
    return this.token;
  }

  // 批量操作任务
  async batchUpdateTasks(
    tasks: UpdateTaskRequest[]
  ): Promise<ApiResponse<null>> {
    return this.request<null>("/tasks/batch", {
      method: "PUT",
      body: JSON.stringify({ tasks }),
    });
  }

  // 按优先级获取任务
  async getTasksByPriority(priority: string): Promise<ApiResponse<Task[]>> {
    return this.request<Task[]>(
      `/tasks?priority=${encodeURIComponent(priority)}`
    );
  }

  // 获取已完成的任务
  async getCompletedTasks(): Promise<ApiResponse<Task[]>> {
    return this.request<Task[]>("/tasks?completed=true");
  }

  // 获取未完成的任务
  async getPendingTasks(): Promise<ApiResponse<Task[]>> {
    return this.request<Task[]>("/tasks?completed=false");
  }

  // 标记任务为完成
  async completeTask(taskId: number): Promise<ApiResponse<null>> {
    return this.request<null>(`/tasks/${taskId}/complete`, {
      method: "PATCH",
    });
  }

  // 取消任务完成状态
  async uncompleteTask(taskId: number): Promise<ApiResponse<null>> {
    return this.request<null>(`/tasks/${taskId}/uncomplete`, {
      method: "PATCH",
    });
  }
}

export const apiClient = new ApiClient(API_BASE_URL);

// 导出一些便捷的工具函数
export const authUtils = {
  isLoggedIn: () => apiClient.isAuthenticated(),
  getToken: () => apiClient.getToken(),
  logout: () => apiClient.logout(),
};
