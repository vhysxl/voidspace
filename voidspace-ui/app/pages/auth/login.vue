<template>
  <div class="flex flex-col lg:flex-row min-h-screen">
    <div class="lg:w-1/2 flex items-center justify-center p-8">
      <h1>Logo</h1>
    </div>
    <div class="h-screen lg:w-1/2 flex items-center justify-center p-8">
      <div class="flex flex-col items-center justify-center w-full gap-1">
        <UModal v-model="isRegisterOpen">
          <UButton
            label="Register"
            color="neutral"
            class="w-2/4 justify-center"
            size="xl" />
          <template #content>
            <div
              class="min-h-fit flex items-center justify-center py-12 px-2 sm:px-6 lg:px-6">
              <div class="max-w-sm w-full space-y-4">
                <div class="text-center">
                  <h2 class="text-2xl font-bold">Join VoidSpace</h2>
                </div>
                <div class="py-8 px-6">
                  <UForm
                    :schema="registerSchema"
                    :state="registerState"
                    class="space-y-6"
                    @submit="onRegisterSubmit">
                    <UFormField label="Full Name" name="fullName">
                      <UInput
                        v-model="registerState.fullName"
                        color="neutral"
                        placeholder="Enter your full name"
                        class="w-full"
                        size="lg" />
                    </UFormField>
                    <UFormField label="Username" name="username">
                      <UInput
                        v-model="registerState.username"
                        color="neutral"
                        placeholder="Choose a username"
                        class="w-full"
                        size="lg" />
                    </UFormField>
                    <UFormField label="Email" name="email">
                      <UInput
                        v-model="registerState.email"
                        color="neutral"
                        placeholder="Enter your email"
                        class="w-full"
                        size="lg" />
                    </UFormField>
                    <UFormField label="Password" name="password">
                      <UInput
                        v-model="registerState.password"
                        color="neutral"
                        placeholder="Create a password"
                        :type="showRegister ? 'text' : 'password'"
                        class="w-full"
                        size="lg" />
                    </UFormField>
                    <UFormField label="Confirm Password" name="confirmPassword">
                      <UInput
                        v-model="registerState.confirmPassword"
                        color="neutral"
                        placeholder="Confirm your password"
                        :type="showRegister ? 'text' : 'password'"
                        class="w-full"
                        size="lg" />
                    </UFormField>
                    <UButton
                      type="submit"
                      color="neutral"
                      size="lg"
                      class="w-full justify-center">
                      Create Account
                    </UButton>
                  </UForm>
                  <div class="mt-6 text-center">
                    <p class="text-sm text-gray-600">
                      Already have an account?
                      <a
                        href="#"
                        class="text-blue-600 font-medium"
                        @click.prevent="switchToLogin">
                        Sign in here
                      </a>
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </template>
        </UModal>
        <div class="relative flex items-center justify-center w-2/4 my-4">
          <hr class="w-full border-gray-300" />
          <span class="absolute px-1 bg-ui text-sm">OR</span>
        </div>
        <UModal v-model="isLoginOpen">
          <UButton
            label="Login"
            color="neutral"
            variant="outline"
            class="w-2/4 justify-center"
            size="xl" />
          <template #content>
            <div
              class="min-h-fit flex items-center justify-center py-12 px-2 sm:px-6 lg:px-6">
              <div class="max-w-sm w-full space-y-4">
                <div class="text-center">
                  <h2 class="text-2xl font-bold">Welcome back to VoidSpace</h2>
                </div>
                <div class="py-8 px-6">
                  <UForm
                    :schema="schema"
                    :state="state"
                    class="space-y-6"
                    @submit="onSubmit">
                    <UFormField label="Username or Email" name="credential">
                      <UInput
                        v-model="state.credential"
                        color="neutral"
                        placeholder="Enter username or email"
                        class="w-full"
                        size="lg" />
                    </UFormField>
                    <UFormField label="Password" name="password">
                      <UInput
                        v-model="state.password"
                        color="neutral"
                        placeholder="Enter your password"
                        :type="show ? 'text' : 'password'"
                        class="w-full"
                        size="lg" />
                    </UFormField>
                    <UButton
                      type="submit"
                      color="neutral"
                      size="lg"
                      class="w-full justify-center">
                      Sign In
                    </UButton>
                  </UForm>
                  <div class="mt-6 text-center">
                    <p class="text-sm text-gray-600">
                      Don't have an account?
                      <a
                        href="#"
                        class="text-blue-600 font-medium"
                        @click.prevent="switchToRegister">
                        Sign up here
                      </a>
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </template>
        </UModal>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import * as v from "valibot";
import type { FormSubmitEvent } from "@nuxt/ui";

definePageMeta({
  layout: "auth",
});

const show = ref(false);
const showRegister = ref(false);
const isLoginOpen = ref(false);
const isRegisterOpen = ref(false);

// Login schema
const schema = v.object({
  credential: v.pipe(v.string(), v.nonEmpty("Username or Email is required")),
  password: v.pipe(v.string(), v.nonEmpty("Password is required")),
});

// Register schema
const registerSchema = v.object({
  fullName: v.pipe(v.string(), v.nonEmpty("Full name is required")),
  username: v.pipe(
    v.string(),
    v.nonEmpty("Username is required"),
    v.minLength(3, "Username must be at least 3 characters"),
  ),
  email: v.pipe(
    v.string(),
    v.nonEmpty("Email is required"),
    v.email("Please enter a valid email"),
  ),
  password: v.pipe(
    v.string(),
    v.nonEmpty("Password is required"),
    v.minLength(8, "Password must be at least 8 characters"),
  ),
  confirmPassword: v.pipe(
    v.string(),
    v.nonEmpty("Please confirm your password"),
  ),
});

type Schema = v.InferOutput<typeof schema>;
type RegisterSchema = v.InferOutput<typeof registerSchema>;

const state = reactive({
  credential: "",
  password: "",
});

const registerState = reactive({
  fullName: "",
  username: "",
  email: "",
  password: "",
  confirmPassword: "",
});

const toast = useToast();

// Switch between modals
function switchToLogin() {
  isRegisterOpen.value = false;
  isLoginOpen.value = true;
}

function switchToRegister() {
  isLoginOpen.value = false;
  isRegisterOpen.value = true;
}

async function onSubmit(event: FormSubmitEvent<Schema>) {
  toast.add({
    title: "Success",
    description: "Login successful.",
    color: "success",
  });
  console.log(event.data);
}

async function onRegisterSubmit(event: FormSubmitEvent<RegisterSchema>) {
  if (registerState.password !== registerState.confirmPassword) {
    toast.add({
      title: "Error",
      description: "Passwords do not match.",
      color: "error",
    });
    return;
  }

  toast.add({
    title: "Success",
    description: "Account created successfully.",
    color: "success",
  });
  console.log(event.data);
}
</script>
