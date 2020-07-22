<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增图片</el-button>
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

    <el-table-column label="图片路径" prop="img_path" width="120">
      <template slot-scope="scope">
        <el-image
                style="width: 100px; height: 100px"
                :src="scope.row.img_path"
                :preview-src-list="[scope.row.img_path]">
        </el-image>
      </template>
    </el-table-column>

    <el-table-column label="图片描述" prop="description" width="120"></el-table-column>

    <el-table-column label="所属类别id" prop="cat_id" width="120"></el-table-column>

    <el-table-column label="排序" prop="display_order" width="120"></el-table-column>

    <el-table-column label="状态" prop="status" width="120"></el-table-column>

    <el-table-column label="用户id" prop="user_id" width="120"></el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updatePicture(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deletePicture(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
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
      <el-form ref="elForm" :model="formData" size="medium" label-width="100px">
        <el-col :span="23">
          <el-form-item label="上传" prop="imgPath" required>
            <el-upload ref="imgPath" :file-list="imgPathfileList" :action="imgPathAction"
                       :before-upload="imgPathBeforeUpload" :limit="1" :on-success="UploadSuccess" :on-remove="UploadRemove">
              <el-button size="small" type="primary" icon="el-icon-upload">点击上传</el-button>
            </el-upload>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="图片描述" prop="description">
            <el-input v-model="formData.description" type="textarea" placeholder="请输入图片描述"
                      :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="类别id" prop="cat_id">
            <el-input v-model="formData.cat_id" placeholder="请输入类别id" clearable :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="排序" prop="display_order">
            <el-input-number v-model="formData.display_order" placeholder="排序"></el-input-number>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item size="large">
            <el-button type="primary" @click="submitForm">提交</el-button>
            <el-button @click="resetForm">重置</el-button>
          </el-form-item>
        </el-col>
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
    createPicture,
    deletePicture,
    deletePictureByIds,
    updatePicture,
    findPicture,
    getPictureList
} from "@/api/picture";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Picture",
  mixins: [infoList],
  data() {
    return {
      listApi: getPictureList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
        img_path:null,description:null,cat_id:null,display_order:null,status:null,user_id:null,
      },
      imgPathAction: '/api/fileUploadAndDownload/upload',
      imgPathfileList: [],
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
      async onDelete() {
        const ids = []
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deletePictureByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updatePicture(row) {
      const res = await findPicture({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.repicture;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {

          img_path:null,
          description:null,
          cat_id:null,
          display_order:null,
          status:null,
          user_id:null,
      };
    },
    async deletePicture(row) {
      this.visible = false;
      const res = await deletePicture({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createPicture(this.formData);
          break;
        case "update":
          res = await updatePicture(this.formData);
          break;
        default:
          res = await createPicture(this.formData);
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
    },
    submitForm() {
      this.$refs['elForm'].validate(valid => {
        if (!valid) return
        // TODO 提交表单
      })
    },
    resetForm() {
      this.$refs['elForm'].resetFields()
    },
    imgPathBeforeUpload(file) {
      let isRightSize = file.size / 1024 / 1024 < 20
      if (!isRightSize) {
        this.$message.error('文件大小超过 20MB')
      }
      return isRightSize
    },
    UploadSuccess(response, file, fileList) {
      console.log(response)
      if (fileList.length > 0) {
        let file = fileList[0].response.data.file
        this.formData.img_path = file.url
      } else {
        this.formData.img_path = ''
      }
      console.log(this.formData.img_path)
    },
    UploadRemove(file, fileList) {
      this.formData.img_path = ''
      console.log(this.formData.img_path)
    }
  },
  async created() {
    await this.getTableData();}
};
</script>

<style>
</style>
