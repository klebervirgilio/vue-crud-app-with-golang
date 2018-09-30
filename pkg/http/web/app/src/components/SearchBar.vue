<template>
    <v-toolbar dark color="teal">
      <v-spacer></v-spacer>
      <v-text-field
        solo-inverted
        flat
        hide-details
        label="Search for your OOS project on Github + Press Enter"
        prepend-inner-icon="search"
        v-model="query"
        @keyup.enter="onSearchSubmition"
      ></v-text-field>
      <v-spacer></v-spacer>
    </v-toolbar>
</template>

<script>
export default {
    data() {
        return {
            query: 'okta',
        };
    },
    props: ['defaultQuery'],
    mounted() {
        if(!this.defaultQuery) return;
        
        this.query = this.defaultQuery;
        this.onSearchSubmition();
    },
    methods: {
        onSearchSubmition() {
            this.$emit('search-submitted', this.query);
        }
    }
}
</script>