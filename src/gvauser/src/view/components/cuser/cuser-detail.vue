<template>
    <div>
        <router-link to="/cuser"><button>返回个人中心</button></router-link></div>
    <div>用户详情</div>
    <div>头像：<img class="img" src="@/assets/logo.jpg"></div>
    <div><el-button type="text" @click="dialogVisible = true">昵称:{{nickName}}</el-button>
        <el-button type="text" @click="dialogVisible = true">性别:{{sex}}</el-button>
        <el-button type="text" @click="dialogVisible = true">称呼:{{name}}</el-button>
        <el-button type="text" @click="dialogVisible = true">手机号:{{phone}}</el-button>
        <el-button type="text" @click="dialogVisible = true">微信号:{{wechat}}</el-button>
        <el-button type="text" @click="dialogVisible = true">地区:{{region}}</el-button>
        <el-button type="text" @click="dialogVisible = true">地址:{{address}}</el-button>
        <div>用户名:{{userName}}</div>
        <el-button type="text" @click="dialogVisible = true">登录密码:{{password}}</el-button>
        <el-dialog
                v-model="dialogVisible"
                width="30%"
                title="修改个人信息"
        >
            <el-form :model="form" :rules="rules" ref="form1"> <el-form-item label="昵称" :label-width="formLabelWidth">
                <el-input v-model="nickName"></el-input>
            </el-form-item>
                <el-form-item label="性别" :label-width="formLabelWidth">
                    <el-select v-model="sex" placeholder="请选择">
                        <el-option v-for="item in options" :key="item.value"  :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="称呼(示例：张先生/女士 或者姓名)" :label-width="formLabelWidth">
                    <el-input v-model="name"></el-input>
                </el-form-item>
                <el-form-item label="手机号(请认真核对，以免影响配送)" :label-width="formLabelWidth">
                    <el-input v-model="phone"></el-input>
                </el-form-item>
                <el-form-item label="微信号" :label-width="formLabelWidth">
                    <el-input v-model="wechat"></el-input>
                </el-form-item>
                <el-form-item label="地区(格式：XX省XX市XX县/区 示例：福建省厦门市思明区)" :label-width="formLabelWidth">
                    <el-input v-model="region"></el-input>
                </el-form-item>
                <el-form-item label="详细地址(XX路/街道XX号XX室 示例：南普陀路11号288室)" :label-width="formLabelWidth">
                    <el-input v-model="address"></el-input>
                </el-form-item>
                <el-form-item label="登录密码" :label-width="formLabelWidth">
                    <el-input v-model="password"></el-input>
                </el-form-item></el-form>
            <template #footer>
    <span class="dialog-footer">
      <el-button @click="dialogVisible = false">取 消</el-button>
      <el-button type="primary" @click="confirm">确 定</el-button>
    </span>
            </template>
        </el-dialog>
        </div>
    <div>伊瑟拉编号：{{number}}</div>

</template>

<script>
    import { reactive, toRefs, onMounted, computed, getCurrentInstance } from 'vue'
    import { getCuserDetailList, createCuserDetail, updateCuserDetail } from '@/api/cuserDetail'
    import { useStore } from 'vuex'
    import { changePassword } from '@/api/user'
    export default {
        name: "cuser-detail",
        setup(){
            const { ctx } = getCurrentInstance()
            const store = useStore()
            const info = computed(() => store.getters["user/userInfo"])
            const username = info.value.userName
            const nickname = info.value.nickName
            const password = info.value.password
            const state = reactive({
                options:[{
                    value:'男',
                },
                    {
                        value:'女',
                    }],
                dialogVisible : false,
                hasChanged: false,
                pwdChanged: false,
                firstData:{},
                form:{
                    userName:'',
                    nickName:'',
                    phone:'',
                    sex:'',
                    name:'',
                    region:'',
                    address:'',
                    password:'',
                    wechat:"",
                    number:'',
                    ID:'',
                }
            })
            const checkPassword1 = (rule, value, callback) => {
                if (value.length < 6 || value.length > 12) {
                    return callback(new Error("密码应为6-12位"));
                } else {
                    callback();
                }
            };
            const rules = reactive({
                password: [{ validator: checkPassword1, trigger: "blur" }],
            });

            onMounted (async() => {
                const res = await getCuserDetailList({Username:username})
                if (res.data.list.length == 0 ){
                    ctx.form.userName = username
                    ctx.form.nickName = nickname
                    ctx.form.password = password
                    await createCuserDetail(ctx.form)
                }else {
                    const data = res.data.list[0]
                    ctx.form.userName = data.userName
                    ctx.form.nickName = data.nickName
                    ctx.form.phone = data.phone
                    ctx.form.sex = data.sex
                    ctx.form.name = data.name
                    ctx.form.region = data.region
                    ctx.form.address = data.address
                    ctx.form.password = data.password
                    ctx.form.wechat = data.wechat
                    ctx.form.number = "ysl" + data.ID
                    ctx.form.ID = data.ID
                    ctx.firstData.value = data

                }

            }
         )
            const checkChanged = () => {
                if (state.form.nickName != state.firstData.value.nickName || state.form.sex != state.firstData.value.sex ||state.form.name != state.firstData.value.name ||state.form.phone != state.firstData.value.phone ||state.form.wechat != state.firstData.value.wechat ||state.form.region != state.firstData.value.region ||state.form.address != state.firstData.value.address || state.form.password != state.firstData.value.password){
                    state.hasChanged = true
                }
            }
            const checkPassword = () => {
                if (state.form.password != state.firstData.value.password){
                    state.pwdChanged = true
                }
            }
            const updateUserinfo = async () => {
                await updateCuserDetail(ctx.form)
            }
            const confirm = async () => {
                checkChanged()
                if (state.hasChanged == true) {
                    ctx.$refs.form1.validate(async (v) => {
                        if (v) {
                            updateUserinfo({ID: ctx.ID})
                            checkPassword()
                            if (state.pwdChanged) {
                                changePassword({
                                    username: state.form.userName,
                                    password: state.firstData.value.password,
                                    newPassword: state.form.password
                                })
                                state.firstData.value.password = state.form.password
                            }
                        }
                    })
                }
                state.dialogVisible = false
            }
            return {
                ...toRefs(state),
                ...toRefs(state.form),
                confirm,
            }
        }
    }
</script>

<style scoped>
     .img {
         width: 40px;
         height: 40px;
         border-radius: 20px;
     }
</style>