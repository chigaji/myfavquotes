<template lang="">
<div>
 <v-app-bar
      app
      color="#2c3e50"
      darken-4
      dark
      prominent
      extended
      fade-img-on-scroll
      shrink-on-scroll
      src="../assets/dew.jpeg"
    >
    <!-- src="https://picsum.photos/1920/1080?random" -->
    <!-- src="../assets/cover2.jpeg" -->
    <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
    <!-- <v-spacer></v-spacer> -->
    <v-toolbar-title
      class="text-h2 font-weight-bold white--text text--darken-4">
  
        <router-link to="/" tag="span" style="cursor: pointer">{{appTitle}}</router-link>
      <v-btn
              icon
              color="white"
            >
              <v-icon>mdi-crown-outline</v-icon>
            </v-btn>

    </v-toolbar-title>
    <!-- <v-spacer></v-spacer> -->
  
    <!-- <Profile v-bind:user="user"/> -->
      <template v-slot:extension>
        <v-tabs
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
    <v-navigation-drawer v-model="drawer"
    app
    bottom
    temporary
    src="../assets/cloud.jpeg"
    class="white"

    >
      <v-list nav dense>
        <v-list-item
        v-for="item in items"
        :key="item.tab"
        :to="item.link">
        <v-list-item-icon>
            <v-icon v-text="item.icon"></v-icon>
        </v-list-item-icon>
        <v-list-item-content>
          <v-list-item-title v-text="item.tab"></v-list-item-title>
          <!-- {{item.tab}} -->
        </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    </div>
</template>
<script>
// import Profile from '@/components/Profile'
export default {
  data() {
    return {
      drawer: false,
      tab: null,
      appTitle: "MY FAVORITE QUOTES",
      items: [
        { tab: "Home", link: "/", icon: "mdi-home" },
        { tab: "Finance Quotes", link: "/finance", icon: "mdi-account-plus" },
        { tab: "Life Quotes", link: "/life", icon: "mdi-account-check" },
        // { tab: 'Create Asset User', link: '/CreateAssetUser' },
        { tab: "Love Quotes", link: "/love", icon: "mdi-bullseye-arrow" },
        { tab: "Random Quotes", link: "/random", icon: "mdi-history" },
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
