import request from '@/utils/request'

// 查询ScbPost列表
export function listPost(query) {
  return request({
    url: '/api/v1/scbpostList',
    method: 'get',
    params: query
  })
}

// 查询ScbPost详细
export function getPost(postId) {
  return request({
    url: '/api/v1/scbpost/' + postId,
    method: 'get'
  })
}

// 新增ScbPost
export function addPost(data) {
  return request({
    url: '/api/v1/scbpost',
    method: 'post',
    data: data
  })
}

// 修改ScbPost
export function updatePost(data) {
  return request({
    url: '/api/v1/scbpost',
    method: 'put',
    data: data
  })
}

// 删除ScbPost
export function delPost(postId) {
  return request({
    url: '/api/v1/scbpost/' + postId,
    method: 'delete'
  })
}
export function getScbpostAll() {
  return request({
    url: '/api/v1/scbpostAll',
    method: 'get'
  })
}
