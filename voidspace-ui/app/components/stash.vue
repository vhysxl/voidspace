<template>
    <!-- Content Area -->
    <div class="min-h-96">
        <!-- Tweets Tab -->
        <div v-if="activeTab === 'tweets'" class="divide-y divide-gray-100 dark:divide-gray-800">
            <article v-for="tweet in tweets" :key="tweet.id"
                class="px-4 py-3 hover:bg-gray-50/50 dark:hover:bg-gray-950/50 transition-colors cursor-pointer">
                <div class="flex gap-3">
                    <!-- Avatar -->
                    <img :src="user.avatar" :alt="user.displayName" class="w-10 h-10 rounded-full flex-shrink-0" />

                    <!-- Tweet Content -->
                    <div class="flex-1 min-w-0">
                        <!-- Header -->
                        <div class="flex items-center gap-2 text-sm">
                            <span class="font-bold text-black dark:text-white truncate">{{ user.displayName
                            }}</span>
                            <svg v-if="user.isVerified" class="w-4 h-4 text-blue-500 flex-shrink-0" fill="currentColor"
                                viewBox="0 0 24 24">
                                <path
                                    d="M22.46 6c-.77.35-1.6.58-2.46.69.88-.53 1.56-1.37 1.88-2.38-.83.5-1.75.85-2.72 1.05C18.37 4.5 17.26 4 16 4c-2.35 0-4.27 1.92-4.27 4.29 0 .34.04.67.11.98C8.28 9.09 5.11 7.38 3 4.79c-.37.63-.58 1.37-.58 2.15 0 1.49.75 2.81 1.91 3.56-.71 0-1.37-.2-1.95-.5v.03c0 2.08 1.48 3.82 3.44 4.21a4.22 4.22 0 0 1-1.93.07 4.28 4.28 0 0 0 4 2.98 8.521 8.521 0 0 1-5.33 1.84c-.34 0-.68-.02-1.02-.06C3.44 20.29 5.7 21 8.12 21 16 21 20.33 14.46 20.33 8.79c0-.19 0-.37-.01-.56.84-.6 1.56-1.36 2.14-2.23z" />
                            </svg>
                            <span class="text-gray-500 dark:text-gray-400 truncate">@{{ user.username }}</span>
                            <span class="text-gray-500 dark:text-gray-400">·</span>
                            <span class="text-gray-500 dark:text-gray-400 flex-shrink-0">{{ tweet.timestamp
                            }}</span>
                        </div>

                        <!-- Tweet Text -->
                        <p class="mt-1 text-black dark:text-white leading-normal">{{ tweet.content }}</p>

                        <!-- Actions -->
                        <div class="flex items-center justify-between mt-3 max-w-md">
                            <!-- Reply -->
                            <button
                                class="flex items-center gap-2 text-gray-500 dark:text-gray-400 hover:text-blue-500 transition-colors group">
                                <div
                                    class="p-2 rounded-full group-hover:bg-blue-50 dark:group-hover:bg-blue-950 transition-colors">
                                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z">
                                        </path>
                                    </svg>
                                </div>
                                <span class="text-sm">{{ tweet.replies }}</span>
                            </button>

                            <!-- Retweet -->
                            <button @click.stop="toggleRetweet(tweet)" :class="[
                                'flex items-center gap-2 transition-colors group',
                                tweet.isRetweeted
                                    ? 'text-green-500'
                                    : 'text-gray-500 dark:text-gray-400 hover:text-green-500'
                            ]">
                                <div
                                    class="p-2 rounded-full group-hover:bg-green-50 dark:group-hover:bg-green-950 transition-colors">
                                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15">
                                        </path>
                                    </svg>
                                </div>
                                <span class="text-sm">{{ tweet.retweets }}</span>
                            </button>

                            <!-- Like -->
                            <button @click.stop="toggleLike(tweet)" :class="[
                                'flex items-center gap-2 transition-colors group',
                                tweet.isLiked
                                    ? 'text-red-500'
                                    : 'text-gray-500 dark:text-gray-400 hover:text-red-500'
                            ]">
                                <div
                                    class="p-2 rounded-full group-hover:bg-red-50 dark:group-hover:bg-red-950 transition-colors">
                                    <svg class="w-4 h-4" :fill="tweet.isLiked ? 'currentColor' : 'none'"
                                        stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z">
                                        </path>
                                    </svg>
                                </div>
                                <span class="text-sm">{{ tweet.likes }}</span>
                            </button>

                            <!-- Share -->
                            <button
                                class="flex items-center gap-2 text-gray-500 dark:text-gray-400 hover:text-blue-500 transition-colors group">
                                <div
                                    class="p-2 rounded-full group-hover:bg-blue-50 dark:group-hover:bg-blue-950 transition-colors">
                                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                            d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.367 2.684 3 3 0 00-5.367-2.684z">
                                        </path>
                                    </svg>
                                </div>
                            </button>
                        </div>
                    </div>
                </div>
            </article>
        </div>

    </div>
</template>