<!-- 此项目用来 后台管理->商品管理->分类管理 -->

<template>
  <div class="sys">
    <div class="navbar">
      <a
        style="
          color: black;
          margin-top: 6px;
          margin-bottom: 6px;
          font-weight: bold;
          font-size: 18px;
        "
        >分类管理</a
      >
    </div>

    <div class="topButtonBox">
      <el-button color="#626aef" @click="dialogFormVisible = true">添加</el-button>
      <el-button type="danger" @click="deleteSelectionChange">删除</el-button>
    </div>

    <div>
      <el-table
        ref="multipleTableRef"
        :data="tableData"
        :table-layout="tableLayout"
        style="width: 100%"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="40" />
        <el-table-column prop="name" label="分类名称" />
        <el-table-column prop="info" label="描述" />
        <el-table-column prop="sort" label="排序(越小越前)" />
        <el-table-column fixed="right" label="状态">
          <template #default="scope">
            <el-tag v-if="scope.row.tag === 'true'" class="ml-2" type="success"
              >运行</el-tag
            >
            <el-tag v-else-if="scope.row.tag === 'false'" class="ml-2" type="danger"
              >停止</el-tag
            >
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作">
          <template #default="scope">
            <el-button size="small" @click="handleEdit(scope.row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- 页数管理 -->
    <div class="endPagination">
      <el-pagination
        background
        v-model="currentPage"
        @current-change="getNowPageData"
        layout="prev, pager, next"
        :total="tablepage * 10"
      />
    </div>
  </div>

  <!-- 用与添加分类的 from  -->
  <el-dialog v-model="dialogFormVisible" title="添加新的分类">
    <el-form ref="form" :model="form">
      <el-form-item label="名称" :label-width="formLabelWidth">
        <el-input v-model="form.name" autocomplete="off" />
      </el-form-item>
      <el-form-item label="描述" :label-width="formLabelWidth">
        <el-input v-model="form.info" autocomplete="off" />
      </el-form-item>

      <el-form-item label="权重" :label-width="formLabelWidth">
        <el-input-number v-model="form.sort" :min="0" :max="10" />
      </el-form-item>
      <el-form-item label="状态" :label-width="formLabelWidth">
        <el-select v-model="form.tag" placeholder="请选择状态">
          <el-option label="开启" value="true" />
          <el-option label="关闭" value="false" />
        </el-select>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="noAddNewClass">取消</el-button>
        <el-button type="primary" @click="addNewClass"> 确认 </el-button>
      </span>
    </template>
  </el-dialog>

  <!-- 用于编辑分类的 form -->
  <el-dialog v-model="editFormVisible" title="编辑此分类">
    <el-form ref="form" :model="form">
      <el-form-item label="名称" :label-width="formLabelWidth">
        <el-input v-model="form.name" autocomplete="off" />
      </el-form-item>
      <el-form-item label="描述" :label-width="formLabelWidth">
        <el-input v-model="form.info" autocomplete="off" />
      </el-form-item>

      <el-form-item label="权重" :label-width="formLabelWidth">
        <el-input-number v-model="form.sort" :min="0" :max="10" />
      </el-form-item>
      <el-form-item label="状态" :label-width="formLabelWidth">
        <el-select v-model="form.tag" placeholder="请选择状态">
          <el-option label="开启" value="true" />
          <el-option label="关闭" value="false" />
        </el-select>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="noEditClass">取消</el-button>
        <el-button type="primary" @click="confirmEditClass"> 确认 </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script>
import { reactive, ref } from "vue";
import { ElTable } from "element-plus";
import axios from "axios";
import { ElNotification } from "element-plus";
import { ElMessage, ElMessageBox } from "element-plus";

