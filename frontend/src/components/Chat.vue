<template>
  <div class="chat-container">
    <div class="input-container">
      <q-input
        rounded
        filled
        v-model="text"
        placeholder="Type your message here..."
        color="primary"
      />
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const text = ref("");
</script>

<style>
.chat-container {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: flex-end; /* Aligns the input to the bottom */
}

.input-container {
  width: 100%; /* Makes the input container full width */
}
</style>
