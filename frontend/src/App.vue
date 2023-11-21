<template>
  <title>{{ this.title }}</title>
  <router-view></router-view>
</template>


<script >
import axios from 'axios';
export default{
  data(){
    return{
      title: ""
    }
    
  },
  mounted() {
        axios.get('http://localhost:8082/api/user/title')
            .then(response => {
                this.title = response.data.title;
            })
            .catch(error => {
                console.error('Error:', error);
            });
    }
}

</script>



<style scoped></style>
