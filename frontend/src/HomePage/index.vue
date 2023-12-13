<template>

  <div>


  <div v-show="ifshow">
      <!-- 回到顶部 -->
      <el-backtop :right="100" :bottom="100" />
      <!-- 上导航栏 -->
      <div class="Nav">
        <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" :ellipsis="false"
          @select="handleSelect">
          <div style="margin-right: 5px;" />
          <img style="width: 100px " src="./public/logo.png" alt="logo" />

          <!-- 起一个 隔断 的 作用 -->
          <div class="flex-grow" />
          <el-menu-item index="1" class="menuItem">网站首页</el-menu-item>
          <el-menu-item index="2">卡密查询</el-menu-item>
          <el-menu-item index="3">投诉进度</el-menu-item>
          <el-menu-item index="4">联系我们</el-menu-item>
          <el-menu-item index="5">常见问题</el-menu-item>
          <div class="flex-grow" />
          <el-button class="loginButton" @click="Sign">登录/注册</el-button>

        </el-menu>
      </div>
      <!-- 内容 -->
      <div style="max-width: 1080px;  margin: 0 auto;">
        <!-- page1 -->
        <div class="page1">
          <div class="pageCon1">
            <div class="tagCon1">
              <el-tag round style="height: 35px;">
                <el-tag effect="dark" type="danger" round class="tag1">
                  v1.0
                </el-tag>
                全平台Go开发,0.1ms响应
              </el-tag>
              <div class="text1">
                <el-text tag="b" style="font-size: 250%;">为 </el-text>
                <el-text type="primary" tag="b" style="font-size: 250%;">虚拟产品 </el-text>
                <el-text tag="b" style="font-size: 200%;">销售</el-text>

                <br>
                <el-text tag="b" style="font-size: 250%;">打造完美的服务平台</el-text>

              </div>
              <div style="margin-top: 5%;margin-bottom: 5%;">
                <el-text style="font-size: large; ">
                  致力于虚拟商品的快捷销售服务-安全,可靠
                </el-text>

              </div>
              <div>
                <el-button color="#626aef" round size="large" style="font-weight: bold;">快速入住 成为商户</el-button>
              </div>
            </div>
          </div>
          <div class="pageCon1">
            <img src="./public/homePage1.png" alt="" class="pageImg1">
          </div>
        </div>

        <!-- page 2 -->
        <div class="page2">
          <div style="margin-right: 2%;">
            <img src="./public/homePage2.png" alt="" class="pageImg2">
          </div>
          <div>
            <el-text tag="b" style="font-size: 200%;">开心发卡网 </el-text>
            <br>
            <el-text type="primary" tag="b" style="font-size: 200%;">虚拟产品 </el-text>

            <el-text tag="b" style="font-size: 200%;">销售服务</el-text>
            <br>
            <div style="margin-top: 5%;margin-bottom: 5%;">
              <el-text style="font-size: large; ">
                注册成为商家,将卡密托管在平台,(软件注册码、游戏帐号、虚拟商品卡密等) 为商家提供安全、便捷、绿色、稳定虚拟商品卡密自动交易服务

              </el-text>

            </div>



            <el-button color="#626aef" size="large" style="font-weight: bold;"> 立即体验 > </el-button>
          </div>
        </div>

        <!-- page 3 -->
        <div class="page3">
          <div style="margin-right: 3%;">
            <el-text tag="b" style="font-size: 200%;">企业级发卡项目 </el-text>
            <br>
            <el-text type="primary" tag="b" style="font-size: 200%;">响应时间低至0.1ms </el-text>
            <div style="margin-top: 5%;margin-bottom: 5%;">
              <el-text style="font-size: large; ">
                致力于为用户提供最好的虚拟产品交易服务
              </el-text>
            </div>
            <div>
              <el-text><el-icon>
                  <CircleCheck />
                </el-icon> 全平台Vue+Go开发 </el-text>
              <br>
              <el-text><el-icon>
                  <CircleCheck />
                </el-icon> 为商家提供一站式虚拟商品在线购买及全自动发货 </el-text>
              <br>
              <el-text><el-icon>
                  <CircleCheck />
                </el-icon> 平台系统功能齐全，支付即发卡，数万用户的信赖之选 </el-text>
            </div>

          </div>
          <div>
            <img src="./public/homePage3.png" alt="" class="pageImg3">
          </div>
        </div>
      </div>

      <!-- page4 -->
      <el-divider />
      <div class="page4">
        <el-row style="width: 500px; margin-left: 5%;">
          <el-col :span="6">
            <el-statistic title="总注册人数" :value="2685" />
          </el-col>
          <el-col :span="6">
            <el-statistic title="总商户数" :value="78" />
          </el-col>
          <el-col :span="6">
            <el-statistic title="总商品数" :value="1720" />
          </el-col>
          <el-col :span="6">
            <el-statistic title="平台成交总数" :value="55" />
          </el-col>
        </el-row>
      </div>
      <el-divider />
      <!-- page5 -->
      <div class="page5">
        <div>
          <img src="./public/homePage5.png" alt="" class="pageImg5">
        </div>
        <div>
          <el-text tag="b" style="font-size: 200%;">开心发卡网 </el-text>
          <br>
          <el-text type="primary" tag="b" style="font-size: 200%;">平台公告 </el-text>
          <br>
          <div style="margin-top: 5%;margin-bottom: 5%;">
            <el-text style="font-size: large; ">
              我们将一些重要信息公布于此，您可以再此了解最新动态
            </el-text>
          </div>
          <el-divider />
          <div class="pageCon5">
            <div style="margin-top: 8px;margin-right: 10px;">
              <el-icon :size="50">
                <ChatSquare />
              </el-icon>
            </div>
            <div>
              掌握平台最新咨询
              <br>
              <el-link href="#" type="primary">即刻进入,了解详情</el-link>

            </div>

          </div>
        </div>
      </div>

      <!-- page-end -->
      <div class="pageEnd">
        Copyright © 2023 开心发卡网 All rights reserved. 版权所有
      </div>
    </div>

  </div>
