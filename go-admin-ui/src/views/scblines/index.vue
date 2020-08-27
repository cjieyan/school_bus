
<template>
  <div class="app-container">
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="" prop="id">
        <el-input
          v-model="queryParams.id"
          placeholder="请输入"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="线路名称" prop="name">
        <el-input
          v-model="queryParams.name"
          placeholder="请输入线路名称"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="出发时间" prop="departed_at">
        <el-time-picker
          v-model="queryParams.departed_at"
          format="HH:mm:ss"
          value-format="HH:mm:ss"
          :picker-options='{"selectableRange":"00:00:00-23:59:59"}'
          :style="{width: '100%'}"
          placeholder="请选择出发时间"
          clearable
        />
      </el-form-item>

      <el-form-item label="到达时间" prop="arrivedAt">
        <el-time-picker
          v-model="queryParams.arrivedAt"
          format="HH:mm:ss"
          value-format="HH:mm:ss"
          :picker-options='{"selectableRange":"00:00:00-23:59:59"}'
          :style="{width: '100%'}"
          placeholder="请选择到达时间"
          clearable
        />
      </el-form-item>
      <el-form-item label="换站截止时间" prop="changeExpiredAt">
        <el-time-picker
          v-model="queryParams.changeExpiredAt"
          format="HH:mm:ss"
          value-format="HH:mm:ss"
          :picker-options='{"selectableRange":"00:00:00-23:59:59"}'
          :style="{width: '100%'}"
          placeholder="请选择换站截止时间"
          clearable
        />
      </el-form-item>
      <el-form-item label="绑定车辆" prop="carIds">
        <el-input
          v-model="queryParams.carIds"
          placeholder="请选择车辆"
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
      <el-form-item label="0正常 1已删除" prop="isDelete">
        <el-input
          v-model="queryParams.isDelete"
          placeholder="请输入0正常 1已删除"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="" prop="deletedAt">
        <el-input
          v-model="queryParams.deletedAt"
          placeholder="请输入"
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
          v-permisaction="['scblines:scblines:add']"
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
        >新增
        </el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          v-permisaction="['scblines:scblines:edit']"
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
          v-permisaction="['scblines:scblines:remove']"
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
        >删除
        </el-button>
      </el-col>
    </el-row>

    <el-table v-loading="loading" :data="scblinesList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" /><el-table-column
        label=""
        align="center"
        prop="id"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="线路名称"
        align="center"
        prop="name"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="出发时间"
        align="center"
        prop="departed_at"
        :show-overflow-tooltip="true"
      />
      <el-table-column
        label="到达时间"
        align="center"
        prop="arrivedAt"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="换站截止时间"
        align="center"
        prop="changeExpiredAt"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="绑定的车辆"
        align="center"
        prop="carIds"
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
        label="0正常 1已删除"
        align="center"
        prop="isDelete"
        :show-overflow-tooltip="true"
      /><el-table-column
        label=""
        align="center"
        prop="deletedAt"
        :show-overflow-tooltip="true"
      />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            v-permisaction="['scblines:scblines:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
          >修改
          </el-button>
          <el-button
            v-permisaction="['scblines:scblines:remove']"
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

        <el-form-item label="线路名称" prop="name">
          <el-input
            v-model="form.name"
            placeholder="线路名称"
          />
        </el-form-item>
        <el-form-item label="出发时间" prop="departed_at">
          <el-time-picker
            v-model="queryParams.departed_at"
            format="HH:mm:ss"
            value-format="HH:mm:ss"
            :picker-options='{"selectableRange":"00:00:00-23:59:59"}'
            :style="{width: '100%'}"
            placeholder="请选择出发时间"
            clearable
          />
        </el-form-item>
        <el-form-item label="到达时间" prop="arrivedAt">
          <el-time-picker
            v-model="queryParams.arrivedAt"
            format="HH:mm:ss"
            value-format="HH:mm:ss"
            :picker-options='{"selectableRange":"00:00:00-23:59:59"}'
            :style="{width: '100%'}"
            placeholder="请选择到达时间"
            clearable
          />
        </el-form-item>


        <el-form-item label="换站截止时间" prop="changeExpiredAt">
          <el-time-picker
            v-model="queryParams.changeExpiredAt"
            format="HH:mm:ss"
            value-format="HH:mm:ss"
            :picker-options='{"selectableRange":"00:00:00-23:59:59"}'
            :style="{width: '100%'}"
            placeholder="请选择换站截止时间"
            clearable
          />
        </el-form-item>
        <el-form-item label="绑定的车辆">
          <el-tree
            ref="menu"
            :data="carsOptions"
            show-checkbox
            node-key="id"
            empty-text="加载中，请稍后"
            :props="defaultProps"
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
import { addScbLines, delScbLines, getScbLines, listScbLines, updateScbLines } from '@/api/scblines'
import { treeselect as carsTreeselect } from '@/api/scbcars'

export default {
  name: 'Config',
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
      scblinesList: [],
      // 车辆列表
      carsOptions: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        id:
            undefined,
        name:
            undefined,
        departed_at:
            undefined,
        arrivedAt:
            undefined,
        changeExpiredAt:
            undefined,
        carIds:
            undefined,
        createdAt:
            undefined,
        updatedAt:
            undefined,
        isDelete:
            undefined,
        deletedAt:
            undefined

      },
      // 表单参数
      form: {
      },
      defaultProps: {
        children: 'children',
        label: 'label'
      },
      // 表单校验
      rules: { id:
                [
                  { required: true, message: '不能为空', trigger: 'blur' }
                ],
      name:
                [
                  { required: true, message: '线路名称不能为空', trigger: 'blur' }
                ],
      departed_at:
                [
                  { required: false, message: '出发时间不能为空', trigger: 'change' }
                ],
      arrivedAt:
                [
                  { required: false, message: '到达时间不能为空', trigger: 'change' }
                ],
      changeExpiredAt:
                [
                  { required: false, message: '换站截止时间不能为空', trigger: 'change' }
                ],
      carIds:
                [
                  { required: true, message: '绑定的车辆不能为空', trigger: 'blur' }
                ],
      createdAt:
                [
                  { required: true, message: '创建时间不能为空', trigger: 'blur' }
                ],
      updatedAt:
                [
                  { required: true, message: '更新时间不能为空', trigger: 'blur' }
                ],
      isDelete:
                [
                  { required: true, message: '0正常 1已删除不能为空', trigger: 'blur' }
                ],
      deletedAt:
                [
                  { required: true, message: '不能为空', trigger: 'blur' }
                ]
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询菜单树结构 */
    getCarsTreeselect() {
      carsTreeselect().then(response => {
        this.carsOptions = response.data
      })
    },
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listScbLines(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.scblinesList = response.data.list
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
        departed_at: undefined,
        arrivedAt: undefined,
        changeExpiredAt: undefined,
        carIds: undefined,
        isDelete: undefined
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
      this.title = '添加线路管理'
      this.isEdit = false
      this.getCarsTreeselect()
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
      const id = row.id || this.ids
      getScbLines(id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改线路管理'
        this.isEdit = true
      })
      this.getCarsTreeselect()
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateScbLines(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addScbLines(this.form).then(response => {
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
        return delScbLines(Ids)
      }).then(() => {
        this.getList()
        this.msgSuccess('删除成功')
      }).catch(function() {
      })
    }
  }
}
</script>
