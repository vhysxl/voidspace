<script setup lang="ts">
import { loginSchema, registerSchema } from '@/validations'
import { useAuthStore } from '@/stores/AuthStore'
import { useAuth } from "~/composables/useAuth";

const props = defineProps<{
  mode: 'login' | 'register'
}>();

const isLoading = ref(false)
const auth = useAuthStore()
const authCall = useAuth()
const toast = useToast();

const state = reactive({
  credential: "",
  password: "",
  username: "",
  email: "",
  confirmPassword: "",
});

const currentSchema = computed(() => props.mode === 'login' ? loginSchema : registerSchema);

async function onSubmit() {
  if (isLoading.value) return

  isLoading.value = true
  try {
    let res;
    if (props.mode === 'login') {
      res = await authCall.login(state.credential, state.password)
    } else {
      res = await authCall.register(state.username, state.email, state.password)
    }

    const accessToken = res.data?.access_token
    const expiresIn = res.data?.expires_in

    if (typeof accessToken != "string" || typeof expiresIn != "number") {
      throw new Error("Invalid response from server");
    }

    auth.login(accessToken, expiresIn)
    await navigateTo('/')
  } catch (error: any) {
    toast.add({
      title: `${props.mode === 'login' ? 'Login' : 'Registration'} Failed`,
      description: error?.message?.toString() || "An unexpected error occurred. Please try again later.",
      color: "error",
    })
  } finally {
    isLoading.value = false
  }
}

const toggleMode = (newMode: 'login' | 'register') => {
  if (newMode === props.mode) return;
  navigateTo(`/auth/${newMode}`);
};
</script>

