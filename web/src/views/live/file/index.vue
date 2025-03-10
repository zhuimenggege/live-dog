<template>
    <div class="app-container">
        <el-form :model="queryParams" ref="queryRef" v-show="showSearch" :inline="true" label-width="68px">
            <el-form-item label="文件名称" prop="filename">
                <el-input v-model="queryParams.filename" placeholder="请输入文件名称" clearable style="width: 180px"
                    @keyup.enter="handleQuery" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" icon="Search" @click="handleQuery">搜索</el-button>
                <el-button icon="Refresh" @click="resetQuery">重置</el-button>
            </el-form-item>
        </el-form>
        <el-row :gutter="10" class="mb8">
            <el-col :span="1.5">
                <el-button type="danger" plain icon="Delete" :disabled="multiple" @click="handleDelete"
                    v-hasPermi="['file:manage:delete']">删除文件</el-button>
            </el-col>
            <right-toolbar v-model:showSearch="showSearch" @queryTable="getList"></right-toolbar>
        </el-row>

        <!-- 表格数据 -->
        <el-breadcrumb separator="/" class="app-breadcrumb">
            <el-breadcrumb-item v-for="(part, index) in data.pathParts" :key="index">
                <span @click="navigateTo(index)" style="cursor: pointer;">
                    {{ part.name }}
                </span>
            </el-breadcrumb-item>
        </el-breadcrumb>
        <el-table v-loading="loading" :data="fileList" @selection-change="handleSelectionChange">
            <el-table-column type="selection" width="55" align="center" />
            <el-table-column label="类型" width="50" align="center">
                <template #default="scope">
                    <el-icon v-if="scope.row.isFolder">
                        <folder />
                    </el-icon>
                    <el-icon v-else>
                        <component :is="getFileIconType(scope.row.filename)" />
                    </el-icon>
                </template>
            </el-table-column>
            <el-table-column label="文件名" align="center" prop="filename" :show-overflow-tooltip="true" width="580">
                <template #default="scope">
                    <span @click="openFolder(scope.row.filename, scope.row.isFolder)"
                        :style="{ cursor: scope.row.isFolder ? 'pointer' : 'default' }">
                        {{ scope.row.filename }}
                    </span>
                </template>
            </el-table-column>
            <el-table-column label="文件大小" align="center" prop="size" :show-overflow-tooltip="true" width="120">
                <template #default="scope">
                    <span v-show="!scope.row.isFolder">
                        {{ formatSize(scope.row.size) }}
                    </span>
                </template>
            </el-table-column>
            <el-table-column label="最后修改时间" align="center" prop="lastModified" :show-overflow-tooltip="true"
                width="160">
                <template #default="scope">
                    {{ formatDate(scope.row.lastModified) }}
                </template>
            </el-table-column>
            <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                <template #default="scope">
                    <el-tooltip content="播放" placement="top"
                        v-if="isMediaFile(scope.row.filename) && !scope.row.isFolder">
                        <el-button link type="primary" icon="VideoPlay" @click="handlePlay(scope.row)"></el-button>
                    </el-tooltip>
                    <el-tooltip content="删除" placement="top">
                        <el-button link type="primary" icon="Delete" @click="handleDelete(scope.row)"
                            v-hasPermi="['file:manage:delete']"></el-button>
                    </el-tooltip>
                </template>
            </el-table-column>
        </el-table>
        <el-dialog v-model="mediaDialog.visible" :title="mediaDialog.title" width="70%" destroy-on-close>
            <div v-if="mediaDialog.type === 'video'">
                <VideoPlayer :src="mediaDialog.src" :type="mediaDialog.fileType" :autoplay="false" :controls="true" />
            </div>
            <div v-else-if="mediaDialog.type === 'audio'">
                <AudioPlayer :src="mediaDialog.src" :autoplay="false" />
            </div>
        </el-dialog>
    </div>
</template>

<script setup name="File">
import {
    listFile,
    delFile,
} from "@/api/live/file";
import {
    formatDate,
    formatSize,
} from "@/utils/index";

import VideoPlayer from '@/components/VideoPlayer/index.vue';
import AudioPlayer from '@/components/AudioPlayer/index.vue';

const { proxy } = getCurrentInstance();

const fileList = ref([]);
const loading = ref(true);
const showSearch = ref(true);
const filenames = ref([]);
const single = ref(true);
const multiple = ref(true);

