export interface Execution {
  id: number
  process_id: number
  process_code: string
  process_name: string
  business_key: string
  business_type: string
  status: 'running' | 'completed'
  progress: number
  created_by: string
  started_at: string | null
  ended_at: string | null
  created_at: string
  updated_at: string
}

export interface ExecutionQueryBody {
  page: number
  size: number
  process_id?: number
  process_code?: string
  business_key?: string
  business_type?: string
  status?: string
}

export interface ExecutionQueryData {
  total: number
  items: Execution[]
}

export interface ExecutionStartBody {
  process_code: string
  business_key: string
  business_type: string
  created_by: string
}
