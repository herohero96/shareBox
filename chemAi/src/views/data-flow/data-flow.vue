<template>
  <div class="data-flow">
    <iframe width="100%" height="100%" src="http://10.8.2.87:1880" frameborder="0"></iframe>

  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component
export default class DataFlow extends Vue {


}
</script>
<style lang="less" scoped>
.data-flow {
  width: 100%;
  height: 100%;
  iframe {
    width: 100%;
    height: 100%;
  }
}
</style>


