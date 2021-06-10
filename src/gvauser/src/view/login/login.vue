<template>
  <div id="userLayout" class="user-layout-wrapper">
    <div class="container">
      <div class="top">
        <div class="desc">
          <img class="logo_login" src="@/assets/logo_login.png" alt="" />
        </div>
        <div class="header">
          <a href="/">
            <!-- <img src="~@/assets/logo.png" class="logo" alt="logo" /> -->
            <span class="title">Gin-Vue-Admin</span>
          </a>
        </div>
      </div>
      <div class="main">
        <el-form
          :model="form"
          :rules="rules"
          ref="loginForm"
          @keyup.enter="submitLogin"
        >
          <el-form-item prop="username">
            <el-input placeholder="请输入账号"  v-model="form.username">
              <template #suffix>
                <i class="el-input__icon el-icon-user"></i>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
              v-model="form.password"
            >
              <template #suffix>
                <i :class="'el-input__icon el-icon-' + lock" @click="changeLock"></i>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item style="position: relative">
            <el-input
              v-model="form.captcha"
              placeholder="请输入验证码"
              style="width: 60%"
            />
            <div class="vPic">
              <img
                v-if="picPath"
                :src="picPath"
                alt="请输入验证码"
                @click="loginVefify()"
              />
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitLogin" style="width: 100%"
              >登 录</el-button
            >
          </el-form-item>
          <el-form-item>
            <span>  <el-button type="primary" @click="openRegister" style="width: 45%;color:yellow"
            >注册</el-button
            >
            <el-dialog title="注册账号" v-model="dialogVisible1">
              <div> <el-form
                      :model="form1"
                      :rules="rules1"
                      ref="registerForm"
                      @keyup.enter="submitRegister"
                      class="register"
              >
          <el-form-item label="账号（推荐使用数字或字母）" prop="username">
            <el-input placeholder="请输入账号"  v-model="form1.username" class="input">
              <template #suffix>
                <i class="el-input__icon el-icon-user"></i>
              </template>
            </el-input>
          </el-form-item>
                 <el-form-item label="昵称"  prop="nickName">
          <el-input v-model="form1.nickName" placeholder="请输入昵称"></el-input>
        </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input
                    type='text'
                    placeholder="请输入密码"
                    v-model="form1.password"
                    class="input"
            >
            </el-input>
          </el-form-item>
             <el-form-item label="确认密码" prop="password2">
             <el-input
                     type='text'
                     placeholder="请再次输入密码"
                     v-model="form1.password2"
                     class="input"
             >
            </el-input>
          </el-form-item>
          <el-form-item label="验证码" style="position: relative">
            <el-input
                    v-model="form1.captcha"
                    placeholder="请输入验证码"
                    style="width: 60%"
                    class="input"
            />
            <div class="vPic">
              <img
                      v-if="picPath"
                      :src="picPath"
                      alt="请输入验证码"
                      @click="loginVefify1()"
              />
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitRegister" style="width: 100%"
            >注册</el-button
            >
          </el-form-item>
   </el-form>
              </div>

</el-dialog>
            </span>
          </el-form-item>

        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
import { captcha, register } from "@/api/user";
import { getCurrentInstance, ref, reactive, onMounted } from "vue";
import { useStore } from 'vuex';

export default {
  name: "Login",
  setup() {
    const dialogVisible1 = ref(false)
    const { ctx } = getCurrentInstance();
    const store = useStore();
    const curYear = ref(0);
    const lock = ref("lock");
    const picPath = ref("");
    const form1 = reactive({
      username:'',
      nickName:'',
      password:'',
      password2:'',
      authorityId: "666",
      captcha: "",
      captchaId: "",
    })
    const form = reactive({
      username: "admin",
      password: "123456",
      captcha: "",
      captchaId: "",
    });
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5 || value.length > 12) {
        return callback(new Error("请输入正确的账号(长度5-12)"));
      }
      else {
        callback();
      }
    };
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6 || value.length > 12) {
        return callback(new Error("密码应为6-12位"));
      } else {
        callback();
      }
    };
    const openRegister = () => {
      dialogVisible1.value = true;
      loginVefify1();
    }
    const rules = reactive({
      username: [{ validator: checkUsername, trigger: "blur" }],
      password: [{ validator: checkPassword, trigger: "blur" }],
    });
    const checkPassword2 = (rule, value, callback) => {
      if (value !== form1.password ) {
        return callback(new Error("两次密码不一致"));
      } else {
        callback();
      }
    };
    const rules1 = reactive({
      username: [{ validator: checkUsername, trigger: "blur" }],
      password: [{ validator: checkPassword, trigger: "blur" }],
      password2: [{ validator: checkPassword2, trigger:"blur" }],
      nickName:[{ required: true, message: "请输入昵称",trigger:"blur" }],
    })



    const LoginIn = (form) => store.dispatch("user/LoginIn",form);


    const login = async () => {
      return await LoginIn(form)
    };
    const Register = async () => {
      return await register(form1)
    }
    const submitLogin = async () => {
      ctx.$refs.loginForm.validate(async (v) => {
        if (v) {
          const flag = await login()
          if(!flag){
            loginVefify();
          }
        } else {
          ctx.$message({
            type: "error",
            message: "请正确填写登录信息",
            showClose: true,
          });
          loginVefify();  
          return false;
        }
      });

    };

    const submitRegister = async() => {
      loginVefify1();
      ctx.$refs.registerForm.validate(async (v) => {
        if (v) {
          const res = await Register(form1);
          if (res.code == 0) {
            ctx.$message({ type: "success", message: "注册成功" });
            dialogVisible1.value = false
          }
        }


    })
    };
    const changeLock = () => {
      lock.value === "lock" ? (lock.value = "unlock") : (lock.value = "lock");
    };
    const loginVefify = () => {
      captcha({}).then((ele) => {
        picPath.value = ele.data.picPath;
        form.captchaId = ele.data.captchaId;
      });
    };
    const loginVefify1 = () => {
      captcha({}).then((ele) => {
        picPath.value = ele.data.picPath;
        form1.captchaId = ele.data.captchaId;
      });
    };
    onMounted(async () => {
         await loginVefify()
    })
    return {
      lock,
      form,
      rules,
      changeLock,
      picPath,
      loginVefify,
      submitLogin,
      curYear,
      dialogVisible1,
      form1,
      rules1,
      submitRegister,
      loginVefify1,
      openRegister,
    };
  }
};
</script>

<style scoped lang="scss">
@import "@/style/login.scss";
  .input{
    margin: 15px;
  }
</style>
