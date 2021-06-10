<template>
    <div class="userinfo">
        <span>{{username}}</span>
        <span>{{phone}}</span>
        <span>{{region}}</span>
        <span>{{address}}</span>
    </div>
</template>

<script>
    import { useStore } from 'vuex'
    import { computed, onMounted, getCurrentInstance, reactive, toRefs } from 'vue'
    import { getCuserDetailList } from '@/api/cuserDetail'
    import { useRouter } from 'vue-router'
    export default {
        name: "userinfo",
        setup(){
            const router = useRouter()
            const { ctx } = getCurrentInstance()
            const store = useStore()
            const data = computed(() => store.getters["user/userInfo"]).value
            const username = data.userName
            const state = reactive({
                phone:"",
                region:"",
                address:"",
            })
            onMounted(async () => {
                const res = await getCuserDetailList({Username: username})
                const data = res.data.list[0]
                state.phone = data.phone
                state.region = data.region
                state.address = data.address
                if (data.phone == '' || data.region == "" || data.address == "") {
                    ctx.$message({
                        type: "error",
                        message: "请先完善个人信息,手机号、地区和地址为必填项",
                        showClose: true,
                    }),
                    setTimeout(() => router.push({name:"cuser-detail"}), 2000)
                }else {
                }
            })
            return {
                username,
                ...toRefs(state)
            }
        }
    }
</script>

<style scoped>
  .userinfo span{
      margin: 10px;
  }
</style>