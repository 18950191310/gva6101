<template>
    <div><span>备注:</span><span><el-input type="textarea" @click="showRemark = true" v-model="comments1"></el-input></span>
        <el-dialog
                v-model="showRemark"
                width="30%"
                title="添加备注"
        >
            <el-input type="textarea" v-model="comments"></el-input>
            <el-button @click="confirm" type="primary">确定</el-button>
        </el-dialog>
    </div>
</template>

<script>
    import { reactive, toRefs } from 'vue'
    export default {
        name: "comments",
        setup(){
            const state = reactive({
                showRemark : false,
                comments: '',
                comments1:'',
            })
            const confirm = () => {
                state.comments1 = state.comments
                state.showRemark = false
            }
            return {
                ...toRefs(state),
                confirm,
            }
        }
        
    }
</script>

<style scoped>

</style>