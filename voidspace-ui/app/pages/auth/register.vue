<script setup lang="ts">
import * as v from "valibot";
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

// Register schema
const schema = v.pipe(
  v.object({
    username: v.pipe(
      v.string(),
      v.nonEmpty("Username is required")
    ),
    email: v.pipe(
      v.string(),
      v.nonEmpty('Please enter your email.'),
      v.email('The email address is badly formatted.')
    ),
    password: v.pipe(
      v.string(),
      v.nonEmpty('Please enter your password.'),
      v.minLength(8, 'Your password must have 8 characters or more.')
    ),
    confirmPassword: v.pipe(
      v.string(),
      v.nonEmpty('Please confirm your password.')
    ),
  }),
  v.forward(
    v.partialCheck(
      [['password'], ['confirmPassword']],
      (input) => input.password === input.confirmPassword,
      'The two passwords do not match.'
    ),
    ['confirmPassword']
  )
);

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
    auth.login(res.data.access_token, res.data.expires_in)

    await navigateTo('/')
  } catch (error: any) {
    console.log(error)
    toast.add({
      title: "Register failed",
      description: error.message || "Failed to register please try again later.",
      color: "error",
    });
  } finally {
    isLoading.value = false
  }

  return
}
</script>

<template>
  <div class="flex flex-col lg:flex-row min-h-screen">
    <!-- Form - centered on mobile, half width on desktop -->
    <div class="w-full lg:w-1/2 flex items-center justify-center p-8 min-h-screen lg:min-h-0">
      <div class="w-full max-w-md p-8 rounded-lg">
        <div class="flex justify-center mb-6">
          <img :src="logoDark" alt="VoidSpace Logo" class="h-8 w-auto dark:block hidden" />
          <img :src="logoLight" alt="VoidSpace Logo" class="h-8 w-auto dark:hidden block" />
        </div>

        <h2 class="text-2xl font-semibold mb-6 text-center">Join VoidSpace</h2>

        <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit">
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

    <!-- Container gambar - hidden on mobile -->
    <div class="hidden lg:flex lg:w-1/2 lg:h-screen items-center justify-center p-2 overflow-hidden">
      <img :src="registerImage" class="w-full h-full object-cover rounded-r-lg" alt="Register Image" />
    </div>
  </div>
</template>
