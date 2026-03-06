<script setup lang="ts">
import { formatPostDate } from '@/utils/dateFormater';
import type { DropdownMenuItem } from '@nuxt/ui'
import type { Post } from '@/types';

defineProps<{
    posts: Post[]
}>()

const emit = defineEmits<{
    postDeleted: [postId: number]
    postLiked: [postId: number]
}>()

const editModal = ref(false)
const deleteModal = ref(false)
const isSubmitting = ref(false)
const postCall = usePost()
const toast = useToast()
const selectedPostId = ref<number | null>(null)
const selectedPost = ref<Post | null>(null)
const { toggleLike, isSubmitting: isLikeSubmitting } = useToggleLike()

const openDeleteModal = (post: Post) => {
    selectedPostId.value = post.id
    selectedPost.value = post
    deleteModal.value = true
}

const openEditModal = (post: Post) => {
    selectedPost.value = post
    editModal.value = true
}

const navigateToPost = async (postId: number) => {
    await navigateTo(`/post/${postId}`)
}

const navigateToUser = async (username: string, event: Event) => {
    await navigateTo(`/user/${username}`)
}

const getMenuItems = (post: Post): DropdownMenuItem[] => [
    {
        label: 'Edit Post',
        color: "neutral",
        onSelect: () => openEditModal(post)
    },
    {
        label: 'Delete Post',
        color: "error",
        onSelect: () => openDeleteModal(post)
    }
]

const handleDeletePost = async () => {
    if (isSubmitting.value || selectedPostId.value === null) return

    try {
        isSubmitting.value = true
        await postCall.deletePost(selectedPostId.value.toString())

        emit('postDeleted', selectedPostId.value)

        toast.add({
            title: "Post deleted",
            description: "Your post has been removed",
            color: "neutral",
        })

        deleteModal.value = false

    } catch (error: any) {
        toast.add({
            title: "Delete failed",
            description: error.message || "Failed to delete post",
            color: "error",
        })
    } finally {
        selectedPostId.value = null
        selectedPost.value = null
        isSubmitting.value = false
    }
}


const closeModal = () => {
    deleteModal.value = false
    selectedPostId.value = null
    selectedPost.value = null
}
</script>

<template>
    <div class="divide-y border-b border-neutral-500 divide-neutral-500">
        <article @click="navigateToPost(post.id)" v-for="post in posts" :key="post.id"
            class="w-full hover:cursor-pointer hover:bg-neutral-500/10 transition-colors">
            <header class="flex items-center justify-between p-4">
                <div class="flex items-center gap-3">
                    <UserAvatar size="xl" :user="post.author" />
                    <div>
                        <div class="flex items-center gap-2">
                            <h3 class="font-semibold text-sm cursor-pointer hover:underline"
                                @click="navigateToUser(post.author.username, $event)">
                                {{ post.author.profile.display_name || post.author.username }}
                            </h3>
                        </div>
                        <p class="text-xs text-gray-500 dark:text-gray-400">
                            <span class="cursor-pointer hover:underline"
                                @click="navigateToUser(post.author.username, $event)">
                                @{{ post.author.username }}
                            </span>
                            • {{ formatPostDate(post.created_at) }}
                        </p>
                    </div>
                </div>

                <UDropdownMenu :items="getMenuItems(post)" :content="{
                    align: 'center',
                    side: 'bottom',
                }">
                    <UButton color="neutral" variant="ghost" icon="i-heroicons-ellipsis-horizontal" size="sm" />
                </UDropdownMenu>
            </header>

            <div class="px-4 pb-3 cursor-pointer">
                <p class="text-gray-800 dark:text-gray-200 text-sm leading-relaxed">
                    {{ post.content }}
                </p>
            </div>

            <div v-if="post.post_images?.length" class="cursor-pointer" @click.stop>
                <UModal :ui="{
                    content: 'w-fit max-w-[1200px] max-h-[90vh] rounded-lg'
                }">
                    <UCarousel v-slot="{ item }" :dots="post.post_images.length > 1" :items="post.post_images"
                        class="w-full max-w-4xl mx-auto p-5">
                        <div class="w-full">
                            <img :src="item" :alt="`Post image`"
                                class="w-full h-auto max-h-[480px] md:max-h-[600px] object-cover rounded-lg"
                                loading="lazy">
                        </div>
                    </UCarousel>

                    <template #content>
                        <div class="w-full h-full flex items-center justify-center p-4">
                            <UCarousel v-slot="{ item }" :dots="post.post_images.length > 1" :items="post.post_images"
                                class="w-full flex items-center justify-center">
                                <div class="flex items-center justify-center w-full">
                                    <img :src="item" alt="Post image"
                                        class="max-w-full max-h-[80vh] w-auto h-auto object-contain rounded-lg"
                                        loading="lazy" />
                                </div>
                            </UCarousel>
                        </div>
                    </template>
                </UModal>
            </div>


            <footer class="flex items-center justify-between px-4 py-3 cursor-pointer" @click="navigateToPost(post.id)">
                <div class="flex items-center gap-6">
                    <button @click="toggleLike(post, $event)" class="flex items-center gap-2 hover:cursor-pointer group"
                        :disabled="isLikeSubmitting">
                        <UIcon name="i-heroicons-heart-solid"
                            :class="[!post.is_liked ? 'w-5 h-5 text-gray-500 group-hover:text-red-500 transition-colors' : 'w-5 h-5 text-red-500 group-hover:text-red-800 transition-colors']" />
                        <span class="text-sm text-gray-600 dark:text-gray-400">
                            {{ post.likes_count }}
                        </span>
                    </button>

                    <!-- Comment Button -->
                    <button class="flex items-center gap-2 group">
                        <UIcon name="i-heroicons-chat-bubble-oval-left"
                            class="w-5 h-5 text-gray-500 group-hover:text-blue-500 transition-colors" />
                        <span class="text-sm text-gray-600 dark:text-gray-400">
                            {{ post.comments_count }}
                        </span>
                    </button>
                </div>

            </footer>
        </article>
    </div>

    <UModal v-model:open="editModal" :overlay="false">
        <template #content>
            <div class="p-6">
                <h3 class="text-xl font-bold mb-4">Edit Post</h3>
                <p>Edit form here for: {{ selectedPost?.content }}</p>
            </div>
        </template>
    </UModal>

    <UModal v-model:open="deleteModal" :overlay="false">
        <template #content>
            <div class="p-6">
                <h3 class="text-xl font-bold text-gray-900 dark:text-gray-100 mb-3">
                    Delete post?
                </h3>

                <p class="text-gray-600 dark:text-gray-400 mb-6 leading-relaxed">
                    This action cannot be undone. Your post will be permanently removed from your profile and timeline.
                </p>

                <div class="flex gap-3 justify-end">
                    <UButton color="neutral" variant="ghost" @click="closeModal" :disabled="isSubmitting">
                        Cancel
                    </UButton>

                    <UButton @click="handleDeletePost" color="error" :loading="isSubmitting">
                        Delete
                    </UButton>
                </div>

            </div>

        </template>
    </UModal>
</template>