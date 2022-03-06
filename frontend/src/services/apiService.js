import API from "./api";

const baseURL = "api/v1";
export default {
  homePage() {
    return API().get(`${baseURL}/home`);
  },
  // get a random quote of a given type
  getRandomQuote(type) {
    return API().get(`${baseURL}/quote/${type}`);
  },

  // get all quotes of a given type
  // api.GET("/quotes/:type")
  getAllQuotes(type) {
    return API().get(`${baseURL}/quotes/${type}`);
  },

  // get a quote of given type by id
  // api.GET("/quote/:type/:id")
  getQuote(type, id) {
    return API().get(`${baseURL}/quote/${type}/${id}`);
  },

  // create a quote of a given type
  // api.POST("/quotes/:type", createQuote)
  createQuote(type, quote) {
    return API().post(`${baseURL}/quotes/${type}`, {
      quote: quote,
    });
  },
  // edit a given quote of a type using a given id
  // api.PUT("/quotes/:type/:id", editQuote)
  editQuote(type, id, quote) {
    return API().put(`${baseURL}/quotes/${type}/${id}`, {
      quote: quote,
    });
  },

  // Delete a quote of a given type by Id
  // api.DELETE("/quote/:type/:id", deleteQuote)
  deleteQuote(type, id) {
    return API().delete(`${baseURL}/quote/${type}/${id}`);
  },
};
