const axios = require("axios");

const quote = "love";
const res = axios
  .get(`http://127.0.0.1:4000/api/v1/quote/${quote}`)
  .then((data) => console.log(data.data));
