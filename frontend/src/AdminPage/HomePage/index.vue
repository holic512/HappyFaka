<template>
  <div class="common-layout">
    <el-container>

      <el-aside width="auto" style="height: 95vh;">
        <!--滚动条-->
        <el-scrollbar>
          <!--侧边栏-->
          <el-menu default-active="2" style="height: 95vh;" :collapse="isCollapse" @open="handleOpen"
                   @close="handleClose">
            <el-menu-item index="1">
              <el-icon>
                <icon-menu/>
              </el-icon>
              <template #title>仪表盘</template>
            </el-menu-item>

            <el-sub-menu index="2">
              <template #title>
                <el-icon>
                  <User/>
                </el-icon>
                <span>用户管理</span>
              </template>
              <el-menu-item index="2-1">所有商户</el-menu-item>
              <el-menu-item index="2-2">所有买家</el-menu-item>
              <el-menu-item index="2-3">提现管理</el-menu-item>
            </el-sub-menu>

            <el-menu-item index="3">
              <el-icon>
                <Shop/>
              </el-icon>
              <template #title>所有店铺</template>
            </el-menu-item>

            <el-sub-menu index="4">
              <template #title>
                <el-icon>
                  <ShoppingBag/>
                </el-icon>
                <span>商品管理</span>
              </template>
              <el-menu-item index="4-1">所有分类</el-menu-item>
              <el-menu-item index="4-2">所有商品</el-menu-item>
              <el-menu-item index="4-3">所有卡密</el-menu-item>
            </el-sub-menu>

            <el-menu-item index="5">
              <el-icon>
                <Picture/>
              </el-icon>
              <template #title>广告管理</template>
            </el-menu-item>

            <el-menu-item index="6">
              <el-icon>
                <Present/>
              </el-icon>
              <template #title>折扣优惠</template>
            </el-menu-item>

            <el-sub-menu index="7">
              <template #title>
                <el-icon>
                  <location/>
                </el-icon>
                <span>工单管理</span>
              </template>
              <el-menu-item index="7-1">商户工单</el-menu-item>
              <el-menu-item index="7-2">用户投诉</el-menu-item>
            </el-sub-menu>

            <el-menu-item index="8">
              <el-icon>
                <More/>
              </el-icon>
              <template #title>插件配置</template>
            </el-menu-item>

            <el-menu-item index="9">
              <el-icon>
                <Setting/>
              </el-icon>
              <template #title>总站设置</template>
            </el-menu-item>
          </el-menu>
        </el-scrollbar>

      </el-aside>
      <!--缩放-->
      <div class="Scalebar">
        <!--        <button @click="toggleCollapse"> 123</button>-->
        <div @click="toggleCollapse">
          <div class="line"></div>
        </div>

      </div>


      <el-container>
        <el-header class="topNav1">
          <div class="topNav">
            <div class="input1">
              <form class="form">
                <button>
                  <svg width="17" height="16" fill="none" xmlns="http://www.w3.org/2000/svg" role="img"
                       aria-labelledby="search">
                    <path d="M7.667 12.667A5.333 5.333 0 107.667 2a5.333 5.333 0 000 10.667zM14.334 14l-2.9-2.9"
                          stroke="currentColor" stroke-width="1.333" stroke-linecap="round" stroke-linejoin="round">
                    </path>
                  </svg>
                </button>
                <input class="input" placeholder="输入订单编号..." required="" type="text">
                <button class="reset" type="reset">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                       stroke="currentColor"
                       stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12"></path>
                  </svg>
                </button>
              </form>
            </div>
          </div>

        </el-header>
        <el-main>Main</el-main>
      </el-container>
    </el-container>
  </div>
</template>
<script lang="ts">
import {
  Document,
  Location,
  Menu as IconMenu,
  More,
  Picture,
  Present,
  Setting,
  Shop,
  ShoppingBag,
  User
} from '@element-plus/icons-vue'

export default {
  data() {
    return {
      isCollapse: true
    };
  },
  components: {
    Document,
    Location,
    IconMenu,
    Setting,
    User,
    More,
    Present,
    ShoppingBag,
    Shop,
    Picture

  },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath)
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath)
    },
    toggleCollapse() {
      this.isCollapse = !this.isCollapse;
      // 在这里可以添加其他你希望执行的逻辑
    }
  }
}
</script>


<style>
.Scalebar {
  display: grid;
  place-items: center; /* 水平和垂直居中 */
  height: auto; /* 为了垂直居中整个视口可见区域的内容 */


}

.line {
  margin-left: 5px;
  height: 65px;
  width: 5px;
  border-radius: 9999px; /* 圆角 */
  background-color: rgb(137, 135, 135); /* 灰色 */
  cursor: pointer; /* 鼠标悬停时显示指针 */
}

.input1 {

  padding: 3px;
  width: 200px;
  margin-left: 10px;

}

.topNav1 {
  margin-left: -10px;
  background-color: white;
  height: 50px;
  /* 添加下边界的阴影效果 */
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.topNav {
  width: 100%;
  flex: 1;
}


/* From uiverse.io by @satyamchaudharydev */
/* removing default style of button */

.form button {
  border: none;
  background: none;
  color: #8b8ba7;
}

/* styling of whole input container */
.form {
  --timing: 0.3s;
  --width-of-input: 200px;
  --height-of-input: 40px;
  --border-height: 2px;
  --input-bg: rgb(244, 243, 249);
  --border-color: #2f2ee9;
  --border-radius: 30px;
  --after-border-radius: 1px;
  position: relative;
  width: var(--width-of-input);
  height: var(--height-of-input);
  display: flex;
  align-items: center;
  padding-inline: 0.8em;
  border-radius: var(--border-radius);
  transition: border-radius 0.5s ease;
  background: var(--input-bg, #fff);
}

/* styling of Input */
.input {
  font-size: 0.9rem;
  background-color: transparent;
  width: 100%;
  height: 100%;
  padding-inline: 0.5em;
  padding-block: 0.7em;
  border: none;
}

/* styling of animated border */
.form:before {
  content: "";
  position: absolute;
  background: var(--border-color);
  transform: scaleX(0);
  transform-origin: center;
  width: 100%;
  height: var(--border-height);
  left: 0;
  bottom: 0;
  border-radius: 1px;
  transition: transform var(--timing) ease;
}

/* Hover on Input */
.form:focus-within {
  border-radius: var(--after-border-radius);
}

input:focus {
  outline: none;
}

/* here is code of animated border */
.form:focus-within:before {
  transform: scale(1);
}

/* styling of close button */
/* == you can click the close button to remove text == */
.reset {
  border: none;
  background: none;
  opacity: 0;
  visibility: hidden;
}

/* close button shown when typing */
input:not(:placeholder-shown) ~ .reset {
  opacity: 1;
  visibility: visible;
}

/* sizing svg icons */
.form svg {
  width: 17px;
  margin-top: 3px;
}
</style>
