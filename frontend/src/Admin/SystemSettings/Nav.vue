<template>
    <link href='http://cdn.repository.webfont.com/webfonts/nomal/152961/19673/65096f1a6f50f60be06ad7a0.css' rel='stylesheet'
      type='text/css' />
    <div class="navbar">
      <a style="color:black; margin-top: 6px; font-weight: normal;font-family:'jdzhonyuanjian395203495925581';">网站配置</a>
      <el-divider direction="vertical" style="margin-top: 6px;"/>
      <div class="nav-item" v-for="(item, index) in navItems" :key="index">
        <button @click="changeActive(index)" :class="{ active: activeIndex === index }">{{ item }}</button>
      </div>
    </div>
  </template>
    
  <script>
  
  export default {
    data() {
      return {
        navItems: ["基础设置", "短信设置", "邮箱设置", "其他设置"],
        activeIndex: 0,
      };
    },
    methods: {
      changeActive(index) {
        this.activeIndex = index;
        // 在这里可以添加切换选项卡时的逻辑
        
        // 传递给父组件
        this.$emit("Index",this.activeIndex)
      },
    },
  };
  </script>
    
  <style scoped>
  .navbar {
    display: flex;
    justify-content: flex-start;
    background-color: white;
  
    padding: 10px;
    box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
    border-radius: 5px;
  }
  
  .nav-item {
  
    margin-right: 10px;
  }
  
  button {
  
    background-color: transparent;
    border: none;
    color: slategray;
    cursor: pointer;
    padding: 5px 10px;
  }
  
  button.active {
    color: black;
  }
  
  button:hover {
    color: cadetblue;
  }
  </style>
    