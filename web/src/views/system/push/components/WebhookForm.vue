<template>
  <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
    <el-form-item label="渠道名称" prop="name">
      <el-input v-model="form.name" placeholder="请输入渠道名称" />
    </el-form-item>
    <el-form-item label="URL" prop="url">
      <el-input type="textarea" v-model="form.url" placeholder="请输入Webhook URL" />
    </el-form-item>
    <el-form-item label="状态" prop="status">
      <el-radio-group v-model="form.status">
        <el-radio :label="0">禁用</el-radio>
        <el-radio :label="1">启用</el-radio>
      </el-radio-group>
    </el-form-item>
  </el-form>
</template>

<script setup>
import { ref, defineProps } from 'vue';
const formRef = ref(null);
const props = defineProps({
  form: Object,
  rules: Object,
});

// 使用 defineExpose 暴露 validate 方法
defineExpose({
  validate(callback) {
    if (formRef.value) {
      formRef.value.validate((valid) => callback(valid));
    } else {
      console.error('Form reference is not initialized');
      callback(false);
    }
  }
});

// 定义校验规则
const rules = {
  name: [
    { required: true, message: '请输入渠道名称', trigger: 'blur' },
    {
      min: 2,
      max: 12,
      message: '长度在 2 到 12 个字符',
      trigger: 'blur'
    }
  ],
  url: [
    { required: true, message: '请输入Webhook URL', trigger: 'blur' },
    {
      pattern: /^https?:\/\/[-A-Za-z0-9+&@#\/%?=~_|!:,.;]+[-A-Za-z0-9+&@#\/%=~_|]/,
      message: '请输入正确的 URL',
      trigger: 'blur',
    },
  ],
  status: [{ required: true, message: '请选择是否启用', trigger: 'change' }],
};

</script>
