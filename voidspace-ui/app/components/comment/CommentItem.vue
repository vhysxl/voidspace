<script setup lang="ts">
import type { CommentType } from '@/types';
import { formatPostDate } from '@/utils/dateFormater'

interface Comment {
    comment: CommentType
}

const props = defineProps<Comment>()

const navigateToUser = async (username: string) => {
    await navigateTo(`/user/${username}`)
}

</script>

<template>
    <div class="flex gap-3">
        <UserAvatar size="lg" :user="props.comment.author" />
        <div class="flex-1 min-w-0">
            <div class="flex items-center gap-2 mb-1">
                <h3 class="font-semibold text-gray-900 dark:text-gray-100 text-sm hover:underline"
                    @click="navigateToUser(props.comment.author.username)">
                    {{ props.comment.author.profile.display_name || props.comment.author.username }}
                </h3>
                <span class="text-gray-500 dark:text-gray-400 text-xs hover:underline"
                    @click="navigateToUser(props.comment.author.username)">
                    @{{ comment.author.username }}
                </span>
                <span class="text-gray-500 dark:text-gray-400 text-xs">
                    · {{ formatPostDate(comment.created_at) }}
                </span>
            </div>

            <p class="text-gray-800 dark:text-gray-200 text-sm leading-relaxed">
                {{ comment.content }}
            </p>
        </div>
    </div>
</template>
