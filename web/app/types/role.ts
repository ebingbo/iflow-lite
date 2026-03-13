export interface Role {
  id: number
  name: string
  code: string
  created_at: string
  updated_at: string
}

export interface RoleQueryBody {
  page: number
  size: number
  keyword?: string
}

export interface RoleQueryData {
  total: number
  items: Role[]
}

export interface RoleAddBody {
  name: string
  code: string
}

export interface RoleUpdateBody {
  id: number
  name: string
  code: string
}
