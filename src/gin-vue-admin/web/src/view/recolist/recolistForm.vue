<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="id:">
                <el-input v-model="formData.id" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="cpu:">
                <el-input v-model="formData.cpu" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="显卡:">
                <el-input v-model="formData.gpu" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="内存:">
                <el-input v-model="formData.storage" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="硬盘:">
                <el-input v-model="formData.hard" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="显示器:">
                <el-input v-model="formData.monitor" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="总价:">
                  <el-input-number v-model="formData.total" :precision="2" clearable></el-input-number>
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
    createRecolist,
    updateRecolist,
    findRecolist
} from "@/api/recolist";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Recolist",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            id:"",
            cpu:"",
            gpu:"",
            storage:"",
            hard:"",
            monitor:"",
            total:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createRecolist(this.formData);
          break;
        case "update":
          res = await updateRecolist(this.formData);
          break;
        default:
          res = await createRecolist(this.formData);
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
    const res = await findRecolist({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.rerecolist
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