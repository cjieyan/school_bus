<template>
	<view class="u-demo">
		<view class="top">
			<u-navbar :is-back="isback" title="智慧校车" @click="back" class="top" :background="background" back-icon-color="#fff"
			 title-color="#fff"></u-navbar>
		</view>
		<view class="u-demo-wrap">
			<view class="u-demo-area">
				<u-toast ref="uToast"></u-toast>
				<u-verification-code :keep-running="true" :seconds="seconds" @end="end" @start="start" ref="uCode" @change="codeChange"
				 :startText="startText" :changeText="changeText" :endText="endText"></u-verification-code>
				<u-field v-model="mobile" label="手机号" :error-message="errorMessage" placeholder="请填写手机号" :required="required" :icon="icon1"
				 :type="type">
				</u-field>
				<u-field v-model="code" label="验证码" placeholder="请填写验证码" :required="required" :icon="icon2">
					
					<u-button v-if="showBtn" @click="getCode" slot="right" size="mini" type="success">{{tips}}</u-button>
				</u-field>
			</view>
			<u-button class="btn" type="success" @tap="toindex">绑定手机</u-button>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				isback: false,
				background: {
					backgroundColor: '#12c497',
				},
				tips: '',
				seconds: 60,
				refCode: null,
				startText: '获取验证码',
				changeText: 'X秒重新获取',
				endText: '重新获取',
				mobile: '',
				code: '',
				errorMessage: '',
				required: false,
				placeholderColor: this.$u.color['tipsColor'],
				arrow: false,
				showBtn: true,
				icon1: '',
				icon2: '',
				type: 'text',
			}
		},
		onReady() {
			// 注意这里为错误示例，目前微信小程序最新稳定版开发工具如此
			// 赋值会报错，很诡异，其他端无此问题
			// this.refCode = this.$refs.uCode;
		},
		methods: {
			toindex(){
				uni.redirectTo({
					url:"../index/index"
				})
			},
			codeChange(text) {
				// console.log(text);
				this.tips = text;
			},
			getCode() {
				console.log(this.$refs)
				if (this.$refs.uCode.canGetCode) {
					// 模拟向后端请求验证码
					uni.showLoading({
						title: '正在获取验证码'
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
			end() {
				// this.$u.toast('倒计时结束');
			},
			start() {
				// this.$u.toast('倒计时开始');
			},
			reset() {
				this.$refs.uCode.reset();
			}
		}
	}
</script>

<style lang="scss" scoped>
	.u-demo {}
	.btn{
		margin-top: 30rpx;
	}
</style>
