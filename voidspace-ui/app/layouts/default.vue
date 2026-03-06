<script setup lang="ts">
import { HomeIcon, HandThumbUpIcon, UserCircleIcon } from '@heroicons/vue/24/outline'
import { useAuthStore } from '@/stores/AuthStore'
import MobileNavigation from '~/components/nav/mobileNavigation.vue'
import type { DropdownMenuItem } from '@nuxt/ui'

const route = useRoute()
const router = useRouter()
const auth = useAuthStore()

const menuItems = computed(() => [
  { label: "Home", href: "/", icon: HomeIcon },
  { label: "Likes", href: "/likes", icon: HandThumbUpIcon },
  { label: auth.user ? "Profile" : "Login", href: auth.isLoggedIn ? `/user/${auth.user?.username}` : `/auth/login`, icon: UserCircleIcon },
])

const colorMode = useColorMode();

const isDark = computed({
  get() {
    return colorMode.value === "dark";
  },
  set(_isDark) {
    colorMode.preference = _isDark ? "dark" : "light";
  },
});

const items = ref<DropdownMenuItem[]>([
  {
    label: 'Logout',
    icon: 'i-lucide-log-out',
    color: "error",
    async onSelect() {
      await auth.logout()
      router.push('/auth/login')
    }
  },

])

</script>

<template>
  <div class="max-w-screen-xl mx-auto flex min-h-screen">
    <!-- LEFT SIDEBAR -->
    <aside class="hidden z-100 lg:flex w-64 flex-shrink-0">
      <div
        class="sticky top-0 h-screen w-full p-5 border-r border-neutral-500 flex flex-col justify-between overflow-y-auto">
        <NavSidebar :menu-items="menuItems" />
        <div class="flex items-center space-x-2 mt-auto">
          <template v-if="auth.isLoggedIn">
            <UDropdownMenu :items="items" class="mouse-pointer" :content="{
              align: 'start',
              side: 'top',
              sideOffset: 8
            }" :ui="{
              content: 'w-48'
            }">
              <!-- Trigger button untuk dropdown -->
              <div
                class="flex items-center space-x-3 cursor-pointer hover:bg-neutral-100 dark:hover:bg-neutral-800 rounded-lg p-2 transition-colors">
                <UserAvatar v-if="auth.user" size="xl" :user="auth.user" />
                <div class="flex flex-col min-w-0 flex-1">
                  <span class="font-semibold text-sm truncate">
                    {{ auth.user?.profile?.display_name || auth.user?.username }}
                  </span>
                  <span class="text-xs text-neutral-500 truncate">
                    @{{ auth.user?.username }}
                  </span>
                </div>
                <!-- Dropdown arrow -->
                <Icon name="i-lucide-chevron-up" class="w-4 h-4 text-neutral-400" />
              </div>

            </UDropdownMenu>
          </template>

          <template v-else>
            <button @click="router.push('/auth/login')"
              class="bg-neutral-950 dark:bg-white text-white dark:text-black font-bold rounded-full hover:bg-neutral-500 cursor-pointer px-20 py-2 transition-colors">
              Login
            </button>
          </template>
        </div>
      </div>
    </aside>

    <!-- MAIN CONTENT -->
    <main class="flex-1 min-w-0">
      <slot />
    </main>

    <!-- RIGHT SIDEBAR -->
    <aside class="hidden lg:flex w-2/8 flex-shrink-0">
      <div class="sticky top-0 h-screen w-full p-5 border-l border-neutral-500 overflow-y-auto scrollbar-hide">
        <div class="flex flex-col space-y-6">
          <RightSidebarWidgetFollowWidget title="Who to follow" />
        </div>
        <UButton :icon="isDark ? 'i-lucide-moon' : 'i-lucide-sun'" color="neutral" variant="outline"
          @click="isDark = !isDark" />
      </div>
    </aside>

    <!-- Mobile Nav -->
    <MobileNavigation :-user="auth.user" :menu-items="menuItems" class="lg:hidden" />
  </div>
</template>