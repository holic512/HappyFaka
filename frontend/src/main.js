import { createApp } from 'vue'
import App from './App.vue'
import { createRouter, createWebHashHistory } from 'vue-router'; // 导入 createRouter 和 createWebHistory 函数

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import Home from './User/index.vue'
import Admin from './Admin/index.vue'

const app = createApp(App);



const routes = [
    {
        path: '/',
        component: Home
    },
    {
        path: '/admin',
        component: Admin
    }
];

// 创建路由实例
const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

// 将路由实例挂载到Vue实例上
app.use(router);
app.use(ElementPlus);

// 将Vue实例挂载到#app元素
app.mount('#app');