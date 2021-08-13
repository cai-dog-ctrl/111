<template >
  <div class="out" >
    <div class="view_left">

      <el-card shadow="hover" v-show="noSubmit.length !== 0">
        <div slot="header">
          <span>未交</span>
        </div>
        <div v-for="(item,index) in noSubmit" :key="index">
          {{ item }}
        </div>
      </el-card>

      <el-card shadow="hover" v-show="haveSubmit.length !== 0">
        <div slot="header">
          <span>已交</span>
        </div>
        <div v-for="(item,index) in haveSubmit" :key="index">
          {{ item }}
        </div>
      </el-card>
    </div>
    <div class="view_main">
      <el-form class="header" :inline="true">
        <el-form-item label="日期 : "  style="margin-left: -100px;width: 400px">
          <el-date-picker
              v-model="searchDate"
              type="date"
              size="medium"
              placeholder="请选择日期"
              style="left: 0">
          </el-date-picker>
        </el-form-item>
        <el-form-item label="姓名 : " style="margin-left: 0px;width: 350px">
          <el-input
              size="medium"
              placeholder="请输入姓名"
              prefix-icon="el-icon-search"
              v-model="searchName" >
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button icon="el-icon-search" round style="margin-left: 60px;width: 200px" @click="search(1)">查找</el-button>
        </el-form-item>
      </el-form>

      <div class="main">
        <el-table
            tooltip-effect="light"
            :row-style="{height:'47.5px'}"
            :data="data"
            v-loading="loading"
            element-loading-text="拼命加载中"
            stripe
            style="min-width: 1190px;margin: 0 auto"
            @row-click="rowClick"
            highlight-current-row>



          <el-table-column
              prop="date"
              label="日期"
              min-width="170">
          </el-table-column>

          <el-table-column
              prop="name"
              label="姓名"
              min-width="130"
              show-overflow-tooltip>
          </el-table-column>

          <el-table-column
              prop="complete"
              label="完成"
              min-width="200"
              show-overflow-tooltip>
          </el-table-column>

          <el-table-column
              prop="algorithm"
              label="算法"
              min-width="180"
              show-overflow-tooltip>
          </el-table-column>

          <el-table-column
              prop="plan"
              label="计划"
              min-width="200"
              show-overflow-tooltip>
          </el-table-column>

          <el-table-column
              prop="summary"
              label="总结"
              min-width="200"
              show-overflow-tooltip>
          </el-table-column>
        </el-table>
      </div>
      <div class="footer">
        <el-pagination
            hide-on-single-page
            background
            layout="prev, pager, next"
            :current-page="page"
            @current-change="search"
            :total="pageTotal"
        style="margin-top: 20px">
        </el-pagination>
      </div>
    </div>

    <el-drawer
        :with-header="false"
      :visible.sync="drawer"
      :direction="direction"
    size="470px">
          <el-form label-position="left"  class="demo-table-expand" style="margin-top: 50px">
            <el-form-item label="姓名 :">
              <span>{{ drawerRow.name }}</span>
            </el-form-item>
            <el-form-item label="日期 :">
              <span>{{ drawerRow.date }}</span>
            </el-form-item>
            <el-form-item label="完成 :">
              <span>{{ drawerRow.complete }}</span>
            </el-form-item>
            <el-form-item label="算法 :" >
              <span>{{ drawerRow.algorithm }}</span>
            </el-form-item>
            <el-form-item label="计划 :">
              <span>{{ drawerRow.plan }}</span>
            </el-form-item>
            <el-form-item label="总结 :">
              <span>{{ drawerRow.summary }}</span>
            </el-form-item>
          </el-form>
    </el-drawer>


  </div>
</template>
<script>
import axios from 'axios'
import qs from 'qs'
export default {
  name: 'reportView',
  data() {
    return {
      drawerRow: '',
      direction: 'rtl',
      drawer: false,
      haveSubmit: '',
      noSubmit: '',
      page: 0,
      pageSize: 10,
      pageTotal: 0,
      loading: true,
      searchName: '',
      searchDate: '',
      data: [{
        id: '',
        name: '',
        plan: '',
        complete: '',
        summary: '',
        date: '',
        algorithm:''
      }]
    }
  },
  created() {
    while (this.data.length < 10){
      this.data.push({id: '', name: '', plan: '', complete: '', summary: '', date: '',algorithm: ''})
    }
    this.search(1)
    this.loading = false;

    axios({
      method: 'post',
      url: '/report/todaySubmit',
    }).then(res => {
      if(res.data.status === '200') {
        this.noSubmit = res.data.data["notSubmitted"];
        this.haveSubmit = res.data.data["submitted"];
      }
    })


  },
  methods: {
    rowClick(row,){
      this.drawerRow = row
      this.drawer = true
    },
    handleClose(done) {
      this.$confirm('确认关闭？')
          .then(() => {
            done();
          })
          .catch(() => {});
    },
    search(page) {
      this.loading = true;
      let name = this.searchName;
      let date = this.searchDate;
      let rows = this.pageSize;
      this.page = page;
      axios({
        method: 'post',
        url: '/report/list',
        data: qs.stringify({
          name,
          date,
          page,
          rows,
        })
      }).then(res => {
        if (res.data.status === '200') {
          this.pageTotal = res.data.data.total
          this.data = res.data.data.rows
          while (this.data.length < 10){
            this.data.push({id: '', name: '', plan: '', complete: '', summary: '', date: '',algorithm: ''})
          }

          this.loading = false;
        }
      }).catch(()=>{
        this.$message({
          type : 'error',
          message :"加载失败"
        })
        this.loading = false;
      })
    }
  }
}
</script>
<style>

.header{
}
.el-tooltip__popper {
  display: none;
  max-width: 10%;
  background: #000 !important;/*背景色  !important优先级*/
  opacity: 0.4 !important;/*背景色透明度*/
  color: #FFFFFF !important;/*字体颜色*/

}
.view_left{
  margin-right: 10px;
  margin-top: 62px;
  width: 150px;
  float: left;
}
.view_main{
  float: left;
  width: 1100px;
  margin: 0 auto;
  text-align: center;
}
.out{
  width: 1300px;
  height: auto;
  margin: 0 auto;
}
.demo-table-expand {
  font-size: 0;
}
.demo-table-expand label {
  width: 70px;
  color: #99a9bf;
}
.demo-table-expand .el-form-item {
  margin-left: 30px;
  width: 400px;
}

</style>