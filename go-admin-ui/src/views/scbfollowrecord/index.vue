
<template>
  <div class="app-container">
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
      <el-form-item label="线路id" prop="lineId">
        <el-input
          v-model="queryParams.lineId"
          placeholder="请输入线路id"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="跟车员 关联teachers表的id	" prop="attendantId">
        <el-input
          v-model="queryParams.attendantId"
          placeholder="请输入跟车员 关联teachers表的id	"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="未上车" prop="unGetOn">
        <el-input
          v-model="queryParams.unGetOn"
          placeholder="请输入未上车"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="请假" prop="leave">
        <el-input
          v-model="queryParams.leave"
          placeholder="请输入请假"
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
          v-permisaction="['scbfollowrecord:scbfollowrecord:add']"
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
        >新增
        </el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          v-permisaction="['scbfollowrecord:scbfollowrecord:edit']"
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
          v-permisaction="['scbfollowrecord:scbfollowrecord:remove']"
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
        >删除
        </el-button>
      </el-col>
    </el-row>

    <el-table v-loading="loading" :data="scbfollowrecordList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" /><el-table-column
        label=""
        align="center"
        prop="id"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="线路id"
        align="center"
        prop="lineId"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="车辆id"
        align="center"
        prop="carId"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="跟车员 关联teachers表的id	"
        align="center"
        prop="attendantId"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="所有人数"
        align="center"
        prop="allCount"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="已下车"
        align="center"
        prop="getOff"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="已上车"
        align="center"
        prop="getOn"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="未上车"
        align="center"
        prop="unGetOn"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="请假"
        align="center"
        prop="leave"
        :show-overflow-tooltip="true"
      />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            v-permisaction="['scbfollowrecord:scbfollowrecord:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
          >修改
          </el-button>
          <el-button
            v-permisaction="['scbfollowrecord:scbfollowrecord:remove']"
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

        <el-form-item label="线路id" prop="lineId">
          <el-select
            v-model="form.lineId"
            placeholder="请选择"
          >
            <el-option
              v-for="dict in lineIdOptions"
              :key="dict.dictValue"
              :label="dict.dictLabel"
              :value="dict.dictValue"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="车辆id" prop="carId">
          <el-input
            v-model="form.carId"
            placeholder="车辆id"
          />
        </el-form-item>
        <el-form-item label="跟车员 关联teachers表的id	" prop="attendantId">
          <el-select
            v-model="form.attendantId"
            placeholder="请选择"
          >
            <el-option
              v-for="dict in attendantIdOptions"
              :key="dict.dictValue"
              :label="dict.dictLabel"
              :value="dict.dictValue"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="所有人数" prop="allCount">
          <el-input
            v-model="form.allCount"
            placeholder="所有人数"
          />
        </el-form-item>
        <el-form-item label="已下车" prop="getOff">
          <el-input
            v-model="form.getOff"
            placeholder="已下车"
          />
        </el-form-item>
        <el-form-item label="已上车" prop="getOn">
          <el-input
            v-model="form.getOn"
            placeholder="已上车"
          />
        </el-form-item>
        <el-form-item label="未上车" prop="unGetOn">
          <el-input
            v-model="form.unGetOn"
            placeholder="未上车"
          />
        </el-form-item>
        <el-form-item label="请假" prop="leave">
          <el-input
            v-model="form.leave"
            placeholder="请假"
          />
        </el-form-item>
        <el-form-item label="0未删除 1已删除" prop="isDelete">
          <el-input
            v-model="form.isDelete"
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
import { addScbFollowRecord, delScbFollowRecord, getScbFollowRecord, listScbFollowRecord, updateScbFollowRecord } from '@/api/scbfollowrecord'

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
      scbfollowrecordList: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        lineId:
            undefined,
        attendantId:
            undefined,
        unGetOn:
            undefined,
        leave:
            undefined

      },
      // 表单参数
      form: {
      },
      // 表单校验
      rules: { lineId:
                [
                  { required: true, message: '线路id不能为空', trigger: 'blur' }
                ],
      attendantId:
                [
                  { required: true, message: '跟车员 关联teachers表的id	不能为空', trigger: 'blur' }
                ],
      unGetOn:
                [
                  { required: true, message: '未上车不能为空', trigger: 'blur' }
                ],
      leave:
                [
                  { required: true, message: '请假不能为空', trigger: 'blur' }
                ]
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listScbFollowRecord(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.scbfollowrecordList = response.data.list
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
        lineId: undefined,
        carId: undefined,
        attendantId: undefined,
        allCount: undefined,
        getOff: undefined,
        getOn: undefined,
        unGetOn: undefined,
        leave: undefined,
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
      this.title = '添加ScbFollowRecord'
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
      const id =
                row.id || this.ids
      getScbFollowRecord(id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改ScbFollowRecord'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateScbFollowRecord(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addScbFollowRecord(this.form).then(response => {
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
        return delScbFollowRecord(Ids)
      }).then(() => {
        this.getList()
        this.msgSuccess('删除成功')
      }).catch(function() {
      })
    }
  }
}
</script>
