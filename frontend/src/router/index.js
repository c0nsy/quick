// src/router.js

import { createRouter, createWebHistory } from "vue-router";
import LandingPage from "../views/LandingPage.vue";
import Workspace from "../views/Workspace.vue";
import store from "../store/store"; // Import the store

const routes = [
  {
    path: "/",
    name: "Home",
    component: LandingPage,
  },
  {
    path: "/workspace/:sessionId",
    name: "Workspace",
    component: Workspace,
    beforeEnter: (to, from, next) => {
      // Check if the session ID exists in local storage
      const sessionIdExists = store.sessionIds.includes(to.params.sessionId);
      if (sessionIdExists) {
        next(); // Proceed to the Workspace component
      } else {
        next({ name: "Home" }); // Redirect to the home page or any other route
      }
    },
  },
];

const router = createRouter({
  history: createWebHistory(), // Use HTML5 history mode
  routes,
});

export default router;
