<template>
    <div class="console">
        <p class="textP">控制台</p>

        <div class="section welcome">
            <Welcome />
        </div>

        <div class="section jinridd">
            <Jinridd />
        </div>

        <div class="section zongls"><Zongls /></div>
        <div class="section jinrils"><Jinrils /></div>
        <div class="section zuorils"><Zuorils /></div>
        <div class="section jinrilr"><Jinrilr /></div>
        <div class="section zuorilr"><Zuorilr /></div>
        <div class="section shengyilp"></div>
        <div class="section jiaoyitj"></div>
        <div class="section zuixindd"></div>
        <div class="section xitonggg"></div>

    </div>
</template>
<script>
import Welcome from './welcome.vue'
import Jinridd from './jinridd.vue'
import Zongls from './zongls.vue'
import Jinrils from './jinrils.vue'
import Zuorils from './zuorils.vue'
import Jinrilr from './jinrilr.vue'
import Zuorilr from './zuorilr.vue'

export default {
    components: {
        Welcome,
        Jinridd,
        Zongls,
        Jinrils,
        Zuorils,
        Jinrilr,
        Zuorilr,
    }
}
</script>
<style scoped>
.console {

    /* 导航栏下方内容容器的样式 */
    position: absolute;
    /* 使用绝对定位 */
    top: 70px;
    /* 调整内容容器的上边距，使其位于导航栏下方 */
    left: 220px;

}

.textP {
    color: rgb(110, 110, 110);
    font-weight: bold;
}

.section {
    position: absolute;
    background-color: rgb(255, 255, 255);
    box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
    border-radius: 5px;
}

.welcome {
    top: 55px;
    left: 5px;
    width: 450px;
    height: 220px;
}

.jinridd {

    top: 55px;
    left: 470px;
    width: 280px;
    height: 80px;
}

/* 总流水 */
.zongls {

    top: 150px;
    left: 470px;
    width: 280px;
    height: 80px;


}

.jinrils {

    top: 55px;
    left: 765px;
    width: 280px;
    height: 80px;


}

.zuorils {

    top: 150px;
    left: 765px;
    width: 280px;
    height: 80px;


}

.jinrilr {

    top: 55px;
    left: 1060px;
    width: 280px;
    height: 80px;


}

.zuorilr {

    top: 150px;
    left: 1060px;
    width: 280px;
    height: 80px;


}

.shengyilp {

    top: 290px;
    left: 5px;
    width: 450px;
    height: 220px;


}

.jiaoyitj {

    top: 245px;
    left: 470px;
    width: 870px;
    height: 280px;


}

.zuixindd {

    top: 540px;
    left: 5px;
    width: 870px;
    height: 220px;


}

.xitonggg {

    top: 540px;
    left: 890px;
    width: 450px;
    height: 220px;


}
</style>