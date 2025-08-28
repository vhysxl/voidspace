<script setup>
import { HomeIcon, BellIcon, HandThumbUpIcon, UserCircleIcon } from '@heroicons/vue/24/outline'
import MobileNavigation from '~/components/mobileNavigation.vue'
const route = useRoute()

const activePath = route.path
const menuItems = [
  { label: "Home", href: "/", icon: HomeIcon },
  { label: "Notifications", href: "/notifications", icon: BellIcon, badge: 3 },
  { label: "Likes", href: "/likes", icon: HandThumbUpIcon },
  { label: "Profile", href: "/profile", icon: UserCircleIcon },
]
</script>

<template>
  <div class="max-w-screen-xl mx-auto flex min-h-screen">
    <!-- SIDEBAR -->
    <aside class="hidden lg:flex w-2/10 p-5 border-r border-neutral-500 flex-col justify-between">
      <Navigation :menu-items="menuItems" :active-path="activePath" />

      <!-- AVATAR -->
      <div class="flex items-center space-x-2">
        <UAvatar size="xl"
          src="https://images-ext-1.discordapp.net/external/B-T4tAJ-005Oa4ZCibUPCHCqG7ofGne0YxqZhyi9kmI/https/mudae.net/uploads/7266015/Gf1mE5j~4apc5wo66v4.png?format=webp&quality=lossless&width=281&height=438" />
        <div class="flex flex-col">
          <span class="font-semibold">Mizore</span>
          <span class="text-sm text-neutral-500">@mizore</span>
        </div>
      </div>
    </aside>

    <!-- MAIN CONTENT -->
    <main class="flex-1 p-5">
      <slot />
    </main>

    <!-- Mobile Nav muncul otomatis karena ada di <Navigation /> -->
    <MobileNavigation :menu-items="menuItems" :active-path="activePath" class="lg:hidden" />
  </div>
</template>
