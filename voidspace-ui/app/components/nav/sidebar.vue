<script setup lang="ts">
import logoDark from '~/assets/images/logo_dark.png'
import logoLight from '~/assets/images/logo_light.png'

defineProps<{
    menuItems: {
        label: string
        href: string
        icon?: any
        badge?: number
    }[],
}>()

const route = useRoute()
const activePath = computed(() => route.path)
</script>

<template>
    <nav class="hidden lg:flex flex-col gap-2">
        <!-- Logo -->
        <div class="flex mb-6">
            <img :src="logoDark" alt="VoidSpace Logo" class="h-6 w-auto dark:block hidden" />
            <img :src="logoLight" alt="VoidSpace Logo" class="h-6 w-auto dark:hidden block" />
        </div>

        <!-- Menu Items -->
        <template v-for="item in menuItems" :key="item.label">
            <NuxtLink :to="item.href"
                class="flex items-center gap-3 px-4 py-3 rounded-lg transition-colors hover:bg-neutral-100 dark:hover:bg-neutral-800"
                :class="[
                    activePath === item.href
                        ? 'bg-neutral-200 dark:bg-neutral-800 font-bold text-neutral-900 dark:text-neutral-100'
                        : 'text-neutral-600 dark:text-neutral-400 hover:text-neutral-900 dark:hover:text-neutral-100'
                ]">
                <component :is="item.icon" class="w-6 h-6" />
                <span class="text-base">{{ item.label }}</span>

                <!-- Badge untuk desktop -->
                <span v-if="item.badge"
                    class="ml-auto px-2 py-1 rounded-full text-xs bg-red-500 text-white min-w-[18px] h-[18px] flex items-center justify-center font-medium">
                    {{ item.badge > 99 ? '99+' : item.badge }}
                </span>
            </NuxtLink>
        </template>
    </nav>
</template>