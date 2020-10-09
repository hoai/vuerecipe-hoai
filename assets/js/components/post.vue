<template>
<div>
  <h1 class="page-header">{{category.title}}</h1>

  <blockquote>
    {{category.description}}
  </blockquote>

  <ul class="list-unstyled">
    <li v-for="post in posts">
      <h2>
        {{post.title}} - {{post.description}}
      </h2>
    </li>
  </ul>

</div>
</template>

<script charset="utf-8">
export default {
  data() {
    return {
      category: {},
      posts: {}
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
      let id = this.$route.params.id;

      let req = $.getJSON(`/api/categories/${id}`);
      req.done(data => {
        this.category = data;
      });

      req = $.getJSON(`/api/categories/${id}/posts`);
      req.done(data => {
        this.posts = data;
      });
    }
  }
};
</script>
