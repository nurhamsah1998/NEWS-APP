<!-- eslint-disable vue/no-export-in-script-setup -->
<!-- eslint-disable vue/multi-word-component-names -->
<script lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { defineComponent } from 'vue';
import router from './router';

export default {
  data(){
    return {
      name:"NURHAMSAH"
    }
  }, 
  beforeMount(){
    if(!this.$cookies.get('access_token')) router.replace({path:"/login"})
  },
  methods:{
    getToken(){
      console.log(this.$cookies.get('access_token'))
    }
  }
}
</script>

<template>
  <!-- <header>
    <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" />

    <div class="wrapper">
      <HelloWorld msg="You did it!" />

      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </div>
  </header> -->

  <RouterView />
</template>

<style scoped>

</style>
