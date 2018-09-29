<template>
  <v-app >
    <v-toolbar dark color="teal">
      <v-spacer></v-spacer>
      <v-text-field
        solo-inverted
        flat
        hide-details
        label="Search for your OOS project on Github + Press Enter"
        prepend-inner-icon="search"
        v-model="query"
        @keyup.enter="search"
      ></v-text-field>
      <v-spacer></v-spacer>
       <v-tabs
        slot="extension"
        v-model="tabs"
        centered
        color="teal"
        slider-color="white"
      >
        <v-tab :key="1">
          SEARCH
        </v-tab>
        <v-tab :key="2">
          KUDOS
        </v-tab>
      </v-tabs>
    </v-toolbar>
      
    <v-container grid-list-md fluid class="grey lighten-4" >
        <v-tabs-items style="width:100%" v-model="tabs">
          <v-tab-item :key="1">
            <v-layout row wrap>
              <v-flex v-for="repo in repos" :key="repo.id" md4>
                <v-card >
                  <v-card-title primary-title>
                    <div class="repo-card-content">
                      <h3 class="headline mb-0">{{repo.full_name}}</h3>
                      <div>{{repo.description}}</div>
                    </div>
                  </v-card-title>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-dialog v-if="isKudo(repo)">
                      <v-btn slot="activator" icon flat color="orange lighten-2"><v-icon>notes</v-icon></v-btn>
                      <v-card>
                        <v-card-title class="headline white--text teal lighten-2" primary-title>Notes</v-card-title>
                        <v-card-text>
                          <v-textarea></v-textarea>
                        </v-card-text>
                        <v-divider></v-divider>
                        <v-card-actions>
                          <v-spacer></v-spacer>
                          <v-btn color="orange" flat @click="dialog = false">
                            Save
                          </v-btn>
                        </v-card-actions>
                      </v-card>
                    </v-dialog>
                    <v-btn @click.prevent="toggleKudo(repo)"  flat icon color="pink">
                      <v-icon v-if="isKudo(repo)">favorite</v-icon>
                      <v-icon v-else>favorite_border</v-icon>
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-flex>
            </v-layout>
          </v-tab-item>
          <v-tab-item :key="2">
            <v-layout row wrap>
              <v-flex v-for="repo in kudos" md4 >
                <v-card >
                  <v-card-title primary-title>
                    <div class="repo-card-content">
                      <h3 class="headline mb-0">{{repo.full_name}}</h3>
                      <div>{{repo.description}}</div>
                    </div>
                  </v-card-title>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-dialog v-if="isKudo(repo)">
                      <v-btn slot="activator" icon flat color="orange lighten-2"><v-icon>notes</v-icon></v-btn>
                      <v-card>
                        <v-card-title class="headline white--text teal lighten-2" primary-title>Notes</v-card-title>
                        <v-card-text>
                          <v-textarea></v-textarea>
                        </v-card-text>
                        <v-divider></v-divider>
                        <v-card-actions>
                          <v-spacer></v-spacer>
                          <v-btn color="orange" flat @click="dialog = false">
                            Save
                          </v-btn>
                        </v-card-actions>
                      </v-card>
                    </v-dialog>
                    <v-btn @click.prevent="toggleKudo(repo)"  flat icon color="pink">
                      <v-icon v-if="isKudo(repo)">favorite</v-icon>
                      <v-icon v-else>favorite_border</v-icon>
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-flex>
            </v-layout>
          </v-tab-item>
        </v-tabs-items>
    </v-container>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      tabs: null,
      query: 'okta',
      repos: [],
      kudos: [],
    }
  },
  mounted() {
    this.githubQuery(this.query);
  },
  methods: {
    isKudo(repo){
      return this.kudos.find(r => r.id === repo.id) 
    },
    toggleKudo(repo){
      if (this.isKudo(repo)) {
        return this.dislike(repo)
      }
      this.kudos.push(repo);
    },
    dislike(repo){
      const index = this.kudos.indexOf(repo);
      this.kudos.splice(index, 1);
    },
    search() {
      this.githubQuery(this.query);
    },
    githubQuery(query) {
      fetch("https://api.github.com/search/repositories?q=" + query)
        .then(r => r.json())
        .then((r)=> {
          this.repos = r.items;
        })
    }
  }

}
</script>

<style>
 .repo-card-content {
   height: 90px;
   overflow: scroll;
 }
 .v-tabs__content {
   padding-bottom: 2px;
 }
</style>

