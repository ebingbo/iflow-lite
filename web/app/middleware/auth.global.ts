export default defineNuxtRouteMiddleware(async (to) => {
  if (to.path === '/login' || to.path === '/signup') {
    return
  }

  const { token, user, fetchProfile, logout } = useAuth()
  if (!token.value) {
    return navigateTo('/login')
  }

  if (!user.value) {
    try {
      await fetchProfile()
    } catch {
      logout()
      return navigateTo('/login')
    }
  }
})
