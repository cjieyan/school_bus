import request from '@/utils/request'

// 查询ScbTeachers列表
export function listScbTeachers(query) {
  return request({
    url: '/api/v1/scbteachersList',
    method: 'get',
    params: query
  })
}

// 查询ScbTeachers详细
export function getScbTeachers(id) {
  return request({
    url: '/api/v1/scbteachers/' + id,
    method: 'get'
  })
}

// 新增ScbTeachers
export function addScbTeachers(data) {
  return request({
    url: '/api/v1/scbteachers',
    method: 'post',
    data: data
  })
}

// 修改ScbTeachers
export function updateScbTeachers(data) {
  return request({
    url: '/api/v1/scbteachers',
    method: 'put',
    data: data
  })
}

// 删除ScbTeachers
export function delScbTeachers(id) {
  return request({
    url: '/api/v1/scbteachers/' + id,
    method: 'delete'
  })
}

