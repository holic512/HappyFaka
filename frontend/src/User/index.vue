<template>
  <div class="body">
    <Navbar></Navbar>
    <Mymessage></Mymessage>
  </div>


</template>



<script>
import Navbar from './Navbar/index.vue'
import Mymessage from './Mymessage/index.vue'
export default {
  components: {
    Navbar,
    Mymessage
  }
}


</script >
<style scoped>
.body {
  background-color: rgb(248, 248, 248);
  width: 100%;
  height: 100%;
  
}



</style>