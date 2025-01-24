<template>
  <el-form ref="formRef" :model="form" :rules="rules" label-width="100px">
    <el-form-item label="渠道名称" prop="name">
      <el-input v-model="form.name" placeholder="请输入渠道名称" />
    </el-form-item>
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
      <el-input-number v-model="form.email.port" placeholder="请输入smtp端口" width="100%" />
    </el-form-item>
    <el-form-item label="授权码" prop="email.authCode">
      <el-input v-model="form.email.authCode" placeholder="请输入授权码" />
    </el-form-item>
    <el-form-item label="状态" prop="status">
      <el-radio-group v-model="form.status">
        <el-radio :label="0">禁用</el-radio>
        <el-radio :label="1">启用</el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item label="备注" prop="remark">
      <el-input v-model="form.remark" type="textarea" placeholder="请输入备注内容" />
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
  'email.from': [{ required: true, type: 'email', message: '请输入正确的发送邮箱', trigger: 'blur' }],
  'email.to': [
    { required: true, message: '请输入正确的接收邮箱', trigger: 'blur' },
    {
      validator(rule, value, callback) {
        const emails = value.split(',');
        const valid = emails.every(email => {
          const reg = /^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$/;
          return reg.test(email);
        });
        if (!valid) {
          callback(new Error('请输入正确的接收邮箱'));
        } else {
          callback();
        }
      },
      trigger: 'blur'
    }
  ],
  'email.server': [{ required: true, message: '请输入smtp服务器', trigger: 'blur' }],
  'email.port': [{ required: true, message: '请输入smtp端口', trigger: 'blur' },
    { type: 'number', message: 'smtp端口必须为数字', trigger: 'blur' }],
  'email.authCode': [{ required: true, message: '请输入授权码', trigger: 'blur' }],
  status: [{ required: true, message: '请选择是否启用', trigger: 'change' }],
};

</script>
