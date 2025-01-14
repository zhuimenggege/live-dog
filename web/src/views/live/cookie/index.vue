<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryRef" v-show="showSearch" :inline="true" label-width="68px">
      <el-form-item label="平台名称" prop="platform">
        <el-select v-model="queryParams.platform" placeholder="请选择所属平台" clearable style="width: 200px">
          <el-option v-for="dict in sys_internal_assist_live_platform" :key="dict.value" :label="dict.label" :value="dict.value" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
        <el-button icon="Refresh" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" plain icon="Plus" @click="handleAdd" v-hasPermi="['live:cookie:add']">新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" plain icon="Delete" :disabled="multiple" @click="handleDelete"
          v-hasPermi="['live:cookie:delete']">删除</el-button>
      </el-col>
      <right-toolbar v-model:showSearch="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <!-- 表格数据 -->
    <el-table v-loading="loading" :data="cookieList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="数据编号" align="center" prop="id" width="100" />
      <el-table-column label="所属平台" align="center" prop="platform" :show-overflow-tooltip="true" width="120">
        <template #default="scope">
          <dict-tag :options="sys_internal_assist_live_platform" :value="scope.row.platform" />
        </template>
      </el-table-column>
      <el-table-column label="Cookie" align="center" prop="cookie" :show-overflow-tooltip="true" width="500" />
      <el-table-column label="备注" align="center" prop="remark" :show-overflow-tooltip="true" width="200" />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template #default="scope">
          <el-tooltip content="修改" placement="top" v-if="scope.row.roleId !== 1">
            <el-button link type="primary" icon="Edit" @click="handleUpdate(scope.row)"
              v-hasPermi="['live:cookie:update']"></el-button>
          </el-tooltip>
          <el-tooltip content="删除" placement="top">
            <el-button link type="primary" icon="Delete" @click="handleDelete(scope.row)"
              v-hasPermi="['live:cookie:delete']"></el-button>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" v-model:page="queryParams.pageNum"
      v-model:limit="queryParams.pageSize" @pagination="getList" />

    <!-- 添加或修改角色配置对话框 -->
    <el-dialog :title="title" v-model="open" width="500px" append-to-body>
      <el-form ref="cookieRef" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="所属平台" prop="anchor">
          <el-select :disabled="form.id != null" v-model="form.platform" placeholder="请输入平台名称" clearable style="width: 200px">
            <el-option v-for="dict in sys_internal_assist_live_platform" :key="dict.value" :label="dict.label" :value="dict.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="Cookie" prop="displayName">
          <el-input v-model="form.cookie" placeholder="请输入Cookie" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="form.remark" type="textarea" placeholder="请输入内容"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitForm">确 定</el-button>
          <el-button @click="cancel">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup name="Cookie">
import {
  listCookie,
  getCookie,
  addCookie,
  updateCookie,
  delCookie,
} from "@/api/live/cookie";

const { proxy } = getCurrentInstance();
const { sys_internal_assist_live_platform } = proxy.useDict("sys_internal_assist_live_platform");

const cookieList = ref([]);
const open = ref(false);
const loading = ref(true);
const showSearch = ref(true);
const ids = ref([]);
const single = ref(true);
const multiple = ref(true);
const total = ref(0);
const title = ref("");

const data = reactive({
  form: {},
  queryParams: {
    pageNum: 1,
    pageSize: 10,
    platform: undefined,
  },
  rules: {
    platform: [{ required: true, message: "所属平台不能为空", trigger: "blur" }],
    cookie: [
      { required: true, message: "Cookie不能为空", trigger: "blur" },
    ],
  },
});

const { queryParams, form, rules } = toRefs(data);

/** 查询每日统计列表 */
function getList() {
  loading.value = true;
  listCookie(queryParams.value).then(
    (response) => {
      cookieList.value = response.data.rows;
      total.value = response.data.total;
      loading.value = false;
    }
  );
}

/** 重置新增的表单以及其他数据  */
function reset() {
  form.value = {
    id: undefined,
    platform: undefined,
    cookie: undefined,
    remark: undefined,
  };
  proxy.resetForm("cookieRef");
}
/** 添加记录 */
function handleAdd() {
  reset();
  open.value = true;
  title.value = "添加Cookie";
}
/** 修改记录 */
function handleUpdate(row) {
  reset();
  const id = row.id;
  getCookie(id).then(response => {
    form.value = response.data;
    open.value = true;
    title.value = "修改Cookie";
  });
}

/** 提交按钮 */
function submitForm() {
  proxy.$refs["cookieRef"].validate((valid) => {
    if (valid) {
      if (form.value.id != undefined) {
        updateCookie(form.value).then((response) => {
          proxy.$modal.msgSuccess("修改成功");
          open.value = false;
          getList();
        });
      } else {
        addCookie(form.value).then((response) => {
          proxy.$modal.msgSuccess("新增成功");
          open.value = false;
          getList();
        });
      }
    }
  });
}
/** 取消按钮 */
function cancel() {
  open.value = false;
  reset();
}

/** 搜索按钮操作 */
function handleQuery() {
  queryParams.value.pageNum = 1;
  getList();
}
/** 重置按钮操作 */
function resetQuery() {
  proxy.resetForm("queryRef");
  handleQuery();
}
/** 删除按钮操作 */
function handleDelete(row) {
  const dataIds = row.id || ids.value;
  proxy.$modal
    .confirm('是否确认删除角色编号为"' + dataIds + '"的数据项?')
    .then(function () {
      return delCookie(dataIds);
    })
    .then(() => {
      getList();
      proxy.$modal.msgSuccess("删除成功");
    })
    .catch(() => { });
}

/** 多选框选中数据 */
function handleSelectionChange(selection) {
  ids.value = selection.map((item) => item.id);
  single.value = selection.length != 1;
  multiple.value = !selection.length;
}

getList();
</script>
