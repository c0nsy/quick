// src/store/store.js

const STORAGE_KEY = "sessionIds"; // Key to store session IDs in local storage

const store = {
  get sessionIds() {
    // Retrieve session IDs from local storage
    const ids = localStorage.getItem(STORAGE_KEY);
    return ids ? JSON.parse(ids) : []; // Parse JSON or return empty array
  },
  set sessionIds(ids) {
    // Store session IDs in local storage
    localStorage.setItem(STORAGE_KEY, JSON.stringify(ids));
  },
  addSessionId(sessionId) {
    const ids = this.sessionIds; // Get current session IDs
    if (!ids.includes(sessionId)) {
      ids.push(sessionId); // Add new session ID if it doesn't exist
      this.sessionIds = ids; // Update local storage
    }
  },
};

export default store;
