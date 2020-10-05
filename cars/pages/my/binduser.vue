<template>
	<view>
		<u-navbar back-text="返回" title="绑定手机"></u-navbar>
		<u-image width="100%" height="300rpx" src="../../static/login-banner.jpg"></u-image>
		<view class="page-content">
			<u-form ref="uForm">
				<u-form-item label="手机号码" prop="phone" label-width="150">
					<u-input :border="border" placeholder="请输入手机号" v-model="model.phone" type="number"></u-input>
				</u-form-item>
				<u-form-item label="验证码" prop="code" label-width="150">
					<u-input placeholder="请输入验证码" type="text"  v-model="model.code"></u-input>
					<u-button slot="right" type="success" size="mini" @click="getCode">{{codeTips}}</u-button>
				</u-form-item>
				<u-verification-code seconds="60" ref="uCode" @change="codeChange"></u-verification-code>
				<u-button type="primary" class="btn" shape="square" size="medium" @click="bindUser">绑定</u-button>
			</u-form>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				border: false,
				codeTips: '',
				model: {
					name: '',
					sex: '',
					likeFruit: '',
					intro: '',
					payType: '支付宝',
					agreement: false,
					region: '',
					goodsType: '',
					phone: '',
					code: '',
					password: '',
					rePassword: '',
					remember: false,
					photo: ''
				},
				rules: {
					name: [{
							required: true,
							message: '请输入姓名',
							trigger: 'blur',
						},
						{
							min: 3,
							max: 5,
							message: '姓名长度在3到5个字符',
							trigger: ['change', 'blur'],
						},
						{
							// 此为同步验证，可以直接返回true或者false，如果是异步验证，稍微不同，见下方说明
							validator: (rule, value, callback) => {
								// 调用uView自带的js验证规则，详见：https://www.uviewui.com/js/test.html
								return this.$u.test.chinese(value);
							},
							message: '姓名必须为中文',
							// 触发器可以同时用blur和change，二者之间用英文逗号隔开
							trigger: ['change', 'blur'],
						},
						// 异步验证，用途：比如用户注册时输入完账号，后端检查账号是否已存在
						// {
						// 	trigger: ['blur'],
						// 	// 异步验证需要通过调用callback()，并且在里面抛出new Error()
						// 	// 抛出的内容为需要提示的信息，和其他方式的message属性的提示一样
						// 	asyncValidator: (rule, value, callback) => {
						// 		this.$u.post('/ebapi/public_api/index').then(res => {
						// 			// 如果验证出错，需要在callback()抛出new Error('错误提示信息')
						// 			if(res.error) {
						// 				callback(new Error('姓名重复'));
						// 			} else {
						// 				// 如果没有错误，也要执行callback()回调
						// 				callback();
						// 			}
						// 		})
						// 	},
						// }
					],
					sex: [{
						required: true,
						message: '请选择性别',
						trigger: 'change'
					}, ],
					intro: [{
							required: true,
							message: '请填写简介'
						},
						{
							min: 5,
							message: '简介不能少于5个字',
							trigger: 'change',
						},
						// 正则校验示例，此处用正则校验是否中文，此处仅为示例，因为uView有this.$u.test.chinese可以判断是否中文
						{
							pattern: /^[\u4e00-\u9fa5]+$/gi,
							message: '简介只能为中文',
							trigger: 'change',
						},
					],
					likeFruit: [{
						required: true,
						message: '请选择您喜欢的水果',
						trigger: 'change',
						type: 'array',
					}],
					payType: [{
						required: true,
						message: '请选择任意一种支付方式',
						trigger: 'change',
					}],
					region: [{
						required: true,
						message: '请选择地区',
						trigger: 'change',
					}],
					goodsType: [{
						required: true,
						message: '请选择商品类型',
						trigger: 'change',
					}],
					phone: [{
							required: true,
							message: '请输入手机号',
							trigger: ['change', 'blur'],
						},
						{
							validator: (rule, value, callback) => {
								// 调用uView自带的js验证规则，详见：https://www.uviewui.com/js/test.html
								return this.$u.test.mobile(value);
							},
							message: '手机号码不正确',
							// 触发器可以同时用blur和change，二者之间用英文逗号隔开
							trigger: ['change', 'blur'],
						}
					],
					code: [{
							required: true,
							message: '请输入验证码',
							trigger: ['change', 'blur'],
						},
						{
							type: 'number',
							message: '验证码只能为数字',
							trigger: ['change', 'blur'],
						}
					],
					password: [{
							required: true,
							message: '请输入密码',
							trigger: ['change', 'blur'],
						},
						{
							// 正则不能含有两边的引号
							pattern: /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]+\S{5,12}$/,
							message: '需同时含有字母和数字，长度在6-12之间',
							trigger: ['change', 'blur'],
						}
					],
					rePassword: [{
							required: true,
							message: '请重新输入密码',
							trigger: ['change', 'blur'],
						},
						{
							validator: (rule, value, callback) => {
								return value === this.model.password;
							},
							message: '两次输入的密码不相等',
							trigger: ['change', 'blur'],
						}
					],
				},
			}
		},
		onReady() {
			this.$refs.uForm.setRules(this.rules);
		},
		methods: {
			codeChange(text) {
				this.codeTips = text;
			},
			bindUser: function() {
				uni.showLoading({
					title: "加载中",
					icon: "none"
				})
				uni.redirectTo({
					url: "./user"
				})
				uni.hideLoading()
			},
			getCode() {
				console.log(this.$refs)
				if (this.$refs.uCode.canGetCode) {
					// 模拟向后端请求验证码
					uni.showLoading({
						title: '正在获取验证码',
						mask: true
					})
					setTimeout(() => {
						uni.hideLoading();
						// 这里此提示会被this.start()方法中的提示覆盖
						this.$u.toast('验证码已发送');
						// 通知验证码组件内部开始倒计时
						this.$refs.uCode.start();
					}, 2000);
				} else {
					this.$u.toast('倒计时结束后再发送');
				}
			},
		}
	}
</script>

<style lang="css">
	page {
		background-color: #fff;
	}

	.page-content {
		padding: 20px 20px;
		margin: 0 auto;
		text-align: center;
	}

	.btn {
		margin-top: 20px;
		background: linear-gradient(to right, rgba(8, 189, 175), rgb(247, 218, 99));
	}
	.u-form-item__message{
		text-align: left !important;
	}
</style>
