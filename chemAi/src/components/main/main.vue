<template>
  <Layout class="main">
      <Sider hide-trigger>
        <side-menu :menuList='getMenuList()' :selectMenuName='this.$route.name'></side-menu>
      </Sider>
      <Layout>
          <Header>
            <header-bar></header-bar>
          </Header>
          <Content class="main-content-con" id="mainContentCon">
            <router-view/>
          </Content>
      </Layout>
  </Layout>

</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import SideMenu from './components/side-menu/side-menu.vue';
import HeaderBar from './components/header-bar/header-bar.vue';

@Component({
  components: {
    SideMenu,
    HeaderBar,
  },
})
export default class Mian extends Vue {

  public getMenuList(): any[] {
    return (this as any).$router.options.routes.filter( (val: any, i: number) => {
      return !(val.meta && val.meta.hideInMenu);
    });
  }
}

</script>

<style escoped lang="less">
.main {
  height: 100%;
  .ivu-layout-sider-children {
    background: #2A2C35;
  }
  .ivu-layout-header {
    height:40px;
    background:rgba(255,255,255,1);
    box-shadow:0px 0px 8px 0px rgba(136,136,136,0.59);
    padding: 0 30px;
  }
  .ivu-layout-content {
    background-color: #ECECEE;
    padding: 0 32px 0 30px;
  }
  .main-content-con{
    height: ~"calc(100% - 40px)";
    overflow: auto;
    padding-bottom: 20px;
  }
}
</style>
