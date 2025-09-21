<script setup lang="ts">
import type { User } from '@/types'
import { resolveBanner } from '@/utils/userResolver';

interface Props {
    user: User
}

const props = defineProps<Props>()
</script>

<template>
    <div class="relative h-48 md:h-52 overflow-hidden shadow-md">
        <img :src="resolveBanner(props.user.profile.banner_url)" alt="Cover" class="w-full h-full object-cover" />
        <div class="absolute inset-0 bg-black/20"></div>
    </div>
</template>
