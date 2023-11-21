<template>
            <!-- 上半部分 -->
            <div class="top-section">
                <!-- 背景图片 -->
                <div class="background-image">
                    <!-- 内容 -->
                    <div class="welcometexttop">
                        <br><a style="font-size: larger;font-weight: bold; margin-left: 4%;">欢迎回来</a><br><br>
                        <a style="font-size: 16px;font-weight:normal; margin-left: 4%;;">我们无法定义一个瞬间的价值，直到它成为回忆.</a>
                    </div>
                </div>
            </div>

            <!-- 下半部分 -->
            <div class="bottom-section">

                <!-- 内容 -->
                <div class="content">
                    <span class="itemCount">商品数: {{ itemCount }}</span>
                    <span class="totalSales">总交易额: {{ totalSales }}</span>
                    <span class="totalProfit">总利润: {{ totalProfit }}</span>
                </div>

            </div>
</template>
<script>
</script>
<style scoped>

.top-section {
    width: 100%;
    height: 50%;
}

.bottom-section {
    width: 100%;
    height: 50%;
    display: flex;
    flex-direction: row;
    justify-content: space-evenly;
    align-items: center;
    background-color: white;
    /* 设置背景颜色 */
    border-top: 1px solid #ccc;
    /* 上边框 */
    padding: 0px;
    /* 添加内边距以增加间距 */
    
}

.content {
    text-align: center;
}

/* 美化文本 */
.content span {
    display: block;
    font-size: 14px;
    font-weight: bold;
    margin-bottom: 8px;
    /* 增加文本之间的垂直间距 */
    color: rgb(129, 129, 129);
    /* 设置文本颜色 */
}

/* 根据需要为每个变量定义不同的样式 */
.itemCount {
    color: #007bff;
    /* 蓝色文本颜色 */
    font-size: 16px;
}

.totalSales {
    color: #28a745;
    /* 绿色文本颜色 */
    font-size: 16px;
}

.totalProfit {
    color: #dc3545;
    /* 红色文本颜色 */
    font-size: 16px;
}

.background-image {
    background-image: url('./image/welcome.jpg');
    width: 100%;
    height: 100%;
    border-radius: 5px;
}

.welcometexttop {
    color: white;
}
</style>