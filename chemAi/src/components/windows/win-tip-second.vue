<template>
  <div class="win">
    <div class="win-content">
      <p>{{content}}</p>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';

@Component
export default class WinTipSecond extends Vue {
  @Prop()
  public readonly content!: string;
}
</script>

<style lang="less" scoped>

  .win-content {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);     
    width:450px;
    height:120px;
    line-height: 120px;
    color: #fff;
    text-align: center;
    font-size: 16px;
    background:rgba(0,0,0,1);
    box-shadow:0px 0px 3px 0px rgba(136,136,136,0.59);
    opacity:0.4;
    border-radius:5px;       
  }



</style>


