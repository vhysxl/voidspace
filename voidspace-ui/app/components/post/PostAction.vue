<script setup lang="ts">
import type { Post } from '@/types'

interface Props {
    post: Post
    isLikeSubmitting: boolean
}

const props = defineProps<Props>()
const emit = defineEmits<{
    (e: 'toggle-like', post: Post): void
    (e: 'focus-comment'): void
}>()

function handleLike() {
    emit('toggle-like', props.post)
}

function handleFocusComment() {
    emit('focus-comment')
}
</script>

<template>
    <footer class="flex items-center justify-between px-4 py-4 border-t border-neutral-200 dark:border-neutral-700">
        <div class="flex items-center gap-6">
            <!-- Like -->
            <button @click="handleLike" :disabled="props.isLikeSubmitting"
                class="flex hover:cursor-pointer items-center gap-2 group">
                <UIcon name="i-heroicons-heart-solid" :class="[
                    !props.post.is_liked
                        ? 'w-5 h-5 text-gray-500 group-hover:text-red-500 transition-colors'
                        : 'w-5 h-5 text-red-500 group-hover:text-red-800 transition-colors'
                ]" />
                <span class="text-sm text-gray-600 dark:text-gray-400">
                    {{ props.post.likes_count }}
                </span>
            </button>

            <!-- Reply -->
            <button @click="handleFocusComment" class="hover:cursor-pointer flex items-center gap-2 group">
                <UIcon name="i-heroicons-chat-bubble-oval-left"
                    class="w-5 h-5 text-gray-500 group-hover:text-blue-500 transition-colors" />
                <span class="text-sm text-gray-600 dark:text-gray-400">
                    {{ props.post.comments_count }}
                </span>
            </button>
        </div>
    </footer>
</template>
