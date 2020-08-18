import request from '@/utils/request'

// 查询AkCars列表
export function listAkCars(query) {
  return request({
    url: '/api/v1/akcarsList',
    method: 'get',
    params: query
  })
}

// 查询AkCars详细
export function getAkCars(id) {
  return request({
    url: '/api/v1/akcars/' + id,
    method: 'get'
  })
}

// 新增AkCars
export function addAkCars(data) {
  return request({
    url: '/api/v1/akcars',
    method: 'post',
    data: data
  })
}

// 修改AkCars
export function updateAkCars(data) {
  return request({
    url: '/api/v1/akcars',
    method: 'put',
    data: data
  })
}

// 删除AkCars
export function delAkCars(id) {
  return request({
    url: '/api/v1/akcars/' + id,
    method: 'delete'
  })
}

