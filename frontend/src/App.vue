<template>
  <v-app>
    <Appbar />
    <!-- <v-main class="main-content"> -->
    <AppFab />
    <!-- page wrapper -->
    <v-sheet
      id="scrolling-techniques-3"
      class="overflow-y-auto"
      max-height="600"
    >
      <v-container style="height: 1000px"></v-container>
    </v-sheet>
    <v-fade-transition mode="out-in">
      <div class="page-wrapper"><router-view /></div>
    </v-fade-transition>
    <!-- footer here -->
    <Footer />
    <!-- <h1>this is working</h1> -->
    <!-- </v-main> -->
  </v-app>
</template>

<script>
import AppFab from "@/components/AppFab";
import Appbar from "@/components/Appbar";
// import Profile from '@/components/Profile'
import Footer from "@/components/Footer";

export default {
  name: "App",

  components: {
    // Hello
    AppFab,
    // Profile,
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
};
</script>
<style lang = "sass" scoped>
/* .main-content
 margin: 10px; */
.page-wrapper
  min-height: calc(100vh -112px -48px)
.main-content
  margin-top: 100px
</style>
