<script setup>


function test() {
}
</script>

<template>
  <main>
    <div>
      <div style="margin: 50px">
        <n-button ghost dashed color="#ff69b4" type="error" size="large" @click="test"
                  style="width: 150px;height: 80px;font-size: 40px"> 测试
        </n-button>
      </div>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box {
  text-align: center;
}

.input {
  width: 200px;
  margin-right: 10px;
}
</style>
