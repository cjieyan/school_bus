
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
      <el-form-item label="线路id" prop="lineId">
        <el-input
          v-model="queryParams.lineId"
          placeholder="请输入线路id"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="名称" prop="name">
        <el-input
          v-model="queryParams.name"
          placeholder="请输入名称"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="用途" prop="purpose">
        <el-input
          v-model="queryParams.purpose"
          placeholder="请输入用途"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="排序" prop="sort">
        <el-input
          v-model="queryParams.sort"
          placeholder="请输入排序"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="站点属性" prop="prop">
        <el-input
          v-model="queryParams.prop"
          placeholder="请输入站点属性"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="到达时间" prop="arriveAt">
        <el-input
          v-model="queryParams.arriveAt"
          placeholder="请输入到达时间"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="备注" prop="remark">
        <el-input
          v-model="queryParams.remark"
          placeholder="请输入备注"
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

      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          v-permisaction="['schsites:schsites:add']"
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
        >新增
        </el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          v-permisaction="['schsites:schsites:edit']"
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
          v-permisaction="['schsites:schsites:remove']"
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
        >删除
        </el-button>
      </el-col>
    </el-row>

    <el-table v-loading="loading" :data="schsitesList" @selection-change="handleSelectionChange">
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
        label="名称"
        align="center"
        prop="name"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="用途"
        align="center"
        prop="purpose"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="排序"
        align="center"
        prop="sort"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="站点属性"
        align="center"
        prop="prop"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="到达时间"
        align="center"
        prop="arriveAt"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="备注"
        align="center"
        prop="remark"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="经度"
        align="center"
        prop="longitude"
        :show-overflow-tooltip="true"
      /><el-table-column
        label="维度"
        align="center"
        prop="latitude"
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
      />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            v-permisaction="['schsites:schsites:edit']"
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
          >修改
          </el-button>
          <el-button
            v-permisaction="['schsites:schsites:remove']"
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
    <el-dialog :title="title" :visible.sync="open" width="800px">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="线路" prop="lineId">
          <treeselect
            v-model="form.lineId"
            :options="linesOptions"
            :normalizer="normalizer"
            :show-count="true"
            placeholder="选择线路"
            :is-disabled="isEdit"
          />
        </el-form-item>
        <el-form-item label="名称" prop="name">
          <el-input
            v-model="form.name"
            placeholder="名称"
          />
        </el-form-item>
        <el-form-item label="用途" prop="purpose">
          <el-input
            v-model="form.purpose"
            placeholder="用途"
          />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input
            v-model="form.sort"
            placeholder="排序"
          />
        </el-form-item>
        <el-form-item label="站点属性" prop="prop">
          <treeselect
            v-model="form.prop"
            :options="attrsOptions"
            :normalizer="normalizer"
            :show-count="true"
            placeholder="选择站点属性"
            :is-disabled="isEdit"
          />
        </el-form-item>
        <el-form-item label="到达时间" prop="arriveAt">
          <el-time-picker
            v-model="form.arriveAt"
            format="HH:mm:ss"
            value-format="HH:mm:ss"
            :picker-options="{'selectableRange':'00:00:00-23:59:59'}"
            :style="{width: '100%'}"
            placeholder="请选择到达时间"
            clearable
          />
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input
            v-model="form.remark"
            placeholder="备注"
          />
        </el-form-item>

        <el-form-item label="图片" prop="picture">
          <el-upload
            ref="picture"
            :file-list="picturefileList"
            :action="pictureAction"
            :auto-upload="false"
            :before-upload="pictureBeforeUpload"
            list-type="picture-card"
            accept="image/*"
            name="picture"
          >
            <i class="el-icon-plus" />
            <div slot="tip" class="el-upload__tip">只能上传不超过 2MB 的图片文件</div>
          </el-upload>
        </el-form-item>
        <el-form-item label="位置" prop="address">
          <el-input
            v-model="addressKeyword"
            placeholder="搜索地址"
            clearable
          />
          <el-input
            v-model="form.longitude"
            placeholder="经度"
            readonly
            style="width: 100px"
          />
          <el-input
            v-model="form.latitude"
            placeholder="经度"
            readonly
            style="width: 100px"
          />
          <el-input
            v-model="form.address"
            placeholder="详细地址"
            readonly
            style="width: 300px"
          />
          <div class="address">
            <!-- 给地图加点击事件getLocationPoint，点击地图获取位置相关的信息，经纬度啥的 -->
            <!-- scroll-wheel-zoom：是否可以用鼠标滚轮控制地图缩放，zoom是视图比例 -->
            <baidu-map
              class="bmView"
              :scroll-wheel-zoom="true"
              :center="location"
              :zoom="zoom"
              @click="getLocationPoint"
            >
              <bm-geolocation
                anchor="BMAP_ANCHOR_BOTTOM_RIGHT"
                :show-address-bar="true"
                :auto-location="true"
              />
              <bm-view style="width: 100%; height:400px; flex: 1" />
              <bm-local-search
                :keyword="addressKeyword"
                :panel="true"
                :page-capacity="5"
                :select-first-result="true"
                :auto-viewport="true"
              />
              <bm-navigation anchor="BMAP_ANCHOR_TOP_RIGHT" />
            </baidu-map>
          </div>
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
/* eslint-disable */
import { addSchSites, delSchSites, getSchSites, listSchSites, updateSchSites } from '@/api/schsites'
import Treeselect from '@riophae/vue-treeselect'
import { getAllLines } from '@/api/scblines'
import '@riophae/vue-treeselect/dist/vue-treeselect.css'
export default {
  name: 'Schsites',
  components: { Treeselect },
  data() {
    return {
      map: '', // 地图实例
      location: {
        lng: 116.404,
        lat: 39.915
      },
      zoom: 12.8,
      addressKeyword: "",
      pictureAction: 'https://jsonplaceholder.typicode.com/posts/',
      picturefileList: [],
      // 线路列表
      linesOptions: [],
      // 站点属性列表
      attrsOptions: [],
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
      schsitesList: [],

      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        id:
            undefined,
        lineId:
            undefined,
        name:
            undefined,
        purpose:
            undefined,
        sort:
            undefined,
        prop:
            undefined,
        arriveAt:
            undefined,
        remark:
            undefined,
        longitude:
            undefined,
        latitude:
            undefined,
        picture:
            undefined,
        createdAt:
            undefined,
        updatedAt:
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
      lineId:
                [
                  { required: true, message: '线路id不能为空', trigger: 'blur' }
                ],
      name:
                [
                  { required: true, message: '名称不能为空', trigger: 'blur' }
                ],
      purpose:
                [
                  { required: true, message: '用途不能为空', trigger: 'blur' }
                ],
      sort:
                [
                  { required: true, message: '排序不能为空', trigger: 'blur' }
                ],
      prop:
                [
                  { required: true, message: '站点属性不能为空', trigger: 'blur' }
                ],
      arriveAt:
                [
                  { required: true, message: '到达时间不能为空', trigger: 'blur' }
                ],
      remark:
                [
                  { required: true, message: '备注不能为空', trigger: 'blur' }
                ],
      longitude:
                [
                  { required: true, message: '经度不能为空', trigger: 'blur' }
                ],
      latitude:
                [
                  { required: true, message: '维度不能为空', trigger: 'blur' }
                ],
      picture:
                [
                  { required: false, message: '图片不能为空', trigger: 'blur' }
                ],
      createdAt:
                [
                  { required: true, message: '创建时间不能为空', trigger: 'blur' }
                ],
      updatedAt:
                [
                  { required: true, message: '更新时间不能为空', trigger: 'blur' }
                ]
      }
    }
  },
  created() {
    this.getList()
  },
  mounted() {
  },
  methods: {
    getLocationPoint(e) {
      this.lng = e.point.lng;
      this.lat = e.point.lat;
      /* 创建地址解析器的实例 */
      let geoCoder = new BMap.Geocoder();
      /* 获取位置对应的坐标 */
      geoCoder.getPoint(this.addressKeyword, point => {
        this.lng = point.lng;
        this.lat = point.lat;
      });
      /* 利用坐标获取地址的详细信息 */
      geoCoder.getLocation(e.point, res => {
        console.log('getLocation....')
        console.log(res)
        this.form.address = res.address
        // that.addressKeyword = res.address
        this.form.longitude = res.point.lng
        this.form.latitude = res.point.lat
      })
    },
    handler ({BMap, map}) {
      this.center.lng = 116.404
      this.center.lat = 39.915
      this.zoom = 15
    },
    // 8-1、地址搜索
    querySearchAsync(str, cb) {
      this.addressKeyword = str
      // var options = {
      //   onSearchComplete: function(res) {
      //     var s = []
      //     if (local.getStatus() == BMAP_STATUS_SUCCESS) {
      //       for (var i = 0; i < res.getCurrentNumPois(); i++) {
      //         s.push(res.getPoi(i))
      //       }
      //       cb(s)
      //     } else {
      //       cb(s)
      //     }
      //   }
      // }
      // var local = new BMap.LocalSearch(this.map, options)
      // local.search(str)
    },
    /***百度地图***/


    pictureBeforeUpload(file) {
      const isRightSize = file.size / 1024 / 1024 < 2
      if (!isRightSize) {
        this.$message.error('文件大小超过 2MB')
      }
      return isRightSize
    },
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
      listSchSites(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.schsitesList = response.data.list
        this.total = response.data.count
        this.loading = false
      }
      )
    },
    /** 查询班级下拉树结构 */
    getTreeselect(e) {
      getAllLines().then(response => {
        this.linesOptions = []
        const lines = { id: 0, name: '请选择', children: [] }
        lines.children = response.data
        this.linesOptions.push(lines)
      })
    },
    getSitesAttrs(e) {
      this.attrsOptions = []
      const attrs = { id: 0, name: '请选择', children: [] }
      attrs.children = [{ id: 1, name: '上车' }, { id: 2, name: '下车' }]
      this.attrsOptions.push(attrs)
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
        name: undefined,
        purpose: undefined,
        sort: undefined,
        prop: undefined,
        arriveAt: undefined,
        remark: undefined,
        longitude: undefined,
        latitude: undefined,
        picture: undefined,
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
      this.title = '添加站点管理'
      this.isEdit = false
      this.getTreeselect('add')
      this.getSitesAttrs('add')
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
      getSchSites(id).then(response => {
        this.form = response.data
        this.open = true
        this.title = '修改站点管理'
        this.isEdit = true
      })
      this.getTreeselect('update')
      this.getSitesAttrs('update')
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          if (this.form.id !== undefined) {
            updateSchSites(this.form).then(response => {
              if (response.code === 200) {
                this.msgSuccess('修改成功')
                this.open = false
                this.getList()
              } else {
                this.msgError(response.msg)
              }
            })
          } else {
            addSchSites(this.form).then(response => {
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
        return delSchSites(Ids)
      }).then(() => {
        this.getList()
        this.msgSuccess('删除成功')
      }).catch(function() {
      })
    },
  }
}
</script>
<style lang="scss" scoped>
.autoAddressClass{
  li {
    i.el-icon-search {margin-top:11px;}
    .mgr10 {margin-right: 10px;}
    .title {
      text-overflow: ellipsis;
      overflow: hidden;
    }
    .address {
      line-height: 1;
      font-size: 12px;
      color: #b4b4b4;
      margin-bottom: 5px;
      height: 200px;
    }
  }
}
</style>
