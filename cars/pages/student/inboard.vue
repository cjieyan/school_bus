<template>
	<view>
		<view class="top">
			<u-navbar back-text="返回" title="学生列表"></u-navbar>
		</view>
		<view class="student">
			<view class="student-top">
				<span class="location-img">
					<image src="../../static/location.png" style="width: 32px; height: 31px;" class="location-image"></image>
				</span>
				<span class="location-info">
					梧桐小学--六约新村
				</span>
			</view>
			<view class="student-data">
				<view class="student-img">
					<image src="../../static/location.png" style="width: 30px; height: 30px;" class="location-image"></image>
				</view>
				<view class="student-info">
					<view class="info-top">
						<view class="name">张三王五</view>
						<view class="isinboard">请确认该学生是否打卡上车</view>
					</view>
					<view class="info-bottom">
						<view class="tel">家长电话:1233444556</view>
						<view class="time">考勤时间:2020/09/15 15:11:11</view>
						<view class="carno">车牌号码:粤B1234</view>
						<view class="is">是否接送:否</view>
					</view>
					
				</view>
				<view style="clear:both;height:0"></view>
			</view>
		</view>
		<view class="comfirm">
			<button class="comfirm-bottom" @tap="gotoLunBo">上车</button>
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

			}
		},
		methods: {

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
	.student-img {
		width: 40px;
		height: 40px;
		border-radius: 20px;
		text-align: center;
		position: relative;
		float: left;
	}
	.student-info{
		float: left;
		margin-left: 20px;
		line-height: 25px;
		font-size: 12px;
	}
	.info-bottom{
		margin-top: 20px;
	}
	.student-top {
		border-bottom: 1px solid rgb(238 238 238);
		padding-bottom: 10px;
	}
	.student-data{
		margin-top: 20px;
	}
	.location-info {
		margin-left: 15px;
		color: rgb(89 89 89);
		font-weight: bold;
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
