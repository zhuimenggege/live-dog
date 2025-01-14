<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryRef" v-show="showSearch" :inline="true" label-width="68px">
      <el-form-item label="渠道名称" prop="name">
        <el-input v-model="queryParams.name" placeholder="请输入渠道名称" clearable style="width: 180px"
          @keyup.enter="handleQuery" />
      </el-form-item>
      <el-form-item label="渠道类型" prop="type">
        <el-select v-model="queryParams.type" placeholder="渠道类型" clearable style="width: 120px">
          <el-option v-for="type in typeOptions" :key="type.value" :label="type.label" :value="type.value" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
        <el-button icon="Refresh" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" plain icon="Plus" @click="handleAdd" v-hasPermi="['system:push:add']">新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" plain icon="Delete" :disabled="multiple" @click="handleDelete"
          v-hasPermi="['system:push:delete']">删除</el-button>
      </el-col>
      <right-toolbar v-model:showSearch="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <!-- 表格数据 -->
    <el-table v-loading="loading" :data="channelList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="数据编号" align="center" prop="id" width="120" />
      <el-table-column label="渠道名称" align="center" prop="name" :show-overflow-tooltip="true" width="200" />
      <el-table-column label="渠道类型" align="center" prop="type" :show-overflow-tooltip="true" width="200" />
      <el-table-column label="状态" align="center" prop="status" :show-overflow-tooltip="true" width="100" />
      <el-table-column label="节目备注" align="center" prop="remark" :show-overflow-tooltip="true" width="300" />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template #default="scope">
          <el-tooltip content="修改" placement="top" v-if="scope.row.roleId !== 1">
            <el-button link type="primary" icon="Edit" @click="handleUpdate(scope.row)"
              v-hasPermi="['system:push:update']"></el-button>
          </el-tooltip>
          <el-tooltip content="删除" placement="top">
            <el-button link type="primary" icon="Delete" @click="handleDelete(scope.row)"
              v-hasPermi="['system:push:delete']"></el-button>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" v-model:page="queryParams.pageNum"
      v-model:limit="queryParams.pageSize" @pagination="getList" />

    <!-- 添加或修改角色配置对话框 -->
    <el-dialog :title="title" v-model="open" width="500px" append-to-body>

      <el-form ref="emailRef" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="发送邮箱" prop="email.from">
          <el-input v-model="form.email.from" placeholder="请输入发送邮箱" />
        </el-form-item>
        <el-form-item label="接收邮箱" prop="email.to">
          <el-input v-model="form.email.to" placeholder="请输入接收邮箱" />
        </el-form-item>
        <el-form-item label="smtp服务器" prop="email.server">
          <el-input v-model="form.email.server" placeholder="请输入smtp服务器" />
        </el-form-item>
        <el-form-item label="smtp端口" prop="email.port">
          <el-input v-model="form.email.port" placeholder="请输入smtp端口" />
        </el-form-item>
        <el-form-item label="授权码" prop="email.authCode">
          <el-input v-model="form.email.authCode" placeholder="请输入授权码" />
        </el-form-item>
        <el-form-item label="是否启用" prop="status">
          <el-radio-group v-model="form.status">
            <el-radio-button label="0">禁用</el-radio-button>
            <el-radio-button label="1">启用</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="form.remark" type="textarea" placeholder="请输入备注内容" />
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
  listChannel,
  getChannel,
  addChannel,
  updateChannel,
  delChannel
} from "@/api/system/push";

const { proxy } = getCurrentInstance();

const channelList = ref([]);
const open = ref(false);
const loading = ref(true);
const showSearch = ref(true);
const ids = ref([]);
const single = ref(true);
const multiple = ref(true);
const total = ref(0);
const title = ref("");
const dateRange = ref([]);
const typeOptions = [
  {
    value: 'email',
    label: '邮箱'
  }, {
    value: 'Webhook',
    label: 'Webhook'
  }];

const data = reactive({
  form: {},
  queryParams: {
    pageNum: 1,
    pageSize: 10,
    name: undefined,
    type: undefined,
  },
  rules: {
    email: [{ required: true, message: "角色名称不能为空", trigger: "blur" }],
  },
});

const { queryParams, form, rules } = toRefs(data);

/** 查询每日统计列表 */
function getList() {
  loading.value = true;
  listChannel(proxy.addDateRange(queryParams.value, dateRange.value)).then(
    (response) => {
      channelList.value = response.data.rows;
      total.value = response.data.total;
      loading.value = false;
    }
  );
}

/** 重置新增的表单以及其他数据  */
function reset() {
  form.value = {
    id: undefined,
    name: '邮箱',
    type: 'email',
    status: 1,
    remark: "",
    email: {
      id: undefined,
      channelId: undefined,
      from: "",
      to: "",
      server: "",
      port: "",
      authCode: "",
    }
  };
  proxy.resetForm("emailRef");
}
/** 添加记录 */
function handleAdd() {
  reset();
  open.value = true;
  title.value = "添加推送渠道";
}
/** 修改记录 */
function handleUpdate(row) {
  reset();
  const id = row.id;
  getChannel(id).then(response => {
    form.value = response.data;
    open.value = true;
    title.value = "修改推送渠道";
  });
}

/** 提交按钮 */
function submitForm() {
  proxy.$refs["emailRef"].validate((valid) => {
    if (valid) {
      if (form.value.id != undefined) {
        updateChannel(form.value).then((response) => {
          proxy.$modal.msgSuccess("修改成功");
          open.value = false;
          getList();
        });
      } else {
        addChannel(form.value).then((response) => {
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
      return delChannel(dataIds);
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