<template>
  <div id="app">
    <div class=""></div>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
html {
    height: 100%;
    background: #fff;
}
body {
    height: 100%;
    background: #ccc;
}
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  position: relative;
  max-width: 375px;
  margin: 0 auto;
  /* border: 1px solid #ededed; */
  /* border-bottom: 0; */
}
</style>
