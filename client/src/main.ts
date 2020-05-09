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

axios.defaults.baseURL = "http://13.126.188.33/";

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