// 媒体播放对话框数据
const mediaDialog = reactive({
    visible: false,
    title: '',
    src: '',
    type: '',
    fileType: ''
});

const data = reactive({
    pathParts: [{ name: '系统目录', path: '/' }],
    queryParams: {
        filename: undefined,
        path: undefined,
    },
});

const { queryParams } = toRefs(data);

const navigateTo = (index) => {
    data.pathParts = data.pathParts.slice(0, index + 1);
    data.queryParams.path = data.pathParts[index].path;
    getList();
}

function getFileIconType(filename) {
    if (!filename) return 'document';
    const ext = filename.split('.').pop().toLowerCase();
    if (['flv', 'mp4', 'mov', 'ts'].includes(ext)) {
        return 'film';
    }
    if (['mp3', 'aac', 'flac', 'wav'].includes(ext)) {
        return 'headset';
    }
    return 'document';
}

// 判断是否为媒体文件
function isMediaFile(filename) {
    if (!filename) return false;
    const ext = filename.split('.').pop().toLowerCase();
    return ['mp4', 'flv', 'aac', 'mp3'].includes(ext);
}

// 获取文件类型
function getFileType(filename) {
    if (!filename) return '';
    const ext = filename.split('.').pop().toLowerCase();
    if (['mp4'].includes(ext)) return 'mp4';
    if (['flv'].includes(ext)) return 'flv';
    if (['aac', 'mp3'].includes(ext)) return 'aac';
    return '';
}

// 处理媒体播放
async function handlePlay(row) {
    const filename = row.filename;
    const ext = filename.split('.').pop().toLowerCase();
    const currentPath = data.queryParams.path || '/';
    const filePath = currentPath === '/' ? filename : `${currentPath}/${filename}`;

    try {
        const baseUrl = import.meta.env.VITE_APP_BASE_API;
        const mediaUrl = `${baseUrl}/file/manage/play?path=${encodeURIComponent(filePath)}`;

        mediaDialog.title = filename;
        mediaDialog.src = mediaUrl;
        mediaDialog.fileType = getFileType(filename);

        // 根据文件类型决定使用哪种播放器
        if (['mp4', 'flv'].includes(ext)) {
            mediaDialog.type = 'video';
        } else if (['aac', 'mp3'].includes(ext)) {
            mediaDialog.type = 'audio';
        }

        mediaDialog.visible = true;

        // 在对话框关闭时释放 Blob URL
        watch(() => mediaDialog.visible, (newValue) => {
            if (!newValue) {
                URL.revokeObjectURL(mediaUrl);
            }
        });
    } catch (error) {
        console.error('媒体加载失败:', error);
        proxy.$modal.msgError('媒体加载失败');
    }
}

// 打开子目录
function openFolder(folder, isFolder) {
    if (!isFolder) return;
    const currentPath = data.pathParts[data.pathParts.length - 1].path;
    data.queryParams.path = currentPath === "/" ? folder : currentPath + "/" + folder;
    data.pathParts.push({ name: folder, path: data.queryParams.path });
    getList();
};

/** 查询每日统计列表 */
function getList() {
    loading.value = true;
    listFile(queryParams.value).then((response) => {
        fileList.value = response.data.rows;
        loading.value = false;
    });
}

/** 搜索按钮操作 */
function handleQuery() {
    getList();
}

/** 重置按钮操作 */
function resetQuery() {
    proxy.resetForm("queryRef");
    handleQuery();
}
/** 删除按钮操作 */
function handleDelete(row) {
    const deleteData = {
        filenames: row.filename || filenames.value,
        path: data.queryParams.path
    };
    proxy.$modal
        .confirm('是否确认删除?')
        .then(function () {
            return delFile(deleteData);
        })
        .then(() => {
            getList();
            proxy.$modal.msgSuccess("删除成功");
        })
        .catch(() => { });
}

/** 多选框选中数据 */
function handleSelectionChange(selection) {
    filenames.value = selection.map((item) => item.filename);
    single.value = selection.length != 1;
    multiple.value = !selection.length;
}

getList();
</script>

<style lang='scss' scoped>
.app-breadcrumb.el-breadcrumb {
    display: inline-block;
    font-size: 14px;
    line-height: 30px;
    margin-left: 8px;

    .no-redirect {
        color: #97a8be;
        cursor: text;
    }
}
</style>