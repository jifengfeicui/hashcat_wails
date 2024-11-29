<script setup>
import {ref} from "vue";
import {ChooseFile, ConfirmAddFile} from "../../wailsjs/go/apps/App";
import {useRouter} from 'vue-router';
import {NNotificationProvider, useNotification,useMessage} from 'naive-ui'
const message = useMessage()
const notification = useNotification()
const router = useRouter();
var value = ref("cap")
var fileName = ref("")
const options = [
  {
    label: 'wifi握手包',
    value: 'cap'
  },
  {
    label: 'windows NTLM哈希',
    value: 'ntlm'
  },
]

function chooseFile() {
  ChooseFile().then(res => {
    console.log(res)
    fileName.value = res
  })
}

async function confirm() {
  try {
    await ConfirmAddFile(value.value)
    router.push({name: 'File'});
  } catch (error) {
    message.error(error);
  }
  fileName.value = ""
}

</script>

<template>
    <div class="left_class">
      <p>类型:</p>
      <n-select v-model:value="value" :options="options" style="width: 400px"/>
    </div>
    <div class="left_class">
      <p>文件:</p>
      <n-button strong secondary circle type="info" style="width: 150px" @click="chooseFile">
        <p>选择文件</p>
      </n-button>
    </div>
    <div class="left_class" style="margin-left: 150px">
      <p style="color: #aeb1b3">
        {{ fileName }}
      </p>
    </div>
    <div style="text-align: right;margin: 80px">
      <n-button strong secondary circle type="success" style="width:150px" @click="confirm">
        <p>确认</p>
      </n-button>
    </div>
</template>

<style scoped>
.left_class {
  text-align: left;
  margin-left: 50px
}
</style>