
<template>
  <div class="app-container">
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="车牌编号" prop="carNumber">
        <el-input
          v-model="queryParams.carNumber"
          placeholder="请输入车牌编号"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="车牌号" prop="carNo">
        <el-input
          v-model="queryParams.carNo"
          placeholder="请输入车牌号"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="司机" prop="driver">
        <el-input
          v-model="queryParams.driver"
          placeholder="请输入司机"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="手机号" prop="phone">
        <el-input
          v-model="queryParams.phone"
          placeholder="请输入手机号"
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
          v-permisaction="['scbcars:scbcars:add']"
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
        >新增
        </el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          v-permisaction="['scbcars:scbcars:edit']"
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
          v-permisaction="['scbcars:scbcars:remove']"
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
        >删除
        </el-button>
      </el-col>
    </el-row>

    <el-table v-loading="loading" :data="scbcarsList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" /><el-table-column
        label=""
        align="center"
        prop="id"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="车牌编号"
        align="center"
        prop="carNumber"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="车牌号"
        align="center"
        prop="carNo"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="座位数"
        align="center"
        prop="seatsNum"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="跟车员"
        align="center"
        prop="attendantId"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="司机"
        align="center"
        prop="driver"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="手机号"
        align="center"
        prop="phone"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="所属部门"
        align="center"
        prop="dept"
        :show-overflow-tooltip="true"
      />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            v-permisaction="['scbcars:scbcars:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
          >修改
          </el-button>
          <el-button
            v-permisaction="['scbcars:scbcars:remove']"
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
        <el-form-item label="车牌编号" prop="carNumber">
          <el-input
            v-model="form.carNumber"
            placeholder="车牌编号"
          />
        </el-form-item>
        <el-form-item label="车牌号" prop="carNo">
          <el-input
            v-model="form.carNo"
            placeholder="车牌号"
          />
        </el-form-item>
        <el-form-item label="座位数" prop="seatsNum">
          <el-input
            v-model="form.seatsNum"
            placeholder="座位数"
          />
        </el-form-item>
        <el-form-item label="跟车员" prop="attendantId">
          <treeselect
            v-model="form.attendantId"
            :options="attendantsOptions"
            :normalizer="normalizer"
            :show-count="true"
            placeholder="选择跟车员"
            :is-disabled="isEdit"
          />

        </el-form-item>
        <el-form-item label="司机" prop="driver">
          <el-input
            v-model="form.driver"
            placeholder="司机"
          />
        </el-form-item>
        <el-form-item label="手机号" prop="phone">
          <el-input
            v-model="form.phone"
            placeholder="手机号"
          />
        </el-form-item>
        <el-form-item label="所属部门" prop="dept">
          <el-input
            v-model="form.dept"
            placeholder="所属部门"
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
import { addScbCars, delScbCars, getScbCars, listScbCars, updateScbCars } from '@/api/scbcars'
import { getAttendants } from '@/api/scbteachers'
import Treeselect from '@riophae/vue-treeselect'
import '@riophae/vue-treeselect/dist/vue-treeselect.css'
export default {
  name: 'Scbcars',
  components: { Treeselect },
  data() {
    return {
      attendantsOptions: [],
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
      scbcarsList: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        carNumber:
            undefined,
        carNo:
            undefined,
        driver:
            undefined,
        phone:
            undefined

      },
      // 表单参数
      form: {
        deptId: undefined,
        deptName: undefined
      },
      // 表单校验
      rules: { carNumber:
                [
                  { required: true, message: '车牌编号不能为空', trigger: 'blur' }
                ],
      carNo:
                [
                  { required: true, message: '车牌号不能为空', trigger: 'blur' }
                ],
      driver:
                [
                  { required: true, message: '司机不能为空', trigger: 'blur' }
                ],
      phone:
                [
                  { required: true, message: '手机号不能为空', trigger: 'blur' }
                ]
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    normalizer(node) {
      if (node.children && !node.children.length) {
        delete node.children
      }
      return {
        id: node.id,
        label: node.name,
        children: node.children
      }
    },
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listScbCars(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.scbcarsList = response.data.list
        this.total = response.data.count
        this.loading = false
      }
      )
    },
    /** 查询班级下拉树结构 */
    getTreeselect(e) {
      getAttendants().then(response => {
        this.attendantsOptions = []
        const attendants = { id: 0, name: '请选择', children: [] }
        attendants.children = response.data
        this.attendantsOptions.push(attendants)
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
        carNumber: undefined,
        carNo: undefined,
        seatsNum: undefined,
        attendantId: undefined,
        driver: undefined,
        phone: undefined,
        dept: undefined,
        isDelete: undefined,
        deptId: undefined,
        deptName: undefined
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
      this.open = true
      this.title = '添加车辆'
      this.isEdit = false

      this.getTreeselect('add')
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.row)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset()
      getScbCars().then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改ScbCars'
        this.isEdit = true
      })
      this.getTreeselect('update')
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateScbCars(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addScbCars(this.form).then(response => {
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
        return delScbCars(Ids)
      }).then(() => {
        this.getList()
        this.msgSuccess('删除成功')
      }).catch(function() {
      })
    }
  }
}
</script>
