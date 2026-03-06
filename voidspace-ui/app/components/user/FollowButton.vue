<script setup lang="ts">
import type { User } from '@/types'

const isHovered = ref(false);
interface Props {
    user: User;
    isLoadingProfile: boolean;
    isSubmitting: boolean;
    handleFollowToggle: () => void
}
const props = defineProps<Props>()
</script>

<template>
    <button :disabled="props.isLoadingProfile || isSubmitting" @click="props.handleFollowToggle"
        @mouseenter="isHovered = true" @mouseleave="isHovered = false" :class="[
            'text-sm rounded-full py-2 px-4 font-bold transition-colors',
            !user.is_followed
                ? 'bg-neutral-950 dark:bg-white hover:cursor-pointer text-white dark:text-black hover:bg-neutral-500'
                : isHovered
                    ? 'bg-red-600 hover:cursor-pointer text-white hover:bg-red-700'
                    : 'bg-gray-400 dark:bg-neutral-500 text-white dark:text-black hover:bg-neutral-500 dark:hover:bg-gray-500'
        ]">
        {{ !user.is_followed ? 'Follow' : isHovered ? 'Unfollow' : 'Followed' }}
    </button>
</template>
