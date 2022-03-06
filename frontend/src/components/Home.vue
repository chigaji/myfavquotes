<template>
  <v-container fluid>
    <!-- <v-sheet class="main" color="white" min-height="110vh" elevation="2"> -->
    <v-row dense>
      <!-- <v-container fluid> -->
      <v-col>
        <v-card color="#F5F5F5" elevation="1">
          <v-card-text class="text-h6">
            <h2 class="text-h4 black--text">
              My <span class="green--text">Favorite</span> Quotes &
              <span class="pink--text">Sayings</span> Collections
            </h2>
            <br />
            Here are My Top Best
            <span class="font-weight-bold blue--text">Finance</span>,
            <span class="font-weight-bold blue--text">Life</span> and
            <span class="font-weight-bold blue--text">Love</span> quotes.
            <div class="ml-0 mb-2 mt-8">You Can:-</div>
            <br />
            <v-timeline align-top dense>
              <v-timeline-item
                v-for="message in messages"
                :key="message.message"
                :color="message.color"
                small
              >
                <div>
                  <div>{{ message.message }}</div>
                </div>
              </v-timeline-item>
            </v-timeline>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col>
        <v-img
          src="https://picsum.photos/1920/1080?random"
          aspect-ratio="1"
          gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
          height="438px"
        >
        </v-img>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-card color="white" elevation="9">
          <v-img
            src="https://picsum.photos/1920/1080?random"
            aspect-ratio="1"
            gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
            height="200px"
          >
            <!-- </v-img> -->
            <v-card-title
              class="
                justify-center
                headline
                display-1
                font-weight-bold
                black--text
                text--darken-4
              "
              >LOVE QUOTES
            </v-card-title>
            <v-card-text class="text-center text-h6">
              Random love quotes Collection
            </v-card-text>
            <v-card-actions class="justify-center">
              <v-btn large color="primary mt-8 px-10 py-5" dark
                ><router-link to="/love" tag="span" style="cursor: pointer"
                  >View Quotes</router-link
                >
              </v-btn>
            </v-card-actions>
          </v-img>
        </v-card>
      </v-col>
      <v-col>
        <v-card color="white" elevation="9">
          <v-img
            src="https://picsum.photos/1920/1080?random"
            aspect-ratio="1"
            gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
            height="200px"
          >
            <!-- </v-img> -->
            <v-card-title
              class="
                justify-center
                headline
                display-1
                font-weight-bold
                black--text
                text--darken-4
              "
              >LIFE QUOTES
            </v-card-title>
            <v-card-text class="text-center text-h6">
              Random life quotes Collection
            </v-card-text>
            <v-card-actions class="justify-center">
              <v-btn color="primary" dark>View Quotes </v-btn>
            </v-card-actions>
          </v-img>
        </v-card>
      </v-col>
      <v-col>
        <v-card color="white" elevation="9">
          <v-img
            src="https://picsum.photos/1920/1080?random"
            aspect-ratio="1"
            gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
            height="200px"
          >
            <!-- </v-img> -->
            <v-card-title
              class="
                justify-center
                headline
                display-1
                font-weight-bold
                black--text
                text--darken-4
              "
              >POWER FINANCIAL QUOTES
            </v-card-title>
            <v-card-text class="text-center text-h6">
              Awesome financial quotes Collection
            </v-card-text>
            <v-card-actions class="justify-center">
              <v-btn color="primary" dark>View Quotes </v-btn>
            </v-card-actions>
          </v-img>
        </v-card>
      </v-col>

      <!-- </v-container> -->
    </v-row>
    <!-- </v-sheet> -->
  </v-container>
</template>

<script>
export default {
  name: "Home",
  data() {
    return {
      messages: [
        {
          message: `Create Quotes`,
          color: "deep-purple lighten-1",
        },
        {
          message: "View Quotes",
          color: "green",
        },
        {
          message: "Edit Quotes",
          color: "deep-purple lighten-1",
        },
      ],
    };
  },
};
</script>
<style lang = "sass" scoped>
.main
  margin-top: 0px
  padding: 20px
</style>
