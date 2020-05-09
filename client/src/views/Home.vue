<template>
  <div class="home">
    <h1>Search</h1>
    <div class="app-search">
      <form @submit.prevent="search">
        <div class="form-group">
          <label for="form-name">Search</label>
          <input
            type="text"
            class="form-control"
            id="form-search-query"
            v-model="query"
            placeholder="Search your meme"
          />
        </div>
        <button type="submit" class="btn btn-primary">Submit</button>
      </form>
      <template v-if="showResults">
        <p>
        Found {{ hits }} results in {{ responseTime }}ms
        </p>
        <ul style="list-style-type: none; margin: 0; padding: 0; margin-top: 5rem;">
          <li v-for="result in results" v-bind:key="result.id">
            <div class="card" style="width: 18rem;">
              <img class="card-img-top" :src="result.result.file_url" alt="Card image cap" />
              <div class="card-body">
                <h5 class="card-title">{{ result.result.name }}</h5>
                <p
                  class="card-text"
                >{{ result.result.description }}</p>
              </div>
              <ul class="list-group list-group-flush">
                <li class="list-group-item">Tags: {{ result.result.tags }}</li>
                <li class="list-group-item">Origin: {{ result.result.origin }}</li>
                <li class="list-group-item">Image Text: {{ result.result.text }}</li>
                <li class="list-group-item">ElasticSearch Confidence: {{ result.score }}</li>
                <li class="list-group-item">ElasticSearch ID: {{ result.id }}</li>
              </ul>
            </div>
          </li>
        </ul>
      </template>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { AxiosResponse, AxiosError } from "axios";
export default Vue.extend({
  name: "Search",
  data() {
    return {
      query: null,
      hits: null,
      results: null,
      responseTime: null,
      showResults: false
    };
  },
  methods: {
    search() {
      this.showResults = false;

      const searchQuery = this.query;

      this.axios
        .get("/api/search", { params: { q: searchQuery } })
        .then((res: AxiosResponse) => {
          this.showResults = true;
          this.hits = res.data.hits;
          this.responseTime = res.data.time;
          this.results = res.data.results;
        })
        .catch((err: AxiosError) => {
          console.error(err);
        });
    }
  }
});
</script>

<style lang="scss" scoped>
.app-search {
  width: 50%;
  margin: auto;
  margin-top: 2rem;
  text-align: left;
}
</style>
