import Vue from "vue";
import VueRouter from "vue-router";

import Home from "@/components/Home";
import Love from "@/components/Love";
import Life from "@/components/Life";
import Finance from "@/components/Finance";
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
  {
    path: "/life",
    name: "life",
    component: Life,
  },
  {
    path: "/finance",
    name: "finance",
    component: Finance,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
