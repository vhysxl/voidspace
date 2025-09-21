<script setup lang="ts">
import { computed } from 'vue'
import { formatPostDate } from '@/utils/dateFormater'
import type { User } from '@/types'

interface Props {
    author: User
    createdAt: string
}
const props = defineProps<Props>()

const displayName = computed(() => props.author.profile.display_name || props.author.username)
const username = computed(() => props.author.username)
const formattedDate = computed(() => formatPostDate(props.createdAt))
</script>

<template>
    <header class="flex items-center justify-between p-4">
        <div class="flex items-center gap-3">
            <UserAvatar :user="author" />
            <div>
                <h2 class="font-semibold text-gray-900 dark:text-gray-100">{{ displayName }}</h2>
                <div class="flex text-sm text-gray-500 dark:text-gray-400 gap-2">
                    <span>@{{ username }}</span>
                    <span>· {{ formattedDate }}</span>
                </div>
            </div>
        </div>
        <UButton variant="ghost" color="neutral" icon="i-heroicons-arrow-left" @click="$router.back()">Back</UButton>
    </header>
</template>
