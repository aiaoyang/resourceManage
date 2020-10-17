import http from '@/http/http.js'
export const GetAlarm = (config) => http('get', "/alarm", config)