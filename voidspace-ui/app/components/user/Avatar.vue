<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { resolveAvatar } from '@/utils/userResolver'
import type { User } from '@/types'

interface Props {
    user: User
    size?: "md" | "3xs" | "2xs" | "xs" | "sm" | "lg" | "xl" | "2xl" | "3xl"
}
const props = defineProps<Props>()

const displayName = computed(() => props.user.profile.display_name || props.user.username)
const avatarUrl = computed(() => resolveAvatar(props.user.profile.avatar_url, displayName.value))
const username = computed(() => props.user.username)

const router = useRouter()
const navigateToUser = (username: string) => {
    router.push(`/user/${username}`)
}
</script>

<template>
    <UAvatar :src="avatarUrl" :alt="username" :size="props.size
        " class="cursor-pointer hover:opacity-80 transition-opacity" @click="navigateToUser(username)" />
</template>
