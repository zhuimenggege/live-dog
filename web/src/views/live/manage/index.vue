<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryRef" v-show="showSearch" :inline="true" label-width="68px">
      <el-form-item label="主播名称" prop="anchor">
        <el-input v-model="queryParams.anchor" placeholder="请输入主播名称" clearable style="width: 180px"
          @keyup.enter="handleQuery" />
      </el-form-item>
      <el-form-item label="房间名称" prop="displayName">
        <el-input v-model="queryParams.roomName" placeholder="请输入房间名称" clearable style="width: 200px"
          @keyup.enter="handleQuery" />
      </el-form-item>
      <el-form-item label="直播平台" prop="platform">
        <el-select v-model="queryParams.platform" placeholder="请选择直播平台" clearable style="width: 200px">
          <el-option v-for="dict in sys_internal_assist_live_platform" :key="dict.value" :label="dict.label"
            :value="dict.value" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
        <el-button icon="Refresh" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" plain icon="Plus" @click="handleAdd" v-hasPermi="['live:manage:add']">添加房间</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" plain icon="Delete" :disabled="multiple" @click="handleDelete"
          v-hasPermi="['live:manage:delete']">删除房间</el-button>
      </el-col>
      <right-toolbar v-model:showSearch="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <!-- 表格数据 -->
    <el-table v-loading="loading" :data="infoList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="数据编号" align="center" prop="id" width="120" />
      <el-table-column label="主播名称" align="center" prop="anchor" :show-overflow-tooltip="true" width="200" />
      <el-table-column label="房间名称" align="center" prop="roomName" :show-overflow-tooltip="true" width="300" />
      <el-table-column label="直播平台" align="center" prop="platform" :show-overflow-tooltip="true" width="120">
        <template #default="scope">
          <dict-tag :options="sys_internal_assist_live_platform" :value="scope.row.platform" />
        </template>
      </el-table-column>
      <el-table-column label="状态" align="center" prop="status" :show-overflow-tooltip="true" width="120">
        <template #default="scope">
          <el-tag v-if="scope.row.recording === true" type="danger">录制中</el-tag>
          <el-tag v-else-if="scope.row.status === 0">未监控</el-tag>
          <el-tag v-else-if="scope.row.status === 1" type="success">监控中（定时）</el-tag>
          <el-tag v-else-if="scope.row.status === 2" type="success">监控中</el-tag>
          <el-tag v-else type="info">未知状态</el-tag>
        </template></el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template #default="scope">
          <el-tooltip content="修改" placement="top" v-if="scope.row.roleId !== 1">
            <el-button link type="primary" icon="Edit" @click="handleUpdate(scope.row)"
              v-hasPermi="['live:manage:update']"></el-button>
          </el-tooltip>
          <el-tooltip content="删除" placement="top">
            <el-button link type="primary" icon="Delete" @click="handleDelete(scope.row)"
              v-hasPermi="['live:manage:delete']"></el-button>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" v-model:page="queryParams.pageNum"
      v-model:limit="queryParams.pageSize" @pagination="getList" />

    <!-- 添加或修改角色配置对话框 -->
    <el-dialog :title="title" v-model="open" width="500px" append-to-body>
      <el-form ref="manageRef" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="房间URL" prop="roomUrl">
          <el-input :disabled="form.id != null" v-model="form.roomUrl" placeholder="请输入房间URL" />
        </el-form-item>
        <el-form-item label="监控间隔" prop="interval">
          <el-input-number v-model="form.interval" placeholder="请输入监控间隔" />
        </el-form-item>
        <el-form-item label="视频格式" prop="format">
          <el-radio-group v-model="form.format">
            <el-radio-button label="flv">FLV(推荐)</el-radio-button>
            <el-radio-button label="mp4">MP4</el-radio-button>
            <el-radio-button label="aac">仅音频</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="启用通知" prop="enableNotice">
          <el-radio-group v-model="form.enableNotice">
            <el-radio-button label="0">禁用</el-radio-button>
            <el-radio-button label="1">启用</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="监控类型" prop="monitorType">
          <el-radio-group v-model="form.monitorType">
            <el-radio-button label="0">不监控</el-radio-button>
            <el-radio-button label="1">定时监控</el-radio-button>
            <el-radio-button label="2">实时监控</el-radio-button>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="监控开始时间" prop="monitorStart" v-if="form.monitorType == 1">
          <el-time-picker v-model="form.monitorStart" format="HH:mm" value-format="HH:mm"
            placeholder="任意时间点"></el-time-picker>
        </el-form-item>
        <el-form-item label="监控停止时间" prop="monitorStop" v-if="form.monitorType == 1">
          <el-time-picker v-model="form.monitorStop" format="HH:mm" value-format="HH:mm"
            placeholder="任意时间点"></el-time-picker>
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
  listInfo,
  getLiveManage,
  addLiveManage,
  updateLiveManage,
  delLiveManage,
} from "@/api/live/manage";

const { proxy } = getCurrentInstance();
const { sys_internal_assist_live_platform } = proxy.useDict("sys_internal_assist_live_platform");

const infoList = ref([]);
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
    platform: undefined,
    roomName: undefined,
  },
  rules: {
    roomUrl: [{ required: true, message: "房间Url不能为空", trigger: "blur" }],
    format: [{ required: true, message: "视频格式不能为空", trigger: "blur" }],
    interval: [
      {
        required: true,
        type: "number",
        message: "间隔时间必须为数字",
        trigger: "blur",
      },
    ],
    enableNotice: [
      { required: true, message: "通知状态不能为空", trigger: "blur" },
    ],
    monitorType: [
      { required: true, message: "监控类型不能为空", trigger: "blur" },
    ],
  },
});

const { queryParams, form, rules } = toRefs(data);

/** 查询每日统计列表 */
function getList() {
  loading.value = true;
  listInfo(queryParams.value).then((response) => {
    infoList.value = response.data.rows;
    total.value = response.data.total;
    loading.value = false;
  });
}

/** 重置新增的表单以及其他数据  */
function reset() {
  form.value = {
    id: undefined,
    roomUrl: undefined,
    interval: 30,
    format: "flv",
    enableNotice: 0,
    monitorType: 0,
    monitorStart: undefined,
    monitorStop: undefined,
    remark: undefined,
  };
  proxy.resetForm("manageRef");
}
/** 添加记录 */
function handleAdd() {
  reset();
  open.value = true;
  title.value = "添加房间";
}
/** 修改记录 */
function handleUpdate(row) {
  reset();
  const id = row.id;
  getLiveManage(id).then((response) => {
    form.value = response.data;
    open.value = true;
    title.value = "修改房间";
  });
}

/** 提交按钮 */
function submitForm() {
  proxy.$refs["manageRef"].validate((valid) => {
    if (valid) {
      if (form.value.id != undefined) {
        updateLiveManage(form.value).then((response) => {
          proxy.$modal.msgSuccess("修改成功");
          open.value = false;
          getList();
          setTimeout(() => {
            getList();
          }, 1500);
        });
      } else {
        addLiveManage(form.value).then((response) => {
          proxy.$modal.msgSuccess("添加成功");
          open.value = false;
          getList();
          setTimeout(() => {
            getList();
          }, 1500);
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
    .confirm('是否确认删除数据编号为"' + dataIds + '"的数据项?')
    .then(function () {
      return delLiveManage(dataIds);
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
