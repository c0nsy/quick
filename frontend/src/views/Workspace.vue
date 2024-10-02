<template>
  <div class="workspace-container">
    <div class="workspace-container">
      <div class="chat-container">
        <Chat />
      </div>
      <div class="whiteboard-container">
        <Whiteboard />
      </div>
    </div>
  </div>
</template>

<script setup>
import Chat from "../components/Chat.vue";
import Whiteboard from "../components/Whiteboard.vue";
import { ref } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();
const sessionId = ref(route.params.sessionId); // Get the session ID from the route parameters
</script>

<style scoped>
.workspace-container {
  display: flex; /* Use flexbox for layout */
  width: 100%;
  height: 100vh; /* Full height of the viewport */
}

.chat-container {
  flex: 0 0 25%; /* Take up 25% of the width */
  border-right: 1px solid #ccc; /* Optional: add a border to separate the components */
}

.whiteboard-container {
  flex: 1; /* Take up the remaining space */
  display: flex;
  justify-content: center; /* Center content horizontally in the whiteboard */
  align-items: center; /* Center content vertically in the whiteboard */
}
</style>
