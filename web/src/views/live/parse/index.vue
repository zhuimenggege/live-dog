<template>
    <div class="app-container">
        <el-form :model="queryParams" ref="queryRef" v-show="showSearch" :inline="true" label-width="68px">
            <el-form-item label="作者名称" prop="author">
                <el-input v-model="queryParams.author" placeholder="请输入作者名称" clearable style="width: 180px"
                    @keyup.enter="handleQuery" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
                <el-button icon="Refresh" @click="resetQuery">重置</el-button>
            </el-form-item>
        </el-form>
        <el-row :gutter="10" class="mb8">
            <el-col :span="1.5">
                <el-button type="primary" plain icon="Plus" @click="handleAdd"
                    v-hasPermi="['live:parse:add']">解析</el-button>
            </el-col>
            <el-col :span="1.5">
                <el-button type="danger" plain icon="Delete" :disabled="multiple" @click="handleDelete"
                    v-hasPermi="['live:parse:delete']">删除</el-button>
            </el-col>
            <right-toolbar v-model:showSearch="showSearch" @queryTable="getList"></right-toolbar>
        </el-row>

        <!-- 表格数据 -->
        <el-table v-loading="loading" :data="parseList" @selection-change="handleSelectionChange">
            <el-table-column type="selection" width="55" align="center" />
            <el-table-column label="数据ID" align="center" prop="id" width="80" />
            <el-table-column label="平台" align="center" prop="platform" :show-overflow-tooltip="true" width="80">
                <template #default="scope">
                    <dict-tag :options="sys_internal_assist_live_platform" :value="scope.row.platform" />
                </template>
            </el-table-column>
            <el-table-column label="媒体类型" align="center" width="80" :show-overflow-tooltip="true">
                <template #default="scope">
                    <span v-if="scope.row.type === 'video'">视频</span>
                    <span v-if="scope.row.type === 'music'">音乐</span>
                    <span v-if="scope.row.type === 'note'">图集</span>
                </template>
            </el-table-column>
            <el-table-column label="作品封面" align="center" width="240">
                <template #default="scope">
                    <img v-if="scope.row.type === 'video'" :src="scope.row.videoCoverUrl" alt="视频封面"
                        style="width: 100px; height: auto;" />
                    <img v-if="scope.row.type === 'music'" :src="scope.row.musicCoverUrl" alt="音乐封面"
                        style="width: 100px; height: auto;" />
                    <img v-if="scope.row.type === 'note'" :src="scope.row.imagesCoverUrl" alt="图集封面"
                        style="width: 100px; height: auto;" />
                </template>
            </el-table-column>
            <el-table-column label="作者" align="center" prop="author" :show-overflow-tooltip="true" width="140" />
            <el-table-column label="作品描述" align="center" prop="desc" :show-overflow-tooltip="true" width="240" />
            <el-table-column label="解析时间" align="center" prop="createTime" :show-overflow-tooltip="true" width="160" />
            <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                <template #default="scope">
                    <el-tooltip content="详情" placement="top" v-if="scope.row.roleId !== 1">
                        <el-button link type="primary" icon="View" @click="handleDetail(scope.row)"
                            v-hasPermi="['live:parse:get']"></el-button>
                    </el-tooltip>
                    <el-tooltip content="删除" placement="top">
                        <el-button link type="primary" icon="Delete" @click="handleDelete(scope.row)"
                            v-hasPermi="['live:parse:delete']"></el-button>
                    </el-tooltip>
                </template>
            </el-table-column>
        </el-table>

        <pagination v-show="total > 0" :total="total" v-model:page="queryParams.pageNum"
            v-model:limit="queryParams.pageSize" @pagination="getList" />

        <el-dialog title="媒体解析" v-model="parse" align-center close-on-press-escape>
            <template #title>
                媒体解析
                <el-popover placement="right" title="支持平台及类型" :width="300" effect="dark" trigger="hover">
                    <div slot="content">
                        抖音web版分享链接(视频/图集)
                    </div>
                    <template #reference>
                        <el-icon class="m-2">
                            <InfoFilled />
                        </el-icon>
                    </template>
                </el-popover>
            </template>
            <el-form ref="parseRef" :model="form" :rules="rules" label-width="100px">
                <el-form-item label="解析链接" prop="url" class="centered-form-item">
                    <el-input v-model="form.url" placeholder="请输入解析链接" clearable size="large">
                        <template #append>
                            <el-button type="primary" @click="submitForm" style="height: 40px; padding: 0 20px;">
                                解析
                            </el-button>
                        </template>
                    </el-input>
                </el-form-item>
            </el-form>
        </el-dialog>

        <el-dialog title="解析详情" v-model="detailDialog" width="60%">
            <el-descriptions :column="2" border v-if="detailData">
                <el-descriptions-item label="媒体类型" :span="1">
                    <span v-if="detailData.type === 'video'">视频</span>
                    <span v-if="detailData.type === 'music'">音乐</span>
                    <span v-if="detailData.type === 'note'">图集</span>
                </el-descriptions-item>
                <el-descriptions-item label="媒体ID" :span="1">{{ detailData.mediaId }}</el-descriptions-item>
                <el-descriptions-item label="作者" :span="1">{{ detailData.author }}</el-descriptions-item>
                <el-descriptions-item label="解析时间" :span="1">{{ detailData.createTime }}</el-descriptions-item>
                <el-descriptions-item label="作品描述" :span="2">{{ detailData.desc }}</el-descriptions-item>
                <el-descriptions-item label="视频封面" :span="2" v-if="detailData.type === 'video'">
                    <img :src="detailData.videoCoverUrl" alt="视频封面" style="width: 100px; height: auto;" />
                </el-descriptions-item>
                <el-descriptions-item label="音乐封面" :span="2" v-if="detailData.type === 'music'">
                    <img :src="detailData.musicCoverUrl" alt="音乐封面" style="width: 100px; height: auto;" />
                </el-descriptions-item>
                <el-descriptions-item label="图集封面" :span="2" v-if="detailData.type === 'note'">
                    <img :src="detailData.imagesCoverUrl" alt="图集封面" style="width: 100px; height: auto;" />
                </el-descriptions-item>
                <el-descriptions-item label="视频链接" :span="2" v-if="detailData.type === 'video'">
                    <a :href="detailData.videoUrl" target="_blank">{{ detailData.videoUrl }}</a>
                </el-descriptions-item>
                <el-descriptions-item label="音乐链接" :span="2" v-if="detailData.type === 'music'">
                    <a :href="detailData.musicUrl" target="_blank">{{ detailData.musicUrl }}</a>
                </el-descriptions-item>
                <el-descriptions-item label="图集链接" :span="2" v-if="detailData.type === 'note'">
                    <el-scrollbar style="max-height: 100px; overflow-y: auto;">
                        <div v-for="(url, index) in detailData.imagesUrl.split(',')" :key="index"
                            class="link-container">
                            <a :href="url.trim()" target="_blank" class="link" :title="url.trim()">{{ url.trim() }}</a>
                        </div>
                    </el-scrollbar>
                </el-descriptions-item>
            </el-descriptions>
        </el-dialog>
    </div>
