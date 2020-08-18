import request from '@/utils/request'

// 查询AkCarsRecord列表
export function listAkCarsRecord(query) {
return request({
url: '/api/v1/akcarsrecordList',
method: 'get',
params: query
})
}

// 查询AkCarsRecord详细
export function getAkCarsRecord (id) {
return request({
url: '/api/v1/akcarsrecord/' + id,
method: 'get'
})
}


// 新增AkCarsRecord
export function addAkCarsRecord(data) {
return request({
url: '/api/v1/akcarsrecord',
method: 'post',
data: data
})
}

// 修改AkCarsRecord
export function updateAkCarsRecord(data) {
return request({
url: '/api/v1/akcarsrecord',
method: 'put',
data: data
})
}

// 删除AkCarsRecord
export function delAkCarsRecord(id) {
return request({
url: '/api/v1/akcarsrecord/' + id,
method: 'delete'
})
}

