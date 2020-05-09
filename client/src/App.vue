<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Search</router-link>|
      <router-link to="/upload">Upload</router-link>
    </div>
    <router-view />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { AxiosResponse } from "axios";
export default Vue.extend({
  mounted() {
    this.axios.get("/").then((res: AxiosResponse) => {
      console.log("Successfully connected backend", res);
    });
  }
});
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
