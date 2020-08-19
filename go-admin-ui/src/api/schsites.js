import request from '@/utils/request'

// 查询SchSites列表
export function listSchSites(query) {
  return request({
    url: '/api/v1/schsitesList',
    method: 'get',
    params: query
  })
}

// 查询SchSites详细
export function getSchSites(id) {
  return request({
    url: '/api/v1/schsites/' + id,
    method: 'get'
  })
}

// 新增SchSites
export function addSchSites(data) {
  return request({
    url: '/api/v1/schsites',
    method: 'post',
    data: data
  })
}

// 修改SchSites
export function updateSchSites(data) {
  return request({
    url: '/api/v1/schsites',
    method: 'put',
    data: data
  })
}

// 删除SchSites
export function delSchSites(id) {
  return request({
    url: '/api/v1/schsites/' + id,
    method: 'delete'
  })
}

