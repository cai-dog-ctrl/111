<template>
  <el-container style="min-width: 1400px">
    <el-header style="text-align: right;margin-top: 25px">
      <el-menu :default-active="activeIndex"  class="el-menu-demo" mode="horizontal" @select="handleSelect" style="padding-left: 100px">
        <el-menu-item index="1">提交</el-menu-item>
        <el-menu-item index="2">查看</el-menu-item>
        <el-menu-item index="3">管理</el-menu-item>
      </el-menu>
    </el-header>

    <el-main class="main">
      <router-view/>
    </el-main>
    <el-footer>

    </el-footer>


  </el-container>
</template>
<script>

import router from "@/router";

export default {
  name: 'report',
  data() {
    return {
      activeIndex : this.getNow(),
    }
  },
  mounted() {

  }
  ,
  methods: {
    back(){
      router.back();
    }
    ,
    getNow(){
      switch (this.$route.path){
        case '/report/submit': return this.activeIndex = "1";
        case '/report/view': return this.activeIndex = "2";
        case '/report/manager': return this.activeIndex = "3";
      }
    },
    handleSelect(key) {
      switch (key) {
        case '1': router.push("/report/submit"); break;
        case '2': router.push("/report/view"); break;
        case '3': router.push("/report/manager"); break;
      }
    }
  }
}
</script>
<style>
  .el-header{
    width: 100%;
    min-width: 375px;
    box-shadow:    0 0 0 0 #e5e5e5,   /*上边阴影 */
    0 0 0 0 #e5e5e5,    /*左边阴影  */
    0 0 0 0 #e5e5e5,     /*右边阴影 */
    0 1px 0 0 #e5e5e5;     /*下边阴影 */

  }
  .main{
   min-width: 375px;
  }
  .el-menu-item{
    margin-left: 500px;
    font-size: 18px;
    text-align: center;
    width: 150px;
  }
</style>