<template>
  <div class="flex flex-col items-center w-full animate-in">
    <!-- Header Logo Branding -->
    <div class="flex flex-col items-center mb-10 w-full">
      <div class="mb-4">
        <div class="w-16 h-16 relative">
          <svg viewBox="0 0 64 64" fill="none" xmlns="http://www.w3.org/2000/svg" class="w-full h-full text-white">
            <path d="M32 0L64 32L32 64L0 32L32 0Z" fill="currentColor" fill-opacity="0.1" />
            <circle cx="32" cy="32" r="14" stroke="currentColor" stroke-width="1.5" stroke-dasharray="4 4" class="animate-[spin_20s_linear_infinite]" />
            <circle cx="32" cy="32" r="8" fill="currentColor" />
            <path d="M32 12V18M32 46V52M12 32H18M46 32H52" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" />
          </svg>
        </div>
      </div>
      <h1 class="text-[36px] md:text-[48px] font-black text-white tracking-[12px] uppercase mb-2 leading-none">VOIDSPACE</h1>
      <div class="w-24 h-px bg-white/30"></div>
    </div>

    <!-- Form Section -->
    <div class="w-full bg-black/60 border border-white/10 relative shadow-2xl overflow-hidden">
      <!-- Removed aesthetic corner glows for a matte look -->
      
      <!-- Form Content -->
      <div class="p-6 md:p-8 flex flex-col">
        <div class="text-center mb-8">
          <h2 class="text-xl md:text-2xl font-bold text-white tracking-[4px] uppercase mb-1">ENTER THE VOID</h2>
          <p class="text-[9px] md:text-[10px] font-medium text-white/40 tracking-[3px] uppercase italic">
            Awaiting cosmic credentials...
          </p>
        </div>

        <!-- Tabs -->
        <div class="flex border-b border-white/10 mb-8">
          <button 
            @click="toggleMode('login')"
            class="flex-1 pb-4 text-xs font-bold tracking-[2.5px] uppercase transition-all relative"
            :class="mode === 'login' ? 'text-white' : 'text-white/30 hover:text-white/50'"
          >
            Sign In
            <div v-if="mode === 'login'" class="absolute bottom-[-1px] left-0 right-0 h-0.5 bg-white shadow-[0_0_8px_rgba(255,255,255,0.5)]"></div>
          </button>
          <button 
            @click="toggleMode('register')"
            class="flex-1 pb-4 text-xs font-bold tracking-[2.5px] uppercase transition-all relative"
            :class="mode === 'register' ? 'text-white' : 'text-white/30 hover:text-white/50'"
          >
            Sign Up
            <div v-if="mode === 'register'" class="absolute bottom-[-1px] left-0 right-0 h-0.5 bg-white shadow-[0_0_8px_rgba(255,255,255,0.5)]"></div>
          </button>
        </div>

        <!-- Form Fields -->
        <UForm :schema="currentSchema" :state="state" @submit="onSubmit" class="space-y-5">
          <template v-if="mode === 'login'">
            <UFormField name="credential" :ui="{ label: 'text-[10px] font-bold text-white/60 tracking-[2px] uppercase pl-1 mb-1.5' }">
              <template #label>Username or Email</template>
              <div class="relative group">
                <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none z-10">
                  <UIcon name="i-heroicons-user" class="w-4 h-4 text-white/20 group-focus-within:text-white/50 transition-colors" />
                </div>
                <input 
                  v-model="state.credential"
                  type="text" 
                  placeholder="Enter your username or email"
                  class="block w-full bg-white/[0.03] border border-white/10 py-3.5 pl-11 pr-4 text-white text-sm tracking-[0.5px] focus:outline-none focus:border-white/30 focus:bg-white/[0.06] transition-all placeholder:text-white/10"
                />
              </div>
            </UFormField>

            <UFormField name="password" :ui="{ label: 'text-[10px] font-bold text-white/60 tracking-[2px] uppercase pl-1 mb-1.5' }">
              <template #label>Password</template>
              <div class="relative group">
                <div class="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none z-10">
                  <UIcon name="i-heroicons-lock-closed" class="w-4 h-4 text-white/20 group-focus-within:text-white/50 transition-colors" />
                </div>
                <input 
                  v-model="state.password"
                  type="password" 
                  placeholder="Enter your password"
                  class="block w-full bg-white/[0.03] border border-white/10 py-3.5 pl-11 pr-4 text-white text-sm tracking-[0.5px] focus:outline-none focus:border-white/30 focus:bg-white/[0.06] transition-all placeholder:text-white/10"
                />
              </div>
            </UFormField>
          </template>

          <template v-else>
            <div class="grid grid-cols-1 gap-5">
              <UFormField name="username" :ui="{ label: 'text-[10px] font-bold text-white/60 tracking-[2px] uppercase pl-1 mb-1.5' }">
                <template #label>Username</template>
                <input 
                  v-model="state.username"
                  type="text" 
                  placeholder="Choose a username"
                  class="block w-full bg-white/[0.03] border border-white/10 py-3.5 px-4 text-white text-sm tracking-[0.5px] focus:outline-none focus:border-white/30 focus:bg-white/[0.06] transition-all placeholder:text-white/10"
                />
              </UFormField>

              <UFormField name="email" :ui="{ label: 'text-[10px] font-bold text-white/60 tracking-[2px] uppercase pl-1 mb-1.5' }">
                <template #label>Email Address</template>
                <input 
                  v-model="state.email"
                  type="email" 
                  placeholder="Enter your email"
                  class="block w-full bg-white/[0.03] border border-white/10 py-3.5 px-4 text-white text-sm tracking-[0.5px] focus:outline-none focus:border-white/30 focus:bg-white/[0.06] transition-all placeholder:text-white/10"
                />
              </UFormField>

              <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
                <UFormField name="password" :ui="{ label: 'text-[10px] font-bold text-white/60 tracking-[2px] uppercase pl-1 mb-1.5' }">
                  <template #label>Password</template>
                  <input 
                    v-model="state.password"
                    type="password" 
                    placeholder="Create password"
                    class="block w-full bg-white/[0.03] border border-white/10 py-3.5 px-4 text-white text-sm tracking-[0.5px] focus:outline-none focus:border-white/30 focus:bg-white/[0.06] transition-all placeholder:text-white/10"
                  />
                </UFormField>

                <UFormField name="confirmPassword" :ui="{ label: 'text-[10px] font-bold text-white/60 tracking-[2px] uppercase pl-1 mb-1.5' }">
                  <template #label>Confirm Password</template>
                  <input 
                    v-model="state.confirmPassword"
                    type="password" 
                    placeholder="Repeat password"
                    class="block w-full bg-white/[0.03] border border-white/10 py-3.5 px-4 text-white text-sm tracking-[0.5px] focus:outline-none focus:border-white/30 focus:bg-white/[0.06] transition-all placeholder:text-white/10"
                  />
                </UFormField>
              </div>
            </div>
          </template>

          <button 
            type="submit" 
            :disabled="isLoading"
            class="w-full bg-white text-black py-4 mt-4 font-black text-sm tracking-[2px] uppercase hover:bg-neutral-200 active:bg-neutral-300 transition-all flex items-center justify-center gap-3 active:scale-[0.99] disabled:opacity-50"
          >
            <span v-if="!isLoading">{{ mode === 'login' ? 'Sign In' : 'Create Account' }}</span>
            <span v-else class="animate-pulse">Processing...</span>
            <UIcon v-if="!isLoading" name="i-heroicons-arrow-right-20-solid" class="w-4 h-4" />
          </button>
        </UForm>

        <!-- Footer Links -->
        <div class="mt-10 flex items-center justify-between text-[9px] font-bold tracking-[2px] uppercase">
          <button class="text-white/30 hover:text-white/60 transition-colors">Lost Transmission?</button>
          <div class="flex items-center gap-4">
            <div class="w-10 h-px bg-white/10"></div>
            <span class="text-white/20 flex items-center gap-2">
              <span class="w-1.5 h-1.5 rounded-full bg-white/40 animate-pulse"></span>
              Status: Connected
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@reference "../../assets/css/main.css";

.animate-in {
  animation: authIn 0.8s cubic-bezier(0.16, 1, 0.3, 1);
}

@keyframes authIn {
  from { opacity: 0; transform: translateY(20px) scale(0.98); }
  to { opacity: 1; transform: translateY(0) scale(1); }
}

:deep(.u-form-field-error) {
  @apply text-[9px] font-bold tracking-[1px] uppercase text-red-500/80 mt-1 pl-1;
}
</style>
