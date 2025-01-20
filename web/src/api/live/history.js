import request from '@/utils/request'

export function listHistory(query) {
  return request({
    url: '/live/history/list',
    method: 'get',
    params: query
  })
}
export function getHistory(id) {
    return request({
      url: '/live/history/' + id,
      method: 'get'
    })
  }
  
  // 新增角色
  export function addHistory(data) {
    return request({
      url: '/live/history',
      method: 'post',
      data: data
    })
  }
  
  export function updateHistory(data) {
    return request({
      url: '/live/history',
      method: 'put',
      data: data
    })
  }
  
  export function delHistory(id) {
    return request({
      url: '/live/history/' + id,
      method: 'delete'
    })
  }