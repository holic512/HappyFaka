<template>
    <div class="card">
        <div class="left-content">
            <div class="today-profit">
                <span class="label">今日订单</span>
                <span class="value">{{ profitVariable | formatCurrency }}</span>
            </div>
        </div>
        <div class="right-content">
            <svg t="1695048343888" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6772" width="48" height="48"><path d="M716.8 0H307.2C137.5488 0 0 137.5488 0 307.2v409.6c0 169.6512 137.5488 307.2 307.2 307.2h409.6c169.6512 0 307.2-137.5488 307.2-307.2V307.2c0-169.6512-137.5488-307.2-307.2-307.2z" fill="#1296db" p-id="6773" data-spm-anchor-id="a313x.collections_detail.0.i28.692c3a81dznddB" class=""></path><path d="M332.8 716.8a51.2 51.2 0 0 0 51.2 51.2h307.2a51.2 51.2 0 0 0 51.2-51.2V358.4a51.2 51.2 0 0 0-51.2-51.2v358.4a51.2 51.2 0 0 1-51.2 51.2H332.8z" fill="#ffffff" p-id="6774" data-spm-anchor-id="a313x.collections_detail.0.i29.692c3a81dznddB" class="selected"></path><path d="M614.4 256a51.2 51.2 0 0 1 51.2 51.2v51.2h-51.2a51.2 51.2 0 0 1-51.2-51.2v-51.2h51.2z m0 128h51.2v256a51.2 51.2 0 0 1-51.2 51.2H332.8a51.2 51.2 0 0 1-51.2-51.2V307.2a51.2 51.2 0 0 1 51.2-51.2h204.8v51.2a76.8 76.8 0 0 0 76.8 76.8z m-268.8-51.2a12.8 12.8 0 0 0 0 25.6h102.4a12.8 12.8 0 0 0 0-25.6h-102.4z m-12.8 89.6a12.8 12.8 0 0 0 12.8 12.8h102.4a12.8 12.8 0 0 0 0-25.6h-102.4a12.8 12.8 0 0 0-12.8 12.8z m12.8 64a12.8 12.8 0 0 0 0 25.6h256a12.8 12.8 0 0 0 0-25.6h-256z" fill="#ffffff" p-id="6775" data-spm-anchor-id="a313x.collections_detail.0.i30.692c3a81dznddB" class="selected"></path><path d="M307.2 614.4a25.6 25.6 0 0 1 25.6-25.6h204.8a25.6 25.6 0 1 1 0 51.2h-204.8a25.6 25.6 0 0 1-25.6-25.6zM588.8 614.4a25.6 25.6 0 1 1 51.2 0 25.6 25.6 0 1 1-51.2 0z" fill="#FFFFFF" p-id="6776"></path></svg>
        </div>
    </div>
</template>
  
<script>
export default {

};
</script>
  
<style>
.card {
    height: 75%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-radius: 5px;
    padding: 10px;
}
.left-content {
  display: flex;
  flex-direction: column;
}

.today-profit {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  margin-left: 5px;
}

.label {
  font-size: 14px;
  color: #666;
}

.value {
  font-size: 18px;
  font-weight: bold;
}

.right-content {
  display: flex;
  align-items: center;
  justify-content: flex-end;
}

svg {
  fill: #333;
}
.label{
    margin-bottom: 10px;
  
}
</style>
  