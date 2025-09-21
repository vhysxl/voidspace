<script setup lang="ts">
import { resolveAvatar } from '@/utils/userResolver'
import { ref } from 'vue'
import * as v from "valibot";
import type { CreatePostReq, Post } from '@/types';

const { uploadFile } = useUpload()
const { createPost } = usePost()
const router = useRouter()
const maxLength = 240
const remainingChars = computed(() => maxLength - state.content.length)
const auth = useAuthStore()
const isPosting = ref(false)
const toast = useToast();

const schema = v.object({
    content: v.optional(
        v.pipe(
            v.string(),
            v.maxLength(240, 'Display name must be 240 characters or less')
        )
    ),
    postImages: v.optional(
        v.array(
            v.pipe(
                v.file(),
                v.mimeType(['image/jpeg', 'image/png', 'image/webp'], 'Format must be JPEG, PNG, or WebP'),
                v.maxSize(5 * 1024 * 1024, 'File must be maximum 5MB'),
            )
        )
    )
});

//TODO: STUPID VALIBOT NEXT MAKE CUSTOM ERROR

const emit = defineEmits<{
    postCreated: [post: Post]
}>()

const handleSubmit = async () => {
    if (isPosting.value) return
    if (state.postImages) {
        if (state.postImages.length > 5) {
            //return error
        }
    }

    try {
        isPosting.value = true
        await v.parseAsync(schema, state);

        const uploadedUrls: string[] = []

        if (state.postImages) {
            const uploadPromises = state.postImages.map(image => uploadFile(image))
            const results = await Promise.all(uploadPromises)
            uploadedUrls.push(...results)
        }

        const createPostReq: CreatePostReq = {
            content: state.content || "",
            post_images: uploadedUrls
        }

        const newPost = await createPost(createPostReq)

        // fix later
        emit('postCreated', newPost.data as Post)

        state.content = '';
        state.postImages = undefined;
    } catch (error: any) {
        toast.add({
            title: "Update Failed",
            description: error.message || "Failed to update profile, please try again later",
            color: "error",
        })
    } finally {
        isPosting.value = false
    }
}

const handleLoginRedirect = async () => {
    await router.push('/auth/login')
}

const state = reactive({
    content: "",
    postImages: undefined as File[] | undefined,
});
</script>

<template>
    <div v-if="auth.isLoggedIn" :class="{ 'opacity-50': isPosting }"
        class="border-t border-neutral-500 bg-white dark:bg-black">
        <div class="flex items-start gap-3 p-4">
            <UAvatar :src="resolveAvatar(auth.user?.profile.avatar_url, auth.user?.username!)"
                :alt="auth.user?.username" size="md" />

            <div class="flex-1">
                <UForm :schema="schema" :state="state" @submit="handleSubmit">
                    <textarea :disabled="isPosting" v-model="state.content" placeholder="Post Something to the void..."
                        class="w-full resize-none border-0 bg-transparent text-gray-800 dark:text-gray-200 text-base placeholder-gray-500 dark:placeholder-gray-400 focus:ring-0 focus:outline-none"
                        :class="{ 'cursor-not-allowed': isPosting, 'cursor-text': !isPosting }"></textarea>

                    <div class="flex items-center justify-between ">
                        <UFileUpload accept="image/*" :loading="isPosting" multiple v-model="state.postImages" size="xs"
                            variant="button" />

                        <div class="flex items-center gap-3">
                            <span :class="[
                                'text-xs',
                                remainingChars < 0 ? 'text-red-500' : remainingChars <= 20 ? 'text-yellow-500' : 'text-gray-400'
                            ]">
                                {{ remainingChars }}
                            </span>

                            <UButton color="neutral" size="sm" :disabled="remainingChars < 0 || isPosting"
                                type="submit">
                                Post
                            </UButton>
                        </div>
                    </div>
                </UForm>
            </div>
        </div>
    </div>

    <div v-else
        class="border-t border-neutral-500 bg-white dark:bg-black flex justify-center items-center min-h-[150px]">
        <div class="flex flex-col items-center gap-3 p-4">
            <p class="text-gray-500 dark:text-gray-400 text-base mb-3 text-center">
                Sign in to start posting to the void
            </p>
            <UButton class="hover:cursor-pointer" size="md" color="neutral" @click="handleLoginRedirect">
                Sign in to post
            </UButton>
        </div>
    </div>

</template>