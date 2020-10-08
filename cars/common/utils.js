const formatTime = date => {
	// date = new Date(date);
	// const year = date.getFullYear()
	// const month = date.getMonth() + 1
	// const day = date.getDate()
	// const hour = date.getHours()
	// const minute = date.getMinutes()
	// const second = date.getSeconds()

	// return [year, month, day].map(formatNumber).join('/') + ' ' + [hour, minute, second].map(formatNumber).join(':')
	console.log("test")
}

const getRad = (d) => {
	var PI = Math.PI;
	return d * PI / 180.0;
}

const getGreatCircleDistance = (lat1, lng1, lat2, lng2) => {
	var EARTH_RADIUS = 6378137.0;
	var radLat1 = getRad(lat1);
	var radLat2 = getRad(lat2);

	var a = radLat1 - radLat2;
	var b = getRad(lng1) - getRad(lng2);

	var s = 2 * Math.asin(Math.sqrt(Math.pow(Math.sin(a / 2), 2) + Math.cos(radLat1) * Math.cos(radLat2) * Math.pow(Math.sin(
		b / 2), 2)));
	s = s * EARTH_RADIUS;
	s = Math.round(s * 10000) / 10000.0;

	return s;
}

const getFlatternDistance = (lat1, lng1, lat2, lng2) => {
	var EARTH_RADIUS = 6378137.0;
	var f = getRad((lat1 + lat2) / 2);
	var g = getRad((lat1 - lat2) / 2);
	var l = getRad((lng1 - lng2) / 2);

	var sg = Math.sin(g);
	var sl = Math.sin(l);
	var sf = Math.sin(f);

	var s, c, w, r, d, h1, h2;
	var a = EARTH_RADIUS;
	var fl = 1 / 298.257;

	sg = sg * sg;
	sl = sl * sl;
	sf = sf * sf;

	s = sg * (1 - sl) + (1 - sf) * sl;
	c = (1 - sg) * (1 - sl) + sf * sl;

	w = Math.atan(Math.sqrt(s / c));
	r = Math.sqrt(s * c) / w;
	d = 2 * w * a;
	h1 = (3 * r - 1) / 2 / c;
	h2 = (3 * r + 1) / 2 / s;

	return d * (1 + fl * (h1 * sf * (1 - sg) - h2 * (1 - sf) * sg));
}

module.exports = {
	formatTime,
	getFlatternDistance,
	getGreatCircleDistance,
	getRad,
}
