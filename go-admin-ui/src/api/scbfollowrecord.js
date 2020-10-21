import request from '@/utils/request'

// 查询ScbFollowRecord列表
export function listScbFollowRecord(query) {
  return request({
    url: '/api/v1/scbfollowrecordList',
    method: 'get',
    params: query
  })
}

// 查询ScbFollowRecord详细
export function getScbFollowRecord(id) {
  return request({
    url: '/api/v1/scbfollowrecord/' + id,
    method: 'get'
  })
}

// 新增ScbFollowRecord
export function addScbFollowRecord(data) {
  return request({
    url: '/api/v1/scbfollowrecord',
    method: 'post',
    data: data
  })
}

// 修改ScbFollowRecord
export function updateScbFollowRecord(data) {
  return request({
    url: '/api/v1/scbfollowrecord',
    method: 'put',
    data: data
  })
}

// 删除ScbFollowRecord
export function delScbFollowRecord(id) {
  return request({
    url: '/api/v1/scbfollowrecord/' + id,
    method: 'delete'
  })
}

