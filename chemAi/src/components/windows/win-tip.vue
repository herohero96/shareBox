<template>
  <div class="win">
    <div class="mark">
    </div>
    <div class="win-con">
      <p>{{title}}</p>
      <p>{{content}}</p>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';

@Component
export default class WinTip extends Vue {
  @Prop()
  public readonly title!: string;
  @Prop()
  public readonly content!: string;
}
</script>

<style lang="less" scoped>

  .mark {
    position: fixed;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    background: rgba(0, 0, 0, 0.4)
  }

  .win-con {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 450px;
    height: 120px;
    text-align: center;
    border-radius: 5px;
    background-color: #fff;
    box-shadow:0px 0px 3px 0px rgba(136,136,136,0.59);
    z-index: 2;
    p:first-child {
      height: 50px;
      line-height: 50px;
      font-size: 16px;
      border-bottom: 1px solid #888888;
    }
    p:last-child {
      height: 69px;
      line-height: 50px;
    }

  }


</style>


