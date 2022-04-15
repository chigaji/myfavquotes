<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" persistent max-width="800px">
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          x-large
          outlined
          rounded
          color="pink mt-12 px-10 py-5"
          dark
          v-bind="attrs"
          v-on="on"
        >
          Add Quote
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="text-h5">Create New Quote</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-text-field
                  label="Enter Your Quote"
                  required
                  v-model="quote"
                ></v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="dialog = false">
            Close
          </v-btn>
          <v-btn color="blue darken-1" text @click="save"> Add Quote </v-btn>
        </v-card-actions>
        <span v-if="finalM">
          <!-- <b class="red--text">{{ errorMessage.data }}</b> -->
          <v-alert type="info">😢 {{ finalM }}</v-alert>
        </span>
        <span v-if="errorMessage">
          <!-- <b class="red--text">{{ errorMessage.data }}</b> -->
          <v-alert type="error">😢 {{ errorMessage }}</v-alert>
        </span>
        <span>
          <!-- <b class="red--text">{{ errorMessage.data }}</b> -->
          <v-alert type="info">😢The recieved quote => {{ rQuote }}</v-alert>
        </span>
      </v-card>
    </v-dialog>
  </v-row>
</template>
<script>
export default {
  data: () => ({
    dialog: false,
    errorMessage: "",
    quote: "",
    finalM: ""
  }),
  props: {
    rQuote: Object
  },
  methods: {
    close() {
      this.dialog = false;
      setTimeout(() => {}, 300);
    },
    save() {
      if (this.quote === "") {
        this.errorMessage = "Field cannot be empty";
        return;
      } else {
        // const finalM = this.quote
        this.finalM = this.quote
        console.log(this.quote);

        this.$emit('quoteCreated', this.quote)
        console.log('quoteCreated');
      }
      this.errorMessage = ""
      this.close;
    },
  },
  mounted () {
    console.log('child created');
    this.$emit('child created');
  },
  created () {
    
  },
 
 
};
</script>