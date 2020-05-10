import Vue from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
import store from "./store";
import VueAxios from "vue-axios";
import axios from "axios";

import "bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";

Vue.config.productionTip = false;

Vue.use(VueAxios, axios);

console.log(process.env)

axios.defaults.baseURL = process.env.API_URL || "http://localhost:5000";

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
