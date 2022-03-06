import Vue from "vue";
import VueRouter from "vue-router";

import Home from "@/components/Home";
import Love from "@/components/Love";
Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    component: Home,
  },
  {
    path: "/love",
    name: "love",
    component: Love,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
