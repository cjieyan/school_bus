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
export function getScbLines(id) {
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

export function carsTreeselect(lineId) {
  return request({
    url: '/api/v1/carsTreeselect/' + lineId,
    method: 'get'
  })
}
// 获取所有线路
export function getAllLines() {
  return request({
    url: '/api/v1/getAllLines',
    method: 'get'
  })
}
// 获取线路的所有车辆
export function getLineCars(lineId) {
  return request({
    url: '/api/v1/getLines/cars?id=' + lineId,
    method: 'get'
  })
}
// 获取线路的所有站点
export function getLineSites(lineId) {
  return request({
    url: '/api/v1/getLines/sites?id=' + lineId,
    method: 'get'
  })
}
