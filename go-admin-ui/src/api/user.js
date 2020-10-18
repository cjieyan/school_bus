import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function refreshtoken(data) {
  return request({
    url: '/refreshtoken',
    method: 'post',
    data
  })
}

export function getInfo() {
  return request({
    url: '/api/v1/getinfo',
    method: 'get'
  })
}

export function logout() {
  return request({
    url: '/api/v1/logout',
    method: 'post'
  })
}
export function getRecent() {
  return request({
    url: '/api/v1/cars-record/recent',
    method: 'get'
  })
}
export function getPreview() {
  return request({
    url: '/api/v1/get-preview',
    method: 'get'
  })
}
