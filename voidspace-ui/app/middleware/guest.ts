export default defineNuxtRouteMiddleware((to, from) => {
  const authStore = useAuthStore();

  // Check if user is already logged in
  if (authStore.isLoggedIn) {
    // Check if token is still valid
    if (authStore.expiresIn && Date.now() < authStore.expiresIn) {
      // User is logged in and token is valid, redirect to dashboard
      return navigateTo("/");
    } else {
      // Token expired, logout user
      authStore.logout();
    }
  }
});
