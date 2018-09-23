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
        @keyup.enter="search"
        value="vuejs"
      ></v-text-field>
      <v-spacer></v-spacer>
    </v-toolbar>
    <v-container fluid class="teal lighten-4">
      <v-layout>
        <v-flex> 
          <div class="text-xs-center">
            <v-btn slot="activator" outline color="white lighten-2">Show My Kudos</v-btn>
          </div>
        </v-flex>
      </v-layout>
    </v-container>
      
    <v-container grid-list-md fluid class="grey lighten-4">
      <v-layout row wrap>
        <v-flex v-for="repo in repos" :key="repo.id" md4 xs12 sm12>
          <v-card >
            <v-card-title primary-title>
              <div class="repo-card-content">
                <h3 class="headline mb-0">{{repo.full_name}}</h3>
                <div>{{repo.description}}</div>
              </div>
            </v-card-title>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn @click.prevent="like(repo)" v-if="isLiked(repo)" flat icon color="pink"><v-icon>favorite</v-icon></v-btn>
              <v-btn @click.prevent="dislike(repo)" v-else flat icon color="pink"><v-icon>favorite_border</v-icon></v-btn>
              <v-dialog>
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
            </v-card-actions>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      repos: [],
      favs: [],
    }
  },
  mounted() {
    this.search({target: {value: 'vuejs'}})
  },
  methods: {
    isLiked(repo){
      return this.favs.includes(repo.id)
    },
    like(repo){
      console.log(repo);
      this.favs.push(repo.id);
    },
    dislike(repo){
      this.favs = this.favs.filter( fav => fav !== repo.id );
    },
    search(evt) {
      fetch("https://api.github.com/search/repositories?q=" + evt.target.value)
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
</style>

