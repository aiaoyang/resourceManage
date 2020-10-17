import http from '@/http/http.js';
export const GetResource = (config) => http("get", "/resource", config);
export const TestResource = () => {
	return [
		{
			index: "1",
			type: "ecs",
			name: "ecs1",
			size: "4c16g",
			end: "20200101",
			belong: "碧蓝海外服务器"
		},
		{
			index: "2",
			type: "rds",
			name: "rds1",
			size: "4c16g",
			end: "20200101",
			belong: "yongshiwl"
		}
	]
}

export const TestResourceLabel = () => {
	return [
		"type",
		"endOfTime",
		"belongTo"
	]
}

// 用这个
export const GetECS = (config) => http("get", "/ecs", config)

export const GetRDS = (config) => http("get", "/rds", config)
