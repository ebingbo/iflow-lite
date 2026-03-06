import type { AssignType } from '~/types/assignment'

export type NodeType = 'start' | 'end' | 'user_task' | 'join'

export interface Node {
  id?: number
  process_id?: number
  process_code?: string
  tag?: string
  name: string
  code?: string
  type?: string
  description?: string | null
  created_by?: number
  created_by_name?: string
  updated_by?: number
  updated_by_name?: string
  created_at?: string
  updated_at?: string
  x: number
  y: number
  assignType: AssignType
  assignTo: string[]
}

export interface NodeTypeConfig {
  name: string
  color: string
  bgColor: string
  borderColor: string
}

export const nodeTypeConfigs: Record<NodeType, NodeTypeConfig> = {
  start: {
    name: '开始',
    color: 'text-success',
    bgColor: 'bg-success',
    borderColor: 'border-success'
  },
  end: {
    name: '结束',
    color: 'text-error',
    bgColor: 'bg-error',
    borderColor: 'border-error'
  },
  user_task: {
    name: '用户任务',
    color: 'text-primary',
    bgColor: 'bg-primary',
    borderColor: 'border-primary'
  },
  join: {
    name: '汇聚',
    color: 'text-secondary',
    bgColor: 'bg-secondary',
    borderColor: 'border-secondary'
  }
}

export interface NodeListBody {
  process_id?: number
  process_code?: string
  code: string
  type: string
}

export interface NodeAddBody {
  process_id: number
  process_code?: string
  tag: string
  name: string
  code: string
  type: string
  description: string
}
