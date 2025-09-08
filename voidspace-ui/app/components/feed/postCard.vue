<script setup lang="ts">
import { resolveAvatar } from '@/utils/userResolver';
import { formatPostDate } from '@/utils/dateFormater';
import type { Post } from '~/composables/useFeed';
import type { DropdownMenuItem } from '@nuxt/ui'

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
const postCall = usePosts()
const toast = useToast()
const selectedPostId = ref<number | null>(null)
const selectedPost = ref<Post | null>(null)

const openDeleteModal = (post: Post) => {
    selectedPostId.value = post.id
    selectedPost.value = post
    deleteModal.value = true
}

const openEditModal = (post: Post) => {
    selectedPost.value = post
    editModal.value = true
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


const handleLikePost = async () => {
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
        <article v-for="post in posts" :key="post.id" class="w-full hover:bg-neutral-500/10 transition-colors">
            <!-- Header -->
            <header class="flex items-center justify-between p-4">
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

                <UDropdownMenu :items="getMenuItems(post)" :content="{
                    align: 'center',
                    side: 'bottom',
                }">
                    <UButton color="neutral" variant="ghost" icon="i-heroicons-ellipsis-horizontal" size="sm" />
                </UDropdownMenu>
            </header>

            <!-- Content -->
            <div class="px-4 pb-3">
                <p class="text-gray-800 dark:text-gray-200 text-sm leading-relaxed">
                    {{ post.content }}
                </p>
            </div>

            <!-- Images -->
            <div v-if="post.post_images?.length" class="grid grid-cols-2 gap-2">
                <img v-for="(img, index) in post.post_images" :key="index" :src="img"
                    :alt="`Post by ${post.author.username} image ${index + 1}`" class="w-full h-64 object-cover" />
            </div>

            <!-- Actions -->
            <footer class="flex items-center justify-between px-4 py-3">
                <div class="flex items-center gap-6">
                    <!-- Like -->
                    <button class="flex items-center gap-2 group">
                        <UIcon name="i-heroicons-heart-solid"
                            class="w-5 h-5 text-gray-500 group-hover:text-red-500 transition-colors" />
                        <span class="text-sm text-gray-600 dark:text-gray-400">
                            {{ post.likes_count }}
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
    </div>

    <!-- Modals (outside loop) -->
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