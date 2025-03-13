import { createApp } from "vue";
import App from "./App.vue";
import router from "./router.ts";
import vuetify from "./plugins/vuetify";
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faProjectDiagram,
  faSignOutAlt,
  faUser,
} from "@fortawesome/free-solid-svg-icons";
import { createPinia } from "pinia";
import "./assets/styles/global.css";
import "@mdi/font/css/materialdesignicons.css";

library.add(faProjectDiagram, faSignOutAlt, faUser);

createApp(App)
  .component("font-awesome-icon", FontAwesomeIcon)
  .use(createPinia())
  .use(router)
  .use(vuetify)
  .mount("#app");
