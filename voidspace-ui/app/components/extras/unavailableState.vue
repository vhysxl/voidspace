<script setup lang="ts">
defineProps<{
    readonly title?: string
    readonly message?: string
}>()
</script>

<template>
    <div class="flex justify-center items-center min-h-screen">
        <div class="text-center space-y-5 p-6">
            <!-- Icon / Emoji -->
            <div class="text-7xl animate-bounce">🚧</div>

            <!-- Title -->
            <h2 class="text-2xl font-extrabold text-black dark:text-white">
                {{ title || "Page Under Development" }}
            </h2>

            <!-- Message -->
            <p class="text-gray-500 dark:text-gray-400 max-w-lg mx-auto leading-relaxed">
                {{ `${message} page is under development. We'll be online soon 🚀 Stay tuned!` }}
            </p>

            <!-- Action -->
            <div class="flex gap-4 justify-center pt-4">
                <button @click="$router.go(-1)"
                    class="px-5 py-2 rounded-full border font-semibold text-sm text-black dark:text-white hover:bg-neutral-200 dark:hover:bg-neutral-800 transition-colors shadow-sm">
                    Go Back
                </button>
                <NuxtLink to="/"
                    class="px-5 py-2 rounded-full font-semibold text-sm bg-neutral-950 dark:bg-white text-white dark:text-black hover:bg-neutral-700 dark:hover:bg-neutral-200 transition-colors shadow-sm">
                    Back to Home
                </NuxtLink>
            </div>
        </div>
    </div>
</template>
