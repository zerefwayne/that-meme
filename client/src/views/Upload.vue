<template>
  <div class="about">
    <h1>Upload Page</h1>
    <div class="app-upload">
      <form
        enctype="multipart/form-data"
        novalidate
        autocomplete="off"
        @submit.prevent="submitForm"
      >
        <div class="form-group">
          <label for="form-name">Name</label>
          <input
            type="text"
            class="form-control"
            id="form-name"
            v-model="uploadForm.name"
            placeholder="Enter meme name"
          />
        </div>
        <div class="form-group">
          <label for="form-tags">Tags</label>
          <input
            type="text"
            class="form-control"
            id="form-tags"
            v-model="uploadForm.tags"
            placeholder="Enter relevant tags space seperated"
          />
        </div>
        <div class="form-group">
          <label for="form-description">Description</label>
          <input
            type="text"
            class="form-control"
            id="form-description"
            v-model="uploadForm.description"
            placeholder="Enter description"
          />
        </div>
        <div class="form-group">
          <label for="form-text">Text</label>
          <input
            type="text"
            class="form-control"
            id="form-text"
            v-model="uploadForm.text"
            placeholder="Enter image text"
          />
        </div>
        <div class="form-group">
          <label for="form-origin">Origin</label>
          <input
            type="text"
            class="form-control"
            id="form-origin"
            v-model="uploadForm.origin"
            placeholder="Enter meme origin"
          />
        </div>
        <div class="form-group">
          <label for="exampleFormControlFile1">Upload the meme</label>
          <input
            type="file"
            ref="file"
            class="form-control-file"
            name="newMeme"
            accept="image/*"
            id="exampleFormControlFile1"
            @change="handleFileUpload"
          />
        </div>
        <button type="submit" class="btn btn-primary">Submit</button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { AxiosResponse, AxiosError } from "axios";
export default Vue.extend({
  data() {
    return {
      uploadForm: {
        name: null,
        tags: null,
        description: null,
        text: null,
        origin: null,
        file: null
      }
    };
  },
  methods: {
    submitForm() {
      console.log("Form submitted");

      const formData = new FormData();

      formData.append("name", this.uploadForm.name || "");
      formData.append("tags", this.uploadForm.tags || "");
      formData.append("description", this.uploadForm.description || "");
      formData.append("text", this.uploadForm.text || "");
      formData.append("origin", this.uploadForm.origin || "");
      formData.append("newMeme", this.uploadForm.file || "");

      console.log(this.uploadForm);
      console.log(formData);

      this.axios
        .post("/api/upload", formData, {
          headers: {
            "Content-Type": "multipart/form-data"
          }
        })
        .then((res: AxiosResponse) => {
          console.log(res);
        })
        .catch((err: AxiosError) => {
          console.log(err);
        });
    },
    handleFileUpload() {
      // eslint-disable-next-line
      this.uploadForm.file = (this.$refs.file as any).files[0];
    }
  }
});
</script>


<style lang="scss" scoped>
.app-upload {
  width: 50%;
  margin: 0 auto;
  text-align: left;
}
</style>