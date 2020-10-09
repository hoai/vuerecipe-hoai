require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap-sass/assets/javascripts/bootstrap.js");

import Vue from "vue";
import VueRouter from "router";
Vue.use(VueRouter);

import BandComponent from "./components/band.vue";
import MembersComponent from "./components/members.vue";

import CategroyComponent from "./components/category.vue";
import PostComponent from "./components/post.vue";

const routes = [
  {path: "/band/:id", component: MembersComponent, name: "showBand"},
  {path: "/category/:id", component: PostComponent, name: "showCategory"},
  {path: "/bands", component: BandComponent},
  {path: "/categories", component: CategroyComponent},
  {path: "/", component: CategroyComponent}
];

const router = new VueRouter({
  mode: "history",
  routes
});

const app = new Vue({
  router
}).$mount("#app");
