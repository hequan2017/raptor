<template>
<!--  <div class="bottom-info">-->
<!--    <div>-->
<!--      <span>Powered by</span>-->
<!--      <span>-->
<!--        <a href="https://github.com/hequan2017/raptor">{{ $GIN_VUE_ADMIN.appName }}</a>-->
<!--      </span>-->
<!--      <el-divider direction="vertical" />-->
<!--      <span>Copyright</span>-->
<!--      <span>-->
<!--        <a href="https://github.com/hequan2017/raptor">raptor</a>-->
<!--      </span>-->
<!--    </div>-->
<!--  </div>-->
</template>

<script>
// 项目为apatch协议 请遵守版权协议内容
export default {
  name: 'BottomInfo'
}
</script>

<style lang="scss">
.bottom-info {
  color: #888;
  height: 30px;
  line-height: 12px;
  a {
    color: #888;
  }
  div {
    display: flex;
    justify-content: center;
    span{
      margin: 0 3px;
    }
  }
}
</style>
