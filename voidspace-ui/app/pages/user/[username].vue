<script setup lang="ts">
import { UModal } from '#components'
import type { FormSubmitEvent } from '@nuxt/ui';
import * as v from "valibot";
import { resolveAvatar, resolveBanner } from "@/utils/userResolver"

definePageMeta({
    layout: "default",
})

const route = useRoute()
const usersApi = useUsers()
const user = useAuthStore()
const toast = useToast();
const { uploadFile } = useUpload()
const { updateProfile } = useUsers()
const username = route.params.username as string
const userData = ref<User | null>(null)
const isLoadingProfile = ref(false)
const isSubmitting = ref(false)
const error = ref<string | null>(null)

const isOwner = computed(() => {
    return user.user?.username === username
})

//todo: move to utilities
const formattedJoinDate = computed(() => {
    if (!userData.value?.created_at) return ''

    const date = new Date(userData.value.created_at)
    return date.toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'long'
    })
})

const profileTabs = [
    { label: "Posts", href: `/user/${username}` },
    { label: "Media", href: `/user/${username}/media` },
    { label: "Likes", href: `/user/${username}/likes` },
]

const fetchUserProfile = async () => {
    isLoadingProfile.value = true
    error.value = null

    try {
        const res = await usersApi.getUser(username)
        if (res.success && res.data) {
            userData.value = res.data
        } else {
            throw new Error('Failed to fetch user data')
        }
    } catch (err) {
        console.error("Failed to fetch user profile")
        error.value = err instanceof Error ? err.message : 'An error occurred'
    } finally {
        isLoadingProfile.value = false
    }
}

const retry = async () => {
    await fetchUserProfile()
}

// function to handle ownership data
const currentProfile = computed(() => {
    if (!userData.value) return null
    if (isOwner.value && user.user) {
        return {
            displayName: user.user.profile.displayName || userData.value.profile.display_name || userData.value.username,
            bio: user.user.profile.bio || userData.value.profile.bio,
            avatarUrl: resolveAvatar(user.user.profile.avatarUrl || userData.value.profile.avatar_url, userData.value.username),
            bannerUrl: resolveBanner(user.user.profile.bannerUrl || userData.value.profile.banner_url),
            location: user.user.profile.location || userData.value.profile.location,
        }
    }

    return {
        displayName: userData.value.profile.display_name || userData.value.username,
        bio: userData.value.profile.bio,
        avatarUrl: resolveAvatar(userData.value.profile.avatar_url, userData.value.username),
        bannerUrl: resolveBanner(userData.value.profile.banner_url),
        location: userData.value.profile.location,
    }
})

onMounted(() => {
    fetchUserProfile()
})

const schema = v.object({
    displayName: v.optional(
        v.pipe(
            v.string(),
            v.maxLength(20, 'Display name must be 20 characters or less')
        )
    ),
    bio: v.optional(
        v.pipe(
            v.string(),
            v.maxLength(160, 'Bio must be 160 characters or less')
        )
    ),
    location: v.optional(
        v.pipe(
            v.string(),
            v.maxLength(50, 'Location must be 50 characters or less')
        )
    ),
    avatar: v.optional(
        v.pipe(
            v.file(),
            v.mimeType(['image/jpeg', 'image/png', 'image/webp'], 'Format must be JPEG, PNG, or WebP'),
            v.maxSize(2 * 1024 * 1024, 'Avatar must be maximum 2MB')
        )
    ),
    banner: v.optional(
        v.pipe(
            v.file(),
            v.mimeType(['image/jpeg', 'image/png', 'image/webp'], 'Format must be JPEG, PNG, or WebP'),
            v.maxSize(5 * 1024 * 1024, 'Banner must be maximum 5MB')
        )
    )
});

const state = reactive({
    displayName: userData.value?.profile.display_name || "",
    bio: userData.value?.profile.bio || "",
    location: userData.value?.profile.location || "",
    avatar: undefined as File | undefined,
    banner: undefined as File | undefined,
});

const resetForm = () => {
    if (!userData.value) return
    state.displayName = userData.value.profile.display_name || ""
    state.bio = userData.value.profile.bio || ""
    state.location = userData.value.profile.location || ""
    state.avatar = undefined
    state.banner = undefined
}

