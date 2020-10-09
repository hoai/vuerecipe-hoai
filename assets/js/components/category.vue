<template>
<div>
  <h1 class="page-header">Categories</h1>

  <ul class="list-unstyled">
    <li v-for="category in categories">
      <router-link :to='{name: "showCategory", params: {id: category.id}}'>
        <h2>
          {{ category.title }}
        </h2>
      </router-link>
    </li>
  </ul>
</div>
</template>

<script charset="utf-8">
export default {
  data() {
    return {
      categories: []
    };
  },

  created() {
    this.fetchData();
  },

  watch: {
    $route: "fetchData"
  },

  methods: {
    fetchData: function() {
      let req = $.getJSON("/api/categories");
      req.done(data => {
        this.categories = data;
      });
    }
  }
};
</script>
