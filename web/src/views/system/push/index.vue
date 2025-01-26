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
      <el-table-column label="状态" align="center" prop="status" :show-overflow-tooltip="true" width="100" >
        <template #default="scope">
          <el-tag v-if="scope.row.status === 1">启用</el-tag>
          <el-tag v-else-if="scope.row.status === 0" type="danger">禁用</el-tag>
          <el-tag v-else type="info">未知</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="备注" align="center" prop="remark" :show-overflow-tooltip="true" width="300" />
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

    <!-- 添加或修改对话框 -->
    <el-dialog :title="title" v-model="open" width="500px" append-to-body>
      <el-tabs v-model="activeTab" type="border-card">
        <el-tab-pane label="Email" name="email" :disabled="isEditing && activeTab !== 'email'">
          <EmailForm ref="emailRef" :form="form" :rules="rules" />
        </el-tab-pane>
        <el-tab-pane label="Gotify" name="gotify" :disabled="isEditing && activeTab !== 'gotify'">
          <GotifyForm ref="gotifyRef" :form="form" :rules="rules" />
        </el-tab-pane>
      </el-tabs>

      <template #footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="submitForm">确 定</el-button>
          <el-button @click="cancel">取 消</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup name="Push">
import {
  listChannel,
  getChannel,
  addChannel,
  updateChannel,
  delChannel
} from "@/api/system/push";
import EmailForm from '@/views/system/push/components/EmailForm.vue';
import GotifyForm from '@/views/system/push/components/GotifyForm.vue';

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
const isEditing = ref(false);
const typeOptions = [
  {
    value: 'email',
    label: '邮箱'
  }, {
    value: 'gotify',
    label: 'Gotify'
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

  }
});

const { queryParams, form, rules } = toRefs(data);

const activeTab = ref('email');

/** 查询推送渠道列表 */
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
    status: 0,
    email: {
      port: 25,
    },
    web: {}
  };
  proxy.resetForm("formRef");
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
    // 确保 email 属性被正确初始化
    if (!form.value.email) {
      form.value.email = {
        port: 25,
      };
    }
    if (!form.value.web) {
      form.value.web = {};
    }
    open.value = true;
    title.value = "修改推送渠道";
    isEditing.value = true; // 设置为编辑模式
    activeTab.value = form.value.type || 'email';
  });
}

const emailRef = ref(null);
const gotifyRef = ref(null);
/** 提交按钮 */
function submitForm() {
  const formRef = {
    'email': emailRef,
    'gotify': gotifyRef
  }[activeTab.value];

  if (formRef) {
    formRef.value.validate((valid) => {
      if (valid) {
        form.value.type = activeTab.value;
        if (form.value.id != undefined) {
          updateChannel(form.value).then((response) => {
            proxy.$modal.msgSuccess("修改成功");
            open.value = false;
            getList();
            resetTab(); // 重置选项卡
          });
        } else {
          addChannel(form.value).then((response) => {
            proxy.$modal.msgSuccess("新增成功");
            open.value = false;
            getList();
            resetTab(); // 重置选项卡
          });
        }
      }
    });
  } else {
    console.error('formRef is undefined');
  }
}
/** 取消按钮 */
function cancel() {
  open.value = false;
  reset();
  nextTick(() => {
    resetTab();
  });
}

function resetTab() {
  activeTab.value = 'email'; // 重置为默认选项卡
  isEditing.value = false; // 退出编辑模式
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