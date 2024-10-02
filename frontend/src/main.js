import { createApp } from "vue";
import { Quasar } from "quasar";
import "@quasar/extras/material-icons/material-icons.css";
import "quasar/src/css/index.sass";
import "./styles/main.css";
import router from "./router/index.js";
import App from "./App.vue";

const myApp = createApp(App);

myApp.use(Quasar, {
  plugins: {}, // Import Quasar plugins and add here
});

myApp.use(router);

myApp.mount("#app");
