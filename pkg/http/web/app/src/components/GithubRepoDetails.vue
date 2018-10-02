<template>
  <v-container grid-list-md fluid class="grey lighten-4" >
    <v-layout align-center justify-space-around wrap>
      <v-flex md6>
        <v-img
          :src="repo.owner.avatar_url"
          :alt="repo.owner.login"
          class="grey darken-4"
          width="200"
        ></v-img>
        <h1 class="primary--text">
          <a :href="repo.html_url">{{repo.full_name}}</a>
        </h1>

        <v-chip class="text-xs-center">
          <v-avatar class="teal">
            <v-icon class="white--text">star</v-icon>
          </v-avatar>
          Stars: {{repo.stargazers_count}}
        </v-chip>

        <v-chip class="text-xs-center">
          <v-avatar class="teal white--text">L</v-avatar>
          Language: {{repo.language}}
        </v-chip>

        <v-chip class="text-xs-center">
          <v-avatar class="teal white--text">O</v-avatar>
          Open Issues: {{repo.open_issues_count}}
        </v-chip>

        <v-textarea
          name="input-7-1"
          label="Show some love"
          value=""
          hint="Describe why you love this project"
        ></v-textarea>
        <v-btn @click.prevent="submit"> Kudo </v-btn>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>

export default {
  data() {
    return {
      repo: {}
    }
  },
  watch: {
    '$route': 'fetchData'
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      // FIXME: fetch from server
      fetch('https://api.github.com/repositories/' + this.$route.params.id)
        .then(r => r.json())
        .then(r => this.repo = r )
    }
  }
}
</script>

<style>
</style>

