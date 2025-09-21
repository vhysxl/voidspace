<script setup lang="ts">
import { ref, watch } from 'vue'
import CreatePostInput from '~/components/feed/CreatePostInput.vue'
import type { Post } from '@/types'
import PostCard from '~/components/feed/PostCard.vue'
import { useFeedAction } from '~/composables/useFeedAction'

const posts = ref<Post[] | null>(null)
const { feedState, activeTab, refreshFeed, setupInfiniteScroll, loadInitialPosts, loadMorePosts } = useFeedAction(posts)
const scrollContainer = ref<HTMLElement | null>(null);
const route = useRoute()
const router = useRouter()
setupInfiniteScroll(scrollContainer)

const tabs = [
  {
    key: 'for-you',
    label: 'For You'
  },
  {
    key: 'following',
    label: 'Following'
  }
]

const switchTab = (tabKey: string) => {
  router.push({
    query: {
      ...route.query,
      tab: tabKey === 'for-you' ? undefined : tabKey
    }
  })
}

onMounted(() => {
  loadInitialPosts();
});

// Watch for tab changes
watch(() => activeTab.value, (newTab) => {
  try {
    feedState.initialLoading = true;
  } catch (error) {

  }

  refreshFeed();
}, { immediate: true })

const handlePostCreated = (newPost: Post) => {
  // Only add to current active tab if it makes sense
  if (posts.value && activeTab.value === 'for-you') {
    posts.value.unshift(newPost)
  }
  // Don't add to Following tab unless user follows the author
  // OR refresh the current tab to get updated data
  else if (activeTab.value === 'following') {
    // Option 1: Refresh following feed to get proper data
    refreshFeed()

    // Option 2: Or just don't add it locally, let next refresh handle it
  }
}
const handlePostDeleted = (postId: number) => {
  // Remove from local array
  if (posts.value) {
    posts.value = posts.value.filter(p => p.id !== postId)
  }
}
</script>

<template>
  <ClientOnly v-if="!$colorMode?.forced">
    <!-- Sticky Nav -->
    <nav class="sticky top-0 z-20 bg-white dark:bg-black border-b border-gray-200 dark:border-gray-800">
      <div class="flex">
        <button v-for="tab in tabs" :key="tab.key" @click="switchTab(tab.key)" :class="[
          'flex-1 px-4 py-4 text-center font-medium text-sm transition-all relative hover:bg-neutral-50 dark:hover:bg-neutral-950',
          activeTab === tab.key
            ? 'text-black dark:text-white'
            : 'text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-300'
        ]">
          {{ tab.label }}
          <span v-if="activeTab === tab.key"
            class="absolute bottom-0 left-1/2 transform -translate-x-1/2 w-12 h-1 bg-black dark:bg-white rounded-full transition-all duration-200"></span>
        </button>
      </div>
    </nav>

    <div ref="scrollContainer" class="min-h-screen scrollbar-hide">

      <div v-if="feedState.initialLoading" class="p-8 text-center">
        <ExtrasLoadingState :title="`Loading ${activeTab === 'following' ? 'Following' : 'For You'} Feed`" />
      </div>

      <div v-else-if="!posts?.length" class="p-8 text-center">
        <div class="max-w-md mx-auto">
          <h3 class="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-2">
            {{ activeTab === 'following' ? 'No posts from people you follow' : 'No posts available' }}
          </h3>
          <p class="text-gray-500 dark:text-gray-400 mb-4">
            {{ activeTab === 'following'
              ? 'Follow some people to see their posts here.'
              : 'Check back later for new content.'
            }}
          </p>
          <UButton v-if="activeTab === 'following'" to="/user/yogi" color="neutral">
            Find People to Follow
          </UButton>
        </div>
      </div>

      <div v-else>
        <PostCard :posts="posts || []" @post-deleted="handlePostDeleted" />

        <div v-if="feedState.loadingMore" class="p-8 text-center">
          <div class="flex items-center justify-center space-x-2">
            <div class="animate-spin rounded-full h-5 w-5 border-b-2 border-neutral-500"></div>
            <span class="text-gray-500">Loading more posts...</span>
          </div>
        </div>

        <div v-else-if="feedState.isEndOfFeed" class="p-8 text-center">
          <p class="text-gray-500 text-sm">You've reached the end of the void</p>
        </div>
      </div>

      <div class="pb-14 lg:pb-36"></div>
    </div>

    <button
      class="lg:hidden fixed bottom-24 right-6 z-30 bg-black dark:bg-white text-white dark:text-black rounded-full shadow-lg w-14 h-14 flex items-center justify-center transition-colors">
      <Icon name="i-lucide-plus" class="w-6 h-6" />
    </button>

    <div class="hidden lg:flex fixed bottom-0 inset-x-0 z-20">
      <div class="mx-auto w-full max-w-screen-xl flex">
        <div class="hidden lg:block w-64"></div>

        <div class="flex-1">
          <CreatePostInput @post-created="handlePostCreated" />
        </div>

        <div class="hidden lg:block w-2/8"></div>
      </div>
    </div>

    <template #fallback>
      <ExtrasLoadingState title="Loading Feed" />
    </template>
  </ClientOnly>
</template>