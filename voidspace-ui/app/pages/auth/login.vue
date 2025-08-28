<template>
  <div class="flex flex-col lg:flex-row min-h-screen">
    <!-- Container gambar - hidden on mobile -->


    <!-- Form - centered on mobile, half width on desktop -->
    <div class="w-full lg:w-1/2 flex items-center justify-center p-8 min-h-screen lg:min-h-0">
      <div class="w-full max-w-md p-8 rounded-lg">
        <div class="flex justify-center mb-6">
          <img :src="logoDark" alt="VoidSpace Logo" class="h-8 w-auto dark:block hidden" />
          <img :src="logoLight" alt="VoidSpace Logo" class="h-8 w-auto dark:hidden block" />
        </div>

        <h2 class="text-2xl font-semibold mb-6 text-center">Login</h2>
        <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit">
          <UFormField label="Username or Email" name="credential">
            <UInput v-model="state.credential" color="neutral" placeholder="Enter username or email" class="w-full"
              size="lg" />
          </UFormField>
          <UFormField label="Password" name="password">
            <UInput v-model="state.password" color="neutral" placeholder="Enter your password"
              :type="show ? 'text' : 'password'" class="w-full" size="lg" />
          </UFormField>
          <UButton type="submit" color="neutral" size="lg" class="w-full mt-2 justify-center">
            Sign In
          </UButton>
        </UForm>
        <p class="text-center mt-4">Don't have an account?<NuxtLink class="text-neutral-500 hover:text-neutral-400"
            to="/auth/register"> Register</NuxtLink>
        </p>
      </div>
    </div>
    <div class="hidden lg:flex lg:w-1/2 lg:h-screen items-center justify-center p-2 overflow-hidden">
      <img :src="loginImage" class="w-full h-full object-cover rounded-l-lg" alt="Login Image" />
    </div>
  </div>

</template>

<script setup lang="ts">
import * as v from "valibot";
import type { FormSubmitEvent } from "@nuxt/ui";
import loginImage from '~/assets/images/auth.jpg'
import logoDark from '~/assets/images/logo_dark.png'
import logoLight from '~/assets/images/logo_light.png'

const show = ref(false);

definePageMeta({
  layout: "auth",
});

// Login schema
const schema = v.object({
  credential: v.pipe(v.string(), v.nonEmpty("Username or Email is required")),
  password: v.pipe(v.string(), v.nonEmpty("Password is required")),
});

type Schema = v.InferOutput<typeof schema>;

const state = reactive({
  credential: "",
  password: "",
});

const toast = useToast();

async function onSubmit(event: FormSubmitEvent<Schema>) {
  toast.add({
    title: "Success",
    description: "Login successful.",
    color: "neutral",
  });
  console.log(event.data);
}
</script>