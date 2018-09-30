<template>
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
</template>

<script>

import { mapActions } from 'vuex'

export default {
  data() {
    return {}
  },
  props: ['repo'],
  methods: {
    isKudo(repo) {
      return this.$store.getters.isKudo(repo);
    },
    ...mapActions(['toggleKudo'])
  }
}
</script>

<style>
 .repo-card-content {
   height: 90px;
   overflow: scroll;
 }
</style>

