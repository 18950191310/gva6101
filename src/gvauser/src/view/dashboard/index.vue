<template>
    <el-container class="layout-cont">
        <el-button @click="loginOut" icon="el-icon-table-lamp"
        >登 出</el-button>
    </el-container>
</template>

<script>

    import { useStore } from "vuex";

    export default {
        name: "LoginOut",
        setup() {
            const store = useStore()
            const loginOut = () => store.dispatch("user/LoginOut");

            return {
                loginOut,
            };
        },
    };
</script>
<style lang="scss">
    /*@import "@/style/mobile.scss";*/
</style>
