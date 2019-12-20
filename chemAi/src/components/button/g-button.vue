<template>
  <Button 
    class="gateway-btn"
    :type="type"
    :ghost="ghost"
    :size="size"
    :shape="shape"
    :long="long"
    :htmlType="htmlType"
    :disabled="disabled"
    :unusable="unusable"
    :loading="loading"
    :icon="icon"
    :customIcon="customIcon"
    :to="to"
    :replace="replace"
    :target="target"
    :append="append"
    :style="{width: width + 'px'}"
    @click="$emit('click')"
  >
    <slot></slot>
  </Button>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class GButton extends Vue {
  @Prop(String) private type?: string;
  @Prop(Boolean) private ghost!: boolean;
  @Prop(String) private size!: string;
  @Prop(String) private shape!: string;
  @Prop(String) private long!: string;
  @Prop(String) private htmlType!: string;
  @Prop(Boolean) private disabled!: boolean;
  @Prop(Boolean) private unusable!: boolean;
  @Prop(Boolean) private loading!: boolean;
  @Prop(String) private icon!: string;
  @Prop(String) private customIcon!: string;
  @Prop([String, Object]) private to!: string | object;
  @Prop(Boolean) private replace!: boolean;
  @Prop(String) private target!: string;
  @Prop(Boolean) private append!: boolean;
  @Prop(String) private width!: string;
}
</script>
<style lang="less" scoped>
@primaryColor: #4D9DED;
@hoverColor: #2C71B6;
@disabledColor: #A8D0F8;
@unusableColor: #fff;
.gateway-btn {
  &.ivu-btn-primary {
    background-color: @primaryColor;
    border-color: @primaryColor;
    &:hover {
      background-color: @hoverColor;
      border-color: @hoverColor;
    }
    &:focus {
      box-shadow: none;
    }
    &[disabled] {
      background-color: @disabledColor;
      border-color: @disabledColor;
      color: #fff;
    }
    &[unusable]{
      color: #888888;
      background-color: @unusableColor;
      border-color: rgba(136, 136, 136, 1);;
    }
  }
  &.ivu-btn-default {
    border-color: #888888;
    color: #000;
    &:active {
      color: #fff;
      border-color: @primaryColor;
      background-color: @primaryColor;
    }
    &:focus {
      box-shadow: none;
    }
  }
}
</style>
