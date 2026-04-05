<script setup lang="ts">
import { ref } from "vue";
import { createPostSchema } from "@/validations";
import { useRouter } from "vue-router";

const modal = ref(false);
const router = useRouter();
const auth = useAuthStore();
const { state, remainingChars, isSubmitting, submitPost } = usePostForm();

const handleLoginRedirect = async () => {
    await router.push("/auth/login");
};

const handleOpenModal = () => {
    if (!auth.isLoggedIn) {
        handleLoginRedirect();
        return;
    }
    modal.value = true;
};
</script>

<template>
    <UModal v-model="modal" title="Create a post">
        <button @click="handleOpenModal"
            class="fixed bottom-24 right-4 sm:right-6 md:right-8 lg:right-42 xl:right-64 z-30 bg-black dark:bg-white text-white dark:text-black rounded-full w-14 h-14 flex items-center justify-center shadow-lg hover:scale-105 active:scale-95 transition-all">
            <Icon name="i-lucide-plus" class="w-6 h-6" />
        </button>

        <template v-if="isSubmitting" #content>
            <ExtrasLoadingState class="h-fit" title="Posting" />
        </template>
        <template v-else #content>
            <div class="p-4">
                <h2 class="text-xl font-semibold">Create a Post</h2>
                <p class="text-sm text-gray-500 dark:text-gray-400 mb-4">
                    Share something to the void.
                </p>

                <UForm :schema="createPostSchema" :state="state" @submit="submitPost" class="space-y-4">
                    <UFormField name="content">
                        <textarea v-model="state.content" placeholder="What's on your mind?"
                            class="w-full resize-none border rounded-md p-2 text-gray-800 dark:text-gray-200 bg-white dark:bg-black focus:outline-none focus:ring-2 focus:ring-cyan-500"
                            rows="4" :disabled="isSubmitting"></textarea>
                    </UFormField>

                    <UFormField class="w-fit" name="postImages">
                        <UFileUpload accept="image/*" multiple v-model="state.postImages" :loading="isSubmitting"
                            label="Upload Images (max 5)" />
                    </UFormField>

                    <div class="flex items-center justify-between">
                        <span :class="[
                            'text-xs',
                            remainingChars < 0 ? 'text-red-500' : remainingChars <= 20 ? 'text-yellow-500' : 'text-gray-400'
                        ]">
                            {{ remainingChars }} characters left
                        </span>
                        <UButton type="submit" color="neutral" :disabled="remainingChars < 0 || isSubmitting">
                            Post
                        </UButton>
                    </div>
                </UForm>
            </div>
        </template>
    </UModal>
</template>
