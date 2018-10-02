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
        return state.kudos = { [repo.id]: repo, ...state.kudos };
      }

      return state.kudos =  Object.entries(state.kudos).reduce((kudos, [repoId, kudo]) => {
                              (repoId == repo.id) ? kudos
                                                  : { [repoId]: kudo, ...kudos };
                            }, {});
    },
    resetRepos (state, repos) {
      state.repos = repos;
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