export default {
  data() {
    return {
      search: ref(""),
      multipleTableRef: ref(null),
      tableLayout: ref("fixed"),
      multipleSelection: ref([]),

      dialogFormVisible: ref(false),
      formLabelWidth: "80px",

      editFormVisible: ref(false),

      tableData: [], // 用来存储 分类数据
      tablepage: "", // 总页数
      currentPage: "", // 当前页数

      form: {
        id: 0,
        name: "",
        info: "",
        sort: 0,
        tag: "",
      },
    };
  },
  methods: {
    deleteSelectionChange() {
      this.SelectionChange = [];
      this.multipleSelection.forEach((selectedRow) => {
        this.SelectionChange.push(selectedRow.id);
      });
      if (this.SelectionChange.length > 0) {
        axios
          .post(
            "http://localhost:8084/api/admin/deleteSelectionClass",
            this.SelectionChange
          )
          .then((response) => {
            const statusCode = response.status;
            if (statusCode == 200) {
              return axios.get("http://localhost:8084/api/admin/allClass/page=1");
            } else {
              throw new Error("删除选中失败");
            }
          })
          .then((response) => {
            // 更新表格数据
            this.tableData = response.data;
            // 添加分类成功，显示成功通知
            ElNotification({
              title: "删除选中分类成功",
              message: "删除选中分类成功",
              type: "success",
            });
          })
          .catch((error) => {
            ElNotification({
              title: "删除失败",
              message: "删除选中分类失败",
              type: "error",
            });
          });
      }else{
        ElNotification({
              title: "请至少选中一个选项",
              message: "删除选中分类失败",
              type: "error",
            });
      }
    },
    // 用于监听 选择了 哪些行
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },

    getNowPageData(currentPage) {
      axios
        .get(`http://localhost:8084/api/admin/allClass/page=${currentPage}`)
        .then((response) => {
          this.tableData = response.data;
        });
    },

    addNewClass() {
      // 触发表单验证

      if (
        this.form.name.trim() !== "" &&
        this.form.info.trim() !== "" &&
        this.form.tag.trim() !== ""
      ) {
        // 表单验证通过，发送添加分类的请求
        axios
          .post("http://localhost:8084/api/admin/addClass", this.form)
          .then((response) => {
            const statusCode = response.status;
            if (statusCode === 200) {
              // 添加分类成功后，获取最新数据
              return axios.get("http://localhost:8084/api/admin/allClass/page=1");
            } else {
              throw new Error("添加分类失败");
            }
          })
          .then((response) => {
            // 更新表格数据
            this.tableData = response.data;
            // 添加分类成功，显示成功通知
            ElNotification({
              title: "添加分类成功",
              message: "添加分类成功",
              type: "success",
            });
            // 重新初始化 form
            this.form.name = "";
            this.form.info = "";
            this.form.sort = 0;
            this.form.tag = "";
            // 关闭对话框
            this.dialogFormVisible = false;
          })
          .catch((error) => {
            // 处理添加分类失败的情况，显示错误通知
            ElNotification({
              title: "添加分类失败",
              message: error.message,
              type: "error",
            });
            // 关闭对话框
            this.dialogFormVisible = false;
          });
      } else {
        // 表单验证不通过，显示警告通知
        ElNotification({
          title: "请填写完整数据",
          message: "请填写完整数据",
          type: "warning",
        });
      }
    },

    noAddNewClass() {
      this.dialogFormVisible = false;

      this.form.name = "";
      this.form.info = "";
      this.form.sort = 0;
      this.form.tag = "";
    },
    // 用于编辑此分类
    handleEdit(row) {
      // 将 row 对象的属性拷贝到一个新的对象中
      this.form = { ...row };
      this.editFormVisible = true;
    },
    // 取消编辑此分类
    noEditClass() {
      // 重新初始化 form
      this.form.id = 0;
      this.form.name = "";
      this.form.info = "";
      this.form.sort = 0;
      this.form.tag = "";
      // 关闭对话框
      this.editFormVisible = false;
    },
    // 确认编辑此分类
    confirmEditClass() {
      if (
        this.form.name.trim() !== "" &&
        this.form.info.trim() !== "" &&
        this.form.tag.trim() !== ""
      ) {
        axios
          .post("http://localhost:8084/api/admin/editClass", this.form)
          .then((response) => {
            const statusCode = response.status;

            if (statusCode === 200) {
              // 添加分类成功后，获取最新数据
              return axios.get("http://localhost:8084/api/admin/allClass/page=1");
            } else {
              throw new Error("添加分类失败");
            }
          })
          .then((response) => {
            // 更新表格数据
            this.tableData = response.data;
            // 添加分类成功，显示成功通知
            ElNotification({
              title: "修改成功",
              message: "修改成功",
              type: "success",
            });
            // 重新初始化 form
            this.form.name = "";
            this.form.info = "";
            this.form.sort = 0;
            this.form.tag = "";
            // 关闭对话框
            this.editFormVisible = false;
          })
          .catch((error) => {
            // 处理添加分类失败的情况，显示错误通知
            ElNotification({
              title: "修改失败",
              message: error.message,
              type: "error",
            });
            // 关闭对话框
            this.editFormVisible = false;
          });
      }
    },

    handleDelete(row) {
      ElMessageBox.confirm("此操作会永久删除此分类(无法找回)", "警告", {
        confirmButtonText: "确认",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          axios
            .post("http://localhost:8084/api/admin/deleteClass", row.id)
            .then((response) => {
              const statusCode = response.status;
              if (statusCode === 200) {
                // 添加分类成功后，获取最新数据
                return axios.get("http://localhost:8084/api/admin/allClass/page=1");
              } else {
                throw new Error("删除分类失败");
              }
            })
            .then((response) => {
              // 更新表格数据
              this.tableData = response.data;
              // 添加分类成功，显示成功通知
              ElNotification({
                title: "删除成功",
                message: "删除成功",
                type: "success",
              });
            })
            .catch((error) => {
              // 处理添加分类失败的情况，显示错误通知
              ElNotification({
                title: "删除失败",
                message: error.message,
                type: "error",
              });
            });
        })
        .catch(() => {
          ElMessage({
            type: "info",
            message: "Delete canceled",
          });
        });
    },
  },
  components: {
    ElTable,
  },
  mounted() {
    // 用来获取 初始  页面数据 为 第一页
    axios.get("http://localhost:8084/api/admin/allClass/page=1").then((response) => {
      this.tableData = response.data;
    });
    // 用来获取 页数
    axios.get("http://localhost:8084/api/admin/classPageNum").then((response) => {
      this.tablepage = response.data;
    });
  },
};
</script>

<style scoped>
.sys {
  margin: 1%;
  height: auto;
}

.navbar {
  display: flex;
  justify-content: flex-start;
  background-color: white;

  padding: 10px;
  box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
  border-radius: 5px;
}

.topButtonBox {
  margin-top: 1%;
  display: flex;
  /* 靠右 */
  justify-content: flex-end;
  align-items: center;
  height: 40px;
  background-color: white;
  padding-right: 5%;
  padding-top: 1%;
}
.endPagination {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60px;
  background-color: white;
}
</style>
