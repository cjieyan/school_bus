import request from '@/utils/request'

// 查询ScbDept列表
export function getDeptList(query) {
  return request({
    url: '/api/v1/scbdeptList',
    method: 'get',
    params: query
  })
}

// 查询ScbDept详细
export function getScbDept(deptId) {
  return request({
    url: '/api/v1/scbdept/' + deptId,
    method: 'get'
  })
}

// 查询部门下拉树结构
export function treeselect() {
  return request({
    url: '/api/v1/scbdeptTree',
    method: 'get'
  })
}

// 根据角色ID查询部门树结构
export function roleDeptTreeselect(roleId) {
  return request({
    url: '/api/v1/roleScbDeptTreeselect/' + roleId,
    method: 'get'
  })
}

// 新增ScbDept
export function addScbDept(data) {
  return request({
    url: '/api/v1/scbdept',
    method: 'post',
    data: data
  })
}

// 修改ScbDept
export function updateScbDept(data) {
  return request({
    url: '/api/v1/scbdept',
    method: 'put',
    data: data
  })
}

// 删除ScbDept
export function delScbDept(deptId) {
  return request({
    url: '/api/v1/scbdept/' + deptId,
    method: 'delete'
  })
}

