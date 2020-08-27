import request from '@/utils/request'

// 查询ScbCars列表
export function listScbCars(query) {
  return request({
    url: '/api/v1/scbcarsList',
    method: 'get',
    params: query
  })
}

// 查询ScbCars详细
export function getScbCars() {
  return request({
    url: '/api/v1/scbcars/',
    method: 'get'
  })
}

// 新增ScbCars
export function addScbCars(data) {
  return request({
    url: '/api/v1/scbcars',
    method: 'post',
    data: data
  })
}

// 修改ScbCars
export function updateScbCars(data) {
  return request({
    url: '/api/v1/scbcars',
    method: 'put',
    data: data
  })
}

// 删除ScbCars
export function delScbCars() {
  return request({
    url: '/api/v1/scbcars/',
    method: 'delete'
  })
}
export function treeselect() {
  return request({
    url: '/api/v1/scbcarsAll',
    method: 'get'
  })
}
