<template>
    <div class="card">
        <div class="left-content">
            <div class="today-profit">
                <span class="label">总流水</span>
                <span class="value">{{ profitVariable | formatCurrency }}</span>
            </div>
        </div>
        <div class="right-content">
          <svg t="1695098355226" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2121" width="48" height="48"><path d="M716.8 0H307.2C137.5488 0 0 137.5488 0 307.2v409.6c0 169.6512 137.5488 307.2 307.2 307.2h409.6c169.6512 0 307.2-137.5488 307.2-307.2V307.2c0-169.6512-137.5488-307.2-307.2-307.2z" fill="#1296db" p-id="2122" data-spm-anchor-id="a313x.collections_detail.0.i10.237b3a81wBPhCG" class="selected"></path><path d="M179.2 332.8a76.8 76.8 0 0 1 76.8-76.8h478.8224a76.8 76.8 0 0 1 76.8 76.8v30.72H179.2V332.8z m0 57.6h632.4224V486.4H674.1248c-18.8416 0-34.1248 15.5904-34.1248 34.816v211.8912c0 2.4832 0.2048 4.7104 0.6144 6.7328H256a76.8 76.8 0 0 1-76.8-76.8V390.4z m247.4752 188.16h-164.992v26.88h164.992V578.56z m-164.992 53.76v26.88h164.992V632.32h-164.992z" fill="#ffffff" p-id="2123" data-spm-anchor-id="a313x.collections_detail.0.i11.237b3a81wBPhCG" class=""></path><path d="M695.4752 512c-16.512 0-29.8752 14.1824-29.8752 31.6416v192.64c0 28.2112 32.1792 42.3168 50.9952 22.3744l17.4848-18.5344a28.672 28.672 0 0 1 42.24 0l17.4848 18.5344c18.816 19.968 50.9952 5.8368 50.9952-22.3744v-192.64c0-17.4592-13.3632-31.6416-29.8752-31.6416h-119.4496z m90.8288 133.9648l30.336 30.336a8.7296 8.7296 0 1 1-12.3392 12.3392l-30.336-30.336a52.4032 52.4032 0 1 1 12.3392-12.3392z m-7.7824-30.3616a34.944 34.944 0 1 1-69.8624 0 34.944 34.944 0 0 1 69.888 0z" fill="#ffffff" p-id="2124" data-spm-anchor-id="a313x.collections_detail.0.i12.237b3a81wBPhCG" class=""></path><path d="M256 486.4a25.6 25.6 0 0 1 25.6-25.6h230.4a25.6 25.6 0 1 1 0 51.2h-230.4a25.6 25.6 0 0 1-25.6-25.6zM563.2 486.4a25.6 25.6 0 1 1 51.2 0 25.6 25.6 0 1 1-51.2 0z" fill="#FFFFFF" p-id="2125"></path></svg>        </div>
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
  