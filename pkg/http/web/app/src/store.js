import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    kudos: {},
    repos: [],
  },
  mutations: {
    toggleKudo(state, repo) {
      if (!state.kudos[repo.id]) {
        let kudo = {}
        kudo[repo.id] = repo;
        return state.kudos = Object.assign({}, kudo, state.kudos);
      }

      let kudos = {};
      Object.entries(state.kudos).forEach(([repoId, kudo]) => {
        if (repoId !== repo.id) {
          kudos[repoId] = kudo;
        }
      });

      return state.kudos = kudos;
    },
    addRepos (state, repos) {
      state.repos = [...state.repos, ...repos];
    }
  },
  getters: {
    kudos(state) {
      return Object.values(state.kudos);
    },
    repos(state) {
      return state.repos;
    },
    isKudo(state) {
      return (repo)=> {
        return !!state.kudos[repo.id];
      };
    }
  },
  actions: {
    toggleKudo(context, repo) {
      setTimeout(() => {
        context.commit('toggleKudo', repo);
      }, 100)
    }
  }
});

export default store;