export interface Process {
  id: number
  name: string
  code: string
  description: string
  status: number
  created_by: number
  created_by_name: string
  updated_by: number
  updated_by_name: string
  created_at: string
  updated_at: string
}

export interface ProcessQueryBody {
  page: number
  size: number
  name: string
  code: string
  status: number
}

export interface ProcessQueryData {
  total: number
  items: Process[]
}

export interface ProcessAddBody {
  name: string
  code: string
  description: string
}

export interface ProcessDeleteBody {
  ids: number[]
}
