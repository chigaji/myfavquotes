<template>
  <v-container fluid>
    <!-- <v-sheet class="main" color="white" min-height="110vh" elevation="2"> -->
    <v-row no-gutters>
      <!-- <v-container fluid> -->
      <v-col fluid dense>
        <v-card class="rounded-xl" elevation="6">
          <v-img height="150px" src="../assets/money.jpg">
            <!-- </v-img> -->
            <v-card-title class="text-h4 font-weight-bold white--text mt-4">
              Finance <span class="red--text">.</span> Quotes
            </v-card-title>
          </v-img>

          <v-card-text class="d-flex justify-center text-h5">
            <form>
              <span v-if="errorMessage">
                <!-- <b class="red--text">{{ errorMessage.data }}</b> -->
                <v-alert type="error">😢 {{ errorMessage }}</v-alert>
              </span>
              <span class="font-weight-bold blue--text" v-if="randomQuote">
                <v-img
                  aspect-ratio="1"
                  width="200"
                  height="300"
                  src="../assets/robot.png"
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
              <!-- get random quotes here  -->
              <span class="font-weight-bold blue--text" v-if="quotes">
                <ol>
                  <li v-for="quote in quotes" :key="quote.id">
                    {{ quote.quote }}
                    <br />
                    <br />
                  </li>
                </ol>
              </span>
            </form>
          </v-card-text>
          <v-card-actions class="justify-center">
            <v-spacer></v-spacer>
            <div class="my-2">
              <v-btn
                x-large
                outlined
                rounded
                color="pink mt-12 px-10 py-5"
                dark
                @click="getAllQuotes"
              >
                All Quotes
              </v-btn>
            </div>
            <v-spacer></v-spacer>
            <div class="my-2">
             <Dialog class="ml-3" :rQuote="quoteToSend" @quoteCreated="handleCreated"/>
            </div>
            <v-spacer></v-spacer>
          </v-card-actions>
          
        </v-card>
      </v-col>
    </v-row>
    <!-- </v-sheet> -->
  </v-container>
</template>

<script>
import PostsService from "@/services/apiService";
import Dialog from "@/components/Dialog";
export default {
  name: "Home",
  data() {
    return {
      quotes: [],
      quoteToSend: {},
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
          this.quoteType = "Finance";
        }
        const resp = await PostsService.getAllQuotes(this.quoteType);
        this.quotes = resp.data;
        console.log(resp.data);

        this.randomQuote = "";
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
          this.quoteType = "Finance";
        }
        const resp = await PostsService.getRandomQuote(this.quoteType);
        this.randomQuote = resp.data;
        console.log(resp.data);
        this.quoteToSend = resp.data
        this.quoteType = "";
        this.quotes = "";
      } catch (error) {
        this.errorMessage = error;
        this.randomQuote = "";
        console.log(error);
      }
    },
    handleCreated(value){
      console.log("received child created:")
      console.log("value=>", value)
    }
  },
  components: {
    Dialog,
  },
  beforeMount() {
    this.getRandomQuote();
    this.errorMessage = "" 
  },
};
</script>
<style lang = "sass" scoped>
.main
  margin-top: 0px
  padding: 20px
</style>
