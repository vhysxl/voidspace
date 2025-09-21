<script setup lang="ts">
import { ref, computed } from 'vue'
import { resolveAvatar } from '@/utils/userResolver'
import type { User } from '@/types'

interface Props {
    user: User,
    value: string,
    isSubmitting: boolean
}
const props = defineProps<Props>()
const emit = defineEmits<{
    (e: 'submit', content: string): void
}>()

const value = ref(props.value)
const textarea = ref<HTMLTextAreaElement | null>(null)
const avatarUrl = computed(() => resolveAvatar(props.user.profile.avatar_url, props.user.profile.display_name || props.user.username))

const submit = () => {
    if (!value.value.trim()) return
    emit('submit', value.value)
}
</script>


<template>
    <div class="flex gap-3">
        <UAvatar :src="avatarUrl" size="md" />
        <div class="flex-1">
            <textarea v-model="value" ref="textarea" placeholder="Write a reply..."
                class="w-full resize-none border-0 bg-transparent text-gray-900 dark:text-gray-100 placeholder-gray-500 focus:outline-none text-sm"
                rows="3" @keydown.ctrl.enter="submit" @keydown.meta.enter="submit"></textarea>
            <div class="flex items-center justify-between mt-3">
                <span class="text-xs text-gray-400">Ctrl+Enter to reply</span>
                <div class="flex gap-2">
                    <UButton color="neutral" size="xs" :disabled="!value.trim() || isSubmitting" :loading="isSubmitting"
                        @click="submit">Reply</UButton>
                </div>
            </div>
        </div>
    </div>
</template>
