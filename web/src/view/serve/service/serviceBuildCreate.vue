<template>
  <div>
    <div class="gva-form-box">
      <el-form     ref="ruleFormRef"  :model="formData" label-position="right" label-width="80px"   :rules="rules"    >
        <el-form-item label="构建项目:">
          {{ formData.name }}
        </el-form-item>
        <el-form-item label="类型:">
          {{ formData.type }}
        </el-form-item>
        <el-form-item label="库地址:">
          {{ formData.url }}
        </el-form-item>
        <el-form-item label="分支:">
          {{ formData.branch }}
        </el-form-item>
        <el-form-item label="版本号:" prop="version">
          <el-input v-model="formData.version" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="其他:" >
          <el-input v-model="formData.other" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="saveBuild">提交构建</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Build'
}
</script>

<script setup>
import {
  createBuild,
} from '@/api/build'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, defineComponent, reactive, unref } from 'vue'
import {findService} from "@/api/service";
import {register, setUserInfo} from "@/api/user";
const route = useRoute()
const router = useRouter()
const type = ref('')

const ruleFormRef = ref(null)

const formData = ref({
  name: '',
  type: '',
  url: '',
  branch: '',
  start: '',
  stop: '',
  activity: '',
  env: '',
  shell: '',
  products:[],
  static: '',
  version: '',
  info: '',
  history: '',
  other: '',
})


const rules = ref({
  version: [
    { required: true, message: '请输入版本号', trigger: 'blur' },
    { min: 5, message: '最低5位字符', trigger: 'blur' }
  ]
})


const saveBuild = async () => {
  ruleFormRef.value.validate(async valid => {
    if (valid) {
      console.log("验证通过")
      const req = {
        ...formData.value
      }
      const res = await createBuild(req)
      if (res.code === 0) {
        ElMessage({type: 'success', message: '创建成功'})
      }
    }
  })
}



// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findService({ ID: route.query.id })
    if (res.code === 0) {
      console.log(res.data.reservice)
      formData.value = res.data.reservice
    }
  }
}

init()

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
