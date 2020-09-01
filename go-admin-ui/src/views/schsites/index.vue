
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
    <el-dialog :title="title" :visible.sync="open" width="500px">
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
            v-model="form.arrivedAt"
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
        <el-form-item label="详细地址：" prop="address">
          <el-autocomplete
            v-model="form.address"
            style="width:100%;"
            popper-class="autoAddressClass"
            :fetch-suggestions="querySearchAsync"
            :trigger-on-focus="false"
            placeholder="详细地址"
            clearable
            @select="handleSelect"
          >
            <template slot-scope="{ item }">
              <i class="el-icon-search fl mgr10" />
              <div style="overflow:hidden;">
                <div class="title">{{ item.title }}</div>
                <span class="address ellipsis">{{ item.address }}</span>
              </div>
            </template>
          </el-autocomplete>
        </el-form-item>
        <el-form-item label="地图定位：">
          <div id="map-container" style="width:100%;height:500px;" />
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
/* eslint-disable */
import { addSchSites, delSchSites, getSchSites, listSchSites, updateSchSites } from '@/api/schsites'
import Treeselect from '@riophae/vue-treeselect'
import { getAllLines } from '@/api/scblines'
import loadBMap from '@/utils/loadBMap'
import '@riophae/vue-treeselect/dist/vue-treeselect.css'
export default {
  name: 'Schsites',
  components: { Treeselect },
  data() {
    return {
      map: '', // 地图实例
      mk: '', // Marker实例
      locationPoint: null,
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
                  { required: true, message: '图片不能为空', trigger: 'blur' }
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
  async mounted() {
    await loadBMap('vTW8ITU3lpEfHCoZGV3StGLchcQwsjZA') // 加载引入BMap
    this.initMap()
  },
  methods: {
    initMap() {
      var that = this
      // 1、挂载地图
      this.map = new BMap.Map('map-container', { enableMapClick: false })
      console.log('this.map...1..')
      var point = new BMap.Point(113.3324436, 23.1315381)
      this.map.centerAndZoom(point, 12);
      // 3、设置图像标注并绑定拖拽标注结束后事件
      this.mk = new BMap.Marker(point, { enableDragging: true })
      this.map.addOverlay(this.mk)
      this.mk.addEventListener('dragend', function(e) {
        that.getAddrByPoint(e.point)
      })
      // 4、添加（右上角）平移缩放控件
      this.map.addControl(new BMap.NavigationControl({ anchor: BMAP_ANCHOR_TOP_RIGHT, type: BMAP_NAVIGATION_CONTROL_SMALL }))
      // 5、添加（左下角）定位控件
      var geolocationControl = new BMap.GeolocationControl({ anchor: BMAP_ANCHOR_BOTTOM_LEFT })
      geolocationControl.addEventListener('locationSuccess', function(e) {
        that.getAddrByPoint(e.point)
      })
      geolocationControl.addEventListener('locationError', function(e) {
        alert(e.message)
      })
      this.map.addControl(geolocationControl)
      // 6、浏览器定位
      this.geolocation()
      // 7、绑定点击地图任意点事件
      this.map.addEventListener('click', function(e) {
        that.getAddrByPoint(e.point)
      })
    },
    // 获取两点间的距离
    getDistancs(pointA, pointB) {
      return this.map.getDistance(pointA, pointB).toFixed(2)
    },
    // 浏览器定位函数
    geolocation() {
      var that = this
      var geolocation = new BMap.Geolocation()
      geolocation.getCurrentPosition(function(res) {
        if (this.getStatus() == BMAP_STATUS_SUCCESS) {
          that.getAddrByPoint(res.point)
          that.locationPoint = res.point
        } else {
          alert('failed' + this.getStatus())
          that.locationPoint = new BMap.Point(113.3324436, 23.1315381)
        }
      }, { enableHighAccuracy: true })
    },// 2、逆地址解析函数
    getAddrByPoint(point) {
      var that = this
      var geco = new BMap.Geocoder()
      geco.getLocation(point, function(res) {
        console.log(res)
        that.mk.setPosition(point)
        that.map.panTo(point)
        that.form.address = res.address
        that.form.addrPoint = point
      })
    },
    // 8-1、地址搜索
    querySearchAsync(str, cb) {
      var options = {
        onSearchComplete: function(res) {
          var s = []
          if (local.getStatus() == BMAP_STATUS_SUCCESS) {
            for (var i = 0; i < res.getCurrentNumPois(); i++) {
              s.push(res.getPoi(i))
            }
            cb(s)
          } else {
            cb(s)
          }
        }
      }
      var local = new BMap.LocalSearch(this.map, options)
      local.search(str)
    }, // 8-2、选择地址
    handleSelect(item) {
      this.form.address = item.address + item.title
      this.form.longitude = item.point.lng
      this.form.latitude = item.point.lat
      this.map.clearOverlays()
      this.mk = new BMap.Marker(item.point)
      this.map.addOverlay(this.mk)
      this.map.panTo(item.point)
    },
    // 以上baidu map

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
    }
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
    }
  }
}
</style>
