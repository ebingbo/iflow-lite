import type { AssignType } from '~/types/assignment'

export type NodeType = 'start' | 'end' | 'user_task' | 'join'

export interface Node {
  id: number
  process_id: number
  process_code: string
  tag: string
  name: string
  code: string
  type: string
  description: string | null
  created_by: number
  created_by_name: string
  updated_by: number
  updated_by_name: string
  created_at: string
  updated_at: string
  x: number
  y: number
  assignType: AssignType
  assignTo: string[]
}

export interface NodeTypeConfig {
  name: string
  icon: string
  uiColor: 'primary' | 'secondary' | 'success' | 'info' | 'warning' | 'error' | 'neutral'
  canvasFill: string
}

export const nodeTypeConfigs: Record<NodeType, NodeTypeConfig> = {
  start: {
    name: '开始',
    icon: 'i-lucide-play',
    uiColor: 'success',
    canvasFill: 'fill-emerald-500'
  },
  end: {
    name: '结束',
    icon: 'i-lucide-flag',
    uiColor: 'error',
    canvasFill: 'fill-rose-500'
  },
  user_task: {
    name: '用户任务',
    icon: 'i-lucide-user-check',
    uiColor: 'primary',
    canvasFill: 'fill-sky-500'
  },
  join: {
    name: '汇聚',
    icon: 'i-lucide-git-merge',
    uiColor: 'warning',
    canvasFill: 'fill-amber-500'
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

export interface NodeUpdateBody {
  id: number
  tag: string
  description: string
}
