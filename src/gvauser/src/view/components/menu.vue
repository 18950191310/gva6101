<template>
    <el-menu  :default-active="cidx"
              class="el-menu-demo"
              mode="horizontal"
              background-color="#545c64"
              text-color="#fff"
              active-text-color="#ffd04b"
              router="true"
    >
        <el-menu-item index="1" route="/chome">
            >首页
        </el-menu-item>
        <el-menu-item index="2" route="/cart">
            购物车</el-menu-item>
        <el-menu-item index="3" route="/cuser">
            个人中心</el-menu-item>
    </el-menu>
</template>

<script>
    import { toRefs, reactive } from "vue";


    export default {
        name: "Menu",
        props:{
            cidx: Number,
        },
        setup(){
            const state = reactive({

                })
            return{
                ...toRefs(state),

            }
        },
    }
</script>

<style scoped>

</style>