<template lang="">
<!-- src="../assets/appbar3.jpg" -->
 <v-app-bar
    absolute
    color="#6A76AB"
    dark
    shrink-on-scroll
    prominent
    extended
    src="../assets/cover2.jpeg"
    fade-img-on-scroll
    scroll-target="#scrolling-techniques-3"
    scroll-threshold="200"
    >
    <!-- src="https://picsum.photos/1920/1080?random" -->
     <!-- src="../assets/cover.png" -->
    <!-- <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon> -->
     <template v-slot:img="{ props }">
        <v-img
          v-bind="props"
          gradient="to top right, rgba(100,115,201,.7), rgba(25,32,72,.7)"
        ></v-img>
      </template>
      <v-spacer></v-spacer>
    <v-toolbar-title
      class="text-center display-1 font-weight-bold white--text text--lighten-4">
      <router-link to="/" tag="span" style="cursor: pointer">{{appTitle}}</router-link>
      <!-- <v-icon color="yellow lighten-4">mdi-fingerprint</v-icon> -->
    </v-toolbar-title>
   <v-spacer></v-spacer>
    </v-app-bar>
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
        { tab: "Sign Up", link: "/register", icon: "mdi-account-plus" },
        { tab: "Login", link: "/login", icon: "mdi-account-check" },
        // { tab: 'Create Asset User', link: '/CreateAssetUser' },
        { tab: "Create An Asset", link: "/CreateAsset", icon: "mdi-buffer" },
        {
          tab: "Query My Assets",
          link: "/getAssetsByUser",
          icon: "mdi-clipboard-text-search",
        },
        // { tab: 'Update Asset', link: '/updateAsset' },
        { tab: "Asset Details", link: "/readAsset", icon: "mdi-library" },
        {
          tab: "ChangeAssetOwnership",
          link: "/changeAssetOwnership",
          icon: "mdi-share-all",
        },
        {
          tab: "Asset History",
          link: "/getHistoryForAsset",
          icon: "mdi-history",
        },
        {
          tab: "Download An Asset",
          link: "/downloadAsset",
          icon: "mdi-cloud-download",
        },
        {
          tab: "Verify Timestamp",
          link: "/verifyAsset",
          icon: "mdi-fingerprint",
        },
        // { tab: 'Delete An Asset', link: '/deleteAsset' }
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
