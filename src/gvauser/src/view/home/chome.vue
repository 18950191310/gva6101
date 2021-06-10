<template>
    <el-container class="layout-cont">
        <div>
            <div><Menu cidx="1"></Menu></div>
            欢迎来到伊瑟拉电脑自选
            <div>
                <el-button @click="reco">推荐配置</el-button>
                <el-button @click="select">我要自选</el-button>
            </div>
            <div v-if="choose"><Reco /></div>
            <div v-else><Sel /></div>
            <div class="loginout"><LoginOut/></div>
        </div>
    </el-container>
</template>

<script>

    import LoginOut from '@/view/dashboard/index.vue';
    import { ref, reactive, toRefs } from 'vue';
    import  Reco  from './reco.vue';
    import Sel from './select.vue';
    import Menu from '@/view/components/menu.vue';

    // import { getCpuList } from '@/api/cpu'

    export default {
        name: "chome",
        components: {
            Reco,
            Sel,
            Menu,
            LoginOut,
        },
        setup() {
            const state = reactive({
            })
            const choose= ref(true)
            const reco = () => choose.value = true
            const select = () => choose.value = false

            return {
                choose,
                reco,
                select,
                ...toRefs(state)
            };
        },
    };
</script>
<style lang="scss">
    /*@import "@/style/mobile.scss";*/
    .layout-cont {
        display: flex;
        height: 100%;
        overflow: auto;
    }
    .loginout {
        margin: 20px;
    }
</style>
