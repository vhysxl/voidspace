<script setup lang="ts">
import type { User } from '@/types'
import { formatJoinDate } from '@/utils/dateFormater';

interface Props {
    user: User;
}

const props = defineProps<Props>()

const joinDate = computed(() =>
    props.user.created_at ? formatJoinDate(props.user.created_at) : ''
)
</script>

<template>
    <div>
        <h1 class="text-2xl font-extrabold text-black dark:text-white flex items-center gap-2">
            {{ props.user?.profile.display_name }}
        </h1>
        <p class="text-gray-500 dark:text-gray-400">
            @{{ props.user?.username }}
        </p>
    </div>

    <p v-if="user.profile.bio" class="text-black dark:text-white leading-relaxed max-w-xl">
        {{ props.user.profile.bio }}
    </p>

    <div class="flex flex-wrap gap-x-6 gap-y-3 text-sm text-gray-500 dark:text-gray-400">
        <div v-if="props.user.profile.location" class="flex items-center gap-2">
            <UIcon name="i-ion-location" class="w-4 h-4" />
            <span>{{ props.user.profile.location }}</span>
        </div>
        <div class="flex items-center gap-2">
            <UIcon name="i-ion-calendar" class="w-4 h-4" />
            <span>Joined {{ joinDate }}</span>
        </div>
    </div>

    <div class="flex gap-8 text-sm">
        <button class="hover:underline transition-all">
            <span class="font-bold text-black dark:text-white">{{
                props.user.profile.following?.toLocaleString() || '0'
            }}</span>
            <span class="text-gray-500 dark:text-gray-400 ml-1">Following</span>
        </button>
        <button class="hover:underline transition-all">
            <span class="font-bold text-black dark:text-white">{{
                props.user.profile.followers?.toLocaleString() || '0'
            }}</span>
            <span class="text-gray-500 dark:text-gray-400 ml-1">Followers</span>
        </button>
    </div>
</template>
