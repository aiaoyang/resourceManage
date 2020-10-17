import axios from "axios";
var instance = axios.create({
  baseURL: "http://10.10.0.134:9090",
  timeout: 3000,
});

// 请求
instance.interceptors.request.use((request) => {
  return request;
}),
  (error) => {
    console.log(error);
    return Promise.resolve(error);
  };

// 响应
instance.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    return Promise.resolve(error.response);
  }
);
/**
 * 使用es6的export default导出了一个函数，导出的函数代替axios去帮我们请求数据，
 * 函数的参数及返回值如下：
 * @param {String} method  请求的方法：get、post、delete、put
 * @param {String} url     请求的url:
 * @param {Object} data    请求的参数
 * @returns {Promise}     返回一个promise对象，其实就相当于axios请求数据的返回值
 */
export default function (method, url, data = null) {
  method = method.toLowerCase();
  if (method == "post") {
    return instance.post(url, data);
  } else if (method == "get") {
    return instance.get(url, { params: data });
  } else if (method == "put") {
    return instance.put(url, data);
  } else {
    console.error("未知请求方法" + method);
    return false;
  }
}
