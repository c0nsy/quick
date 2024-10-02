<template>
  <div class="container">
    <LandingPage />
  </div>
</template>

<script setup>
import LandingPage from "./views/LandingPage.vue";
</script>

<style scoped>
.container {
  display: flex; /* Use flexbox for layout */
  height: 100vh; /* Make the container full height */
  justify-content: center; /* Center horizontally */
  align-items: center; /* Center vertically */
}
</style>
