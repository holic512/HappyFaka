<template>
  <div class="sys">
    <!-- 上导航栏 -->
    <Nav @Index="getindex" />


    <!-- 当 index 为 0 的时候 为基础设置 -->
    <BasicSet v-if="this.index === 0"/>


    <!-- 当 index 为 0 的时候 为基础设置 -->



    <!-- 当 index 为 2 的时候 为邮箱设置 -->
    <EmailSet v-else-if="this.index === 2"/>



    
    <!-- 当 index 为 0 的时候 为基础设置 -->

  </div>
</template>
  
<script>
import Nav from './Nav.vue'
import BasicSet from './BasicSetup.vue'
import EmailSet from './EmailSetup.vue'
export default {
  data() {
    return {
      index : 0
    };
  },
  methods: {
    getindex(data) {
      this.index = data;
    }
  },
  components: {
    Nav,
    BasicSet,
    EmailSet
  },
};
</script>
  
<style scoped>
.sys {
  margin: 1%;
  height: auto;
}

</style>
  