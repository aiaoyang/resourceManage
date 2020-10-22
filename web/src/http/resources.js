import http from '@/http/http.js';

export const GetResource = (config) => http("get", "/resource", config);

export const GetECS = (config) => http("get", "/ecs", config)

export const GetRDS = (config) => http("get", "/rds", config)

export const GetDomain = (config) => http("get", "/domain", config)
