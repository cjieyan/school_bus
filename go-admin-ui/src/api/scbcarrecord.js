import request from '@/utils/request'

// 查询ScbCarRecord列表
export function listScbCarRecord(query) {
return request({
url: '/api/v1/scbcarrecordList',
method: 'get',
params: query
})
}

// 查询ScbCarRecord详细
export function getScbCarRecord (id) {
return request({
url: '/api/v1/scbcarrecord/' + id,
method: 'get'
})
}


// 新增ScbCarRecord
export function addScbCarRecord(data) {
return request({
url: '/api/v1/scbcarrecord',
method: 'post',
data: data
})
}

// 修改ScbCarRecord
export function updateScbCarRecord(data) {
return request({
url: '/api/v1/scbcarrecord',
method: 'put',
data: data
})
}

// 删除ScbCarRecord
export function delScbCarRecord(id) {
return request({
url: '/api/v1/scbcarrecord/' + id,
method: 'delete'
})
}

