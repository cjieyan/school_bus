import request from '@/utils/request'

// 查询UserAccount列表
export function listUserAccount(query) {
  return request({
    url: '/api/v1/useraccountList',
    method: 'get',
    params: query
  })
}

// 查询UserAccount详细
export function getUserAccount(id) {
  return request({
    url: '/api/v1/useraccount/' + id,
    method: 'get'
  })
}

// 新增UserAccount
export function addUserAccount(data) {
  return request({
    url: '/api/v1/useraccount',
    method: 'post',
    data: data
  })
}

// 修改UserAccount
export function updateUserAccount(data) {
  return request({
    url: '/api/v1/useraccount',
    method: 'put',
    data: data
  })
}

// 删除UserAccount
export function delUserAccount(id) {
  return request({
    url: '/api/v1/useraccount/' + id,
    method: 'delete'
  })
}

