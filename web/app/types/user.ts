export interface UserProfile {
  id: number
  name: string
  email: string
  status?: number
  created_at: string
  updated_at: string
}

export interface UserListBody {
  keyword?: string
  size?: number
}

export interface UserOption {
  id: number
  name: string
  email: string
}

export interface RoleListBody {
  keyword?: string
  size?: number
}

export interface RoleOption {
  id: number
  name: string
  code: string
}

export interface User {
  id: number
  name: string
  email: string
  status: number | null
  created_at: string
  updated_at: string
}

export interface UserQueryBody {
  page: number
  size: number
  email?: string
  name?: string
  status?: number
}

export interface UserQueryData {
  total: number
  items: User[]
}

export interface UserStatusUpdateBody {
  id: number
  status: number
}

export interface UserRoleListBody {
  user_id: number
}

export type UserRoleListData = number[]

export interface UserRoleUpdateBody {
  user_id: number
  role_ids: number[]
}

export interface UserLoginBody {
  email: string
  password: string
}

export interface UserSignupBody {
  email: string
  name: string
  password: string
}

export interface UserLoginData {
  token: string
  user: UserProfile
}

export interface UserProfileUpdateBody {
  name?: string
}

export interface UserPasswordUpdateBody {
  old_password: string
  new_password: string
}

export interface UserPasswordForgotBody {
  email: string
}
