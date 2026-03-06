export type AssignType = 'user' | 'role' | null

export interface Assignment {
  id: number
  process_id: number
  process_code: string
  node_id: number
  node_code: string
  type: string
  value: string
  priority: number
  strategy: string
  created_at: string
  updated_at: string
}
