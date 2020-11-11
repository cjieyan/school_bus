
<template>
  <div class="app-container">
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="关键词" prop="name">
        <el-input
          v-model="queryParams.name"
          placeholder="请输入学号/姓名"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          v-permisaction="['scbstudents:scbstudents:add']"
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
        >新增
        </el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          v-permisaction="['scbstudents:scbstudents:edit']"
          type="success"
          icon="el-icon-edit"
          size="mini"
          :disabled="single"
          @click="handleUpdate"
        >修改
        </el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          v-permisaction="['scbstudents:scbstudents:remove']"
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
        >删除
        </el-button>
      </el-col>
    </el-row>

    <el-table v-loading="loading" :data="scbstudentsList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column
        label="学号"
        align="center"
        prop="number"
        :show-overflow-tooltip="true"
      />
      <el-table-column
        label="姓名"
        align="center"
        prop="name"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="班级"
        align="center"
        prop="className"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="线路"
        align="center"
        prop="lineName"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="上车站点名称"
        align="center"
        prop="siteUpName"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="下车站点名称"
        align="center"
        prop="siteDownName"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="车辆"
        align="center"
        prop="carNo"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="家长电话"
        align="center"
        prop="parentPhone"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="图片"
        align="center"
        prop="picture"
        :show-overflow-tooltip="true"
      />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            v-permisaction="['scbstudents:scbstudents:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
          >修改
          </el-button>
          <el-button
            v-permisaction="['scbstudents:scbstudents:remove']"
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
          >删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryParams.pageIndex"
      :limit.sync="queryParams.pageSize"
      @pagination="getList"
    />
    <!-- 添加或修改对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="700px">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-row :gutter="15">
          <el-col :span="9">
            <el-form-item label="姓名" prop="name">
              <el-input
                v-model="form.name"
                placeholder="姓名"
              />
            </el-form-item>
          </el-col>
          <el-col :span="15" style="position: absolute ;right: 10px;z-index: 5">
            <el-form-item label="图片" prop="picture">
              <el-upload
                ref="picture"
                :file-list="picturefileList"
                :class="{disabled:uploadDisabled}"
                :limit="1"
                :action="pictureAction"
                :auto-upload="false"
                :before-upload="pictureBeforeUpload"
                :on-change="onUploadChange"
                :on-success="onUploadSuccess"
                :on-remove="handleRemove"
                :disabled="showdisabled"
                list-type="picture-card"
                accept="image/*"
                name="picture"
              >
                <i class="el-icon-plus" />
                <div slot="tip" class="el-upload__tip">只能上传不超过 2MB 的图片文件</div>
              </el-upload>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="15">
          <el-col :span="9">
            <el-form-item label="学号" prop="number">
              <el-input
                v-model="form.number"
                placeholder="学号"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="15">
          <el-col :span="9">
            <el-form-item label="班级" prop="classId">
              <el-cascader
                v-model="form.classId"
                :options="deptOptions"
                :props="classIdProps"
                :style="{width: '100%'}"
                placeholder="请选择"
                clearable
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="15">
          <el-col :span="9">
            <el-form-item label="线路" prop="lineId">
              <el-select v-model="form.lineId" placeholder="选择线路" clearable :style="{width: '100%'}">
                <el-option
                  v-for="(item, index) in linesOptions"
                  :key="index"
                  :label="item.label"
                  :value="item.value"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="15">
          <el-col :span="24">
            <el-form-item label="车辆" prop="carId">
              <el-radio-group v-model="form.carId" size="medium">
                <el-radio
                  v-for="(item, index) in carIdsOptions"
                  :key="index"
                  :label="item.value"
                  :disabled="item.disabled"
                >{{ item.label }}</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="15">
          <el-col :span="9">
            <el-form-item label="上车站点" prop="siteIdUp">
              <el-select v-model="form.siteIdUp" placeholder="上车站点" clearable :style="{width: '100%'}">
                <el-option
                  v-for="(item, index) in siteIdsOptions"
                  :key="index"
                  :label="item.label"
                  :value="item.value"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="9">
            <el-form-item label="下车站点" prop="siteIdDown">
              <el-select v-model="form.siteIdDown" placeholder="上车站点" clearable :style="{width: '100%'}">
                <el-option
                  v-for="(item, index) in siteIdsOptions"
                  :key="index"
                  :label="item.label"
                  :value="item.value"
                  :disabled="item.disabled"
                />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="15">
          <el-col :span="9">
            <el-form-item label="是否接送" prop="isPickUp">
              <el-switch v-model="form.isPickUp" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="15">
          <el-col :span="9">
            <el-form-item label="家长电话" prop="parentPhone">
              <el-input
                v-model="form.parentPhone"
                placeholder="家长电话"
              />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { addScbStudents, delScbStudents, getScbStudents, listScbStudents, updateScbStudents } from '@/api/scbstudents'
