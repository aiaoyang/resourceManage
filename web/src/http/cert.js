import http from '@/http/http.js'
export const GetCert = (config) => http('get', "/cert", config)