const onSubmit = async (event: FormSubmitEvent<any>) => {
    if (isSubmitting.value) return

    isSubmitting.value = true
    // Prepare an array of promises, each promise resolves to a [key, value] tuple
    // Example: ["avatar", uploadResult] or ["banner", uploadResult]
    try {
        const uploadPromises: Promise<[string, any]>[] = [];

        // If avatar exists, push its upload promise with identifier "avatar"
        if (event.data.avatar) {
            uploadPromises.push(
                uploadFile(event.data.avatar).then(res => ["avatar", res])
            );
        }

        // If banner exists, push its upload promise with identifier "banner"
        if (event.data.banner) {
            uploadPromises.push(
                uploadFile(event.data.banner).then(res => ["banner", res])
            );
        }

        // Wait for all uploads to finish, results will be an array of [key, value] pairs
        // Example: [["avatar", {...}], ["banner", {...}]]
        const results = await Promise.all(uploadPromises);

        // convert it to an object (hashmap):
        const mapped = Object.fromEntries(results);

        // Use state.displayName if it's not null/undefined,
        // otherwise fall back to user.user?.profile.displayName.
        // Unlike `||`, this allows empty string ("") to be a valid value.
        const profileData: UpdateProfileRequest = {
            display_name: state.displayName ?? user.user?.profile.displayName,
            bio: state.bio ?? user.user?.profile.bio,
            avatar_url: mapped.avatar ?? user.user?.profile.avatarUrl,
            banner_url: mapped.banner ?? user.user?.profile.bannerUrl,
            location: state.location ?? user.user?.profile.location,
        }

        await updateProfile(profileData)
        await user.refreshUser()

    } catch (error: any) {
        toast.add({
            title: "Update Failed",
            description: error.message || "Failed to update profile, please try again later",
            color: "error",
        })
    } finally {
        isSubmitting.value = false
    }

}

</script>

