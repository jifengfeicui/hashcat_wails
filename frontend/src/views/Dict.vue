<script setup>
import {ref} from 'vue'
import {useMessage} from "naive-ui";
import {GetDictContent} from "../../wailsjs/go/apps/App";
import hljs from 'highlight.js/lib/core'
import plaintext from 'highlight.js/lib/languages/plaintext'

hljs.registerLanguage('plaintext', plaintext)

const message = useMessage()
const dictContent = ref("")

async function getDictContent() {
  try {
    await GetDictContent().then(res => {
      dictContent.value = res
    })
  } catch (e) {
    message.error(e)
  }
}

getDictContent()
</script>

<template>
  <div>
    <n-config-provider  :hljs="hljs">
      <div class="custom-scrollbar" style="overflow: auto; height:700px;width: 500px;text-align: center">
        <n-code :code="dictContent" language="plaintext" show-line-numbers />
      </div>
    </n-config-provider>

  </div>
</template>

<style scoped>
.custom-scrollbar {
  scrollbar-width:auto;
  scrollbar-color: #3f453f transparent; /* 滚动条颜色 (前景/背景) */
}

</style>