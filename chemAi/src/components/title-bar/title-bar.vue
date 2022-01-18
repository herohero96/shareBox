<template>
  <div class="title-bar">
    <h1>{{titleName}}</h1>
    <slot></slot>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';

@Component
export default class TitleBar extends Vue {
  @Prop()
  public readonly titleName!: string;
}
</script>

<style escoped lang="less">
.title-bar {
  h1 {
    display: inline;
    color: #4D9DED;
    font-size: 16px;
    line-height: 32px;
    font-weight: normal;
  }
  border-bottom: 1px solid #888888;
}
</style>
