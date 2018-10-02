import Vue from 'vue';
import VueRouter from 'vue-router';

import Home from './components/Home';
import GitHubRepoDetails from './components/GithubRepoDetails';

Vue.use(VueRouter);

export default new VueRouter(
    { routes: [
        { path: '/', component: Home },
        { name: 'repo-details', path: '/repo/:id', component: GitHubRepoDetails }
    ]}
);