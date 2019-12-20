<template>
  <div class="side-menu">
    <p class="side-menu-title">格创东智网关</p>
    <ul>
      <router-link
        tag="li"
        :class="{'cur': selectMenuName === item.children[0].name }"
        :to='{name: item.children[0].name}'
        v-for="(item, index) in menuList"
        :key="index"
      >
        <i :class="['iconfont', item.children[0].meta.icon]"></i>
        <span>{{item.children[0].meta.title}}</span>
      </router-link>
    </ul> 
  </div>
</template> 

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';

@Component
export default class SideMenu extends Vue {
  @Prop()
  public readonly menuList!: object [];
  @Prop()
  public readonly selectMenuName!: string;
}
</script>

<style escoped lang="less">
.side-menu {
  color: #fff;
  .side-menu-title {
    padding: 10px 0 40px 0;
    text-align: center;
    font-size:20px;
    line-height: 20px;
    font-family:PingFangSC-Light;
    color: #fff;
  }
  ul {
    li {
      height: 40px;
      padding-left: 32px;
      line-height: 40px;
      font-size: 14px;
      cursor: pointer;
      opacity:0.5;
      &:hover {
        background-color:#181A23;
        color:#fff;
        opacity: 0.4;
      }
      i {
        font-size: 15px;
        padding-right: 15px;
      }
      &.cur {
        background-color:#181A23;
        color: #4D9DED;
        opacity: 1;
      }
      span {
        float: right;
        width: 130px;

      }
    }
  }
}

</style>


