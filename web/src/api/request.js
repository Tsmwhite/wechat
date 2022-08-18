import axios from 'axios'
import Qs from 'qs'
import {Toast} from 'vant';
import {GetToken, LogoutAfterCache} from "../utis/cache";
import {LogoutAfterStore} from "../stores";

const LoginAfter = () => {
    LogoutAfterCache()
    LogoutAfterStore()
}

let _token
export const SetToken = (token) => {
    _token = token
}

const request = axios.create({
    headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
        'X-Requested-With': 'XMLHttpRequest'
    },
    timeout: 600000, // 请求超时时间
    transformRequest: [function (data) {
        data = Qs.stringify(data)
        return data
    }]
});

/**
 * 请求拦截器
 * 每次请求前，如果存在token则在请求头中携带token
 */
request.interceptors.request.use(
    config => {
        // 登录流程控制中，根据本地是否存在token判断用户的登录情况
        // 但是即使token存在，也有可能token是过期的，所以在每次的请求头中携带token
        // 后台根据携带的token判断用户的登录情况，并返回给我们对应的状态码
        // 而后我们可以在响应拦截器中，根据状态码进行一些统一的操作。
        const token = GetToken();
        if (token || _token) {
            config.headers["authorization"] = token || _token
        }
        return config;
    },
    error => Promise.error(error)
);
// respone拦截器
request.interceptors.response.use(
    response => {
        /** res为非0是抛错 **/
        const status = response.status
        const res = response.data
        switch (status) {
            case 200:
            case 201:
                // 内层switch@start
                switch (res.err_code) {
                    case 0:
                        return Promise.resolve(res)
                    case -9:
                        //当错误需要强提醒时
                        Toast('错误');
                        return Promise.reject(res)
                    case 401://登录态失效
                        Toast('请重新登录');
                        LoginAfter()
                        setTimeout(() => {

                        }, 1000);
                        break;
                    default:
                        Toast(res.err_msg);
                        return Promise.reject(res)
                }
                // 内层switch@end
                break
            case 500:
                Toast('服务器错误');
                return Promise.reject(res);
                break
            default:
                Toast('网络连接出错, error_code: ' + status);
                return Promise.reject(res)
        }

    }, error => {
        Toast('网络连接出错')
        return Promise.reject(error)
    }
);

export default request
