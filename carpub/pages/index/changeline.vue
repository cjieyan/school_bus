<template>
	<view class="u-demo">
		<view class="top">
			<u-navbar :is-back="isback" title="智慧校车" @click="back" class="top" :background="background" back-icon-color="#fff"
			 title-color="#fff"></u-navbar>
		</view>
		<view class="u-demo-wrap">
			<u-form :model="model" :rules="rules" ref="uForm" :errorType="errorType">
				<u-form-item :leftIconStyle="{color: '#888', fontSize: '32rpx'}" label-width="240"
				 :label-position="labelPosition" label="姓名" prop="name">
					<text>李明</text>
				</u-form-item>
				<u-form-item :leftIconStyle="{color: '#888', fontSize: '32rpx'}" label-width="240"
				 :label-position="labelPosition" label="现乘车路线" prop="name">
					<text>1号线（粤B1234）</text>
				</u-form-item>
				<u-form-item :leftIconStyle="{color: '#888', fontSize: '32rpx'}" label-width="240"
				 :label-position="labelPosition" label="现上车站点" prop="name">
					<text>硅谷人才公寓</text>
				</u-form-item>
				<u-form-item :leftIconStyle="{color: '#888', fontSize: '32rpx'}" label-width="240"
				 :label-position="labelPosition" label="现下车站点" prop="name">
					<text>西丽湖地铁站</text>
				</u-form-item>
				<u-form-item :label-position="labelPosition" label="更换路线" label-width="240" prop="sex">
					<u-input :border="border" type="select" :select-open="actionSheetShow" v-model="model.sex" placeholder="请选择性别"
					 @click="actionSheetShow = true"></u-input>
				</u-form-item>
				<u-form-item :label-position="labelPosition" label="上车站点" label-width="240" prop="sex">
					<u-input :border="border" type="select" :select-open="actionSheetShow" v-model="model.sex" placeholder="请选择性别"
					 @click="actionSheetShow = true"></u-input>
				</u-form-item>
				<u-form-item :label-position="labelPosition" label="下车站点" label-width="240" prop="sex">
					<u-input :border="border" type="select" :select-open="actionSheetShow" v-model="model.sex" placeholder="请选择性别"
					 @click="actionSheetShow = true"></u-input>
				</u-form-item>
				<u-form-item :label-position="labelPosition" label="备注" label-width="240" prop="intro">
					<u-input type="textarea" :border="border" placeholder="请填写备注" v-model="model.intro" />
				</u-form-item>
				<u-form-item :label-position="labelPosition" label="更换方式" label-width="240" prop="payType">
					<u-radio-group v-model="radio" @change="radioGroupChange" :width="radioCheckWidth" :wrap="radioCheckWrap">
						<u-radio shape="circle" v-for="(item, index) in radioList" :key="index" :name="item.name">{{ item.name }}</u-radio>
					</u-radio-group>
				</u-form-item>
			</u-form>
			<u-button @click="submit">提交</u-button>
		</view>
		<u-action-sheet :list="actionSheetList" v-model="actionSheetShow" @click="actionSheetCallback"></u-action-sheet>
		<u-select mode="single-column" :list="selectList" v-model="selectShow" @confirm="selectConfirm"></u-select>
		<u-picker mode="region" v-model="pickerShow" @confirm="regionConfirm"></u-picker>
		<u-verification-code seconds="60" ref="uCode" @change="codeChange"></u-verification-code>
	</view>
</template>

<script>
	export default {
		data() {
			let that = this;
			return {
				isback: false,
				background: {
					backgroundColor: '#12c497',
				},
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
				selectList: [{
						value: '电子产品',
						label: '电子产品'
					},
					{
						value: '服装',
						label: '服装'
					},
					{
						value: '工艺品',
						label: '工艺品'
					}
				],
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
				border: false,
				check: false,
				selectStatus: 'close',
				checkboxList: [{
						name: '荔枝',
						checked: false,
						disabled: false
					},
					{
						name: '香蕉',
						checked: false,
						disabled: false
					},
					{
						name: '橙子',
						checked: false,
						disabled: false
					},
					{
						name: '草莓',
						checked: false,
						disabled: false
					}
				],
				radioList: [{
						name: '临时更换',
						checked: true,
						disabled: false
					},
					{
						name: '永久更换',
						checked: false,
						disabled: false
					}
				],
				radio: '支付宝',
				actionSheetList: [{
						text: '1号线（粤B1234）'
					},
					{
						text: '2号线（粤B5678）'
					}
				],
				actionSheetShow: false,
				pickerShow: false,
				selectShow: false,
				radioCheckWidth: 'auto',
				radioCheckWrap: false,
				labelPosition: 'left',
				codeTips: '',
				errorType: ['message'],
			};
		},
		onLoad() {

		},
		computed: {
			borderCurrent() {
				return this.border ? 0 : 1;
			}
		},
		onReady() {
			this.$refs.uForm.setRules(this.rules);
		},
		methods: {
			submit() {
				this.$refs.uForm.validate(valid => {
					if (valid) {
						if (!this.model.agreement) return this.$u.toast('请勾选协议');
						console.log('验证通过');
					} else {
						console.log('验证失败');
					}
				});
			},
			// 点击actionSheet回调
			actionSheetCallback(index) {
				uni.hideKeyboard();
				this.model.sex = this.actionSheetList[index].text;
			},
			// checkbox选择发生变化
			checkboxGroupChange(e) {
				this.model.likeFruit = e;
			},
			// radio选择发生变化
			radioGroupChange(e) {
				this.model.payType = e;
			},
			// 勾选版权协议
			checkboxChange(e) {
				this.model.agreement = e.value;
			},
			// 选择地区回调
			regionConfirm(e) {
				this.model.region = e.province.label + '-' + e.city.label + '-' + e.area.label;
			},
			// 选择商品类型回调
			selectConfirm(e) {
				this.model.goodsType = '';
				e.map((val, index) => {
					this.model.goodsType += this.model.goodsType == '' ? val.label : '-' + val.label;
				})
			},
			borderChange(index) {
				this.border = !index;
			},
			radioCheckboxChange(index) {
				if (index == 0) {
					this.radioCheckWrap = false;
					this.radioCheckWidth = 'auto';
				} else if (index == 1) {
					this.radioCheckWrap = true;
					this.radioCheckWidth = 'auto';
				} else if (index == 2) {
					this.radioCheckWrap = false;
					this.radioCheckWidth = '50%';
				}
			},
			labelPositionChange(index) {
				this.labelPosition = index == 0 ? 'left' : 'top';
			},
			codeChange(text) {
				this.codeTips = text;
			},
			// 获取验证码
			getCode() {
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
			errorChange(index) {
				if (index == 0) this.errorType = ['message'];
				if (index == 1) this.errorType = ['toast'];
				if (index == 2) this.errorType = ['border-bottom'];
				if (index == 3) this.errorType = ['border'];
			}
		}
	};
</script>

<style scoped lang="scss">
	.wrap {
		padding: 30rpx;
	}

	.agreement {
		display: flex;
		align-items: center;
		margin: 40rpx 0;

		.agreement-text {
			padding-left: 8rpx;
			color: $u-tips-color;
		}
	}
</style>
