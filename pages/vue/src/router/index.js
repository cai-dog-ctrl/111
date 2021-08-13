import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)





// 解决ElementUI导航栏中的vue-router在3.0版本以上重复点菜单报错问题
const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err)
}


const report = () => import('./../view/report/index.vue')
const reportSubmit = () => import('./../view/report/submit/index.vue')
const reportView = () => import('./../view/report/view/index.vue')
const reportManager = () => import('./../view/report/manage/index.vue')

const router =  new Router({
    routes: [
        {
            //默认页面
            path: '',
            redirect: '/report'
        },
        {
            path: '/report',
            component: report,
            children: [
                {
                    path: '',
                    redirect: 'submit'
                }
                ,
                {
                    path: 'submit',
                    component: reportSubmit
                }
                ,
                {
                    path: 'view',
                    component: reportView
                }
                ,
                {
                    path: 'manager',
                    component: reportManager
                }

            ]
        }
    ]
})


export default router;