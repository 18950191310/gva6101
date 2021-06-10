<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="昵称:">
                <el-input v-model="formData.nickName" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="手机号:">
                <el-input v-model="formData.phone" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="性别:">
                <el-input v-model="formData.sex" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="称呼:">
                <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="地区:">
                <el-input v-model="formData.region" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="地址:">
                <el-input v-model="formData.address" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="密码:">
                <el-input v-model="formData.password" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="微信号:">
                <el-input v-model="formData.wechat" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="伊瑟拉编号:">
                <el-input v-model="formData.number" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createCuserDetail,
    updateCuserDetail,
    findCuserDetail
} from "@/api/cuserDetail";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "CuserDetail",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            nickName:"",
            phone:"",
            sex:"",
            name:"",
            region:"",
            address:"",
            password:"",
            wechat:"",
            number:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createCuserDetail(this.formData);
          break;
        case "update":
          res = await updateCuserDetail(this.formData);
          break;
        default:
          res = await createCuserDetail(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findCuserDetail({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.recuserDetail
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
}
};
</script>

<style>
</style>