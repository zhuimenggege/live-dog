<template>
    <div class="app-container">
      <el-form :model="queryParams" ref="queryRef" v-show="showSearch" :inline="true" label-width="68px">
        <el-form-item label="主播名称" prop="anchor">
            <el-input v-model="queryParams.anchor" placeholder="请输入主播名称" clearable style="width: 180px" @keyup.enter="handleQuery" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
          <el-button icon="Refresh" @click="resetQuery">重置</el-button>
        </el-form-item>
      </el-form>
      <el-row :gutter="10" class="mb8">
        <el-col :span="1.5">
          <el-button type="primary" plain icon="Plus" @click="handleAdd" v-hasPermi="['live:history:add']">新增</el-button>
        </el-col>
        <el-col :span="1.5">
          <el-button type="danger" plain icon="Delete" :disabled="multiple" @click="handleDelete"
            v-hasPermi="['live:history:delete']">删除</el-button>
        </el-col>
        <right-toolbar v-model:showSearch="showSearch" @queryTable="getList"></right-toolbar>
      </el-row>
  
      <!-- 表格数据 -->
      <el-table v-loading="loading" :data="historyList" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center" />
        <el-table-column label="数据编号" align="center" prop="id" width="100" />
        <el-table-column label="主播名称" align="center" prop="anchor" :show-overflow-tooltip="true" width="250" />
        <el-table-column label="开播时间" align="center" prop="startTime" :show-overflow-tooltip="true" width="250" />
        <el-table-column label="下播时间" align="center" prop="endTime" :show-overflow-tooltip="true" width="250" />
        <el-table-column label="直播时长" align="center" prop="duration" :show-overflow-tooltip="true" width="200" />
        <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
          <template #default="scope">
            <el-tooltip content="修改" placement="top" v-if="scope.row.roleId !== 1">
              <el-button link type="primary" icon="Edit" @click="handleUpdate(scope.row)"
                v-hasPermi="['live:history:update']"></el-button>
            </el-tooltip>
            <el-tooltip content="删除" placement="top">
              <el-button link type="primary" icon="Delete" @click="handleDelete(scope.row)"
                v-hasPermi="['live:history:delete']"></el-button>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>
  
      <pagination v-show="total > 0" :total="total" v-model:page="queryParams.pageNum"
        v-model:limit="queryParams.pageSize" @pagination="getList" />
  
      <!-- 添加或修改直播历史对话框 -->
      <el-dialog :title="title" v-model="open" width="500px" append-to-body>
        <el-form ref="historyRef" :model="form" :rules="rules" label-width="100px">
          <el-form-item label="主播名称" prop="liveId">
            <el-select
              :disabled="form.id != null"
              v-model="form.liveId"
              placeholder="请选择主播"
              clearable
              style="width: 200px"
            >
              <el-option
                v-for="room in roomList"
                :key="room.liveId"
                :label="room.anchor"
                :value="room.liveId"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="开播时间" prop="startTime">
            <el-date-picker
              v-model="form.startTime"
              type="datetime"
              placeholder="请选择开播时间"
              style="width: 100%"
            />
          </el-form-item>
          <el-form-item label="下播时间" prop="endTime">
            <el-date-picker
              v-model="form.endTime"
              type="datetime"
              placeholder="请选择下播时间"
              style="width: 100%"
            />
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
  
  <script setup name="History">
  import {
    listHistory,
    getHistory,
    addHistory,
    updateHistory,
    delHistory,
  } from "@/api/live/history";
  import {
    listInfo,
  } from "@/api/live/manage";
  
  const { proxy } = getCurrentInstance();
  
  const historyList = ref([]);
  const roomList = ref([]);
  const open = ref(false);
  const loading = ref(true);
  const showSearch = ref(true);
  const ids = ref([]);
  const single = ref(true);
  const multiple = ref(true);
  const total = ref(0);
  const title = ref("");
  
  const data = reactive({
    form: {
      anchor: null,
      liveId: null,
    },
    queryParams: {
      pageNum: 1,
      pageSize: 10,
      anchor: undefined,
      liveId: undefined,
    },
    roomQueryParams: {
      pageNum: 1,
      pageSize: 10,
      anchor: undefined,
    },
    rules: {
      liveId: [{ required: true, message: "主播名称不能为空", trigger: "blur" }],
      startTime: [{ required: true, message: "开播时间不能为空", trigger: "blur" }],
      endTime: [{ required: true, message: "下播时间不能为空", trigger: "blur" }],
    },
  });
  
  const { queryParams, form, rules, roomQueryParams } = toRefs(data);
  
  /** 查询每日统计列表 */
  function getList() {
    loading.value = true;
    listHistory(queryParams.value).then(
      (response) => {
        historyList.value = response.data.rows;
        total.value = response.data.total;
        loading.value = false;
      }
    );
  }

  function getRoomList() {
    loading.value = true;
    listInfo(roomQueryParams.value).then(
      (response) => {
        roomList.value = response.data.rows;
        total.value = response.data.total;
        loading.value = false;
      }
    );
  }
  
  /** 重置新增的表单以及其他数据  */
  function reset() {
    form.value = {
      id: undefined,
      anchor: null,
      liveId: null,
    };
    proxy.resetForm("historyRef");
  }
  /** 添加记录 */
  function handleAdd() {
    reset();
    open.value = true;
    title.value = "添加直播历史";
  }
  /** 修改记录 */
  function handleUpdate(row) {
    reset();
    const id = row.id;
    getHistory(id).then(response => {
      form.value = response.data;
      open.value = true;
      title.value = "修改直播历史";
    });
  }
  
  /** 提交按钮 */
  function submitForm() {
    proxy.$refs["historyRef"].validate((valid) => {
      if (valid) {
        if (form.value.id != undefined) {
          updateHistory(form.value).then((response) => {
            proxy.$modal.msgSuccess("修改成功");
            open.value = false;
            getList();
          });
        } else {
          addHistory(form.value).then((response) => {
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
        return delHistory(dataIds);
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
  
  getRoomList();
  getList();
  </script>
  