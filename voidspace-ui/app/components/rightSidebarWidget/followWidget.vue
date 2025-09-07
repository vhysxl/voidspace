<script setup lang="ts">
import { ref } from 'vue'

defineProps<{
    title: string
}>()

// Dummy data untuk MVP
const suggestedUsers = ref([
    {
        id: 1,
        displayName: 'Alice Johnson',
        username: 'alicej',
        avatarUrl: 'https://i.pravatar.cc/100?img=1',
        isFollowing: false
    },
    {
        id: 2,
        displayName: 'Bob Martin',
        username: 'bobmartin',
        avatarUrl: 'https://i.pravatar.cc/100?img=2',
        isFollowing: false
    },
    {
        id: 3,
        displayName: 'Charlie Kim',
        username: 'charliek',
        avatarUrl: 'https://i.pravatar.cc/100?img=3',
        isFollowing: true
    }
])

function followUser(userId: number) {
    const user = suggestedUsers.value.find(u => u.id === userId)
    if (user) {
        user.isFollowing = !user.isFollowing
        // TODO: call API follow/unfollow kalau backend sudah siap
    }
}

</script>

<template>
    <div class="bg-neutral-50 dark:bg-neutral-900 rounded-xl p-4">
        <h2 class="text-xl font-bold mb-4">{{ title }}</h2>
        <div class="space-y-4">
            <div v-for="user in suggestedUsers" :key="user.id" class="flex items-center justify-between">
                <!-- User info -->
                <div class="flex items-center space-x-3 flex-1 min-w-0">
                    <UAvatar size="md" :src="user.avatarUrl" :alt="user.displayName" />
                    <div class="flex flex-col min-w-0 flex-1">
                        <span class="font-semibold text-sm truncate">
                            {{ user.displayName }}
                        </span>
                        <span class="text-xs text-neutral-500 truncate">
                            @{{ user.username }}
                        </span>
                    </div>
                </div>

                <!-- Follow button -->
                <button @click="followUser(user.id)" :class="[
                    'px-4 py-1.5 text-sm font-medium rounded-full transition-colors',
                    user.isFollowing
                        ? 'bg-neutral-200 dark:bg-neutral-700 text-neutral-700 dark:text-neutral-300 hover:bg-neutral-300 dark:hover:bg-neutral-600'
                        : 'bg-neutral-950 dark:bg-white text-white dark:text-black hover:bg-neutral-800 dark:hover:bg-neutral-200'
                ]">
                    {{ user.isFollowing ? 'Following' : 'Follow' }}
                </button>
            </div>
        </div>

        <button class="text-blue-500 hover:text-blue-600 text-sm mt-3 font-medium">
            Show more
        </button>
    </div>
</template>
