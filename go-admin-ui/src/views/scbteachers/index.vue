
<template>
    <div class="app-container">
        <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
            <el-form-item label="" prop="id">
                        <el-input v-model="queryParams.id" placeholder="请输入" clearable
                                      size="small" @keyup.enter.native="handleQuery"/>
                    </el-form-item>
                <el-form-item label="名称" prop="name">
                        <el-input v-model="queryParams.name" placeholder="请输入名称" clearable
                                      size="small" @keyup.enter.native="handleQuery"/>
                    </el-form-item>
                <el-form-item label="手机号" prop="phone">
                        <el-input v-model="queryParams.phone" placeholder="请输入手机号" clearable
                                      size="small" @keyup.enter.native="handleQuery"/>
                    </el-form-item>
                <el-form-item label="班级id" prop="classId">
                        <el-input v-model="queryParams.classId" placeholder="请输入班级id" clearable
                                      size="small" @keyup.enter.native="handleQuery"/>
                    </el-form-item>
                <el-form-item label="岗位id" prop="postId">
                        <el-input v-model="queryParams.postId" placeholder="请输入岗位id" clearable
                                      size="small" @keyup.enter.native="handleQuery"/>
                    </el-form-item>
                <el-form-item label="0未删除 1已删除" prop="isDeleted">
                        <el-input v-model="queryParams.isDeleted" placeholder="请输入0未删除 1已删除" clearable
                                      size="small" @keyup.enter.native="handleQuery"/>
                    </el-form-item>
                
            <el-form-item>
                <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
            </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
            <el-col :span="1.5">
                <el-button
                        v-permisaction="['scbteachers:scbteachers:add']"
                        type="primary"
                        icon="el-icon-plus"
                        size="mini"
                        @click="handleAdd"
                >新增
                </el-button>
            </el-col>
            <el-col :span="1.5">
                <el-button
                        v-permisaction="['scbteachers:scbteachers:edit']"
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
                        v-permisaction="['scbteachers:scbteachers:remove']"
                        type="danger"
                        icon="el-icon-delete"
                        size="mini"
                        :disabled="multiple"
                        @click="handleDelete"
                >删除
                </el-button>
            </el-col>
        </el-row>

        <el-table v-loading="loading" :data="scbteachersList" @selection-change="handleSelectionChange">
            <el-table-column type="selection" width="55" align="center"/><el-table-column label="" align="center" prop="id"
                                         :show-overflow-tooltip="true"/><el-table-column label="名称" align="center" prop="name"
                                         :show-overflow-tooltip="true"/><el-table-column label="手机号" align="center" prop="phone"
                                         :show-overflow-tooltip="true"/><el-table-column label="班级id" align="center" prop="classId"
                                         :show-overflow-tooltip="true"/><el-table-column label="岗位id" align="center" prop="postId"
                                         :show-overflow-tooltip="true"/><el-table-column label="备注" align="center" prop="remark"
                                         :show-overflow-tooltip="true"/><el-table-column label="创建时间" align="center" prop="createdAt"
                                         :show-overflow-tooltip="true"/><el-table-column label="更新时间" align="center" prop="updatedAt"
                                         :show-overflow-tooltip="true"/><el-table-column label="0未删除 1已删除" align="center" prop="isDeleted"
                                         :show-overflow-tooltip="true"/>
            <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                    <el-button
                            v-permisaction="['scbteachers:scbteachers:edit']"
                            size="mini"
                            type="text"
                            icon="el-icon-edit"
                            @click="handleUpdate(scope.row)"
                    >修改
                    </el-button>
                    <el-button
                            v-permisaction="['scbteachers:scbteachers:remove']"
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
                                <el-input v-model="form.name" placeholder="名称"
                                              />
                            </el-form-item>
                            <el-form-item label="手机号" prop="phone">
                                <el-input v-model="form.phone" placeholder="手机号"
                                              />
                            </el-form-item>
                            <el-form-item label="班级id" prop="classId">
                                <el-input v-model="form.classId" placeholder="班级id"
                                              />
                            </el-form-item>
                            <el-form-item label="岗位id" prop="postId">
                                <el-input v-model="form.postId" placeholder="岗位id"
                                              />
                            </el-form-item>
                            <el-form-item label="备注" prop="remark">
                                <el-input v-model="form.remark" placeholder="备注"
                                              />
                            </el-form-item>
                            <el-form-item label="0未删除 1已删除" prop="isDeleted">
                                <el-input v-model="form.isDeleted" placeholder="0未删除 1已删除"
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
    import {addScbTeachers, delScbTeachers, getScbTeachers, listScbTeachers, updateScbTeachers} from '@/api/scbteachers'

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
                scbteachersList: [],
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
            id:
            undefined,
            name:
            undefined,
            phone:
            undefined,
            classId:
            undefined,
            postId:
            undefined,
            isDeleted:
            undefined,
            
        },
            // 表单参数
            form: {
            }
        ,
            // 表单校验
            rules: {id:
                [
                    {required: true, message: '不能为空', trigger: 'blur'}
                ],
                name:
                [
                    {required: true, message: '名称不能为空', trigger: 'blur'}
                ],
                phone:
                [
                    {required: true, message: '手机号不能为空', trigger: 'blur'}
                ],
                classId:
                [
                    {required: true, message: '班级id不能为空', trigger: 'blur'}
                ],
                postId:
                [
                    {required: true, message: '岗位id不能为空', trigger: 'blur'}
                ],
                isDeleted:
                [
                    {required: true, message: '0未删除 1已删除不能为空', trigger: 'blur'}
                ],
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
                listScbTeachers(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.scbteachersList = response.data.list
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
                phone: undefined,
                classId: undefined,
                postId: undefined,
                remark: undefined,
                isDeleted: undefined,
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
                this.title = '添加教职员工表'
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
                getScbTeachers(id).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改教职员工表'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateScbTeachers(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess('修改成功')
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addScbTeachers(this.form).then(response => {
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
                }).then(function () {
                    return delScbTeachers(Ids)
                }).then(() => {
                    this.getList()
                    this.msgSuccess('删除成功')
                }).catch(function () {
                })
            }
        }
    }
</script>
