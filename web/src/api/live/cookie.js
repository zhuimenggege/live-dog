import request from '@/utils/request'

// 查询角色列表
export function listCookie(query) {
  return request({
    url: '/live/cookie/list',
    method: 'get',
    params: query
  })
}

export function getCookie(id) {
  return request({
    url: '/live/cookie/' + id,
    method: 'get'
  })
}

// 新增角色
export function addCookie(data) {
  return request({
    url: '/live/cookie',
    method: 'post',
    data: data
  })
}

export function updateCookie(data) {
  return request({
    url: '/live/cookie',
    method: 'put',
    data: data
  })
}

export function delCookie(id) {
  return request({
    url: '/live/cookie/' + id,
    method: 'delete'
  })
}