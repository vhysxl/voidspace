<script setup lang="ts">

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
    <nav
        class="lg:hidden fixed bottom-0 left-0 right-0 bg-white dark:bg-neutral-900 border-t border-neutral-200 dark:border-neutral-800 px-2 py-3 z-50">
        <div class="flex justify-around items-center">
            <template v-for="item in menuItems" :key="item.label">
                <a v-if="item.label !== 'Profile'" :href="item.href" :title="item.label"
                    class="flex items-center justify-center p-3 rounded-xl transition-colors relative" :class="[
                        activePath === item.href
                            ? 'text-blue-600 dark:text-blue-400 bg-blue-50 dark:bg-blue-900/20'
                            : 'text-neutral-600 dark:text-neutral-400 hover:text-neutral-900 dark:hover:text-neutral-100 hover:bg-neutral-100 dark:hover:bg-neutral-800'
                    ]">
                    <div class="relative">
                        <component :is="item.icon" class="w-7 h-7" />
                        <!-- Badge untuk mobile -->
                        <span v-if="item.badge"
                            class="absolute -top-2 -right-2 px-1.5 py-0.5 rounded-full text-xs bg-red-500 text-white min-w-[18px] h-[18px] flex items-center justify-center font-medium">
                            {{ item.badge > 99 ? '99+' : item.badge }}
                        </span>
                    </div>
                </a>
                <div v-else class="flex items-center justify-center">
                    <UAvatar size="lg"
                        src="https://images-ext-1.discordapp.net/external/B-T4tAJ-005Oa4ZCibUPCHCqG7ofGne0YxqZhyi9kmI/https/mudae.net/uploads/7266015/Gf1mE5j~4apc5wo66v4.png?format=webp&quality=lossless&width=281&height=438"
                        alt="Profile Avatar" />
                </div>
            </template>
        </div>
    </nav>
</template>