<template>
    <div class="max-w-screen bg-white dark:bg-black min-h-screen">
        <!-- Loading State -->
        <extras-loading-state v-if="isLoadingProfile" title="Loading profile" />
        <!-- Error State -->
        <extras-error-state v-if="error && !userData" :retry="retry" error="Failed to load user" />

        <!-- Profile Content -->
        <div v-else-if="userData">
            <!-- Cover Image -->
            <div class="relative h-48 md:h-52 bg-gray-300 dark:bg-gray-700 overflow-hidden  shadow-md">
                <img :src="currentProfile?.bannerUrl" alt="Cover" class="w-full h-full object-cover" />
                <div class="absolute inset-0 bg-black/20"></div>
            </div>

            <!-- Profile Header -->
            <div class="relative px-6 pb-6">
                <!-- Avatar + Actions -->
                <div class="flex justify-between items-end -mt-20 mb-6">
                    <!-- Avatar -->
                    <div class="relative">
                        <img :src="currentProfile?.avatarUrl" alt="User avatar"
                            class="w-32 h-32 md:w-36 md:h-36 rounded-md border-4 border-white dark:border-black bg-white dark:bg-black shadow-lg hover:scale-105 transition-transform duration-300" />
                        <!-- Loading overlay for avatar -->
                        <div v-if="isLoadingProfile"
                            class="absolute inset-0 bg-gray-200 dark:bg-gray-700 rounded-full animate-pulse"></div>
                    </div>

                    <!-- Action Buttons -->
                    <div class="flex gap-3 mb-3">
                        <template v-if="isOwner">
                            <UModal>
                                <button :disabled="isSubmitting" @click="resetForm"
                                    class="px-5 py-2 rounded-full font-semibold text-sm border dark:border-gray-600 text-black dark:text-white hover:bg-gray-100 dark:hover:bg-neutral-800 transition-colors shadow-sm disabled:opacity-50 disabled:cursor-not-allowed">
                                    Edit profile
                                </button>
                                <template #content>
                                    <UForm :schema="schema" :state="state" @submit="onSubmit"
                                        class="space-y-4 overflow-auto p-8">
                                        <UFormField label="Display Name" name="displayName">
                                            <UInput v-model="state.displayName" color="neutral"
                                                placeholder="Enter your display name" class="w-full" size="lg" />
                                        </UFormField>

                                        <UFormField label="Bio" name="bio">
                                            <UTextarea v-model="state.bio" color="neutral"
                                                placeholder="Write something about yourself" class="w-full" size="lg" />
                                        </UFormField>

                                        <UFormField label="Avatar" name="avatarUrl">
                                            <div class="space-y-2">
                                                <UFileUpload accept="image/*" label="Upload Avatar"
                                                    class="min-h-36 w-1/2" :loading="isSubmitting"
                                                    v-model="state.avatar" />
                                            </div>
                                        </UFormField>

                                        <UFormField label="Banner" name="bannerUrl">
                                            <div class="space-y-2">
                                                <UFileUpload accept="image/*" label="Upload Banner"
                                                    class="min-h-36 w-full" :loading="isSubmitting"
                                                    v-model="state.banner" />
                                            </div>
                                        </UFormField>

                                        <UFormField label="Location" name="location">
                                            <UInput v-model="state.location" color="neutral" placeholder="Your location"
                                                class="w-full" size="lg" />
                                        </UFormField>

                                        <UButton :disabled="isSubmitting" :loading="isSubmitting" type="submit"
                                            color="neutral" size="lg" class="w-full mt-2 justify-center">
                                            {{ isSubmitting ? 'Saving...' : 'Save Changes' }}
                                        </UButton>
                                    </UForm>
                                </template>

                            </UModal>
                        </template>
                        <template v-else>
                            <button :disabled="isLoadingProfile" :class="[
                                'px-5 py-2 rounded-full font-semibold text-sm transition-all duration-200 shadow-sm disabled:opacity-50 disabled:cursor-not-allowed',
                            ]">
                                <!-- Dynamic follow state nanti bisa ditambahkan -->
                                Follow
                            </button>
                        </template>
                    </div>
                </div>

                <!-- Profile Info -->
                <div class="space-y-4">
                    <!-- Name & Username -->
                    <div>
                        <h1 class="text-2xl font-extrabold text-black dark:text-white flex items-center gap-2">
                            {{ currentProfile?.displayName }}
                        </h1>
                        <p class="text-gray-500 dark:text-gray-400">
                            @{{ userData.username }}
                        </p>
                    </div>

                    <!-- Bio -->
                    <p v-if="userData.profile.bio" class="text-black dark:text-white leading-relaxed max-w-xl">
                        {{ currentProfile?.bio }}
                    </p>

                    <!-- Metadata -->
                    <div class="flex flex-wrap gap-x-6 gap-y-3 text-sm text-gray-500 dark:text-gray-400">
                        <div v-if="userData.profile.location" class="flex items-center gap-2">
                            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                                <path d="M12 2C8.13 2 5 5.13 5 9c0 5.25 7 13 7 13s7-7.75 7-13c0-3.87-3.13-7-7-7zm0 9.5c-1.38 
                    0-2.5-1.12-2.5-2.5s1.12-2.5 2.5-2.5 
                    2.5 1.12 2.5 2.5-1.12 2.5-2.5 2.5z" />
                            </svg>
                            <span>{{ currentProfile?.location }}</span>
                        </div>
                        <div class="flex items-center gap-2">
                            <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                                <path d="M7 4V2C7 1.45 7.45 1 8 1S9 1.45 9 
                    2V4H15V2C15 1.45 15.45 1 16 1S17 
                    1.45 17 2V4H19C20.1 4 21 4.9 21 
                    6V20C21 21.1 20.1 22 19 22H5C3.9 
                    22 3 21.1 3 20V6C3 4.9 3.9 4 
                    5 4H7ZM5 10V20H19V10H5ZM5 6V8H19V6H5Z" />
                            </svg>
                            <span>Joined {{ formattedJoinDate }}</span>
                        </div>
                    </div>

                    <!-- Following Stats -->
                    <div class="flex gap-8 text-sm">
                        <button class="hover:underline transition-all">
                            <span class="font-bold text-black dark:text-white">{{
                                userData.profile.following?.toLocaleString() || '0'
                            }}</span>
                            <span class="text-gray-500 dark:text-gray-400 ml-1">Following</span>
                        </button>
                        <button class="hover:underline transition-all">
                            <span class="font-bold text-black dark:text-white">{{
                                userData.profile.followers?.toLocaleString() || '0'
                            }}</span>
                            <span class="text-gray-500 dark:text-gray-400 ml-1">Followers</span>
                        </button>
                    </div>
                </div>
            </div>

            <!-- Navigation Tabs -->
            <NavNavigation :tabs="profileTabs" class="border-b dark:border-gray-700" />
            <NuxtPage />
        </div>
    </div>
</template>