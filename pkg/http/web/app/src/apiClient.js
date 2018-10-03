import Vue from 'vue';
import axios from 'axios';

const client = axios.create({
  baseURL: process.env.NODE_ENV == 'production' ? '' : 'http://localhost:4444',
  json: true
});

const APIClient =  {
  toggleKudo(repo) {
    return this.perform('patch', `/kudos/${repo.id}`, repo);
  },
  
  createKudo(repo) {
    return this.perform('post', '/kudos', repo);
  },

  deleteKudo(repo) {
    return this.perform('post', '/kudos', repo);
  },

  updateKudo(repo) {
    return this.perform('put', `/kudos/${repo.id}`, repo);
  },

  getKudos() {
    return this.perform('get', '/kudos');
  },

  getKudo(repo) {
    return this.perform('get', `/kudo/${repo.id}`);
  },

  async perform (method, resource, data) {
    let accessToken = await Vue.prototype.$auth.getAccessToken()
    return client({
      method,
      url: resource,
      data,
      headers: {
        Authorization: `Bearer ${accessToken}`
      }
    }).then(req => {
      return req.data
    })
  }
}

export default APIClient;