</template>

<script setup name="Parse">
import {
    listParseInfo,
    getParseInfo,
    parseUrl,
    delParseInfo,
} from "@/api/live/parse";

const { proxy } = getCurrentInstance();
const { sys_internal_assist_live_platform } = proxy.useDict("sys_internal_assist_live_platform");

const parseList = ref([]);
const parse = ref(false);
const loading = ref(true);
const showSearch = ref(true);
const ids = ref([]);
const single = ref(true);
const multiple = ref(true);
const total = ref(0);
const detailDialog = ref(false);
const detailData = ref(null);

const data = reactive({
    form: {},
    queryParams: {
        pageNum: 1,
        pageSize: 10,
        author: undefined,
    },
    rules: {
        url: [{ required: true, message: "解析链接不能为空", trigger: "blur" }],
    },
});

const { queryParams, form, rules } = toRefs(data);

/** 查询每日统计列表 */
function getList() {
    loading.value = true;
    listParseInfo(queryParams.value).then(
        (response) => {
            parseList.value = response.data.rows;
            total.value = response.data.total;
            loading.value = false;
        }
    );
}

/** 重置新增的表单以及其他数据  */
function reset() {
    form.value = {
        id: undefined,
        url: undefined,
    };
    proxy.resetForm("parseRef");
}
/** 添加记录 */
function handleAdd() {
    reset();
    parse.value = true;
}

function handleDetail(row) {
    getParseInfo(row.id).then(response => {
        detailData.value = response.data;
        detailDialog.value = true;
    });
}

/** 提交按钮 */
function submitForm() {
    proxy.$refs["parseRef"].validate((valid) => {
        if (valid) {
            parseUrl(form.value).then((response) => {
                proxy.$modal.msgSuccess("解析成功");
                parse.value = false;
                getList();
            });
        }
    });
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
        .confirm('是否确认删除数据编号为"' + dataIds + '"的数据项?')
        .then(function () {
            return delParseInfo(dataIds);
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

<style scoped>
.link-container {
    width: 100%;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
}

.link {
    display: inline-block;
    max-width: 80ch;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
}
</style>