<!-- 此页面用于 后台管理系统 登录验证 -->



<template>
    <div class="app">
        <form class="form" @submit.prevent="submitForm">
            <p class="form-title">后台管理系统</p>
            <div class="input-container">
                <input v-model="username" type="username" placeholder="输入您的账号">
            </div>
            <div class="input-container">
                <input v-model="password" type="password" placeholder="输入您的密码">
            </div>
            <div style="display: grid; place-items: center;">
                <button type="submit" class="buttonlogin" @click="loginAdmin">
                    登录
                </button>
            </div>

            <p class="signup-link">
                忘记密码?
                <a href="#">找回密码</a>
            </p>
        </form>
        <div>
            <transition name="bounce" mode="out-in">
                <div v-show="loginSuccess" key="success" class="success-message">
                    <i class="fas fa-check-circle"></i>
                    登录成功！欢迎回来，{{ username }}！
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

            //防止无限调用接口 引起卡顿
            time: true
        };
    },
    methods: {
        async loginAdmin() {
            if (this.time === true) {
                this.time = false
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

        },
    }
};
</script>

<style scoped>
.buttonlogin {
    height: 3em;
    width: 20em;
    border: none;
    border-radius: 10em;
    background: #ea82ff;
    font-size: 17px;
    color: #ffffff;
    font-family: inherit;
    font-weight: 500;
    margin-top: 2%;
}

.buttonlogin:hover {
    animation: shake3856 0.3s linear infinite both;
}

@keyframes shake3856 {
    0% {
        -webkit-transform: translate(0);
        transform: translate(0);
    }

    20% {
        -webkit-transform: translate(-2px, 2px);
        transform: translate(-2px, 2px);
    }

    40% {
        -webkit-transform: translate(-2px, -2px);
        transform: translate(-2px, -2px);
    }

    60% {
        -webkit-transform: translate(2px, 2px);
        transform: translate(2px, 2px);
    }

    80% {
        -webkit-transform: translate(2px, -2px);
        transform: translate(2px, -2px);
    }

    100% {
        -webkit-transform: translate(0);
        transform: translate(0);
    }
}

.app {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
}

.form {

    background-color: #fff;
    display: block;
    padding: 1rem;
    max-width: 450px;
    border-radius: 0.5rem;
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

.form-title {

    caret-color: transparent;
    /* or any color you prefer */

    font-size: 1.25rem;
    line-height: 1.75rem;
    font-weight: 600;
    text-align: center;
    color: #000;
}

.input-container {
    position: relative;
    caret-color: tr;
}

.input-container input {
    outline: none;
    border: 1px solid #e5e7eb;
    margin: 8px 0;
}

.input-container input {
    background-color: #fff;
    padding: 1rem;
    padding-right: 3rem;
    font-size: 0.875rem;
    line-height: 1.25rem;
    width: 300px;
    border-radius: 0.5rem;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);

}





.signup-link {
    color: #6B7280;
    font-size: 0.875rem;
    line-height: 1.25rem;
    text-align: center;
    caret-color: transparent;
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