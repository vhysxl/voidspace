<script setup lang="ts">
import { ref, computed } from 'vue'
import CreatePostInput from '~/components/feed/createPostInput.vue'

const { getVanillaFeed } = useFeed()
const posts = ref<Post[] | null>(null)
const loading = ref(false)
const route = useRoute()
const router = useRouter()

// Tab configuration with query params
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

// Get current active tab from query params
const activeTab = computed(() => {
  return route.query.tab === 'following' ? 'following' : 'for-you'
})

const switchTab = (tabKey: string) => {
  router.push({
    query: {
      ...route.query,
      tab: tabKey === 'for-you' ? undefined : tabKey
    }
  })
}

const fetchVanillaFeed = async () => {
  loading.value = true
  try {
    const res = await getVanillaFeed()
    posts.value = res.data.posts
  } catch (error) {
    console.error('Error fetching vanilla feed:', error)
  } finally {
    loading.value = false
  }
}

const fetchFollowingFeed = async () => {
  loading.value = true
  try {
    const res = await getVanillaFeed()
    posts.value = res.data.posts
  } catch (error) {
    console.error('Error fetching following feed:', error)
  } finally {
    loading.value = false
  }
}


// Watch for tab changes
watch(() => activeTab.value, (newTab) => {
  if (newTab === 'following') {
    fetchFollowingFeed()
  } else {
    fetchVanillaFeed()
  }
}, { immediate: true })
</script>

<template>
  <ClientOnly v-if="!$colorMode?.forced">
    <!-- Sticky Nav -->
    <nav class="sticky  top-0 z-20 bg-white dark:bg-black border-b border-gray-200 dark:border-gray-800">
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

    <!-- Feed Container -->
    <div class="min-h-screen scrollbar-hide">


      <!-- Loading -->
      <div v-if="loading" class="p-8 text-center">
        <ExtrasLoadingState :title="`Loading ${activeTab === 'following' ? 'Following' : 'For You'} Feed`" />
      </div>

      <!-- Empty -->
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
          <UButton v-if="activeTab === 'following'" to="/profile/yogi" color="neutral">
            Find People to Follow
          </UButton>
        </div>
      </div>


      <!-- Has posts -->
      <FeedPostCard v-else :posts="posts || []" />
    </div>

    <!-- Floating Button (Mobile Only) -->
    <button
      class="lg:hidden fixed bottom-24 right-6 z-30 bg-black dark:bg-white text-white dark:text-black rounded-full shadow-lg w-14 h-14 flex items-center justify-center transition-colors">
      <Icon name="i-lucide-plus" class="w-6 h-6" />
    </button>

    <!--Create Post -->
    <div class="hidden lg:flex fixed bottom-0 inset-x-0 z-20">
      <div class="mx-auto w-full max-w-screen-xl flex">
        <!-- Spacer left (biar sejajar sidebar) -->
        <div class="hidden lg:block w-64"></div>

        <!-- Actual input -->
        <div class="flex-1">
          <CreatePostInput />
        </div>

        <!-- Spacer right (biar sejajar sidebar kanan) -->
        <div class="hidden lg:block w-2/8"></div>
      </div>
    </div>

    <template #fallback>
      <ExtrasLoadingState title="Loading Feed" />
    </template>
  </ClientOnly>
</template>
