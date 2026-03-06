<script setup lang="ts">
import { resolveAvatar } from '@/utils/userResolver'

const commentInput = ref<HTMLTextAreaElement | null>(null)
const isLikeSubmitting = ref(false)

const auth = useAuthStore()
const { toggleLike } = useToggleLike()
const route = useRoute()
const router = useRouter()
const { fetchComments, submitComment, state: commentState } = useCommentState()
const { fetchPost, state: postState } = usePostState()

const postIdParam = Array.isArray(route.params.id) ? route.params.id[0] : route.params.id
const postId = Number(postIdParam)

if (isNaN(postId)) {
    throw createError({
        statusCode: 400,
        statusMessage: 'Invalid post ID'
    })
}

const focusComment = () => {
    commentInput.value?.focus()
}

const handleSubmitComment = async () => {
    if (postState.post) {
        postState.post.comments_count++;
    }
    try {
        await submitComment({
            post_id: postId,
            content: commentState.newComment
        });
    } catch (error) {
        if (postState.post) {
            postState.post.comments_count--;
        }
    }
};

onMounted(async () => {
    await Promise.all([
        fetchPost(postId.toString()),
        fetchComments(postId.toString())
    ])
})
</script>

<template>
    <div class="max-w-2xl mx-auto">
        <div v-if="postState.pending" class="p-8">
            <ExtrasLoadingState title="Loading post" />
        </div>

        <article v-else-if="postState.post" class="border-b border-neutral-500">
            <PostHeader :author="postState.post.author" :created-at="postState.post.created_at" />
            <PostContent :post="postState.post" />
            <PostAction :post="postState.post" :is-like-submitting="isLikeSubmitting" @focus-comment="focusComment"
                @toggle-like="toggleLike" />
        </article>

        <div v-else class="p-8 text-center text-gray-500 dark:text-gray-400">
            Post not found or has been deleted.
        </div>

        <div v-if="auth.isLoggedIn && auth.user" class="p-4 border-b border-neutral-200 dark:border-neutral-700">
            <div class="flex gap-3">
                <UserAvatar size="lg" :user="auth.user" />
                <div class="flex-1">
                    <textarea id="comment-textarea" ref="commentInput" v-model="commentState.newComment"
                        placeholder="Write a reply..."
                        class="w-full resize-none border-0 bg-transparent text-gray-900 dark:text-gray-100 placeholder-gray-500 focus:outline-none text-sm"
                        rows="3" @keydown.ctrl.enter="handleSubmitComment" @keydown.meta.enter="handleSubmitComment">
                    </textarea>

                    <div class="flex items-center justify-between mt-3">
                        <span class="text-xs text-gray-400">Ctrl+Enter to reply</span>
                        <div class="flex gap-2">
                            <UButton color="neutral" size="xs"
                                :disabled="!commentState.newComment.trim() || commentState.isSubmittingComment"
                                :loading="commentState.isSubmittingComment" @click="handleSubmitComment">
                                Reply
                            </UButton>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div v-else class="p-4 border-b border-neutral-200 dark:border-neutral-700 ">
            <div class="text-center py-4">
                <p class="text-gray-600 dark:text-gray-400 mb-3">Join the conversation</p>
                <UButton class="hover:cursor-pointer" color="neutral" @click="router.push('/auth/login')" size="sm">
                    Sign in to reply
                </UButton>
            </div>
        </div>


        <!-- Comments Section -->
        <div class="border-b border-neutral-500">
            <div class="p-4">
                <h2 class="font-semibold text-lg mb-4">Replies</h2>

                <!-- Loading Comments -->
                <div v-if="commentState.pending" class="py-4">
                    <ExtrasLoadingState title="Loading Comments" />
                </div>

                <div v-else-if="commentState.comments.length > 0" class="space-y-4">
                    <CommentItem
                        v-for="comment in commentState.comments.slice().sort((a, b) => +new Date(b.created_at) - +new Date(a.created_at))"
                        :key="comment.id" :comment="comment" />
                </div>

                <div v-else class="text-center text-gray-500 dark:text-gray-400 py-8">
                    <p>No replies yet. Be the first to reply!</p>
                </div>
            </div>
        </div>
    </div>
</template>