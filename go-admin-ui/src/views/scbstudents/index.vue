
<template>
  <div class="app-container">
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="名称" prop="name">
        <el-input
          v-model="queryParams.name"
          placeholder="请输入名称"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="学号" prop="number">
        <el-input
          v-model="queryParams.number"
          placeholder="请输入学号"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="班级id" prop="classId">
        <treeselect
          v-model="form.classId"
          :options="deptOptions"
          :normalizer="normalizer"
          :show-count="true"
          placeholder="选择班级"
          :is-disabled="isEdit"
        />
      </el-form-item>
      <el-form-item label="线路id" prop="lineId">
        <el-input
          v-model="queryParams.lineId"
          placeholder="请输入线路id"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="站点名称" prop="siteName">
        <el-input
          v-model="queryParams.siteName"
          placeholder="请输入站点名称"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="站点id" prop="siteId">
        <el-input
          v-model="queryParams.siteId"
          placeholder="请输入站点id"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="车辆id" prop="carId">
        <el-input
          v-model="queryParams.carId"
          placeholder="请输入车辆id"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="家长电话" prop="parentPhone">
        <el-input
          v-model="queryParams.parentPhone"
          placeholder="请输入家长电话"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="图片" prop="picture">
        <el-input
          v-model="queryParams.picture"
          placeholder="请输入图片"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="创建时间" prop="createdAt">
        <el-input
          v-model="queryParams.createdAt"
          placeholder="请输入创建时间"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="更新时间" prop="updatedAt">
        <el-input
          v-model="queryParams.updatedAt"
          placeholder="请输入更新时间"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="0未删除 1已删除" prop="isDeleted">
        <el-input
          v-model="queryParams.isDeleted"
          placeholder="请输入0未删除 1已删除"
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
      <el-table-column type="selection" width="55" align="center" /><el-table-column
        label=""
        align="center"
        prop="id"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="名称"
        align="center"
        prop="name"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="学号"
        align="center"
        prop="number"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="班级id"
        align="center"
        prop="classId"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="线路id"
        align="center"
        prop="lineId"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="站点名称"
        align="center"
        prop="siteName"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="站点id"
        align="center"
        prop="siteId"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="车辆id"
        align="center"
        prop="carId"
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
      /><el-table-column
        label="创建时间"
        align="center"
        prop="createdAt"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="更新时间"
        align="center"
        prop="updatedAt"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="0未删除 1已删除"
        align="center"
        prop="isDeleted"
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
    <el-dialog :title="title" :visible.sync="open" width="500px">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">

        <el-form-item label="名称" prop="name">
          <el-input
            v-model="form.name"
            placeholder="名称"
          />
        </el-form-item>
        <el-form-item label="学号" prop="number">
          <el-input
            v-model="form.number"
            placeholder="学号"
          />
        </el-form-item>
        <el-form-item label="班级" prop="classId">
          <treeselect
            v-model="form.classId"
            :options="deptOptions"
            :normalizer="normalizer"
            :show-count="true"
            placeholder="选择班级"
            :is-disabled="isEdit"
          />
        </el-form-item>
        <el-form-item label="线路id" prop="lineId">
          <el-input
            v-model="form.lineId"
            placeholder="线路id"
          />
        </el-form-item>
        <el-form-item label="站点名称" prop="siteName">
          <el-input
            v-model="form.siteName"
            placeholder="站点名称"
          />
        </el-form-item>
        <el-form-item label="站点id" prop="siteId">
          <el-input
            v-model="form.siteId"
            placeholder="站点id"
          />
        </el-form-item>
        <el-form-item label="车辆id" prop="carId">
          <el-input
            v-model="form.carId"
            placeholder="车辆id"
          />
        </el-form-item>
        <el-form-item label="家长电话" prop="parentPhone">
          <el-input
            v-model="form.parentPhone"
            placeholder="家长电话"
          />
        </el-form-item>
        <el-form-item label="图片" prop="picture">
          <el-input
            v-model="form.picture"
            placeholder="图片"
          />
        </el-form-item>
        <el-form-item label="0未删除 1已删除" prop="isDeleted">
          <el-input
            v-model="form.isDeleted"
            placeholder="0未删除 1已删除"
          />
        </el-form-item>
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
import Treeselect from '@riophae/vue-treeselect'
import '@riophae/vue-treeselect/dist/vue-treeselect.css'

export default {
  name: 'Scbstudents',
  components: { Treeselect },
  data() {
    return {
      // 遮罩层
      loading: true,
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
            undefined

      },
      // 表单参数
      form: {
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
                  { required: true, message: '班级id不能为空', trigger: 'blur' }
                ],
      lineId:
                [
                  { required: true, message: '线路id不能为空', trigger: 'blur' }
                ],
      siteName:
                [
                  { required: true, message: '站点名称不能为空', trigger: 'blur' }
                ],
      siteId:
                [
                  { required: true, message: '站点id不能为空', trigger: 'blur' }
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
                  { required: true, message: '图片不能为空', trigger: 'blur' }
                ],
      createdAt:
                [
                  { required: true, message: '创建时间不能为空', trigger: 'blur' }
                ],
      updatedAt:
                [
                  { required: true, message: '更新时间不能为空', trigger: 'blur' }
                ],
      isDeleted:
                [
                  { required: true, message: '0未删除 1已删除不能为空', trigger: 'blur' }
                ]
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 转换班级数据结构 */
    normalizer(node) {
      if (node.children && !node.children.length) {
        delete node.children
      }
      return {
        id: node.deptId,
        label: node.deptName,
        children: node.children
      }
    },
    /** 查询班级下拉树结构 */
    getTreeselect(e) {
      getDeptList().then(response => {
        this.deptOptions = []

        if (e === 'update') {
          const dept = { deptId: 0, deptName: '请选择', children: [] }
          dept.children = response.data
          this.deptOptions.push(dept)
        } else {
          const dept = { deptId: 0, deptName: '请选择', children: [] }
          dept.children = response.data
          this.deptOptions.push(dept)
        }
      })
    },
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listScbStudents(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.scbstudentsList = response.data.list
        this.total = response.data.count
        this.loading = false
      }
      )
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
      this.open = true
      this.title = '添加学生信息表'
      this.isEdit = false
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
      const id = row.id || this.ids
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
