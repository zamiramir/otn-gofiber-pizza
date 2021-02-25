<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <img alt="Fiber logo" src="./assets/fiber_v2_logo.svg">
    <img alt="Go logo" src="./assets/go-logo-blue.svg">
    <HelloWorld msg="Welcome to Open Tech Night!"/>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',
  components: {
    HelloWorld
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
img {
  height:50px;
}
</style>
