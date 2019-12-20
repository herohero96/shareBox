<template>
<div class="error-wrap">
  <div class="error">
    <div class="img">
      <img src="@/assets/images/status_404.png" alt="">
    </div>
    <p class="text">您要找的页面不存在</p>
    <g-button type="primary" @click="goHomePage()">返回首页</g-button>
  </div>
</div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import GButton from '@/components/button/g-button.vue';

@Component({
  components: {
    GButton,
  },
})
export default class Error extends Vue {

  private goHomePage(): void {
    this.$router.replace({name: 'home'});
  }

}
</script>
<style lang="less" scoped>
.error-wrap {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  .error {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    .img {
      width: 384px;
      height: auto;
      img {
        width: 100%;
      }
    }
    .text {
      font-size: 18px;
      font-weight: 400;
      color: #0B0202;
      margin: 33px 0 70px 0;
    }
  }
}
</style>
