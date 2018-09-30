<template>
  <v-app >   
    <SearchBar defaultQuery='okta' v-on:search-submitted="githubQuery" />
    <v-container grid-list-md fluid class="grey lighten-4" >
         <v-tabs
        slot="extension"
        v-model="tabs"
        centered
        color="teal"
        text-color="white"
        slider-color="white"
      >
        <v-tab class="white--text" :key="1">
          SEARCH
        </v-tab>
        <v-tab class="white--text" :key="2">
          KUDOS
        </v-tab>
      </v-tabs>
        <v-tabs-items style="width:100%" v-model="tabs">
          <v-tab-item :key="1">
            <v-layout row wrap>
              <v-flex v-for="repo in repos" :key="repo.id" md4>
                <GitHubRepo :repo="repo" />
              </v-flex>
            </v-layout>
          </v-tab-item>
          <v-tab-item :key="2">
            <v-layout row wrap>
              <v-flex v-for="kudo in kudos" :key="kudo.id" md4 >
                <GitHubRepo :repo="kudo" />
              </v-flex>
            </v-layout>
          </v-tab-item>
        </v-tabs-items>
    </v-container>
  </v-app>
</template>

<script>
import SearchBar from './components/SearchBar.vue'
import GitHubRepo from './components/GithubRepo.vue'
import githubClient from './githubClient'
import { mapMutations, mapGetters } from 'vuex'

export default {
  name: 'App',
  components: { SearchBar, GitHubRepo },
  data() {
    return {
      tabs: null
    }
  },
  computed: mapGetters(['kudos', 'repos']),
  methods: {
    githubQuery(query) {
      githubClient
        .getJSONRepos(query)
        .then(response => this.addRepos(response.items) )
    },
    ...mapMutations(['addRepos'])
  },
}
</script>

<style>
 .v-tabs__content {
   padding-bottom: 2px;
 }
</style>

