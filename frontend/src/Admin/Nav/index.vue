<template>
    <el-row class="menu-container">
        <el-col :span="12">
            <h3 class="h3">开心发卡网</h3>
            <div class="menu-container">
                <el-menu active-text-color="#ffd04b" background-color="transparent"
                    class="el-menu-vertical-demo custom-menu" default-active="2" text-color="#fff" @open="handleOpen"
                    @close="handleClose">
                    <h4 class="h4">MAIN</h4>

                    <el-menu-item index="1" @click="changeActive(1)" >
                        <el-icon><icon-menu /></el-icon>
                        <span>控制台</span>
                    </el-menu-item>

                    <h4 class="h41">USER</h4>
                    <el-sub-menu index="2">
                        <template #title>
                            <el-icon>
                                <ShoppingBag />
                            </el-icon>
                            <span>商品管理</span>
                        </template>
                        <el-menu-item index="2-1" @click="changeActive(2.1)">分类管理</el-menu-item>
                        <el-menu-item index="2-2" @click="changeActive(2.2)">商品管理</el-menu-item>
                        <el-menu-item index="2-3">卡密管理</el-menu-item>
                    </el-sub-menu>

                    <el-sub-menu index="3">
                        <template #title>
                            <el-icon>
                                <Tickets />
                            </el-icon>
                            <span>订单管理</span>
                        </template>
                        <el-menu-item index="3-1">已完成</el-menu-item>
                        <el-menu-item index="3-2">未完成</el-menu-item>

                    </el-sub-menu>

                    <!-- <el-menu-item index="4" disabled> 禁用此组件 -->
                    <el-menu-item index="4">
                        <el-icon>
                            <Present />
                        </el-icon>
                        <span>优惠券</span>
                    </el-menu-item>
                    <h4 class="h41">CONFIG</h4>

                    <el-menu-item index="5">
                        <el-icon>
                            <Money />
                        </el-icon>
                        <span>支付配置</span>
                    </el-menu-item>
                    <el-menu-item index="6" @click="changeActive(6)">
                        <el-icon>
                            <setting />
                        </el-icon>
                        <span>网站设置</span>
                    </el-menu-item>
                </el-menu>
            </div>

        </el-col>
    </el-row>
</template>
  
<script>
import {
    Present,
    Menu as IconMenu,
    ShoppingBag,
    Setting,
    Tickets,
    Money,
} from '@element-plus/icons-vue'

const handleOpen = (key, keyPath) => {
    console.log(key, keyPath)
}

const handleClose = (key, keyPath) => {
    console.log(key, keyPath)
}

export default {
    setup() {
        return {
            handleOpen,
            handleClose,
            activeIndex: 0,
        }
    },
    components: {
        Present,
        IconMenu,
        ShoppingBag,
        Setting,
        Tickets,
        Money,

    },
    methods: {
        changeActive(index) {
            this.activeIndex = index;
            // 在这里可以添加切换选项卡时的逻辑

            // 传递给父组件
            this.$emit("Index", this.activeIndex)
        },
    },
}
</script>




<style scoped>
.menu-container {
    width: 200px;
    /* 设置宽度为100%以自适应父容器的大小 */
    height: 100%;
    /* 设置高度为100%以自适应父容器的大小 */
    min-width: 200px;
    /* 设置最小宽度 */
    min-height: 200px;
    /* 设置最小高度 */
    /* 可以根据需要进一步自定义样式 */
    background-color: rgb(42, 47, 66);

    margin-left: -2px;

    caret-color: transparent;
}

.custom-menu {
    width: 100%;
    /* 填满整个父容器的宽度 */
    background-color: rgb(42, 47, 66);
    /* 设置自定义背景颜色 */
    border-right: none;
    /* 取消右侧的白色条纹 */
}

.h3 {
    color: rgb(255, 255, 255);
    font-size: 14px;
    font-weight: bold;
    margin-bottom: 10px;
    padding-left: 10%;
    padding-bottom: 10%;
    padding-top: 5%;
}

.h4 {
    color: rgb(189, 199, 235);
    font-size: 12px;
    font-weight: normal;
    margin-top: 50px;
    margin-left: 25px;
    margin-bottom: 0px;

}

.h41 {
    color: rgb(189, 199, 235);
    font-size: 12px;
    font-weight: normal;
    margin-top: 20px;
    margin-left: 25px;
    margin-bottom: 0px;

}

/* 取消导航按钮悬停时的背景颜色改变 */
.el-menu-item:hover {
    background-color: transparent !important;
    color: #ffd04b !important;
}
</style>