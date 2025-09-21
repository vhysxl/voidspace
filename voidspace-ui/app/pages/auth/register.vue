<script setup lang="ts">
import { registerSchema } from "@/validations";
import registerImage from '~/assets/images/auth2.jpg'
import logoDark from '~/assets/images/logo_dark.png'
import logoLight from '~/assets/images/logo_light.png'

definePageMeta({
  layout: "auth",
  middleware: "guest"
});

const show = ref(false);
const isLoading = ref(false)
const auth = useAuthStore()
const authCall = useAuth()

const state = reactive({
  username: "",
  email: "",
  password: "",
  confirmPassword: "",
});

const toast = useToast();

async function onSubmit() {
  if (isLoading.value) return

  isLoading.value = true
  try {
    const res = await authCall.register(state.username, state.email, state.password)
    const accessToken = res.data?.access_token
    const expiresIn = res.data?.expires_in

    if (typeof accessToken != "string" || typeof expiresIn != "number") {
      toast.add({
        title: "Register Failed",
        description: "Failed to Register please try again later",
        color: "error",
      })
      return
    }

    auth.login(accessToken, expiresIn)

    await navigateTo('/')
  } catch (error: any) {
    toast?.add?.({
      title: "Login Failed",
      description: error?.message?.toString() || "Failed to login please try again later",
      color: "error",
    })
  } finally {
    isLoading.value = false
  }

  return
}
</script>

<template>
  <div class="flex flex-col lg:flex-row min-h-screen">
    <div class="w-full lg:w-1/2 flex items-center justify-center p-8 min-h-screen lg:min-h-0">
      <div class="w-full max-w-md p-8 rounded-lg">
        <div class="flex justify-center mb-6">
          <img :src="logoDark" alt="VoidSpace Logo" class="h-8 w-auto dark:block hidden" />
          <img :src="logoLight" alt="VoidSpace Logo" class="h-8 w-auto dark:hidden block" />
        </div>

        <h2 class="text-2xl font-semibold mb-6 text-center">Join VoidSpace</h2>

        <UForm :schema="registerSchema" :state="state" class="space-y-4" @submit="onSubmit">
          <UFormField label="Username" name="username">
            <UInput v-model="state.username" color="neutral" placeholder="Enter username" class="w-full" size="lg" />
          </UFormField>
          <UFormField label="Email" name="email">
            <UInput v-model="state.email" color="neutral" placeholder="Enter email" class="w-full" size="lg" />
          </UFormField>
          <UFormField label="Password" name="password">
            <UInput v-model="state.password" :type="show ? 'text' : 'password'" color="neutral"
              placeholder="Enter password" class="w-full" size="lg" />
          </UFormField>
          <UFormField label="Confirm Password" name="confirmPassword">
            <UInput v-model="state.confirmPassword" :type="show ? 'text' : 'password'" color="neutral"
              placeholder="Confirm password" class="w-full" size="lg" />
          </UFormField>

          <UButton :disabled="isLoading" type="submit" color="neutral" size="lg" class="w-full justify-center">
            Register
          </UButton>
        </UForm>

        <p class="text-center mt-4">
          Already have an account?
          <NuxtLink class="text-neutral-500 hover:text-neutral-400" to="/auth/login">Login</NuxtLink>
        </p>
      </div>
    </div>
    <div class="hidden lg:flex lg:w-1/2 lg:h-screen items-center justify-center p-2 overflow-hidden">
      <img :src="registerImage" class="w-full h-full object-cover rounded-r-lg" alt="Register Image" />
    </div>
  </div>
</template>
