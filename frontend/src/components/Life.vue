<template>
  <v-container fluid>
    <!-- <v-sheet class="main" color="white" min-height="110vh" elevation="2"> -->
    <v-row no-gutters>
      <!-- <v-container fluid> -->
      <v-col col-6 dense>
        <v-card class="text-center rounded-xl" elevation="6">
          <v-img height="150px" src="../assets/stars_bg.jpeg">
            <!-- </v-img> -->
            <v-card-title class="text-h4 font-weight-bold white--text mt-4">
              Life <span class="red--text">.</span> Quotes
            </v-card-title>
          </v-img>

          <v-card-text class="d-flex justify-center text-h5" >
            <form>
              <span v-if="errorMessage">
                <!-- <b class="red--text">{{ errorMessage.data }}</b> -->
                <v-alert type="error">😢 {{ errorMessage }}</v-alert>
              </span>
              <span class="font-weight-bold blue--text" v-if="randomQuote">
                <v-img
                  max-height="250"
                  max-width="250"
                  src="../assets/love_bg.jpg"
                >
                </v-img>
                <v-icon large color="pink"
                  >mdi-format-quote-open-outline</v-icon
                >
                {{ randomQuote.quote }}
                <v-icon large color="pink"
                  >mdi-format-quote-close-outline</v-icon
                >
              </span>
            </form>
          </v-card-text>
          <!-- </v-img> -->
        </v-card>
      </v-col>
      <!-- </v-container> -->
    </v-row>
    <!-- </v-sheet> -->
  </v-container>
</template>

<script>
import PostsService from "@/services/apiService";
export default {
  name: "Home",
  data() {
    return {
      quotes: [],
      randomQuote: "",
      errorMessage: "",
      quoteType: "",
    };
  },
  methods: {
    async getAllQuotes() {
      try {
        this.errorMessage = "";
        if (this.quoteType === "") {
          this.quoteType = "Life";
        }
        const resp = await PostsService.getAllQuotes(this.quoteType);
        this.quotes = resp.data;
        console.log(resp.data);
        this.quoteType = "";
      } catch (error) {
        this.errorMessage = error;
        this.quotes = "";
        console.log(error);
      }
    },
    async getRandomQuote() {
      try {
        this.errorMessage = "";
        if (this.quoteType === "") {
          this.quoteType = "Life";
        }
        const resp = await PostsService.getRandomQuote(this.quoteType);
        this.randomQuote = resp.data;
        console.log(resp.data);
        this.quoteType = "";
      } catch (error) {
        this.errorMessage = error;
        this.randomQuote = "";
        console.log(error);
      }
    },
  },
  beforeMount() {
    this.getRandomQuote();
  },
};
</script>
<style lang = "sass" scoped>
.main
  margin-top: 0px
  padding: 20px
</style>
