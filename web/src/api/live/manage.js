import request from '@/utils/request'

// 查询角色列表
export function listInfo(query) {
  return request({
    url: '/live/info/list',
    method: 'get',
    params: query
  })
}

export function getLiveManage(id) {
  return request({
    url: '/live/manage/' + id,
    method: 'get'
  })
}

// 新增角色
export function addLiveManage(data) {
  return request({
    url: '/live/manage',
    method: 'post',
    data: data
  })
}

export function updateLiveManage(data) {
  return request({
    url: '/live/manage',
    method: 'put',
    data: data
  })
}

export function delLiveManage(id) {
  return request({
    url: '/live/manage/' + id,
    method: 'delete'
  })
}