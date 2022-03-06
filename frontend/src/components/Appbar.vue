<template lang="">
<div>
 <v-app-bar
      app
      color="#2c3e50"
      dark
      extended
      fade-img-on-scroll
      src="../assets/dew.jpeg"
      elevation="0"
    >
  <v-subheader>
    <router-link to="/" tag="span" style="cursor: pointer">
      <h1 class="blue--text text-h3 font-weight-bold">Power
        <span class="pink--text">.Quotes</span>
      </h1>
    </router-link>
    <v-btn
      icon
      color="white">
      <v-icon>mdi-crown-outline</v-icon>
    </v-btn>
  </v-subheader>
   
  <template v-slot:extension>
        <v-tabs right
          v-model="tab"
          align-with-title
          slider-color="white"
          center-active
          grow
          show-arrows
        >
          <v-tab
              v-for="item in items"
              :key="item.tab"
              @click="handleLink(item.link)"
            >
              {{ item.tab }}
            </v-tab>
        </v-tabs>
  </template>
  </v-app-bar>
  </div>
</template>
<script>
// import Profile from '@/components/Profile'
export default {
  data() {
    return {
      drawer: false,
      tab: null,
      appTitle: "FAVORITE QUOTES",
      items: [
        { tab: "Home", link: "/", icon: "mdi-home" },
        { tab: "Finance Quotes", link: "/finance", icon: "mdi-account-plus" },
        { tab: "Life Quotes", link: "/life", icon: "mdi-account-check" },
        // { tab: 'Create Asset User', link: '/CreateAssetUser' },
        { tab: "Love Quotes", link: "/love", icon: "mdi-bullseye-arrow" },
        // { tab: "Random Quotes", link: "/random", icon: "mdi-history" },
        // { tab: 'Update Asset', link: '/updateAsset' },
      ],
      user: {
        initials: "",
        fullName: "",
        email: "",
      },
    };
  },
  components: {
    // eslint-disable-next-line vue/no-unused-components
    // Profile
  },
  methods: {
    handleLink(link) {
      console.log(link);
      this.$router.push(link);
    },
  },
};
</script>
<style lang="">
</style>
