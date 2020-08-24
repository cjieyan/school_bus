import request from '@/utils/request'

// 查询ScbStudents列表
export function listScbStudents(query) {
  return request({
    url: '/api/v1/scbstudentsList',
    method: 'get',
    params: query
  })
}

// 查询ScbStudents详细
export function getScbStudents(id) {
  return request({
    url: '/api/v1/scbstudents/' + id,
    method: 'get'
  })
}

// 新增ScbStudents
export function addScbStudents(data) {
  return request({
    url: '/api/v1/scbstudents',
    method: 'post',
    data: data
  })
}

// 修改ScbStudents
export function updateScbStudents(data) {
  return request({
    url: '/api/v1/scbstudents',
    method: 'put',
    data: data
  })
}

// 删除ScbStudents
export function delScbStudents(id) {
  return request({
    url: '/api/v1/scbstudents/' + id,
    method: 'delete'
  })
}

