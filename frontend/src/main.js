import { createApp } from 'vue'
import App from './App.vue'
import { createRouter, createWebHashHistory } from 'vue-router';
import axios from 'axios';
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// 设置 Axios 的默认配置
axios.defaults.withCredentials = true;

const app = createApp(App);

// 将 Axios 实例挂载到 Vue 原型上，在组件中通过 this.$axios 使用
app.config.globalProperties.$axios = axios;

// 导入需要使用的组件
// import Home from './User/index.vue'
import Home from './HomePage/index.vue'
import Login from './Sign/login.vue'
import Register from './Sign/register.vue'

import Admin from './AdminPage/login/index.vue'
import AdminPage from './AdminPage/HomePage/index.vue'
const routes = [
    {
        path: '/',
        component: Home
    },
    {
        path: '/login',
        component: Login
    },
    {
        path: '/register',
        component: Register
    },

    {
        path: '/admin',
        component: Admin
    },
    {
        path: '/admin/home',
        component: AdminPage
    }
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

app.provide('$axios', axios);
app.use(router);
app.use(ElementPlus);

app.mount('#app');
