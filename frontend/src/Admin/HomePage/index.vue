<!-- 此页面为 后台管理页面 整体 页面框架 -->

<template>
    <div class="background">
        <!-- 导航栏 -->
        <Nav @Index="getindex" />
        <div class="topnav">

            <NavTop  />

            <!-- 设置页面 -->


            <!-- 控制台 -->
            <Console v-if="index === 1"/>




            <!-- 分类管理 -->

            <SortMm v-else-if="index === 2.1"/>

            <!-- 商品管理 -->
            <CommodityMm v-else-if="index === 2.2" />


            <!-- 网站配置 -->
            <SystemSetting v-else-if="index === 6"/>







        </div>



    </div>
</template>
  
<script>
// 左导航栏
import Nav from "../Nav/index.vue"

// 上导航栏
import NavTop from "../NavTop/index.vue"

// 仪表盘
import Console from "../Console/index.vue"

// 网站配置
import SystemSetting from "../SystemSettings/index.vue"

// 分类管理
import SortMm from "../CommodityMm/SortMm/index.vue"

// 商品管理
import CommodityMm from "../CommodityMm/CommodityMm/index.vue"

export default {
    data() {
        return {
            index: 1,
            subpage: 1

        };
    },
    methods: {
        getindex(data) {
            this.index = data;
        }
    },
    components: {
        Nav,
        NavTop,
        Console,
        SystemSetting,
        SortMm,
        CommodityMm,
    },

};
</script>
  
<style scoped>
.background {
    background-color: rgb(241, 241, 241);
    overflow: hidden;
    /* 防止整个页面滚动 */
    display: flex;
    height: 98vh;
}

.topnav {
    flex: auto;


}
</style>
  