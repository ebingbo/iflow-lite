export interface Task {
  id: number
  process_id: number
  process_code: string
  process_name: string
  execution_id: number
  node_id: number
  node_code: string
  node_name: string
  assignee_id: number
  assignee_name?: string
  status: 'pending' | 'running' | 'completed' | 'skipped'
  started_at: string | null
  claimed_at: string | null
  ended_at: string | null
  remark: string
  created_at: string
  updated_at: string
}

export interface TaskQueryBody {
  page: number
  size: number
  process_id?: number
  process_code?: string
  execution_id?: number
  node_id?: number
  node_code?: string
  assignee_id?: number
  status?: string
}

export interface TaskClaimableQueryBody {
  page: number
  size: number
  status?: string
}

export interface TaskQueryData {
  total: number
  items: Task[]
}

export interface TaskClaimBody {
  id: number
}

export interface TaskCompleteBody {
  id: number
  assignee_id: number
  remark?: string
}

export interface TaskSkipBody {
  id: number
  assignee_id: number
}

export interface TaskCandidate {
  id: number
  task_id: number
  user_id: number
  source_type: 'user' | 'role'
  source_id: number
  user_name?: string
  created_at: string
  updated_at: string
}

export interface TaskCandidateListBody {
  task_id: number
}

export type TaskCandidateListData = TaskCandidate[]

export interface TaskLog {
  id: number
  process_id: number
  process_code: string
  execution_id: number
  node_id: number
  node_code: string
  task_id: number
  action: string
  assignee_id: string
  assignee_name?: string
  remark: string
  created_at: string
}

export interface TaskLogQueryBody {
  page: number
  size: number
  task_id?: number
  execution_id?: number
}

export interface TaskLogQueryData {
  total: number
  items: TaskLog[]
}
