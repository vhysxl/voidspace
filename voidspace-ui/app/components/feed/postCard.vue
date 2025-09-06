<script setup lang="ts">
import { resolveAvatar } from '@/utils/userResolver';
import { formatPostDate } from '@/utils/dateFormater';
import type { Post } from '~/composables/useFeed';


defineProps<{
    posts: Post[]
}>()

// const toggleLike = (post: Post) => {
//     post.liked = !post.liked
//     post.likes = post.liked ? post.likes + 1 : post.likes - 1
// }

</script>

<template>
    <div class="divide-y border-b border-neutral-500  divide-neutral-500">
        <div v-for="post in posts" :key="post.id" class="w-full">
            <!-- Header -->
            <div class="flex items-center justify-between p-4">
                <div class="flex items-center gap-3">
                    <UAvatar :src="resolveAvatar(post.author.profile.avatar_url, post.author.username)"
                        :alt="post.author.username" size="md" />
                    <div>
                        <div class="flex items-center gap-2">
                            <h3 class="font-semibold text-sm">
                                {{ post.author.profile.display_name || post.author.username }}
                            </h3>
                        </div>
                        <p class="text-xs text-gray-500 dark:text-gray-400">
                            @{{ post.author.username }} • {{ formatPostDate(post.created_at) }}
                        </p>
                    </div>
                </div>

                <UButton color="neutral" variant="ghost" icon="i-heroicons-ellipsis-horizontal" size="sm" />
            </div>

            <!-- Content -->
            <div class="px-4 pb-3">
                <p class="text-gray-800 dark:text-gray-200 text-sm leading-relaxed">
                    {{ post.content }}
                </p>
            </div>

            <!-- Image -->
            <div v-if="post.post_images && post.post_images.length" class="grid grid-cols-2 gap-2">
                <div v-for="(img, index) in post.post_images" :key="index">
                    <img :src="img" :alt="'Post by ' + post.author.username + ' image ' + (index + 1)"
                        class="w-full h-64 object-cover " />
                </div>
            </div>

            <!-- Actions -->
            <div class="flex items-center justify-between px-4 py-3">
                <div class="flex items-center gap-6">
                    <!-- Like -->
                    <button class="flex items-center gap-2 group">
                        <UIcon :name="'i-heroicons-heart-solid'" :class="'text-gray-500 group-hover:text-red-500'"
                            class="w-5 h-5 transition-colors" />
                        <span class="text-sm text-gray-600 dark:text-gray-400">
                            {{ post.likes_count }}
                        </span>
                    </button>

                    <!-- Comment
                    <button class="flex items-center gap-2 group">
                        <UIcon name="i-heroicons-chat-bubble-oval-left"
                            class="w-5 h-5 text-gray-500 group-hover:text-blue-500 transition-colors" />
                        <span class="text-sm text-gray-600 dark:text-gray-400">
                            {{ post.come }}
                        </span>
                    </button> -->
                </div>

                <!-- Share -->
                <button class="group">
                    <UIcon name="i-heroicons-share"
                        class="w-5 h-5 text-gray-500 group-hover:text-gray-700 dark:group-hover:text-gray-300 transition-colors" />
                </button>
            </div>
        </div>
    </div>
</template>