<template>
    <div class="user-box">
        <Menu cidx="3"></Menu>
        <s-header :name="'我的'"></s-header>
        <div class="user-info">
            <div class="info">
                <img src="//s.yezgea02.com/1604040746310/aaaddd.png"/>
                <div class="user-desc">
                    <span>昵称：红尘中的智慧</span>
                    <span class="user-detail">
                        <router-link to="/cuser-detail">个人资料</router-link></span>
                </div>
            </div>
        </div>
        <ul class="user-list">
            <li class="van-hairline--bottom" @click="goTo('/order')">
                <span>我的订单</span>
                <van-icon name="arrow" />
            </li>
            <li class="van-hairline--bottom" @click="goTo('/setting')">
                <span>账号管理</span>
                <van-icon name="arrow" />
            </li>
            <li class="van-hairline--bottom" @click="goTo('/address', { from: 'mine' })">
                <span>地址管理</span>
                <van-icon name="arrow" />
            </li>
            <li @click="goTo('/about')">
                <span>关于我们</span>
                <van-icon name="arrow" />
            </li>
        </ul>
    </div>


</template>

<script>
    import Menu from '@/view/components/menu.vue'

    export default {
        name: "cuser",
        components:{
            Menu
        }
    }
</script>

<style scoped>
        .user-detail {
            margin:10px
        }
</style>