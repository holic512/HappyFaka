<template>
    <!-- 在 index.html 或 App.vue 中引入 Google Fonts 字体链接 -->
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap">

    <div class="centered-container">
        <div class="rounded-rectangle">
            <div class="title">{{ name }}</div>
            <el-divider />
            <div class="content">
                <!-- 正常的内容放在这里 -->
                <p v-html="messageWithlineBreaks"></p>
            </div>
        </div>
    </div>
</template>
  
<script>

import axios from 'axios'


export default {
    data() {
        return {
            name: "",
            message: "",
        };
    },
    computed: {
        messageWithlineBreaks() {
            return this.message.replace(/\n/g, '<br>');
        }
    },
    mounted() {
        axios.get('http://localhost:8082/api/user/message')
            .then(response => {
                this.name = response.data.name;
                this.message = response.data.message
            })
            .catch(error => {
                console.error('Error:', error);
            });
    }
};
</script>
  
<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Ma+Shan+Zheng&display=swap');

.centered-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 260px;
    /* 使容器占满整个视口高度，垂直居中 */


    caret-color: transparent;
}

.rounded-rectangle {
    width: 900px;
    height: 200px;
    background-color: white;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    /* 设置为列布局 */
    align-items: flex-start;
    /* 左对齐 */
    padding: 20px;
    /* 添加一些内边距 */
    box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
    /* 添加下边界的阴影效果 */
}

.title {
    font-size: 24px;
    margin-top: 1px;
    /* 与下方的分隔线添加间距 */
    font-family: 'Ma Shan Zheng', cursive;
    letter-spacing: 1px;
    color: forestgreen;

}


.content {
    margin-top: -15px;
    /* 正常的内容样式 */
    font-size: 18px;
    font-family: 'Ma Shan Zheng', cursive;
    /* 更换字体 */
    line-height: 1.6;
    /* 调整行高，提高可读性 */
    letter-spacing: 1.5px;
    color: red;

}
</style>
  