<template>
  <el-form label-position="left" label-width="100px" :model="emailInfo" class="set">
    <p>SMTP配置</p>
    <el-divider />
    <el-form-item label="SMTP服务器">
      <el-select
        v-model="emailInfo.server"
        placeholder="请选择你的服务器类型"
        style="width: 100%"
      >
        <el-option label="QQ服务器" value="QQ" />
      </el-select>
    </el-form-item>
    <div v-if="emailInfo.server === QQ">
      <el-form-item label="用户名">
        <el-input v-model="emailInfo.name" />
      </el-form-item>
      <el-form-item label="授权码">
        <el-input v-model="emailInfo.authcode" />
      </el-form-item>
    </div>

    <div class="label-button">
      <button>
        <span text @click="open"> 发送测试邮件(Send Test Message) </span>
      </button>
      <button>
        <span @click="editEmailSetup"> 保存设置(Save Config) </span>
      </button>
    </div>
  </el-form>
</template>

<script>
import { reactive, onMounted } from "vue";
import { ElMessage, ElMessageBox } from "element-plus";
import axios from "axios";
import { ElNotification } from "element-plus";

export default {
  data() {
    return {
      emailInfo: {
        server: "",
        name: "",
        authcode: "",
      },
      QQ: "QQ",
    };
  },
  mounted() {
    axios
      .get("http://localhost:8083/api/email/getInfo")
      .then((response) => {
        this.emailInfo.server = response.data.server;
        this.emailInfo.name = response.data.name;
        this.emailInfo.authcode = response.data.authcode;
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  },
  methods: {
    open() {
      ElMessageBox.prompt("请输入邮箱地址", "邮箱地址", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        inputPattern: /[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?/,
        inputErrorMessage: "错误的邮箱格式",
      }).then(({ value }) => {
        axios
          .post("http://localhost:8083/api/email/sendEmailIf", {
            sendEmail: value,
          })
          .then((response) => {
            // 获取后端返回的状态码
            const statusCode = response.status;
            if (statusCode === 200) {
              ElNotification({
                title: "测试邮件发送成功",
                message: "已成功发送",
                type: "success",
              });
            } else {
              alert("请联系网站管理员");
            }
          })
          .catch((error) => {
            console.error("请求失败:", error);
            alert("请求失败，请联系网站管理员");
          });
      });
    },
    editEmailSetup() {
      axios
        .post("http://localhost:8083/api/email/editInfo", {
          server: this.emailInfo.server,
          name: this.emailInfo.name,
          authcode: this.emailInfo.authcode,
        })
        .then((response) => {
          // 获取后端返回的状态码
          const statusCode = response.status;
          if (statusCode === 200) {
            ElNotification({
              title: "保存配置成功",
              message: "邮箱设置已更新",
              type: "success",
            });
          } else {
            alert("请联系网站管理员");
          }
        })
        .catch((error) => {
          console.error("请求失败:", error);
          alert("请求失败，请联系网站管理员");
        });
    },
  },
};
</script>

<style scoped>
.set {
  margin-top: 1%;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  background-color: white;
  box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
  border-radius: 5px;
  height: auto;
  padding: 5%;
  padding-top: 20px;
}

.set p {
  font-size: 20px;
  font-weight: bold;
  margin-bottom: 0;
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
</style>
