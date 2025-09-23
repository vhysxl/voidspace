<script setup lang="ts">
import type { User } from '@/types';
import { useRoute } from 'vue-router'

defineProps<{
    menuItems: {
        label: string
        href: string
        icon?: any
        badge?: number
    }[]
    User: User | null;
}>()

const route = useRoute()
const activePath = computed(() => route.path)
</script>

<template>
    <nav
        class="lg:hidden fixed bottom-0 left-0 right-0 border-t border-neutral-200 bg-white dark:bg-black dark:border-neutral-800 px-2 py-2 z-50">
        <div class="flex justify-around items-center">
            <template v-for="item in menuItems" :key="item.label">
                <NuxtLink :to="item.href" :title="item.label"
                    class="flex flex-col items-center justify-center px-4 py-2 rounded-lg transition-colors" :class="[
                        activePath === item.href
                            ? 'bg-neutral-200 dark:bg-neutral-800 font-bold text-neutral-900 dark:text-neutral-100'
                            : 'text-neutral-600 dark:text-neutral-400 hover:text-neutral-900 dark:hover:text-neutral-100 hover:bg-neutral-100 dark:hover:bg-neutral-800'
                    ]">
                    <div class="relative flex flex-col items-center">
                        <!-- Jika ada user -->
                        <template v-if="User">
                            <UAvatar v-if="item.label === 'Profile'" size="md" :src="User?.profile?.avatar_url"
                                :alt="User?.username || 'Profile Avatar'" />
                            <component v-else :is="item.icon" class="w-6 h-6" />
                        </template>

                        <template v-else>
                            <component :is="item.label === 'Profile' ? 'LoginButton' : item.icon" class="w-6 h-6" />
                        </template>
                    </div>
                    <span class="text-xs mt-1 truncate">{{ item.label }}</span>
                </NuxtLink>
            </template>
        </div>
    </nav>
</template>
