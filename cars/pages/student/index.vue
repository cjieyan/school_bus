<template>
	<view>
		<view class="top">
			<u-navbar back-text="返回" @tap="back" title="学生列表"></u-navbar>
		</view>
		<view class="student">
			<view class="student-top">
				<span class="location-img">
					<image src="../../static/location.png" style="width: 32px; height: 31px;" class="location-image"></image>
				</span>
				<span class="location-info">
					梧桐小学
				</span>
			</view>
			<view class="student-list">
				<view class="student-list-info" v-for="item in students">
					<view class="student-img">
						<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
						<view class="inboard">
							<span v-if="item.swipeStatus == '-1'">已上车</span>
							<span v-else>未上车</span>
							<!-- <fa-icon type="check" size="14" color="white" class="font-icon inboard-check"></fa-icon> -->
							<u-icon name="checkbox-mark" color="#fff" size="14" class="inboard-check"></u-icon>
						</view>
					</view>
					<view class="student-name">
						{{item.name}}
					</view>
				</view>
				<!-- <view class="student-list-info">
					<view class="student-img">
						<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
						<view class="outboard">
							<span>未上车</span>
						</view>
					</view>
					<view class="student-name">
						张三
					</view>
				</view>
				<view class="student-list-info">
					<view class="student-img">
						<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
					</view>
					<view class="student-name">
						张三
					</view>
				</view>
				<view class="student-list-info">
					<view class="student-img">
						<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
					</view>
					<view class="student-name">
						张三
					</view>
				</view>
				<view class="student-list-info">
					<view class="student-img">
						<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
					</view>
					<view class="student-name">
						张三
					</view>
				</view> -->
				<view style="clear:both;height:0"></view>
			</view>
		</view>
		<view class="comfirm">
			<button class="comfirm-bottom" @tap="back">确认结束</button>
		</view>
	</view>
</template>

<script>
	// import faIcon from "@/components/kilvn-fa-icon/fa-icon.vue"
	export default {
		components: {
			// faIcon
		},
		data() {
			return {
				students: {}
			}
		},
		methods: {
			back() {
				uni.switchTab({
					url: "../index/index",
					success: (res) => {
						console.log(res)
					},
					fail: (err) => {
						console.log(err)
					}
				})
			},
			studentList() {
				const token = uni.getStorageSync('token')
				uni.request({
					url: "http://localhost:8000/xcx/auth/line-students",
					method:"POST",
					header: {
						'token': token,
					},
					method: "GET",
					success: (res) => {
						console.log(res)
						this.students = res.data.data.studentsDataRet
					},
					fail: (err) => {
						console.log(err)
					}
				})
			}
		},
		onShow() {
			this.studentList()
		},
		onLoad() {
			this.studentList()
		}
	}
</script>

<style>
	page {
		background-color: #f8f8f8;
	}

	.student {
		margin-top: 10px;
		padding: 0 10%;
	}

	.student-top {
		border-bottom: 1px solid rgb(238 238 238);
		padding-bottom: 10px;
	}

	.location-info {
		margin-left: 15px;
		color: rgb(89 89 89);
		font-weight: bold;
	}

	.student-list-info {
		width: 40px;
		float: left;
		margin: 15px 10px;
	}

	.student-img {
		width: 40px;
		height: 40px;
		border-radius: 20px;
		text-align: center;
		position: relative;
	}

	.outboard,
	.inboard {
		background-color: #808080;
		width: 42px;
		height: 42px;
		position: absolute;
		top: 0;
		left: 0;
		border-radius: 20px;
		opacity: 0.9;
	}

	.outboard span,
	.inboard span {
		font-size: 8px;
		color: #fff;
		position: absolute;
		top: 16px;
		left: 9px;
	}

	.inboard-check {
		position: absolute;
		bottom: -6px;
		left: 18px;
		background-color: #4CD964;
		border-radius: 6px;
	}

	.student-name {
		margin-top: 8px;
		font-size: 12px;
		text-align: center;
	}

	.comfirm {
		margin-bottom: 30px;
		padding: 10px 40px;
	}

	.comfirm-bottom {
		background: linear-gradient(to right, rgba(8, 189, 175), rgb(247, 218, 99));
		color: #fff;
		border: none;
		border-radius: 25px;
	}

	.comfirm-bottom::after {
		border: none;
	}
</style>
