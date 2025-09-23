    <script setup lang="ts">
    import type { User } from '@/types'
    import { editProfileSchema } from '@/validations'

    const modal = ref(false)

    interface UpdateState {
        displayName: string;
        bio: string;
        location: string;
        avatar: File | undefined;
        banner: File | undefined;
    }

    interface Props {
        user: User
        isSubmitting: boolean
        resetForm: () => void
        onSubmit: (event: any) => void
        deleteConfirmation: string
        deleteAccount: () => void
        canDelete: boolean
        userData: User | null
        updateState: UpdateState
    }
    const props = defineProps<Props>()
</script>

    <template>
        <UModal>
            <button :disabled="props.isSubmitting" @click="props.resetForm"
                class="px-5 py-2 md:mt-0 rounded-full font-semibold text-sm border dark:border-gray-600 text-black dark:text-white hover:bg-gray-100 dark:hover:bg-neutral-800 transition-colors shadow-sm disabled:opacity-50 disabled:cursor-not-allowed">
                Account Settings
            </button>
            <template #content>
                <div class="max-w-2xl mx-auto max-h-[90vh] overflow-hidden flex flex-col">
                    <!-- Header -->
                    <div class="flex-shrink-0 px-6 py-4 border-b dark:border-gray-700">
                        <h2 class="text-xl font-semibold text-gray-900 dark:text-white">
                            Account Settings
                        </h2>
                        <p class="text-sm text-gray-500 dark:text-gray-400 mt-1">
                            Manage your profile information and account settings
                        </p>
                    </div>
                    <!-- Scrollable Content -->
                    <div class="flex-1 overflow-y-auto">
                        <!-- Profile Settings -->
                        <div class="px-6 py-6">
                            <h3 class="text-lg font-medium text-gray-900 dark:text-white mb-4">
                                Profile Information
                            </h3>
                            <UForm :schema="editProfileSchema" :state="props.updateState" @submit="onSubmit"
                                class="space-y-4">
                                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                                    <UFormField label="Display Name" name="displayName" class="md:col-span-2">
                                        <UInput v-model="updateState.displayName" color="neutral"
                                            placeholder="Enter your display name" class="w-full" size="lg" />
                                    </UFormField>
                                    <UFormField label="Bio" name="bio" class="md:col-span-2">
                                        <UTextarea v-model="updateState.bio" color="neutral"
                                            placeholder="Write something about yourself" class="w-full" size="lg"
                                            :rows="3" />
                                    </UFormField>

                                    <UFormField label="Location" name="location" class="md:col-span-2">
                                        <UInput v-model="updateState.location" color="neutral"
                                            placeholder="Your location" class="w-full" size="lg" />
                                    </UFormField>

                                    <UFormField label="Avatar" name="avatarUrl">
                                        <UFileUpload accept="image/*" label="Upload Avatar" class="min-h-32"
                                            :loading="props.isSubmitting" v-model="updateState.avatar" />
                                        <p class="text-xs text-gray-500 mt-1">Max 2MB • JPEG, PNG, WebP</p>
                                    </UFormField>
                                    <UFormField label="Banner" name="bannerUrl">
                                        <UFileUpload accept="image/*" label="Upload Banner" class="min-h-32"
                                            :loading="props.isSubmitting" v-model="updateState.banner" />
                                        <p class="text-xs text-gray-500 mt-1">Max 5MB • JPEG, PNG, WebP</p>
                                    </UFormField>
                                </div>

                                <div class="pt-4">
                                    <UButton :disabled="props.isSubmitting" :loading="props.isSubmitting" type="submit"
                                        color="neutral" size="lg" class="w-full justify-center">
                                        {{ props.isSubmitting ? 'Saving...' : 'Save Changes' }}
                                    </UButton>
                                </div>
                            </UForm>
                        </div>
                        <!-- Danger Zone -->
                        <div class="px-6 py-6 border-t dark:border-gray-700">
                            <div class="max-w-3xl">
                                <h3 class="text-lg font-medium text-red-600 dark:text-red-400 mb-2">
                                    Danger Zone
                                </h3>
                                <p class="text-sm text-gray-600 dark:text-gray-400 mb-4">
                                    These actions are irreversible. Please proceed with caution.
                                </p>

                                <div
                                    class="bg-red-50 dark:bg-red-950/20 border border-red-200 dark:border-red-800 rounded-lg p-4">
                                    <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3">
                                        <div>
                                            <h4 class="text-sm font-medium text-red-800 dark:text-red-300">
                                                Delete Account
                                            </h4>
                                            <p class="text-xs text-red-600 dark:text-red-400 mt-1">
                                                Permanently remove your account and all associated data
                                            </p>
                                        </div>
                                        <!-- Delete Account Modal -->
                                        <UModal v-model="modal" :dismissible="false">
                                            <UButton variant="outline" color="error" size="sm" class="flex-shrink-0">
                                                Delete Account
                                            </UButton>
                                            <template #content>
                                                <div class="p-6">
                                                    <div class="flex items-center gap-3 mb-4">
                                                        <div
                                                            class="flex-shrink-0 w-10 h-10 bg-red-100 dark:bg-red-900/20 rounded-full flex items-center justify-center">
                                                            <svg class="w-5 h-5 text-red-600 dark:text-red-400"
                                                                fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                                                <path stroke-linecap="round" stroke-linejoin="round"
                                                                    stroke-width="2"
                                                                    d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
                                                            </svg>
                                                        </div>
                                                        <div>
                                                            <h3
                                                                class="text-lg font-semibold text-gray-900 dark:text-white">
                                                                Delete Account
                                                            </h3>
                                                            <p class="text-sm text-gray-500 dark:text-gray-400">
                                                                This action cannot be undone
                                                            </p>
                                                        </div>
                                                    </div>
                                                    <p class="text-gray-600 dark:text-gray-400 text-sm mb-6">
                                                        Are you absolutely sure you want to delete your account?
                                                        Type your username <strong>{{ props.userData?.username
                                                            }}</strong>
                                                        to confirm.
                                                    </p>
                                                    <!-- Confirmation Input -->
                                                    <div class="mb-6">
                                                        <UInput color="neutral"
                                                            placeholder="Enter your username to confirm"
                                                            v-model="props.deleteConfirmation" class="w-full" />
                                                    </div>
                                                    <div class="flex flex-col sm:flex-row gap-3 sm:justify-end">
                                                        <UButton variant="ghost" size="sm" @click="modal = false"
                                                            color="neutral" class="sm:order-1">
                                                            Cancel
                                                        </UButton>
                                                        <UButton color="error" size="sm" :loading="props.isSubmitting"
                                                            :disabled="props.deleteConfirmation !== props.userData?.username || !props.canDelete"
                                                            @click="props.deleteAccount" class="sm:order-2">
                                                            Delete My Account
                                                        </UButton>
                                                    </div>
                                                </div>
                                            </template>
                                        </UModal>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </template>
        </UModal>
    </template>