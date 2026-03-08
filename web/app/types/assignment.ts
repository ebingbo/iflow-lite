export type AssignType = 'user' | 'role' | null

export interface Assignment {
  id: number
  process_id: number
  process_code: string
  node_id: number
  node_code: string
  principal_type: string
  principal_id: number
  priority: number
  strategy: string
  created_at: string
  updated_at: string
}

export interface AssignmentAddBody {
  process_id: number
  node_id: number
  principal_type: string
  principal_id: number
  priority: number
  strategy: string
}

export interface AssignmentUpdateBody {
  id: number
  principal_type: string
  principal_id: number
  priority: number
  strategy: string
}

export interface AssignmentListBody {
  process_id?: number
  node_id?: number
}
