<!-- 此页面用于 后台管理系统 登录验证 -->



<template>
    <div class="app">
        <div class="login-container">
            <div class="login-header">
                <div>后台管理系统</div>
            </div>
            <input type="text" class="login-input" placeholder="请输入用户名" id="username" v-model="this.username">
            <input type="password" class="login-input" placeholder="请输入密码" id="password" v-model="this.password">
            <button class="login-button" id="login-button" @click="loginAdmin">登录</button>
        </div>

        <div>
            <transition name="bounce" mode="out-in">
                <div v-show="loginSuccess" key="success" class="success-message">
                    <i class="fas fa-check-circle"></i>
                    登录成功！欢迎回来{{ username }}！
                </div>
            </transition>
        </div>


        <transition name="bounce" mode="out-in">
            <div v-show="nologinSuccess" key="" class="success-message">
                <i class="fas fa-check-circle"></i>
                密码不正确，请重新输入
            </div>
        </transition>

        <transition name="bounce" mode="out-in">
            <div v-show="noNameOrPassword" key="" class="success-message">
                <i class="fas fa-check-circle"></i>
                用户名和密码不能为空
            </div>
        </transition>
    </div>
</template>
  
<script>
import axios from "axios"

export default {
    components: {

    },
    data() {
        return {
            login: false,
            username: '',
            password: '',

            loginSuccess: false,
            nologinSuccess: false,
            noNameOrPassword: false,
        };
    },
    methods: {
        async loginAdmin() {

            if (!this.username || !this.password) {
                this.noNameOrPassword = true;
                setTimeout(() => {
                    this.noNameOrPassword = false;
                }, 2000);
                return;
            }

            try {
                const response = await axios.post("http://localhost:8080/admin/login", {
                    username: this.username,
                    password: this.password,
                });

                // 获取后端返回的状态码
                const statusCode = response.status;
                if (statusCode === 200) {
                    this.loginSuccess = true;
                    setTimeout(() => {
                        this.loginSuccess = false;
                        this.login = true;
                    }, 500);

                } else {
                    alert("请联系网站管理员");
                }
            } catch (error) {
                this.nologinSuccess = true;
                setTimeout(() => {
                    this.nologinSuccess = false;
                }, 500);
            }
            setTimeout(() => {
                this.time = true;
            }, 500);


        }
    }
};
</script>

<style scoped>
.app {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 97vh;
    background-color: rgb(232, 232, 232);
}

.login-container {
    width: 350px;
    padding: 20px;
    background: #fff;
    border-radius: 8px;
    box-shadow: 0px 10px 20px rgba(0, 0, 0, 0.1), 0px 0px 0px 3px #3498db, inset 0px 0px 10px rgba(0, 0, 0, 0.1);
}

.login-header {
    font-size: x-large;
    font-weight: 600;
    text-align: center;
    margin-bottom: 20px;
}

.login-header h1 {
    color: #333;
    font-size: 24px;
}

.login-input {
    width: 95%;
    padding: 10px;
    margin-bottom: 15px;
    border: none;
    background: #f5f5f5;
    border-radius: 4px;
    box-shadow: inset 0px 1px 3px rgba(0, 0, 0, 0.2);
}

.login-button {
    width: 100%;
    padding: 10px;
    border: none;
    background: #3498db;
    color: #fff;
    font-weight: bold;
    border-radius: 4px;
    cursor: pointer;
    box-shadow: 0px 3px 6px rgba(0, 0, 0, 0.1), inset 0px 1px 3px rgba(255, 255, 255, 0.5);
    transition: background 0.3s ease, transform 0.2s ease;
}

.login-button:hover {
    background: #2980b9;
    transform: translateY(-2px);
}


.success-message {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    padding: 20px;
    background-color: #6a1b9a;
    color: #fff;
    border-radius: 10px;
    box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
    animation: bounceIn 0.8s, fadeOutUp 0.5s 2.5s;
}

@keyframes bounceIn {
    0% {
        opacity: 0;
        transform: translate(-50%, -70%) scale(0.3);
    }

    70% {
        opacity: 1;
        transform: translate(-50%, -50%) scale(1.05);
    }

    100% {
        transform: translate(-50%, -50%);
    }
}

@keyframes fadeOutUp {
    0% {
        opacity: 1;
        transform: translate(-50%, -50%);
    }

    100% {
        opacity: 0;
        transform: translate(-50%, -70%);
    }
}
</style>