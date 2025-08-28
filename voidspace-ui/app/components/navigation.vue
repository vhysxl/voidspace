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
    activePath?: string
}>()


</script>

<template>
    <nav class="hidden lg:flex  flex-col gap-2">
        <!-- Logo -->
        <div class="flex mb-6">
            <img :src="logoDark" alt="VoidSpace Logo" class="h-6 w-auto dark:block hidden" />
            <img :src="logoLight" alt="VoidSpace Logo" class="h-6 w-auto dark:hidden block" />
        </div>

        <!-- Menu Items -->
        <template v-for="item in menuItems" :key="item.label">
            <a :href="item.href" class="flex items-center gap-3 px-4 py-3 rounded-lg transition-colors" :class="[
                activePath === item.href
                    ? 'bg-neutral-200 dark:bg-neutral-800 font-bold'
                    : 'hover:bg-neutral-100 dark:hover:bg-neutral-800'
            ]">
                <component :is="item.icon" class="w-6 h-6" />
                <span class="text-base">{{ item.label }}</span>

                <span v-if="item.badge"
                    class="ml-auto px-2 py-1 rounded-full text-xs bg-neutral-400 dark:bg-neutral-700">
                    {{ item.badge }}
                </span>
            </a>
        </template>
    </nav>
</template>
