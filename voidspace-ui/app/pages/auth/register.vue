<script setup lang="ts">
definePageMeta({
  layout: "auth",
  middleware: "guest"
});
</script>

<template>
  <AuthForm mode="register" />
</template>
