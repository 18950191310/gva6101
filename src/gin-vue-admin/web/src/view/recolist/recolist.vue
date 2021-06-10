<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="name">
          <el-input placeholder="搜索条件" v-model="searchInfo.name"></el-input>
        </el-form-item>    
        <el-form-item label="cpu">
          <el-input placeholder="搜索条件" v-model="searchInfo.cpu"></el-input>
        </el-form-item>    
        <el-form-item label="显卡">
          <el-input placeholder="搜索条件" v-model="searchInfo.gpu"></el-input>
        </el-form-item>    
        <el-form-item label="内存">
          <el-input placeholder="搜索条件" v-model="searchInfo.storage"></el-input>
        </el-form-item>    
        <el-form-item label="硬盘">
          <el-input placeholder="搜索条件" v-model="searchInfo.hard"></el-input>
        </el-form-item>    
        <el-form-item label="显示器">
          <el-input placeholder="搜索条件" v-model="searchInfo.monitor"></el-input>
        </el-form-item>    
        <el-form-item label="总价">
          <el-input placeholder="搜索条件" v-model="searchInfo.total"></el-input>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增推荐列表</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="name" prop="name" width="120"></el-table-column>
    
    <el-table-column label="cpu" prop="cpu" width="120"></el-table-column> 
    
    <el-table-column label="显卡" prop="gpu" width="120"></el-table-column> 
    
    <el-table-column label="内存" prop="storage" width="120"></el-table-column> 
    
    <el-table-column label="硬盘" prop="hard" width="120"></el-table-column> 
    
    <el-table-column label="显示器" prop="monitor" width="120"></el-table-column> 
    
    <el-table-column label="总价" prop="total" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateRecolist(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="name:">
            <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
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
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createRecolist,
    deleteRecolist,
    deleteRecolistByIds,
    updateRecolist,
    findRecolist,
    getRecolistList
} from "@/api/recolist";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Recolist",
  mixins: [infoList],
  data() {
    return {
      listApi: getRecolistList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            name:"",
            cpu:"",
            gpu:"",
            storage:"",
            hard:"",
            monitor:"",
            total:0,
            
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10           
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      deleteRow(row){
        this.$confirm('确定要删除吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
           this.deleteRecolist(row);
        });
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteRecolistByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          if (this.tableData.length == ids.length) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateRecolist(row) {
      const res = await findRecolist({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rerecolist;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          name:"",
          cpu:"",
          gpu:"",
          storage:"",
          hard:"",
          monitor:"",
          total:0,
          
      };
    },
    async deleteRecolist(row) {
      const res = await deleteRecolist({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1) {
            this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
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
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
    console.log(this.tableData)
}
};
</script>

<style>
</style>
