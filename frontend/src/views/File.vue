<script setup>
import {ref} from "vue";
import {CreatTask, GetFileList} from "../../wailsjs/go/apps/App";
import {AddCircleOutlineRound} from '@vicons/material'
import {useMessage} from "naive-ui";

const message = useMessage()
const SelectFileID = ref(0)
const fileList = ref([])
const columns = [
  {title: "任务号", key: "id"},
  {title: "状态", key: "status"},
  {title: "结果", key: "result"},
  // {title: "输出文件", key: "out_file_path"},
  {
    title: "破解类型", key: "attack_mode", render(row) {
      return row.attack_mode === 0 ? '字典' :
          row.attack_mode === 3 ? '掩码' :
              row.attack_mode === 6 ? '字典+掩码' :
                  '未知';
    }
  },
  {
    title: "掩码类型", key: "mask_type", render(row) {
      return row.mask_type === 0 ? '纯数字' :
          row.mask_type === 1 ? '小写字母' :
              row.mask_type === 2 ? '大写字母' :
                  row.mask_type === 3 ? '小写字母+数字' :
                      row.mask_type === 4 ? '大小写字母+数字' :
                          row.mask_type === 5 ? '大小写字母+数字+可见特殊符号' :
                              '未知';
    }
  },
  {title: "最小长度", key: "increment_min"},
  {title: "最大长度", key: "increment_max"},
  {
    title: "执行命令", key: "cmd",
    width: 200,
    ellipsis: {
      tooltip: true
    }
  },
  {
    title: 'Action',
    key: 'actions',
    // render(row) {
    //   return h(
    //       NButton,
    //       {
    //         strong: true,
    //         tertiary: true,
    //         size: 'small',
    //         // onClick: () => play(row)
    //       },
    //       { default: () => 'Play' }
    //   )
    // }
  }
]
const attackValue = ref(3)
const attackOptions = [
  {label: '字典', value: 0},
  {label: '掩码', value: 3},
  {label: '字典+掩码', value: 6},
]
const maskValue = ref(0)
const maskOptions = [
  {label: '纯数字', value: 0},
  {label: '小写字母', value: 1},
  {label: '大写字母', value: 2},
  {label: '小写字母加数字', value: 3},
  {label: '大小写字母加数字', value: 4},
  {label: '大小写字母+数字+可见特殊符号', value: 5},
]
const showModal = ref(false);
const maskLengthMin = ref(1);
const maskLengthMax = ref(12);

function openModal(fileID) {
  console.log("openModal:", fileID)
  showModal.value = !showModal.value
  SelectFileID.value = fileID
}

async function getFileList() {
  GetFileList().then(res => {
    // fileList.value = transformedData(res)
    fileList.value = res
  })
}

async function createTask() {
  var task = {
    file_id: SelectFileID.value,
    attack_mode: attackValue.value,
    mask_type: maskValue.value,
    increment_min: maskLengthMin.value,
    increment_max: maskLengthMax.value
  }
  console.log(task)
  try {
    await CreatTask(task).then(res => {
      console.log(res)
    })
  } catch (c) {
    message.error(error);
  }
  showModal.value = false
  await getFileList()
}

getFileList()
</script>


<template>
  <div>
    <span v-for="file in fileList.values()">
      <n-card :bordered="false" :title="'文件名:  '+file.file_name" style=" text-align: left;">
        <div style=" display: flex;margin-top: -2%">
            <n-collapse style="margin-top: 2%;width:100%" arrow-placement="right">
              <n-collapse-item :title="'文件类型:  '+file.file_type">
                <n-data-table
                    :columns="columns"
                    :data=file.tasks
                    :bordered="false"
                />
              </n-collapse-item>
            </n-collapse>
          <n-button text style="position: absolute; font-size: 24px;margin-top: 1%;right: 10%;">
            <n-icon @click="openModal(file.id)">
              <AddCircleOutlineRound/>
            </n-icon>
          </n-button>
        </div>
      </n-card>
    </span>

    <n-modal v-model:show="showModal">
      <n-card title="创建任务" size="medium" class="modal" style=" text-align: left;">
        <div class="modal-content">
          <p class="modal-text">攻击方式：</p>
          <n-select v-model:value="attackValue" :options="attackOptions" class="select-class"/>
          <span v-if="attackValue!==0">
          <p class="modal-text">掩码类型：</p>
           <n-select v-model:value="maskValue" :options="maskOptions" style="margin-top:20px" class="select-class"/>
            <div style=" display: flex;margin-top: 15%;margin-left: 10%">
             <n-input-number class="input-number"
                             v-model:value="maskLengthMin"
                             :min="1"
                             :max="12"
             />
              <p style="margin: 6px">---------------------</p>
             <n-input-number class="input-number"
                             v-model:value="maskLengthMax"
                             :min="1"
                             :max="12"
             />
            </div>

          </span>
          <n-button style="float: right;margin-top: 20%"
                    @click="createTask">确认
          </n-button>
        </div>

      </n-card>
    </n-modal>

  </div>
</template>


<style scoped>
.modal {
  width: 50%;
  //background-color: #34495e;
}

.modal-content {
  text-align: left;
  padding: 20px;
}

.modal-text {

}

.select-class {
  margin-left: 15%;
  width: 70%;
}

.input-number {
  margin-left: 3%;
  width: 20%
}
</style>