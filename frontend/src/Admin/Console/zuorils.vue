<template>
  <div class="card">
    <div class="left-content">
      <div class="today-profit">
        <span class="label">昨日流水</span>
        <span class="value">{{ profitVariable | formatCurrency }}</span>
      </div>
    </div>
    <div class="right-content">
      <svg t="1695098447644" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"
        p-id="2752" width="48" height="48">
        <path
          d="M716.8 0H307.2C137.5488 0 0 137.5488 0 307.2v409.6c0 169.6512 137.5488 307.2 307.2 307.2h409.6c169.6512 0 307.2-137.5488 307.2-307.2V307.2c0-169.6512-137.5488-307.2-307.2-307.2z"
          fill="#1296db" p-id="2753" data-spm-anchor-id="a313x.collections_detail.0.i27.237b3a81wBPhCG" class=""></path>
        <path
          d="M332.8 382.9504L716.8 384a71.5264 71.5264 0 0 0-25.1904-11.5456l-270.5152-63.5136c-42.0352-9.856-81.5104 23.2448-88.2944 73.984zM410.112 290.432c5.7088 0.128 11.52 0.9472 17.3056 2.4576l243.5328 63.6416a80.896 80.896 0 0 0-23.808-17.28l-147.968-70.4c-32.6912-15.5392-69.2224-5.7088-89.088 21.5552z"
          fill="#ffffff" p-id="2754" data-spm-anchor-id="a313x.collections_detail.0.i29.237b3a81wBPhCG" class="selected">
        </path>
        <path
          d="M330.9056 401.792a51.2 51.2 0 0 0-51.2 51.2V716.8a51.2 51.2 0 0 0 51.2 51.2H716.8a51.2 51.2 0 0 0 51.2-51.2V452.992a51.2 51.2 0 0 0-51.2-51.2H330.9056z m119.7056 48.8192h39.0656l39.0656 97.664 39.0656-97.664h39.0656l-39.0656 97.664h19.5328v19.5328h-39.0656v19.5328h39.0656v19.5328h-39.0656v39.0656H509.184v-39.0656h-39.0656v-19.5328h39.0656v-19.5328h-39.0656v-19.5328h19.5328l-39.0656-97.664z"
          fill="#ffffff" p-id="2755" data-spm-anchor-id="a313x.collections_detail.0.i28.237b3a81wBPhCG" class="selected">
        </path>
        <path
          d="M358.4 716.8m12.8 0l256 0q12.8 0 12.8 12.8l0 0q0 12.8-12.8 12.8l-256 0q-12.8 0-12.8-12.8l0 0q0-12.8 12.8-12.8Z"
          fill="#1296db" p-id="2756" data-spm-anchor-id="a313x.collections_detail.0.i25.237b3a81wBPhCG" class=""></path>
        <path
          d="M665.6 716.8m12.8 0l0 0q12.8 0 12.8 12.8l0 0q0 12.8-12.8 12.8l0 0q-12.8 0-12.8-12.8l0 0q0-12.8 12.8-12.8Z"
          fill="#1296db" p-id="2757" data-spm-anchor-id="a313x.collections_detail.0.i26.237b3a81wBPhCG" class=""></path>
      </svg>
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

.label {
  margin-bottom: 10px;

}</style>
  