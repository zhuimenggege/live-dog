<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryRef" v-show="showSearch" :inline="true" label-width="68px">
      <el-form-item label="主播名称" prop="anchor">
        <el-input v-model="queryParams.anchor" placeholder="请输入主播名称" clearable style="width: 180px"
          @keyup.enter="handleQuery" />
      </el-form-item>
      <el-form-item label="节目名称" prop="displayName">
        <el-input v-model="queryParams.displayName" placeholder="请输入节目名称" clearable style="width: 180px"
          @keyup.enter="handleQuery" />
      </el-form-item>
      <el-form-item label="节目时间" style="width: 220px" prop="displayDate">
        <el-date-picker v-model="queryParams.displayDate" value-format="YYYY-MM-DD" type="date"
          placeholder="选择日期"></el-date-picker>
      </el-form-item>
      <el-form-item label="节目类型" prop="displayType">
        <el-select v-model="queryParams.displayType" placeholder="节目类型" clearable style="width: 100px">
          <el-option v-for="type in anchor_show_type" :key="type.value" :label="type.label" :value="type.value" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
        <el-button icon="Refresh" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" plain icon="Plus" @click="handleAdd" v-hasPermi="['live:daily:add']">新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="success" plain icon="Edit" :disabled="single" @click="handleUpdate"
          v-hasPermi="['live:daily:update']">修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" plain icon="Delete" :disabled="multiple" @click="handleDelete"
          v-hasPermi="['live:daily:delete']">删除</el-button>
      </el-col>
      <right-toolbar v-model:showSearch="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <!-- 表格数据 -->
    <el-table v-loading="loading" :data="dailyList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="数据编号" align="center" prop="id" width="120" />
      <el-table-column label="主播名称" align="center" prop="anchor" :show-overflow-tooltip="true" width="160" />
      <el-table-column label="节目名称" align="center" prop="displayName" :show-overflow-tooltip="true" width="240" />
      <el-table-column label="节目类型" align="center" prop="displayType" :show-overflow-tooltip="true" width="120">
        <template #default="scope">
          <dict-tag :options="anchor_show_type" :value="scope.row.displayType" />
        </template>
      </el-table-column>
      <el-table-column label="节目次数" align="center" prop="count" :show-overflow-tooltip="true" width="100" />
      <el-table-column label="节目时间" align="center" prop="displayDate" :show-overflow-tooltip="true" width="120" />
      <el-table-column label="节目备注" align="center" prop="remark" :show-overflow-tooltip="true" width="240" />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template #default="scope">
          <el-tooltip content="修改" placement="top" v-if="scope.row.roleId !== 1">
            <el-button link type="primary" icon="Edit" @click="handleUpdate(scope.row)"
              v-hasPermi="['live:daily:update']"></el-button>
          </el-tooltip>
          <el-tooltip content="删除" placement="top">
            <el-button link type="primary" icon="Delete" @click="handleDelete(scope.row)"
              v-hasPermi="['live:daily:delete']"></el-button>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" v-model:page="queryParams.pageNum"
      v-model:limit="queryParams.pageSize" @pagination="getList" />

    <!-- 添加或修改角色配置对话框 -->
    <el-dialog :title="title" v-model="open" width="500px" append-to-body>
      <el-form ref="dailyRef" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="主播名称" prop="anchor">
          <el-input v-model="form.anchor" placeholder="请输入主播名称" />
        </el-form-item>
        <el-form-item label="节目名称" prop="displayName">
          <el-input v-model="form.displayName" placeholder="请输入节目名称" />
        </el-form-item>
        <el-form-item label="节目类型" prop="displayType">
          <el-select v-model="form.displayType" placeholder="节目类型">
            <el-option v-for="type in anchor_show_type" :key="type.value" :label="type.label" :value="type.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="节目时间" prop="displayDate">
          <el-date-picker v-model="form.displayDate" value-format="YYYY-MM-DD" type="date"
            placeholder="选择日期"></el-date-picker>
        </el-form-item>
        <el-form-item label="节目次数" prop="count">
          <el-input v-model="form.count" placeholder="请输入节目次数" />
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

<script setup name="Live">
import {
  listDaily,
  getDaily,
  addDaily,
  updateDaily,
  delDaily,
} from "@/api/live/daily";

const { proxy } = getCurrentInstance();
const { anchor_show_type } = proxy.useDict("anchor_show_type");

const dailyList = ref([]);
const open = ref(false);
const loading = ref(true);
const showSearch = ref(true);
const ids = ref([]);
const single = ref(true);
const multiple = ref(true);
const total = ref(0);
const title = ref("");
const dateRange = ref([]);

const data = reactive({
  form: {},
  queryParams: {
    pageNum: 1,
    pageSize: 10,
    anchor: undefined,
    displayName: undefined,
    displayType: undefined,
    displayDate: undefined,
  },
  rules: {
    anchor: [{ required: true, message: "主播名称不能为空", trigger: "blur" }],
    displayName: [
      { required: true, message: "节目名称不能为空", trigger: "blur" },
    ],
    displayType: [
      { required: true, message: "节目类型不能为空", trigger: "blur" },
    ],
    displayDate: [
      { required: true, message: "节目日期不能为空", trigger: "blur" },
    ],
  },
});

const { queryParams, form, rules } = toRefs(data);

/** 查询每日统计列表 */
function getList() {
  loading.value = true;
  listDaily(proxy.addDateRange(queryParams.value, dateRange.value)).then(
    (response) => {
      dailyList.value = response.data.rows;
      total.value = response.data.total;
      loading.value = false;
    }
  );
}

/** 重置新增的表单以及其他数据  */
function reset() {
  form.value = {
    id: undefined,
    anchor: undefined,
    displayName: undefined,
    displayType: undefined,
    displayDate: undefined,
    count: 1,
    remark: undefined,
  };
  proxy.resetForm("dailyRef");
}
/** 添加记录 */
function handleAdd() {
  reset();
  open.value = true;
  title.value = "添加每日统计";
}
/** 修改记录 */
function handleUpdate(row) {
  reset();
  const id = row.id;
  getDaily(id).then(response => {
    form.value = response.data;
    open.value = true;
    title.value = "修改每日统计";
  });
}

/** 提交按钮 */
function submitForm() {
  proxy.$refs["dailyRef"].validate((valid) => {
    if (valid) {
      if (form.value.id != undefined) {
        updateDaily(form.value).then((response) => {
          proxy.$modal.msgSuccess("修改成功");
          open.value = false;
          getList();
        });
      } else {
        addDaily(form.value).then((response) => {
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
  dateRange.value = [];
  proxy.resetForm("queryRef");
  handleQuery();
}
/** 删除按钮操作 */
function handleDelete(row) {
  const dataIds = row.id || ids.value;
  proxy.$modal
    .confirm('是否确认删除角色编号为"' + dataIds + '"的数据项?')
    .then(function () {
      return delDaily(dataIds);
    })
    .then(() => {
      getList();
      proxy.$modal.msgSuccess("删除成功");
    })
    .catch(() => { });
}

/** 多选框选中数据 */
function handleSelectionChange(selection) {
  ids.value = selection.map((item) => item.roleId);
  single.value = selection.length != 1;
  multiple.value = !selection.length;
}

getList();
</script>
