import { createApp } from 'vue'
import App from './App.vue'
// 路由守卫
import ElementPlus from 'element-plus';
import 'element-plus/lib/theme-chalk/index.css';

// 引入封装的router
import router from '@/router/index'

import '@/permission'
import { store } from '@/store/index'


console.log(`
       欢迎使用 Gin-Vue-Admin
       当前版本:V2.3.9
       加群方式:微信号：shouzi_1994 QQ群：622360840
       默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
       默认前端文件运行地址:http://127.0.0.1:8080
       如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.gin-vue-admin.com/docs/coffee
`)
export default
createApp(App)
    .use(store)
    .use(router)
    .use(ElementPlus)
    .mount('#app')