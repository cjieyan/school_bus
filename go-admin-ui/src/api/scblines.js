import request from '@/utils/request'

// 查询ScbLines列表
export function listScbLines(query) {
return request({
url: '/api/v1/scblinesList',
method: 'get',
params: query
})
}

// 查询ScbLines详细
export function getScbLines (id) {
return request({
url: '/api/v1/scblines/' + id,
method: 'get'
})
}


// 新增ScbLines
export function addScbLines(data) {
return request({
url: '/api/v1/scblines',
method: 'post',
data: data
})
}

// 修改ScbLines
export function updateScbLines(data) {
return request({
url: '/api/v1/scblines',
method: 'put',
data: data
})
}

// 删除ScbLines
export function delScbLines(id) {
return request({
url: '/api/v1/scblines/' + id,
method: 'delete'
})
}

