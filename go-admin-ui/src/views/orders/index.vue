
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
                        v-permisaction="['orders:orders:add']"
                        type="primary"
                        icon="el-icon-plus"
                        size="mini"
                        @click="handleAdd"
                >新增
                </el-button>
            </el-col>
            <el-col :span="1.5">
                <el-button
                        v-permisaction="['orders:orders:edit']"
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
                        v-permisaction="['orders:orders:remove']"
                        type="danger"
                        icon="el-icon-delete"
                        size="mini"
                        :disabled="multiple"
                        @click="handleDelete"
                >删除
                </el-button>
            </el-col>
        </el-row>

        <el-table v-loading="loading" :data="ordersList" @selection-change="handleSelectionChange">
            <el-table-column type="selection" width="55" align="center"/>
            <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                    <el-button
                            v-permisaction="['orders:orders:edit']"
                            size="mini"
                            type="text"
                            icon="el-icon-edit"
                            @click="handleUpdate(scope.row)"
                    >修改
                    </el-button>
                    <el-button
                            v-permisaction="['orders:orders:remove']"
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
                
                            <el-form-item label="订单号" prop="orderNo">
                                <el-input v-model="form.orderNo" placeholder="订单号"
                                              />
                            </el-form-item>
                            <el-form-item label="第三方流水号" prop="tradeNo">
                                <el-input v-model="form.tradeNo" placeholder="第三方流水号"
                                              />
                            </el-form-item>
                            <el-form-item label="用户id" prop="userId">
                                <el-input v-model="form.userId" placeholder="用户id"
                                              />
                            </el-form-item>
                            <el-form-item label="商品id" prop="goodsId">
                                <el-input v-model="form.goodsId" placeholder="商品id"
                                              />
                            </el-form-item>
                            <el-form-item label="查询次数" prop="carTimes">
                                <el-input v-model="form.carTimes" placeholder="查询次数"
                                              />
                            </el-form-item>
                            <el-form-item label="订单金额" prop="orderAmount">
                                <el-input v-model="form.orderAmount" placeholder="订单金额"
                                              />
                            </el-form-item>
                            <el-form-item label="总金额" prop="totalAmount">
                                <el-input v-model="form.totalAmount" placeholder="总金额"
                                              />
                            </el-form-item>
                            <el-form-item label="优惠金额" prop="discountAmount">
                                <el-input v-model="form.discountAmount" placeholder="优惠金额"
                                              />
                            </el-form-item>
                            <el-form-item label="实际支付金额" prop="actualAmount">
                                <el-input v-model="form.actualAmount" placeholder="实际支付金额"
                                              />
                            </el-form-item>
                            <el-form-item label="0支付宝, 1微信支付" prop="payType">
                                <el-input v-model="form.payType" placeholder="0支付宝, 1微信支付"
                                              />
                            </el-form-item>
                            <el-form-item label="0充值 1消费" prop="type">
                                <el-input v-model="form.type" placeholder="0充值 1消费"
                                              />
                            </el-form-item>
                            <el-form-item label="优惠券id" prop="couponId">
                                <el-input v-model="form.couponId" placeholder="优惠券id"
                                              />
                            </el-form-item>
                            <el-form-item label="0:未支付 10:已支付 20:订单作废 30:已发货 40: 已退款" prop="status">
                                <el-input v-model="form.status" placeholder="0:未支付 10:已支付 20:订单作废 30:已发货 40: 已退款"
                                              />
                            </el-form-item>
                            <el-form-item label="支付时间" prop="paidAt">
                                <el-input v-model="form.paidAt" placeholder="支付时间"
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
    import {addOrders, delOrders, getOrders, listOrders, updateOrders} from '@/api/orders'

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
                ordersList: [],
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
            
        },
            // 表单参数
            form: {
            }
        ,
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
                listOrders(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.ordersList = response.data.list
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
                orderNo: undefined,
                tradeNo: undefined,
                userId: undefined,
                goodsId: undefined,
                carTimes: undefined,
                orderAmount: undefined,
                totalAmount: undefined,
                discountAmount: undefined,
                actualAmount: undefined,
                payType: undefined,
                type: undefined,
                couponId: undefined,
                status: undefined,
                paidAt: undefined,
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
                this.title = '添加订单表'
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
                getOrders(id).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改订单表'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateOrders(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess('修改成功')
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addOrders(this.form).then(response => {
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
                    return delOrders(Ids)
                }).then(() => {
                    this.getList()
                    this.msgSuccess('删除成功')
                }).catch(function () {
                })
            }
        }
    }
</script>
