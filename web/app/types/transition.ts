export interface Transition {
  id?: number
  process_id?: number
  from_node_id?: number
  to_node_id?: number
  created_at?: string
  updated_at?: string
}

export interface TransitionListBody {
  process_id: number
}

export interface TransitionAddBody {
  process_id: number
  from_node_id: number
  to_node_id: number
}
