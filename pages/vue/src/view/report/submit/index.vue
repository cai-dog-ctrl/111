<template>
  <div class="submit_main" style="width: 1200px;margin: 0 auto">
    <el-card>
      <div slot="header" class="clearfix" >
        <span style="font-weight: normal;font-size: 35px;">日报</span>
        <span v-text="autoSaveTime" style="float: right;color: silver"></span>
      </div>

      <el-form style="margin: 0 auto;text-align: center;"  ref="form" :model="form" :rules="formRules" label-width="60px" v-loading="loading">

        <el-form-item prop="name" label="姓名 : ">
          <el-input placeholder="姓名" v-model="form.name" maxlength="10" show-word-limit> </el-input>
        </el-form-item>

        <el-form-item prop="complete" label="完成 : ">
          <el-input placeholder="今日完成" type="textarea" v-model="form.complete" maxlength="500" show-word-limit :autosize="{ minRows: 5, maxRows: 18}"> </el-input>
        </el-form-item>

        <el-form-item prop="plan" label="计划 : ">
          <el-input placeholder="明日计划" type="textarea" v-model="form.plan" maxlength="500" show-word-limit :autosize="{ minRows: 5, maxRows: 18}" > </el-input>
        </el-form-item>


        <el-form-item prop="summary" label="总结 : ">
          <el-input placeholder="今日总结"  type="textarea" v-model="form.summary" maxlength="500" show-word-limit :autosize="{ minRows: 5, maxRows: 18}"> </el-input>
        </el-form-item>


        <el-form-item prop="algorithm" label="算法 : ">
          <el-input placeholder="题解地址"  type="textarea" v-model="form.algorithm" maxlength="500" show-word-limit :autosize="{ minRows: 5, maxRows: 18}"> </el-input>
        </el-form-item>


          <el-button type="primary" style="margin: 0 auto;text-align: center;width: 40%" @click="submit('form')">提交</el-button>
      </el-form>
      <el-link type="success" style="float:right;margin-bottom: 20px;" @click="saveCache">保存</el-link>

    </el-card>
  </div>
</template>
<script>
import axios from 'axios'
export default {

  name: 'reportSubmit',


  beforeDestroy() {
    this.saveCache('auto')
  },
  created() {
    this.getCache()
    window.setInterval(() => {
      setTimeout(() => {
        this.saveCache('auto')
      }, 0)
    }, 3000)

  },
  data() {
    return {
      loading: true,
      autoSaveTime: '',
      lastForm: '',
      form: {
        name: '',
        plan: '',
        complete: '',
        summary: '',
        algorithm: '',
      },
      formRules: {
        name: [
          {required: true, message: '请输入姓名', trigger: 'blur'},
          {min: 2, max: 10, message: '姓名长度2-10位', trigger: 'blur'}
        ],
        plan: [
          {required: true, message: '请输入计划', trigger: 'blur'},
          {min: 20, max: 500, message: '计划长度为20-500', trigger: 'blur'}
        ],
        summary: [
          {required: true, message: '请输入总结', trigger: 'blur'},
          {min: 20, max: 500, message: '今日总结长度为20-500', trigger: 'blur'}
        ],
        complete: [
          {required: true, message: '请输入完成', trigger: 'blur'},
          {min: 20, max: 500, message: '今日完成字数为20-500', trigger: 'blur'}
        ],
        algorithm: [
          {required: true, message: '请输入题解地址', trigger: 'blur'},
          {min: 1, max: 500, message: '题解地址长度为10-500', trigger: 'blur'}
        ]
      },
    }
  },
  methods: {
    getCache() {
      axios({
        method: 'post',
        url: '/report/getCache',
      }).then((res) => {
        if (res.data.status === '200') {
          //获取状态
          this.form = JSON.parse(JSON.stringify(res.data.data))
          this.lastForm = JSON.parse(JSON.stringify(res.data.data))
          this.loading = false

        } else {
          this.$message({
            message: '读取保存的信息失败!',
            type: 'error'
          });
        }
      }).catch(() => {
        this.$message({
          message: '读取保存的信息失败!',
          type: 'error'
        });
      })
    },
    saveCache(type) {
      if (type === "auto") {
        if (JSON.stringify(this.form) === JSON.stringify(this.lastForm) || this.loading) {
          return;
        }
        let {name, plan, summary, complete, algorithm} = this.form

        axios({
          method: 'post',
          url: '/report/saveCache',
          data: {
            name,
            plan,
            summary,
            complete,
            algorithm
          }
        }).then((res) => {
          if (res.data.status === '200') {
            this.autoSaveTime = "自动保存 : " + this.now();

            this.lastForm = JSON.parse(JSON.stringify(this.form))


          } else {
            this.autoSaveTime = "自动保存 : 失败";
          }
        }).catch(() => {
          this.autoSaveTime = "自动保存 : 失败";
        })

      } else {
        let {name, plan, summary, complete, algorithm} = this.form
        axios({
          method: 'post',
          url: '/report/saveCache',
          data: {
            name,
            plan,
            summary,
            complete,
            algorithm
          }
        }).then((res) => {
          if (res.data.status === '200') {
            this.$message({
              message: '保存成功!',
              type: 'success'
            });
            this.lastForm = JSON.parse(JSON.stringify(this.form))
          } else {
            this.$message({
              message: '保存失败!',
              type: 'error'
            });
          }
        }).catch(() => {
          this.$message({
            message: '保存失败!',
            type: 'error'
          });
        })

      }

    },
    now() {
      let hh = new Date().getHours();
      let mf = new Date().getMinutes() < 10 ? '0' + new Date().getMinutes() : new Date().getMinutes();
      let ss = new Date().getSeconds() < 10 ? '0' + new Date().getSeconds() : new Date().getSeconds();
      return hh + ':' + mf + ':' + ss;
    },

    submit(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          let {name, plan, complete, summary, algorithm} = this.form
          axios({
            method: 'post',
            url: '/report/insert',
            data: ({
              name,
              plan,
              complete,
              summary,
              algorithm
            })
          }).then(res => {

            if (res.data.status === '200') {
              this.$message({
                type: 'success',
                message: `提交成功!`
              });
            } else {
              this.$message({
                type: 'error',
                message: res.data.msg
              });
            }

          }).catch(() => {
            this.$message({
              type: 'error',
              message: "提交失败"
            });
          })
        } else {
          return false;
        }
      })
    }
  }
}
</script>
<style>

</style>
