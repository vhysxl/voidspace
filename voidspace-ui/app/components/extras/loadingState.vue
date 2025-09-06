<script setup lang="ts">
defineProps<{
    title: string
}>()
</script>

<template>
    <div class="flex justify-center items-center min-h-screen">
        <div class="flex flex-col items-center gap-4">
            <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-neutral-500"></div>
            <p class="text-gray-500 dark:text-gray-400">{{ title }}</p>
        </div>
    </div>
</template>