<template>
  <div class="app-container">
    <el-row>
      <el-col :span="12" class="card-box">
        <el-card>
          <template #header>
            <Cpu style="width: 1em; height: 1em; vertical-align: middle;" /> <span
              style="vertical-align: middle;">CPU</span>
          </template>
          <div class="el-table el-table--enable-row-hover el-table--medium">
            <table cellspacing="0" style="width: 100%;">
              <thead>
                <tr>
                  <th class="el-table__cell is-leaf">
                    <div class="cell">属性</div>
                  </th>
                  <th class="el-table__cell is-leaf">
                    <div class="cell">值</div>
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td class="el-table__cell is-leaf">
                    <div class="cell">核心数</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell" v-if="server.cpu">{{ server.cpu.cores }}</div>
                  </td>
                </tr>
                <tr>
                  <td class="el-table__cell is-leaf">
                    <div class="cell">CPU名称</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell" v-if="server.cpu">{{ server.cpu.modelName }}</div>
                  </td>
                </tr>
                <tr>
                  <td class="el-table__cell is-leaf">
                    <div class="cell">CPU频率</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell" v-if="server.cpu">{{ server.cpu.mhz }}Hz</div>
                  </td>
                </tr>
                <tr>
                  <td class="el-table__cell is-leaf">
                    <div class="cell">当前使用率</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell" v-if="server.cpu">{{ parseFloat(server.cpu.percent).toFixed(2) }}%</div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </el-card>
      </el-col>

      <el-col :span="12" class="card-box">
        <el-card>
          <template #header>
            <Tickets style="width: 1em; height: 1em; vertical-align: middle;" /> <span
              style="vertical-align: middle;">内存</span>
          </template>
          <div class="el-table el-table--enable-row-hover el-table--medium">
            <table cellspacing="0" style="width: 100%;">
              <thead>
                <tr>
                  <th class="el-table__cell is-leaf">
                    <div class="cell">属性</div>
                  </th>
                  <th class="el-table__cell is-leaf">
                    <div class="cell">内存</div>
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td class="el-table__cell is-leaf">
                    <div class="cell">总内存</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell" v-if="server.mem">{{ formatSize(server.mem.total) }}</div>
                  </td>
                </tr>
                <tr>
                  <td class="el-table__cell is-leaf">
                    <div class="cell">已用内存</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell" v-if="server.mem">{{ formatSize(server.mem.used) }}</div>
                  </td>
                </tr>
                <tr>
                  <td class="el-table__cell is-leaf">
                    <div class="cell">剩余内存</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell" v-if="server.mem">{{ formatSize(server.mem.available) }}</div>
                  </td>
                </tr>
                <tr>
                  <td class="el-table__cell is-leaf">
                    <div class="cell">使用率</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell" v-if="server.mem" :class="{ 'text-danger': server.mem.usedPercent > 80 }">{{
                      parseFloat(server.mem.usedPercent).toFixed(2) }}%</div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </el-card>
      </el-col>

      <el-col :span="24" class="card-box">
        <el-card>
          <template #header>
            <Monitor style="width: 1em; height: 1em; vertical-align: middle;" /> <span
              style="vertical-align: middle;">服务器信息</span>
          </template>
          <div class="el-table el-table--enable-row-hover el-table--medium">
            <table cellspacing="0" style="width: 100%;">
              <tbody>
                <tr>
                  <th class="el-table__cell is-leaf">
                    <div class="cell">服务器名称</div>
                  </th>
                  <th class="el-table__cell is-leaf">
                    <div class="cell">操作系统</div>
                  </th>
                  <th class="el-table__cell is-leaf">
                    <div class="cell">系统架构</div>
                  </th>
                  <th class="el-table__cell is-leaf">
                    <div class="cell">开机时间</div>
                  </th>
                </tr>
                <tr>
                  <td class="el-table__cell is-leaf">
                    <div class="cell" v-if="server.sys">{{ server.sys.hostname }}</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell" v-if="server.sys">{{ server.sys.os }}</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell" v-if="server.sys">{{ server.sys.kernelArch }}</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell" v-if="server.sys">{{ server.sys.bootTime }}</div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </el-card>
      </el-col>

      <el-col :span="24" class="card-box">
        <el-card>
          <template #header>
            <MessageBox style="width: 1em; height: 1em; vertical-align: middle;" /> <span
              style="vertical-align: middle;">磁盘状态</span>
          </template>
          <div class="el-table el-table--enable-row-hover el-table--medium">
            <table cellspacing="0" style="width: 100%;">
              <thead>
                <tr>
                  <th class="el-table__cell el-table__cell is-leaf">
                    <div class="cell">盘符路径</div>
                  </th>
                  <th class="el-table__cell is-leaf">
                    <div class="cell">文件系统</div>
                  </th>
                  <th class="el-table__cell is-leaf">
                    <div class="cell">总大小</div>
                  </th>
                  <th class="el-table__cell is-leaf">
                    <div class="cell">可用大小</div>
                  </th>
                  <th class="el-table__cell is-leaf">
                    <div class="cell">已用大小</div>
                  </th>
                  <th class="el-table__cell is-leaf">
                    <div class="cell">已用百分比</div>
                  </th>
                </tr>
              </thead>
              <tbody v-if="server.disks">
                <tr v-for="(disk, index) in server.disks" :key="index">
                  <td class="el-table__cell is-leaf">
                    <div class="cell">{{ disk.path }}</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell">{{ disk.fstype }}</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell">{{ formatSize(disk.total) }}</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell">{{ formatSize(disk.free) }}</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell">{{ formatSize(disk.used) }}</div>
                  </td>
                  <td class="el-table__cell is-leaf">
                    <div class="cell" :class="{ 'text-danger': disk.usedPercent > 80 }">{{ parseFloat(disk.usedPercent).toFixed(2) }}%</div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { getServer } from '@/api/monitor/server'
import { formatSize } from '@/utils/index'

const server = ref([]);
const { proxy } = getCurrentInstance();

function getList() {
  proxy.$modal.loading("正在加载服务监控数据，请稍候！");
  getServer().then(response => {
    server.value = response.data;
    proxy.$modal.closeLoading();
  });
}

getList();
</script>
