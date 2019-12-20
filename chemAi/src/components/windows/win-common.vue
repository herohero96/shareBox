<template>
  <div class="win">
    <div class="mark" v-if="isShowMark">
    </div>
    <div class="win-con" :style="{'width': width+'px'}">
      <i v-if="isShowCloseIcon" @click="$emit('closeWin')" class="icon iconfont icon-close"></i>
      <h1 v-if="title">{{title}}</h1>
      <slot></slot>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';

@Component
export default class WinCommon extends Vue {
  @Prop()
  public readonly title!: string;
  @Prop({ default: false })
  public readonly isShowMark!: boolean;
  @Prop({ default: false })
  public readonly isShowCloseIcon!: boolean;
  @Prop({ default: 450 })
  public readonly width!: number;
}
</script>

<style lang="less" scoped>
  .mark {
    position: fixed;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    background: rgba(0, 0, 0, 0.4);
    z-index: 20;
  }
  .win-con {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    text-align: center;
    border-radius: 5px;
    background-color: #fff;
    box-shadow:0px 0px 3px 0px rgba(136,136,136,0.59);
    z-index: 20;
    h1 {
      display: block;
      height: 50px;
      line-height: 50px;
      font-size: 16px;
      border-bottom: 1px solid #888888;
    }
    .icon-close {
      position: absolute;
      font-size: 20px;
      right: -28px;
      top: -5px;
      cursor: pointer;
    }
  }
</style>


