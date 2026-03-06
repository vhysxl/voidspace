<script setup lang="ts">
defineProps<{
    tabs: {
        label: string
        href: string
    }[]
}>()
</script>

<template>
    <nav class="border-b border-gray-200 dark:border-gray-800">
        <div class="flex">
            <NuxtLink v-for="tab in tabs" :key="tab.href" :to="tab.href" :class="[
                'flex-1 px-4 py-4 text-center font-medium text-sm transition-all relative hover:bg-neutral-50 dark:hover:bg-neutral-950',
                $route.path === tab.href
                    ? 'text-black dark:text-white'
                    : 'text-gray-500 dark:text-gray-400'
            ]">
                {{ tab.label }}
                <!-- Active underline -->
                <span v-if="$route.path === tab.href"
                    class="absolute bottom-0 left-1/2 transform -translate-x-1/2 w-12 h-1 bg-black dark:bg-white rounded-full"></span>
            </NuxtLink>
        </div>
    </nav>
</template>