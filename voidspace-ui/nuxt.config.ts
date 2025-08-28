export default defineNuxtConfig({
  compatibilityDate: "2025-07-15",
  devtools: { enabled: true },
  modules: ["@nuxt/ui", "@nuxt/fonts"],
  css: ["~/assets/css/main.css"],
  fonts:{
    families: [
      {name: 'Gugi', provider: 'google'}
    ],
    defaults: {
      weights: [400],
      styles: ['normal'],
    }
  }
});