<template>
  <div id="userLayout">
    <div class="login_panle">
      <div class="login_panle_form">
        <div class="login_panle_form_title">
<!--          <img-->
<!--            class="login_panle_form_title_logo"-->
<!--            :src="$GIN_VUE_ADMIN.appLogo"-->
<!--            alt-->
<!--          >-->
          <p class="login_panle_form_title_p">{{ $GIN_VUE_ADMIN.appName }} 猛禽运维平台</p>
        </div>
        <el-form
          ref="loginForm"
          :model="loginFormData"
          :rules="rules"
          @keyup.enter="submitForm"
        >
          <el-form-item prop="username">
            <el-input
              v-model="loginFormData.username"
              placeholder="请输入用户名"
            >
              <template #suffix>
                <span class="input-icon">
                  <el-icon>
                    <user />
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="loginFormData.password"
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
            >
              <template #suffix>
                <span class="input-icon">
                  <el-icon>
                    <component
                      :is="lock"
                      @click="changeLock"
                    />
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="captcha">
            <div class="vPicBox">
              <el-input
                v-model="loginFormData.captcha"
                placeholder="请输入验证码"
                style="width: 60%"
              />
              <div class="vPic">
                <img
                  v-if="picPath"
                  :src="picPath"
                  alt="请输入验证码"
                  @click="loginVerify()"
                >
              </div>
            </div>
          </el-form-item>
          <el-form-item>
<!--            <el-button-->
<!--              type="primary"-->
<!--              style="width: 46%"-->
<!--              size="large"-->
<!--              @click="checkInit"-->
<!--            >前往初始化</el-button>-->
            <el-button
              type="primary"
              size="large"
              style="width: 46%"
              @click="submitForm"
            >登 录</el-button>
          </el-form-item>
        </el-form>
      </div>
      <div style="" class="qrCode" id="login_container">
      </div>
      <div class="login_panle_right" />
      <div class="login_panle_foot">

      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
}
</script>

<script setup>
import { captcha,dingLogin } from '@/api/user'
import { checkDB } from '@/api/initdb'
import bootomInfo from '@/view/layout/bottomInfo/bottomInfo.vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import {reactive, ref, onMounted, watch } from "vue"
import {useRouter, useRoute} from "vue-router"
import {store} from "@/pinia";
const userStore = useUserStore()

let goto = encodeURIComponent('https://oapi.dingtalk.com/connect/oauth2/sns_authorize?appid=dingoagfeumjr0ellu9si2&response_type=code&scope=snsapi_login&state=STATE&redirect_uri=http://localhost:8080')
let tmpAuthCode = window.location.search.split("&")[0].split("=")[1]
onMounted(() => {

  DDLogin({
    id: "login_container", //这里需要你在自己的页面定义一个HTML标签并设置id，例如<div id="login_container"></div>或<span id="login_container"></span>
    goto: goto, //请参考注释里的方式
    style: "border:none;background-color:#FFFFFF;",
    width: "270",
    height: "370"
  })

  let handleMessage = function (event) {
    let origin = event.origin
    console.log("origin", event.origin)
    if ( origin == "https://login.dingtalk.com" ) { //判断是否来自ddLogin扫码事件。
      let loginTmpCode = event.data
      //获取到loginTmpCode后就可以在这里构造跳转链接进行跳转了

      //将loginTmpCodeloginTmpCode拼接到url末尾用于跳转，不要open打开新的标签
      window.location.href = decodeURIComponent(goto + "&loginTmpCode=" + loginTmpCode)
    }
  }
  if (typeof window.addEventListener != 'undefined') {

    if (tmpAuthCode){
      //参数要用这种方式，不要直接拼接url，后端拿不到，这个坑也踩过了
      console.log(tmpAuthCode)
      userStore.ddLogin({
        code: tmpAuthCode,
        state: 'auth'
      }).then(res => {
        console.log('成功11', res)
        window.location.href = 'http://localhost:8080/' ;// 这个地址要跟redirect_uri的地址一样，不然你跳转后请求后获取的用户信息的丢在redirect_uri的地址了，你肯定拿不到了，记住这个地址要得是redirect_uri的地址，redirect_uri的地址一点得是你项目的地址，不然会有热更新刷新问题【每隔一段时间就会自动刷新页面问题】
      }, err => {
        console.log('失败11', err)
      })
    }
    window.addEventListener('message', handleMessage, false)
  } else if (typeof window.attachEvent != 'undefined') {
    window.attachEvent('onmessage', handleMessage)
  }
})


const router = useRouter()
// 验证函数
const checkUsername = (rule, value, callback) => {
  if (value.length < 5) {
    return callback(new Error('请输入正确的用户名'))
  } else {
    callback()
  }
}
const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error('请输入正确的密码'))
  } else {
    callback()
  }
}

// 获取验证码
const loginVerify = () => {
  captcha({}).then((ele) => {
    rules.captcha[1].max = ele.data.captchaLength
    rules.captcha[1].min = ele.data.captchaLength
    picPath.value = ele.data.picPath
    loginFormData.captchaId = ele.data.captchaId
  })
}
loginVerify()

// 登录相关操作
const lock = ref('lock')
const changeLock = () => {
  lock.value = lock.value === 'lock' ? 'unlock' : 'lock'
}

const loginForm = ref(null)
const picPath = ref('')
const loginFormData = reactive({
  username: 'admin',
  password: '123456',
  captcha: '',
  captchaId: '',
})
const rules = reactive({
  username: [{ validator: checkUsername, trigger: 'blur' }],
  password: [{ validator: checkPassword, trigger: 'blur' }],
  captcha: [
    { required: true, message: '请输入验证码', trigger: 'blur' },
    {
      message: '验证码格式不正确',
      trigger: 'blur',
    },
  ],
})


const login = async() => {
  return await userStore.LoginIn(loginFormData)
}
const submitForm = () => {
  loginForm.value.validate(async(v) => {
    if (v) {
      const flag = await login()
      if (!flag) {
        loginVerify()
      }
    } else {
      ElMessage({
        type: 'error',
        message: '请正确填写登录信息',
        showClose: true,
      })
      loginVerify()
      return false
    }
  })
}

// 跳转初始化
const checkInit = async() => {
  const res = await checkDB()
  if (res.code === 0) {
    if (res.data?.needInit) {
      userStore.NeedInit()
      router.push({ name: 'Init' })
    } else {
      ElMessage({
        type: 'info',
        message: '已配置数据库信息，无法初始化',
      })
    }
  }
}



</script>

<style lang="scss" scoped>
@import "@/style/newLogin.scss";
</style>
