<template>
    <div class="card">
        <div class="left-content">
            <div class="today-profit">
                <span class="label">昨日利润</span>
                <span class="value">{{ profitVariable | formatCurrency }}</span>
            </div>
        </div>
        <div class="right-content">
          <svg t="1695098221551" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
        p-id="1512" width="48" height="48">
        <path
          d="M716.8 0H307.2C137.5488 0 0 137.5488 0 307.2v409.6c0 169.6512 137.5488 307.2 307.2 307.2h409.6c169.6512 0 307.2-137.5488 307.2-307.2V307.2c0-169.6512-137.5488-307.2-307.2-307.2z"
          fill="#1296db" p-id="1513" data-spm-anchor-id="a313x.collections_detail.0.i0.237b3a81wBPhCG" class=""></path>
        <path
          d="M204.8 358.4l438.3744-99.6352A102.1696 102.1696 0 0 1 768 358.4H204.8zM614.4 563.2a25.6 25.6 0 0 1 25.6-25.6h128a25.6 25.6 0 1 1 0 51.2h-128a25.6 25.6 0 0 1-25.6-25.6z"
          fill="#ffffff" p-id="1514" data-spm-anchor-id="a313x.collections_detail.0.i2.237b3a81wBPhCG" class="selected">
        </path>
        <path
          d="M204.8 435.2a51.2 51.2 0 0 1 51.2-51.2h486.4a51.2 51.2 0 0 1 51.2 51.2v76.8h-153.6a51.2 51.2 0 1 0 0 102.4h153.6v102.4a51.2 51.2 0 0 1-51.2 51.2H256a51.2 51.2 0 0 1-51.2-51.2V435.2z m250.2144 154.24a81.4336 81.4336 0 1 0-19.2 19.2l47.1552 47.104a13.568 13.568 0 1 0 19.2-19.1744l-47.1552-47.1296z m-12.0832-47.2064a54.272 54.272 0 1 1-108.5952 0 54.272 54.272 0 0 1 108.5952 0z"
          fill="#ffffff" p-id="1515" data-spm-anchor-id="a313x.collections_detail.0.i1.237b3a81wBPhCG" class=""></path>
        <path d="M345.1648 531.3792a12.8 12.8 0 1 1 25.6 0 12.8 12.8 0 0 1-25.6 0z" fill="#FFFFFF" p-id="1516"></path>
      </svg>        </div>
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
  