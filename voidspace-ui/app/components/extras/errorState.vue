<script setup lang="ts">
defineProps<{
    readonly retry: () => Promise<void>
    readonly error?: string
}>()
</script>

<template>
    <div class="flex justify-center items-center min-h-screen">
        <div class="text-center space-y-4 p-6">
            <div class="text-6xl">😕</div>
            <h2 class="text-xl font-semibold text-black dark:text-white">
                Oops! Something went wrong
            </h2>
            <p class="text-gray-500 dark:text-gray-400 max-w-md">
                {{ error || 'Something unexpected happened' }}
            </p>
            <div class="flex gap-3 justify-center">
                <button @click="retry"
                    class="bg-neutral-950 dark:bg-white text-white dark:text-black font-bold rounded-full hover:bg-neutral-500 cursor-pointer px-4 py-2 transition-colors">
                    Retry
                </button>
                <button @click="$router.go(-1)"
                    class="text-black dark:text-white rounded-full border hover:bg-neutral-500 cursor-pointer px-4 py-2 transition-colors">
                    Go Back
                </button>
            </div>
        </div>
    </div>
</template>