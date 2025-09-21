<script setup lang="ts">
import { isOwner } from "@/utils/isOwner"

definePageMeta({
    layout: "default",
})

const route = useRoute()
const auth = useAuthStore()
const username = route.params.username as string
const owner = computed(() => isOwner(username, auth.user))
const isSubmitting = ref(false)

const { userData, isLoadingProfile, error, fetchUserProfile } = useUserProfile();
const { state: updateState, resetForm, onSubmit } = useProfileForm(userData, isSubmitting)
const { deleteConfirmation, deleteAccount, canDelete } = useAccountDeletion(userData, isSubmitting);
const { handleFollowToggle } = useFollowActions(userData, isSubmitting);

const user = computed(() => owner.value ? auth.user : userData.value)

const retry = async () => {
    await fetchUserProfile(username)
}

onMounted(() => {
    fetchUserProfile(username)
})

const profileTabs = [
    { label: "Posts", href: `/user/${username}` },
    { label: "Media", href: `/user/${username}/media` },
    { label: "Likes", href: `/user/${username}/likes` },
]
</script>

<template>
    <div class="max-w-screen min-h-screen">
        <extras-loading-state v-if="isLoadingProfile" title="Loading profile" />
        <extras-error-state v-if="error && !userData" :retry="retry" error="Failed to load user" />
        <div v-else-if="user">
            <UserBannerImage :user="user" />
            <div class="relative px-6 pb-6">
                <div class="flex justify-between items-end -mt-20 mb-6">
                    <UserAvatarProfile :user="user" />
                    <div class="flex gap-3 mb-3 mt-24 md:mt-0">
                        <template v-if="owner">
                            <UserEditProfileModal :onSubmit="onSubmit" :update-state="updateState" :user="user"
                                :is-submitting="isSubmitting" :reset-form="resetForm"
                                :delete-confirmation="deleteConfirmation" :delete-account="deleteAccount"
                                :can-delete="canDelete" :user-data="userData" />
                        </template>
                        <template v-else>
                            <UserFollowButton :user="user" :isLoadingProfile="isLoadingProfile"
                                :isSubmitting="isSubmitting" :handle-follow-toggle="handleFollowToggle" />
                        </template>
                    </div>
                </div>
                <div class="space-y-4">
                    <UserMetaData :user="user" />
                </div>
            </div>

            <!-- Navigation Tabs -->
            <NavNavigation :tabs="profileTabs" class="border-b dark:border-gray-700" />
            <NuxtPage />
        </div>
    </div>
</template>