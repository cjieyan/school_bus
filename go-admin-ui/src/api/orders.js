import request from '@/utils/request'

// 查询Orders列表
export function listOrders(query) {
  return request({
    url: '/api/v1/ordersList',
    method: 'get',
    params: query
  })
}

// 查询Orders详细
export function getOrders(id) {
  return request({
    url: '/api/v1/orders/' + id,
    method: 'get'
  })
}

// 新增Orders
export function addOrders(data) {
  return request({
    url: '/api/v1/orders',
    method: 'post',
    data: data
  })
}

// 修改Orders
export function updateOrders(data) {
  return request({
    url: '/api/v1/orders',
    method: 'put',
    data: data
  })
}

// 删除Orders
export function delOrders(id) {
  return request({
    url: '/api/v1/orders/' + id,
    method: 'delete'
  })
}

