<template>
  <v-app>
    <Appbar />
    <v-main class="main-content grey lighten-5">
      <!-- page wrapper -->
      <v-fade-transition mode="out-in">
        <div class="page-wrapper"><router-view /></div>
      </v-fade-transition>
      <!-- <Home /> -->
      <!-- footer here -->
    </v-main>
    <AppFab />
    <Footer />
  </v-app>
</template>

<script>
import AppFab from "@/components/AppFab";
import Appbar from "@/components/Appbar";
// import Home from "@/components/Home";
// import Profile from '@/components/Profile'
import Footer from "@/components/Footer";

export default {
  name: "App",

  components: {
    AppFab,
    // Profile,
    // Home,
    Appbar,
    Footer,
  },

  data: () => ({
    //
    drawer: false,
    group: null,
    user: {
      initials: "",
      fullName: "",
      email: "",
    },
  }),

  computed: {
    // this.user.initials = this.$store.getters.getInitial
    // console.log('initials',this.$store.state.user.email)
    // this.user.email = this.$store.state.user.email
    // console.log('user email', this.user.email)
    // this.user.fullName = this.$store.getters.fullName
    // console.log('full name', this.user.fullName)
  },
};
</script>
<style lang = "sass" scoped>
/* .main-content
 margin: 10px; */
.page-wrapper
  min-height: calc(100vh -112px -48px)
.main-content
  margin-top: 50px
  color: #FBF5F5
</style>
