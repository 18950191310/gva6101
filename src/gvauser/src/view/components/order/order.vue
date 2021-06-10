<template>
    <div><router-link to="/chome"><button>返回</button></router-link></div>
    <div><userinfo /></div>
    <div><order-Detail /></div>
    <div><comments /></div>
    <div><pay-mode /></div>
    <div class="pay"><span>实付款：3999</span><span>免运费</span><span>立即支付</span></div>

</template>

<script>
    import Comments from './components/comments'
    import OrderDetail from './components/orderDetail'
    import PayMode from './components/payMode'
    import Userinfo from './components/userinfo'

    export default {
        name: "order.vue",
        components:{
            Comments,
            OrderDetail,
            PayMode,
            Userinfo,
        }
    }
</script>

<style scoped>
   .pay span{
        margin: 10px;
   }
</style>