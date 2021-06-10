<template>
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
                    label="数量/套"
                    width="180">
            </el-table-column>
        </el-table>
          <div>总计:{{total}}元</div>
    </div>
</template>

<script>
    import { useStore } from 'vuex'
    import { reactive, toRefs, ref } from 'vue'
    export default {
        name: "orderDetail",
        setup(){
            const store = useStore()
            const state = reactive({
                table:[{cname:"配置一",count:0},
                    {cname:"配置二",count:0},
                    {cname:"配置三",count:0},
                    {cname:"配置四",count:0}],
            })
            const info = store.getters['user/orderDetail']
            const total = store.getters['user/total']
            for(let i =0;i<4;i++){
                    state.table[i].count = info[i]
            }
            
            return {
                ...toRefs(state),
                total,
            }
        }
    }
</script>

<style scoped>

</style>