<script setup lang="ts">
import type { User } from '@/types'
import { resolveAvatar, resolveBanner } from '@/utils/userResolver';

interface Props {
    user: User
}

const props = defineProps<Props>()
</script>

<template>
    <div class="relative">
        <img :src="resolveAvatar(props.user.profile.avatar_url, props.user.username)" alt="User avatar"
            class="w-32 h-32 md:w-36 md:h-36 rounded-md border-4 border-white dark:border-black bg-white dark:bg-black shadow-lg hover:scale-105 transition-transform duration-300" />
    </div>
</template>
