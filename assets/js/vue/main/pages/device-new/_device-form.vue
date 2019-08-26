<template>
  <div class="device-form">
    <el-form label-position="left" :model="form" label-width="100px" :rules="rules" ref="form">
      <el-form-item label="设备名称" prop="name">
        <el-input v-model="form.name"></el-input>
      </el-form-item>

      <el-form-item label="设备类型" prop="type">
        <el-input v-model="form.type"></el-input>
      </el-form-item>

      <el-form-item label="设备描述" prop="description">
        <el-input type="textarea" v-model="form.description"></el-input>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  name: 'device-form',
  data() {
    return {
      rules: {
        name: [{ required: true, message: '设备名称必填', trigger: 'blur' }],
        type: [{ required: true, message: '设备类型必填', trigger: 'blur' }]
      },
      form: {
        name: '',
        type: '',
        description: ''
      }
    };
  }
};
</script>