import { getDeptList } from '@/api/scbdept'
import '@riophae/vue-treeselect/dist/vue-treeselect.css'
import { getAllLines, getLineCars, getLineSites } from '@/api/scblines'

export default {
  name: 'Scbstudents',
  data() {
    return {
      showdisabled: false,
      // 遮罩层
      loading: true,
      pictureAction: 'https://jsonplaceholder.typicode.com/posts/',
      picturefileList: [],
      // 线路列表
      linesOptions: [],
      // 站点列表
      siteIdsOptions: [],
      // 线路车辆列表
      carIdsOptions: [],
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 弹出层标题
      title: '',
      // 是否显示弹出层
      open: false,
      isEdit: false,
      // 类型数据字典
      typeOptions: [],
      scbstudentsList: [],
      // 班级树选项
      deptOptions: [],
      uploadDisabled: '',
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        id:
            undefined,
        name:
            undefined,
        number:
            undefined,
        classId:
            undefined,
        lineId:
            undefined,
        siteName:
            undefined,
        siteId:
            undefined,
        carId:
            undefined,
        parentPhone:
            undefined,
        picture:
            undefined,
        createdAt:
            undefined,
        updatedAt:
            undefined,
        isDeleted:
            undefined,
        isPickUp:
            undefined

      },
      // 表单参数
      form: {
        picture: ''
      },
      // 表单校验
      rules: { id:
                [
                  { required: true, message: '不能为空', trigger: 'blur' }
                ],
      name:
                [
                  { required: true, message: '名称不能为空', trigger: 'blur' }
                ],
      number:
                [
                  { required: true, message: '学号不能为空', trigger: 'blur' }
                ],
      classId:
                [
                  {
                    required: true,
                    message: '请至少选择一个班级',
                    trigger: 'change'
                  }
                ],
      lineId:
                [
                  { required: true, message: '线路不能为空', trigger: 'blur' }
                ],
      siteIdUp:
                [
                  { required: true, message: '上车站点不能为空', trigger: 'blur' }
                ],
      siteIdDown:
                [
                  { required: true, message: '下车站点不能为空', trigger: 'blur' }
                ],
      carId:
                [
                  { required: true, message: '车辆id不能为空', trigger: 'blur' }
                ],
      parentPhone:
                [
                  { required: true, message: '家长电话不能为空', trigger: 'blur' }
                ],
      picture:
                [
                  // { required: true, message: '图片不能为空', trigger: 'blur' }
                ],
      isDeleted:
                [
                  { required: true, message: '0未删除 1已删除不能为空', trigger: 'blur' }
                ]
      },
      classIdProps: {
        'multiple': false,
        'label': 'label',
        'value': 'value',
        'children': 'children'
      }
    }
  },
  watch: {
    // 监听deptId
    'form.lineId': 'curLineChange'
  },
  created() {
    this.getList()
  },
  methods: {
    // 线路切换
    curLineChange(val) {
      // 获取上下车站点
      getLineCars().then(response => {
        this.carIdsOptions = []
        for (var i = 0; i < response.data.length; i++) {
          const d = response.data[i]
          this.carIdsOptions.push({
            'label': d.carNo,
            'value': d.id
          })
        }
      })
      getLineSites().then(response => {
        this.siteIdsOptions = []
        for (var i = 0; i < response.data.length; i++) {
          const d = response.data[i]
          const formatData = {
            'label': d.name,
            'value': d.id
          }
          this.siteIdsOptions.push(formatData)
        }
      })
      // 车辆列表
    },
    /** 查询班级下拉树结构 */
    getLinesSelect(e) {
      getAllLines().then(response => {
        this.linesOptions = []
        for (var i = 0; i < response.data.length; i++) {
          const d = response.data[i]
          const formatData = {
            'label': d.name,
            'value': d.id
          }
          this.linesOptions.push(formatData)
        }
      })
    },
    getTreeselect(e) {
      getDeptList().then(response => {
        this.deptOptions = this.getTreeData(response.data)
      })
    },
    getTreeData(data) {
      // 循环遍历json数据
      for (var i = 0; i < data.length; i++) {
        if (data[i].children.length < 1) {
          // children若为空数组，则将children设为undefined
          data[i].children = undefined
        } else {
          // children若不为空数组，则继续 递归调用 本方法
          this.getTreeData(data[i].children)
        }
      }
      return data
    },
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listScbStudents(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.scbstudentsList = response.data.list
        this.total = response.data.count
        this.loading = false
      })
    },
    // 取消按钮
    cancel() {
      this.open = false
      this.reset()
    },
    // 表单重置
    reset() {
      this.form = {
        id: undefined,
        name: undefined,
        number: undefined,
        classId: undefined,
        lineId: undefined,
        siteName: undefined,
        siteId: undefined,
        carId: undefined,
        parentPhone: undefined,
        picture: undefined,
        isDeleted: undefined
      }
      this.resetForm('form')
    },
    onUploadChange(file) {
      const isIMAGE = (file.raw.type === 'image/jpeg' || file.raw.type === 'image/png' || file.raw.type === 'image/gif')
      // const isLt1M = file.size / 1024 / 1024 < 1;

      if (!isIMAGE) {
        this.$message.error('上传文件只能是图片格式!')
        return false
      }
      var reader = new FileReader()
      reader.readAsDataURL(file.raw)
      const that = this
      reader.onload = function(e) {
        that.form.picture = this.result
      }
    },
    onUploadSuccess(file) {
      if (this.picturefileList.length > 0) {
        this.uploadDisabled = 'disabled'
      }
    },
    handleRemove(file, fileList) {
      // 将变量置空
      this.uploadDisabled = ''
    },
    pictureBeforeUpload(file) {
      const isRightSize = file.size / 1024 / 1024 < 2
      if (!isRightSize) {
        this.$message.error('文件大小超过 2MB')
      }
      return isRightSize
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1
      this.getList()
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = []
      this.resetForm('queryForm')
      this.handleQuery()
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset()
      this.getTreeselect('add')
      this.getLinesSelect('add')
      this.open = true
      this.title = '添加学生信息表'
      this.isEdit = false
      this.picture = undefined
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      this.getTreeselect('update')
      this.getLinesSelect('update')
      const id = row.id || this.ids
      this.picture = undefined
      getScbStudents(id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改学生信息表'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.isPickUp) {
            this.form.isPickUp = 1
          } else {
            this.form.isPickUp = 0
          }
          if (this.form.classId.length > 0) {
            this.form.classId = this.form.classId[this.form.classId.length - 1]
          }
          if (this.form.id !== undefined) {
            updateScbStudents(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addScbStudents(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess('新增成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          }
        }
      })
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const Ids = row.id || this.ids
      this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(function() {
        return delScbStudents(Ids)
      }).then(() => {
        this.getList()
        this.msgSuccess('删除成功')
      }).catch(function() {
      })
    }
  }
}
</script>
<style rel="stylesheet/scss" lang="scss" scoped>

</style>
<style rel="stylesheet/scss" lang="scss">
.disabled .el-upload--picture-card{
  display: none;
}
</style>
