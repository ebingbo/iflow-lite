export default defineNuxtRouteMiddleware(async (to) => {
  const publicPrefixes = ['/', '/docs', '/pricing', '/blog', '/changelog', '/login', '/signup', '/forgot-password']
  const protectedPrefixes = ['/dashboard', '/process', '/tasks', '/execution', '/org', '/system', '/profile']
  const isPublicRoute = publicPrefixes.some(prefix => prefix === '/'
    ? to.path === '/'
    : to.path === prefix || to.path.startsWith(`${prefix}/`))
  const isProtectedRoute = protectedPrefixes.some(prefix => to.path === prefix || to.path.startsWith(`${prefix}/`))

  const { token, user, fetchProfile, logout } = useAuth()
  if (isPublicRoute) {
    if (token.value && !user.value) {
      try {
        await fetchProfile()
      } catch {
        logout()
      }
    }
    return
  }

  if (isProtectedRoute && !token.value) {
    return navigateTo('/login')
  }

  if (token.value && !user.value) {
    try {
      await fetchProfile()
    } catch {
      logout()
      return navigateTo('/login')
    }
  }
})
