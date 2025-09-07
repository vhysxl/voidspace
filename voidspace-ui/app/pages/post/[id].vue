<script setup lang="ts">
import { resolveAvatar } from '@/utils/userResolver'
import { formatPostDate } from '@/utils/dateFormater'

const route = useRoute()
const postId = route.params.id as string
const { getPost } = usePosts()

const { data: post, pending } = await useLazyFetch<Post>(`/api/posts/${postId}`, {
    key: `post-${postId}`
})



if (!post.value && !pending.value) {
    throw createError({
        statusCode: 404,
        statusMessage: 'Post not found'
    })
}

// SEO meta
useSeoMeta({

    description: () => post.value?.content || 'View this post on VoidSpace'
})
</script>

<template>
    <div class="max-w-2xl mx-auto">
        <!-- Loading -->
        <div v-if="pending" class="p-8">
            <ExtrasLoadingState title="Loading post" />
        </div>

        <!-- Post Content -->
        <article v-else-if="post" class="border-b border-neutral-500">
            <!-- Header -->
            <header class="flex items-center justify-between p-4">
                <div class="flex items-center gap-3">

                    <div>
                        <div class="flex items-center gap-2">

                        </div>

                    </div>
                </div>

                <!-- Back button -->
                <UButton variant="ghost" color="neutral" icon="i-heroicons-arrow-left" @click="$router.back()">
                    Back
                </UButton>
            </header>

            <!-- Content -->
            <div class="px-4 pb-4">
                <p class="text-gray-800 dark:text-gray-200 text-base leading-relaxed mb-4">
                    {{ post.content }}
                </p>

                <!-- Images -->
                <div v-if="post.post_images?.length" class="grid gap-2" :class="{
                    'grid-cols-1': post.post_images.length === 1,
                    'grid-cols-2': post.post_images.length > 1
                }">
                    <img v-for="(img, index) in post.post_images" :key="index" :src="img"
                        class="w-full h-auto object-cover rounded-lg border border-neutral-200 dark:border-neutral-700" />
                </div>
            </div>

            <!-- Actions -->
            <footer
                class="flex items-center justify-between px-4 py-4 border-t border-neutral-200 dark:border-neutral-700">
                <div class="flex items-center gap-6">
                    <!-- Like -->
                    <button class="flex items-center gap-2 group">
                        <UIcon name="i-heroicons-heart-solid"
                            class="w-5 h-5 text-gray-500 group-hover:text-red-500 transition-colors" />
                        <span class="text-sm text-gray-600 dark:text-gray-400">
                            {{ post.likes_count }}
                        </span>
                    </button>

                    <!-- Reply -->
                    <button class="flex items-center gap-2 group">
                        <UIcon name="i-heroicons-chat-bubble-oval-left"
                            class="w-5 h-5 text-gray-500 group-hover:text-blue-500 transition-colors" />
                        <span class="text-sm text-gray-600 dark:text-gray-400">
                            Reply
                        </span>
                    </button>
                </div>

                <!-- Share -->
                <button class="group">
                    <UIcon name="i-heroicons-share"
                        class="w-5 h-5 text-gray-500 group-hover:text-gray-700 dark:group-hover:text-gray-300 transition-colors" />
                </button>
            </footer>
        </article>

        <!-- Comments Section (placeholder) -->
        <div class="p-4 border-b border-neutral-500">
            <h2 class="font-semibold text-lg mb-4">Replies</h2>
            <div class="text-center text-gray-500 dark:text-gray-400 py-8">
                <p>No replies yet. Be the first to reply!</p>
            </div>
        </div>
    </div>
</template>