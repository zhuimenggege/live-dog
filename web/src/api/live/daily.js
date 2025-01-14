import request from '@/utils/request'

// 查询角色列表
export function listDaily(query) {
  return request({
    url: '/live/daily/list',
    method: 'get',
    params: query
  })
}

export function getDaily(id) {
  return request({
    url: '/live/daily/' + id,
    method: 'get'
  })
}

// 新增角色
export function addDaily(data) {
  return request({
    url: '/live/daily',
    method: 'post',
    data: data
  })
}

export function updateDaily(data) {
  return request({
    url: '/live/daily',
    method: 'put',
    data: data
  })
}

export function delDaily(id) {
  return request({
    url: '/live/daily/' + id,
    method: 'delete'
  })
}