<template>
	<view>
		<view class="top">
			<u-navbar back-text="返回" title="学生信息" @tap="back" :background="background" back-icon-color="#fff" title-color="#fff"></u-navbar>
		</view>
		<view class="student">
			<view class="student-top">
				<span class="location-img">
					<image src="../../static/location.png" style="width: 32px; height: 31px;" class="location-image"></image>
				</span>
				<span class="location-info">
					{{lineinfo.name}}
				</span>
			</view>
			<view class="student-data">
				<view class="student-img">
					<image :src="studentinfo.headImg" style="width: 30px; height: 30px;" class="headImg"></image>
				</view>
				<view class="student-info">
					<view class="info-top">
						<view class="name">{{studentinfo.name}}</view>
						<view class="isinboard">请确认该学生是否打卡上车</view>
					</view>
					<view class="info-bottom">
						<view class="tel">家长电话:{{studentinfo.parentPhone}}</view>
						<view class="time">考勤时间:{{studentinfo.timeString}}</view>
						<view class="carno">车牌号码:{{carinfo.carNumber}}</view>
						<view class="isPickUp" v-if="studentinfo.isPickUp == '0'">是否接送:否</view>
						<view class="isPickUp" v-else>是否接送:是</view>
					</view>
				</view>
				<view style="clear:both;height:0"></view>
			</view>
		</view>
		<view class="comfirm">
			<button class="comfirm-bottom" @tap="swipe" :loading="loading" :disabled="disabled">{{inboard}}</button>
		</view>
	</view>
</template>

<script>
	export default {
		components: {
			// faIcon
		},
		data() {
			return {
				studentinfo: {},
				lineinfo: {},
				carinfo:{},
				carid: "",
				lineid: "",
				userid: "",
				swipeStatus: "",
				disabled: true,
				loading: true,
				inboard: "",
				background: {
					backgroundColor: '#12c497',
				},
			}
		},
		methods: {
			back() {
				uni.switchTab({
					url:"../index/index",
					fail: (err) => {
						console.log(err)
					}
				})
			},
			swipe(){
				//手动打卡
				uni.showLoading({
					
				})
				uni.request({
					url: this.$store.state.apihost + "/xcx/auth/swipe",
					header: {
						'token': this.$store.state.token,
					},
					data:{
						"car_id": this.carid,
						"line_id": this.lineid,
						"student_id": this.userid,
					},
					method:"POST",
					success: (res) => {
						if(res.data.code == 200)
						{
							this.studentinfo.swipeStatus == '0'
							this.inboard = res.data.msg
							this.disabled = true
						}
					},
					fail: (err) => {
						uni.showToast({
							title:"请求失败"
						})
					}
				})
				uni.hideLoading()
			}
		},
		onLoad(options) {
			uni.showLoading({
				
			})
			uni.request({
				url: this.$store.state.apihost + "/xcx/auth/student-info",
				method:"POST",
				header: {
					'token': this.$store.state.token,
				},
				data:{
					"student_id": parseInt(options.id)
				},
				success: (res) => {
					this.carid = res.data.data.carId,
					this.lineid = res.data.data.lineId
					this.userid = res.data.data.id
					this.studentinfo = res.data.data
					if(this.studentinfo.swipeStatus == '-1'){
						this.inboard = "未上车"
						this.disabled = false
					} else if(this.studentinfo.swipeStatus == '0'){
						this.inboard = "已上车"
						this.disabled = false
					} else if(this.studentinfo.swipeStatus == '1'){
						this.inboard = "已下车"
						this.disabled = false
					}
					this.loading = false
				},
				fail: (err) => {
					console.log(err)
				}
			})
			this.lineinfo = this.$store.state.liniInfo
			this.carinfo = this.$store.state.carInfo
			uni.hideLoading()
		}
	}
</script>

<style>
	page {
		background-color: #f8f8f8;
	}

	.student {
		margin-top: 20rpx;
		padding: 0 10%;
	}

	.student-img {
		width: 80rpx;
		height: 80rpx;
		border-radius: 40rpx;
		text-align: center;
		position: relative;
		float: left;
	}

	.student-info {
		float: left;
		margin-left: 40rpx;
		line-height: 50rpx;
		font-size: 24rpx;
	}

	.info-bottom {
		margin-top: 20px;
	}

	.student-top {
		border-bottom: 2rpx solid rgb(238 238 238);
		padding-bottom: 20rpx;
	}

	.student-data {
		margin-top: 40rpx;
	}

	.location-info {
		margin-left: 30rpx;
		color: rgb(89 89 89);
		font-weight: bold;
	}

	.comfirm {
		margin-bottom: 60rpx;
		padding: 20rpx 80rpx;
	}

	.comfirm-bottom {
		background: linear-gradient(to right, rgba(8, 189, 175), rgb(247, 218, 99));
		color: #fff;
		border: none;
		border-radius: 50rpx;
	}

	.comfirm-bottom::after {
		border: none;
	}
	image.headImg {
		border-radius: 50rpx;
	}
</style>
