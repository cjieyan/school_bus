
<template>
  <div class="app-container">
    <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">

      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          v-permisaction="['akcars:akcars:add']"
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
        >新增
        </el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          v-permisaction="['akcars:akcars:edit']"
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
          v-permisaction="['akcars:akcars:remove']"
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
        >删除
        </el-button>
      </el-col>
    </el-row>

    <el-table v-loading="loading" :data="akcarsList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            v-permisaction="['akcars:akcars:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
          >修改
          </el-button>
          <el-button
            v-permisaction="['akcars:akcars:remove']"
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

        <el-form-item label="" prop="userId">
          <el-input
            v-model="form.userId"
            placeholder=""
          />
        </el-form-item>
        <el-form-item label="车辆类型" prop="cartype">
          <el-select
            v-model="form.cartype"
            placeholder="请选择"
          >
            <el-option
              v-for="dict in cartypeOptions"
              :key="dict.dictValue"
              :label="dict.dictLabel"
              :value="dict.dictValue"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="车牌号" prop="carno">
          <el-input
            v-model="form.carno"
            placeholder="车牌号"
          />
        </el-form-item>
        <el-form-item label="车架号" prop="frameno">
          <el-input
            v-model="form.frameno"
            placeholder="车架号"
          />
        </el-form-item>
        <el-form-item label="发动机号" prop="engineno">
          <el-input
            v-model="form.engineno"
            placeholder="发动机号"
          />
        </el-form-item>
        <el-form-item label="用途" prop="purpose">
          <el-input
            v-model="form.purpose"
            placeholder="用途"
          />
        </el-form-item>
        <el-form-item label="公司" prop="company">
          <el-input
            v-model="form.company"
            placeholder="公司"
          />
        </el-form-item>
        <el-form-item label="" prop="driver">
          <el-input
            v-model="form.driver"
            placeholder=""
          />
        </el-form-item>
        <el-form-item label="司机号码" prop="driverPhone">
          <el-input
            v-model="form.driverPhone"
            placeholder="司机号码"
          />
        </el-form-item>
        <el-form-item label="" prop="team">
          <el-input
            v-model="form.team"
            placeholder=""
          />
        </el-form-item>
        <el-form-item label="" prop="remark">
          <el-input
            v-model="form.remark"
            placeholder=""
          />
        </el-form-item>
        <el-form-item label="状态0正常 1禁用" prop="status">
          <el-input
            v-model="form.status"
            placeholder="状态0正常 1禁用"
          />
        </el-form-item>
        <el-form-item label="车牌号md5" prop="carnoMd5">
          <el-input
            v-model="form.carnoMd5"
            placeholder="车牌号md5"
          />
        </el-form-item>
        <el-form-item label="最后一次更新时间" prop="renewAt">
          <el-input
            v-model="form.renewAt"
            placeholder="最后一次更新时间"
          />
        </el-form-item>
        <el-form-item label="查询成功的年月" prop="succMonth">
          <el-input
            v-model="form.succMonth"
            placeholder="查询成功的年月"
          />
        </el-form-item>
        <el-form-item label="0 无违章 1有违章" prop="illegalStatus">
          <el-input
            v-model="form.illegalStatus"
            placeholder="0 无违章 1有违章"
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
import { addAkCars, delAkCars, getAkCars, listAkCars, updateAkCars } from '@/api/akcars'

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
      akcarsList: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10

      },
      // 表单参数
      form: {
      },
      // 表单校验
      rules: {}
    }
  },
  created() {
    this.getList()
  },
  methods: {
    /** 查询参数列表 */
    getList() {
      this.loading = true
      listAkCars(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.akcarsList = response.data.list
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
        userId: undefined,
        cartype: undefined,
        carno: undefined,
        frameno: undefined,
        engineno: undefined,
        purpose: undefined,
        company: undefined,
        driver: undefined,
        driverPhone: undefined,
        team: undefined,
        remark: undefined,
        status: undefined,
        carnoMd5: undefined,
        renewAt: undefined,
        succMonth: undefined,
        illegalStatus: undefined
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
      this.title = '添加车辆信息'
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
      getAkCars(id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改车辆信息'
        this.isEdit = true
      })
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateAkCars(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addAkCars(this.form).then(response => {
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
        return delAkCars(Ids)
      }).then(() => {
        this.getList()
        this.msgSuccess('删除成功')
      }).catch(function() {
      })
    }
  }
}
</script>
