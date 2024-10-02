<template>
  <div>
    <q-btn size="xl" color="primary" @click="createSession">New Session</q-btn>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import { v4 as uuidv4 } from "uuid"; // Import UUID for unique session ID
import store from "../store/store";
const router = useRouter();

const createSession = () => {
  const sessionId = uuidv4(); // Generate a unique session ID
  store.addSessionId(sessionId); // Use the addSessionId method to store the session ID
  router.push(`/workspace/${sessionId}`); // Navigate to the new route
};
</script>

<style scoped></style>
