export default defineNuxtConfig({
  compatibilityDate: "2025-07-15",
  devtools: { enabled: true },
  modules: [
    "@nuxt/ui",
    "@nuxt/fonts",
    "@pinia/nuxt",
    "pinia-plugin-persistedstate/nuxt",
    "@vueuse/nuxt",
  ],
  css: ["~/assets/css/main.css"],
  pinia: {
    storesDirs: ["./stores/**"],
  },
  alias: {
    "@": __dirname,
  },
  runtimeConfig: {
    apiSecret: process.env.NUXT_API_KEY,
    apiUrl: process.env.NUXT_API_URL,
  },
  colorMode: {
    preference: "dark",
    fallback: "dark",
    hid: "nuxt-color-mode-script",
    globalName: "__NUXT_COLOR_MODE__",
    componentName: "ColorScheme",
    classSuffix: "",
    storageKey: "nuxt-color-mode",
  },
});