</template>
<script>
import { ref } from 'vue'
import {
  CircleCheck,
  ChatSquare
} from '@element-plus/icons-vue'

export default {
  data() {
    return {
      ifshow: false
    };
  },
  components: {
    CircleCheck,
    ChatSquare
  },
  methods: {
    Sign() {
      this.$router.push('/login');
    }
  },
  mounted() {
    const openFullScreen2 = () => {
      const loading = this.$loading({
        lock: true,
        text: 'Loading',
        background: 'rgb(255, 255, 255)',
      });

      setTimeout(() => {
        loading.close();
        this.ifshow = true
      }, 500);

    };
    openFullScreen2(); // 如果需要在组件挂载后立即执行该函数，可以调用它
  },
}
</script>
<style scoped>
.Nav {
  /* margin-bottom: 50px; */
  position: fixed;
  top: 0;
  width: 100%;
}

.flex-grow {
  flex-grow: 1;
}

.loginButton {
  width: auto;
  height: 80%;
  margin-top: 12px;
  margin-right: 5%;
}

.page1 {

  display: flex;
  /* 使用 Flexbox 布局 */
  justify-content: space-between;
  /* 左右分布 */
  width: auto;
  height: 800px;
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */


}

.pageCon1 {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;


}

.pageImg1 {
  height: 500px;
  width: 500px;
  justify-content: center;
  align-items: center;
  margin-left: -10%;
}

.tag1 {
  margin-right: 5px;
}

.tagCon1 {

  margin-right: -10%;
}

.page2 {
  display: flex;
  /* 使用 Flexbox 布局 */
  justify-content: space-between;
  /* 左右分布 */
  width: auto;
  height: 800px;
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
}

.pageImg2 {
  width: 500px;
  height: 500px;
}

.page3 {
  display: flex;
  /* 使用 Flexbox 布局 */
  justify-content: space-between;
  /* 左右分布 */
  width: auto;
  height: 800px;
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
}

.pageImg3 {
  width: 500px;
  height: 500px;
}

.page4 {
  display: flex;

  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
  height: 100px;

}

.pageEnd {
  display: flex;
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
  height: 50px;
  background-color: rgb(32, 41, 66);
  font-size: 14px;
  color: rgb(177, 186, 195);
}

.page5 {
  display: flex;
  /* 使用 Flexbox 布局 */
  justify-content: space-between;
  /* 左右分布 */
  width: auto;
  height: 800px;
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
}

.pageImg5 {
  width: 400px;
  height: auto;
}

.pageCon5 {
  display: flex;
  /* 使用 Flexbox 布局 */

  justify-content: flex-start;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
}


</style>