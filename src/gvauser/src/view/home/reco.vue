<template>
    <div class="container">
    <div class="form">
    <el-form>
        <el-form-item >
            <div class="card">
            <el-card class="box-card">
                    <div>
                        <div>{{name[0]}}</div>
                        <div>CPU: {{cpu[0]}}</div>
                        <div>显卡：{{gpu[0]}}</div>
                        <div>内存：{{storage[0]}}</div>
                        <div>硬盘：{{hard[0]}}</div>
                        <div>显示器：{{monitor[0]}}</div>
                        <div>总价：{{price[0]}}</div>
                    </div>
                <el-input-number v-model="num0" :min="0" @change="buy0"></el-input-number>
            </el-card>
                <el-card class="box-card">
                    <div>
                        <div>{{name[1]}}</div>
                        <div>CPU: {{cpu[1]}}</div>
                        <div>显卡：{{gpu[1]}}</div>
                        <div>内存：{{storage[1]}}</div>
                        <div>硬盘：{{hard[1]}}</div>
                        <div>显示器：{{monitor[1]}}</div>
                        <div>总价：{{price[1]}}</div>
                    </div>
                    <el-input-number v-model="num1" :min="0" @change="buy1"></el-input-number>
                </el-card>
                <el-card class="box-card">
                    <div>
                        <div>{{name[2]}}</div>
                        <div>CPU: {{cpu[2]}}</div>
                        <div>显卡：{{gpu[2]}}</div>
                        <div>内存：{{storage[2]}}</div>
                        <div>硬盘：{{hard[2]}}</div>
                        <div>显示器：{{monitor[2]}}</div>
                        <div>总价：{{price[2]}}</div>
                    </div>
                    <el-input-number v-model="num2" :min="0" @change="buy2"></el-input-number>
                </el-card>
                <el-card class="box-card">
                    <div>
                        <div>{{name[3]}}</div>
                        <div>CPU: {{cpu[3]}}</div>
                        <div>显卡：{{gpu[3]}}</div>
                        <div>内存：{{storage[3]}}</div>
                        <div>硬盘：{{hard[3]}}</div>
                        <div>显示器：{{monitor[3]}}</div>
                        <div>总价：{{price[3]}}</div>
                    </div>
                    <el-input-number v-model="num3" :min="0" @change="buy3"></el-input-number>
                </el-card>
            </div>
        </el-form-item>
    </el-form>
        <div class="choose">
            <el-table
                    :data="table"
                    height="250"
                    border
                    style="width: 100%">
                <el-table-column
                        prop="cname"
                        label="已选择"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="count"
                        label="数量"
                        width="180">
                </el-table-column>
            </el-table>
        </div>
        <div>总计:{{cprice}}}</div>
        <div>
            <button @click="submit">提交订单</button>
        </div>
    </div>

    </div>
</template>

<script>
    import { reactive, toRefs, getCurrentInstance } from 'vue'
    import { getRecolistList } from '@/api/recolist'
    import { useRouter } from 'vue-router'
    import { useStore } from 'vuex'
    export default {
        name: "Reco",
        setup () {
            const { ctx } = getCurrentInstance()
            const router = useRouter()
            const store = useStore()
            const state = reactive({
                num0:0,
                num1:0,
                num2:0,
                num3:0,
                name:[],
                cpu:[],
                gpu:[],
                storage:[],
                hard:[],
                monitor:[],
                price:[],
                orderCount:[],
                table:[{
                    cname:'配置一',
                    count: 0,
                },{
                    cname:'配置二',
                    count: 0,
                },{
                    cname:'配置三',
                    count: 0,
                },{
                    cname:'配置四',
                    count: 0,
                }],
                cprice:0,
            })
            const tableInfo = reactive({
                page: 1,
                total: 10,
                pageSize: 10,
                tableData: [],
                searchInfo: {}
            })
            const getTableData = async() => {
                const res = await getRecolistList({page: tableInfo.page, pageSize: tableInfo.pageSize, ...tableInfo.searchInfo })
                if (res.code == 0) {
                    tableInfo.tableData = res.data.list
                    tableInfo.total = res.data.total
                    tableInfo.page = res.data.page
                    tableInfo.pageSize = res.data.pageSize
                }
            }
            const submit = () => {
                if(state.cprice == 0){
                    ctx.$message({
                        type: "error",
                        message: "订单不能为空",
                        showClose: true,
                    })
                }else {
                    for (let i = 0;i<4;i++){
                        state.orderCount.push(state.table[i].count)
                    }
                    store.commit('user/setOrderDetail', state.orderCount)
                    store.commit('user/setTotal', state.cprice)
                    router.push({name:'order'})
                }
            }

            return {
                ...toRefs(tableInfo),
                getTableData,
            ...toRefs(state),
                submit,
            }

        },
        data(){
            return {

            }

        },
            async created(){
                await this.getTableData();
                const data = this.tableData
                for(let i = 0;i<4;i++){
                    this.name.push(data[i].name)
                    this.cpu.push(data[i].cpu)
                    this.gpu.push(data[i].gpu)
                    this.storage.push(data[i].storage)
                    this.hard.push(data[i].hard)
                    this.monitor.push(data[i].monitor)
                    this.price.push(data[i].total)
                }
            },
        methods:{
            buy0(){
                this.table[0].count=this.num0
                this.getCprice()

            },
            buy1(){
                this.table[1].count=this.num1
                this.getCprice()
            } ,
            buy2(){
                this.table[2].count=this.num2
                this.getCprice()
            }  ,
            buy3(){
                this.table[3].count=this.num3
                this.getCprice()
            }   ,
            getCprice(){
                this.cprice=this.price[0]*this.num0+this.price[1]*this.num1+this.price[2]*this.num2+this.price[3]*this.num3

            }
        }

    }
</script>

<style scoped>
    .container {
        display: flex;
    }
    .form {
        height: 50%;
    }
    .card {
        display: inline-flex;
    }
    .box-card {
        display: flex;
        position: relative;
        width: 40%;
        margin: 10px;
    }
    .choose {
        display: flex;
        position: relative;
    }
    .choose span {
        margin: 10px;
    }

</style>