import request from '@/utils/request'

export function listFile(query) {
  return request({
    url: '/file/manage/list',
    method: 'get',
    params: query
  })
}

export function delFile(query) {
  return request({
    url: '/file/manage',
    method: 'delete',
    params: query
  })
}