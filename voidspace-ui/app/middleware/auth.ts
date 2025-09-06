export default defineNuxtRouteMiddleware((to, from) => {
  const authStore = useAuthStore();

  // Check if user is logged in
  if (!authStore.isLoggedIn) {
    return navigateTo("/auth/login");
  }

  // Check token expiration
  if (authStore.expiresIn && Date.now() > authStore.expiresIn) {
    authStore.logout();
    return navigateTo("/auth/login");
  }
});
