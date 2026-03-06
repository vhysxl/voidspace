<script setup lang="ts">
import type { Post } from '@/types';

interface Props {
    post: Post
}

const props = defineProps<Props>()
</script>

<template>
    <div class="px-4 pb-4">
        <!-- Post text -->
        <p class="text-gray-800 dark:text-gray-200 text-base leading-relaxed mb-4">
            {{ props.post.content }}
        </p>

        <!-- Images -->
        <div v-if="props.post.post_images?.length" class="cursor-pointer" @click.stop>
            <UModal :ui="{ content: 'w-fit max-w-[1200px] max-h-[90vh] rounded-lg' }">
                <UCarousel v-slot="{ item }" :dots="props.post.post_images.length > 1" :items="props.post.post_images"
                    class="w-full max-w-4xl mx-auto p-5 mb-10">
                    <div class="w-full">
                        <img :src="item" alt="Post image"
                            class="w-full h-auto max-h-[480px] md:max-h-[600px] object-cover rounded-lg"
                            loading="lazy" />
                    </div>
                </UCarousel>

                <template #content>
                    <div class="w-full h-full flex items-center justify-center p-4">
                        <UCarousel v-slot="{ item }" :dots="props.post.post_images.length > 1"
                            :items="props.post.post_images" class="w-full flex items-center justify-center">
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
    </div>
</template>