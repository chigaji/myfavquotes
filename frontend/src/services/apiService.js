import API from "./api";

export default {
    homePage(){
       return API().get('/home')
    }
}