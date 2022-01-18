<template>
  <div id="app" class="app">
    <router-view/>
  </div>
</template>

<style lang="less">
@import url('/src/assets/icons/iconfont.css');
@borderColor: #888888;
@hoverColor: #57a3f3;
@errorColor: #F4333C;

body,ul,ol,li,p,h1,h2,h3,h4,h5,h6,form,fieldset,table,td,img,div{margin:0;padding:0;border:0;}
a{text-decoration:none;}
.size{
  width: 100%;
  height: 100%;
}
html,body{
  .size;
  overflow: hidden;
  overflow-x: auto;
  margin: 0;
  padding: 0;
  min-width: 1200px;
}

.clearfix:after, .clearfix::before {
  content: "";
  visibility: hidden;
  display: block;
  height: 0;
  clear: both;
}

.clearfix {
  zoom: 1;
}

.fl {
  float: left;
}

.fr {
  float: right;
}

li {
  list-style: none;
}

#app {
  .size;
  font-family: "宋体", "Microsoft YaHei";
  box-sizing: border-box;
}

// 覆盖 iview 样式
.app {
  // input
  .ivu-input {
    border-color: @borderColor;
    &:hover {
      border-color: @hoverColor;
    }
  }
  .selected {
    .ivu-input {
      border-color: @hoverColor;
    }
  }
  .error {
    .ivu-input {
      border-color: @errorColor;
    }
  }
  // select
  .ivu-select-selection {
    border-color: @borderColor;
    &:hover {
      border-color: @hoverColor;
    }
    &.error {
      border-color: @errorColor;
    }
  }
  // from
  .ivu-form-item {
    margin-bottom: 20px;
    .ivu-form-item-label {
      padding-right: 0;
    }
  }
  //poptip
  .delete-poptip-set {
    .ivu-poptip-body {
      padding: 10px 10px 4px;
      .ivu-poptip-body-message {
        padding: 0;
        .iconfont {
          vertical-align: middle;
        }
      }
      .ivu-icon-ios-help-circle {
        display: none;
      }
      .iconfont {
        color: #FF5B05;
        padding-right: 12px;
      }
      &:after {
        background-color: transparent;
      }
    }
    .ivu-poptip-footer {
      .ivu-btn {
        &:first-child {
          border: 1px solid rgba(127, 127, 127, 1);
          background: #fff;
          &:hover {
            color: #515a6e;
          }
          &:active {
            color: #fff;
            border-color: #4D9DED;
            background-color: #4D9DED;
          }
        }
        &:last-child {
          background: #4D9DED;
          border-color: #4D9DED;
          &:hover {
            background: #2C71B6;
            border-color: #2C71B6;
          }
        }
      }
    }
  }
}
</style>
