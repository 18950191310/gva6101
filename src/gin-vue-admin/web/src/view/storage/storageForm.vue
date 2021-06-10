<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="名称:">
                <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="类型:">
                <el-input v-model="formData.type" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="价格:">
                  <el-input-number v-model="formData.price" :precision="2" clearable></el-input-number>
           </el-form-item>
           
             <el-form-item label="出售状态:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.status" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="库存:"><el-input v-model.number="formData.stock" clearable placeholder="请输入"></el-input>
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
    createStorage,
    updateStorage,
    findStorage
} from "@/api/storage";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Storage",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            name:"",
            type:"",
            price:0,
            status:false,
            stock:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createStorage(this.formData);
          break;
        case "update":
          res = await updateStorage(this.formData);
          break;
        default:
          res = await createStorage(this.formData);
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
    const res = await findStorage({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.restorage
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