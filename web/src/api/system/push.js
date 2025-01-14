import request from '@/utils/request'

// 查询字典数据列表
export function listChannel(query) {
  return request({
    url: '/system/push/channel/list',
    method: 'get',
    params: query
  })
}

// 查询字典数据详细
export function getChannel(cId) {
  return request({
    url: '/system/push/channel/' + cId,
    method: 'get'
  })
}

// 新增字典数据
export function addChannel(data) {
  return request({
    url: '/system/push/channel',
    method: 'post',
    data: data
  })
}

// 修改字典数据
export function updateChannel(data) {
  return request({
    url: '/system/push/channel',
    method: 'put',
    data: data
  })
}

// 删除字典数据
export function delChannel(cId) {
  return request({
    url: '/system/push/channel/' + cId,
    method: 'delete'
  })
}
