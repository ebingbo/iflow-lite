export interface UserProfile {
  id: number
  name: string
  email: string
  created_at: string
  updated_at: string
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
