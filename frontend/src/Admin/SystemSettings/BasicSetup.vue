<!-- 基础设置页面 -->
<template>
    <div class="set">
        <p>基本设置</p>
        <hr>
        <div class="label-input">
            <span>店铺名称</span>
            <input type="text" id="shopLabel" name="shopLabel" v-model="name">
        </div>
        <div class="label-input">
            <span>网站标题</span>
            <input type="text" id="shopLabel" name="shopLabel" v-model="title">
        </div>
        <div class="label-input">
            <span>关键词</span>
            <input type="text" id="shopLabel" name="shopLabel" v-model="seo">
        </div>
        <div class="label-input">
            <span>网站描述</span>
            <input type="text" id="shopLabel" name="shopLabel" v-model="description">
        </div>
        <div class="label-input">
            <span>店铺公告</span>
            <input type="text" id="shopLabel" name="shopLabel" v-model="autoMessage">
        </div>
        <div class="label-button">
            <button>
                <span @click="editBasicSetup">
                    保存配置
                </span>
            </button>
            <button>
                <span>
                    恢复初始配置
                </span>
            </button>

        </div>

    </div>

</template>
<script>
import axios from 'axios'

import { ElNotification } from 'element-plus'



export default {
    data() {
        return {
            name: "",
            title: "",
            seo: "",
            description: "",
            autoMessage: "",

            ifSetupSuccess: false,
        };
    },
    mounted() {
        axios.get('http://localhost:8081/api/admin/getInfo')
            .then(response => {
                this.name = response.data.name;
                this.title = response.data.title;
                this.seo = response.data.seo;
                this.description = response.data.description;
                this.autoMessage = response.data.autoMessage;
            })
            .catch(error => {
                console.error('Error:', error);
            });


    },
    methods: {
        editBasicSetup() {
            axios.post("http://localhost:8081/api/admin/editInfo", {
                name: this.name,
                title: this.title,
                seo: this.seo,
                description: this.description,
                autoMessage: this.autoMessage
            })
                .then(response => {
                    // 获取后端返回的状态码
                    const statusCode = response.status;
                    if (statusCode === 200) {
                        // this.ifSetupSuccess = true;
                        // setTimeout(() => {
                        //     this.ifSetupSuccess = false;
                        // }, 500);

                        ElNotification({
                            title: '保存配置成功',
                            message: '基础设置已更新',
                            type: 'success',
                        });
                    } else {
                        alert("请联系网站管理员");
                    }
                })
                .catch(error => {
                    console.error("请求失败:", error);
                    alert("请求失败，请联系网站管理员");
                });
        }
    },

    computed: {

    }
}
</script>

<style scoped> /* 添加样式到包含店铺标签和输入框的容器 */
 .label-input {
     display: flex;
     align-items: center;
     margin-top: 10px;
     margin-bottom: 5px;
 }

 /* 标签的样式 */
 .label-input span {
     flex: 1;
     /* 让标签占据容器的一部分 */
     font-size: 15px;
     margin-left: 5%;

 }

 /* 输入框的样式 */
 .label-input input {
     flex: 0.8;
     /* 让输入框占据容器的1部分 */
     padding: 5px;
     border: 1px solid #ccc;
     border-radius: 3px;
     height: 24px;
     margin-right: 10%;
     background-color: rgb(245, 248, 250);

 }

 .set {
     margin-top: 1%;
     display: flex;
     flex-direction: column;
     justify-content: flex-start;
     background-color: white;
     padding: 20px;
     box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
     border-radius: 5px;
     height: auto;
 }

 .set hr {
     color: dimgray;
     margin: 5px 0;
     /* 上下间距 */
     border: 0;
     border-top: 1px solid #ccc;

 }

 .set p {
     font-size: 20px;
     font-weight: bold;
     margin-left: 4%;

 }

 .label-button {
     margin-top: 10%;
     margin-left: 5%;
     margin-bottom: 5%;
 }

 .label-button button {
     border: 2px solid #24b4fb;
     background-color: #24b4fb;
     border-radius: 0.9em;
     padding: 0.8em 1.2em 0.8em 1em;
     transition: all ease-in-out 0.2s;
     font-size: 12px;
     margin-right: 200px;
 }

 .label-button button span {
     display: flex;
     justify-content: center;
     align-items: center;
     color: #fff;
     font-weight: 600;
 }

 .label-button button:hover {
     background-color: #0071e2;
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