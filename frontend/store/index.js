import Vuex from "vuex";
import axios from "axios";

const createStore = () => {
  return new Vuex.Store({
    state: {
      text: "",
    },
    mutations: {
      setText(state, inputText) {
        state.text = inputText;
      },
    },
    actions: {
      nuxtServerInit(vuexContext, context) {
        return axios
          .get("http://api:8080/")
          .then((response) => {
            vuexContext.commit("setText", response.data);
          })
          .catch((e) => context.error(e));
      },
      setText(vuexContext, inputText) {
        vuexContext.commit("setText", inputText);
      },
    },
    getters: {
      text(state) {
        return state.text;
      },
    },
  });
};

export default createStore;
