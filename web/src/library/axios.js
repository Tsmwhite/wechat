import store2 from 'store2'
import axios from 'axios'
const config = {
    proxy:'/api',       //代理配置
}

//axios基本配置
const api = axios.create({
    timeout: 10000,
    withCredentials:true
})
//请求开始拦截
api.interceptors.request.use(conf => {
        //请求带token
        conf.headers['Authorization'] = store2('accesstoken')
        return conf
    },
    error => ({ status: 0, msg: error.message })
)

//请求返回拦截
api.interceptors.response.use(response => {
        return Promise.resolve(response).then(checkCode)
    },
    error => {
        checkStatus(error.response)
        return Promise.reject(error)
    }
)
// http状态码错误处理
const checkStatus = (res) => {
    window.$pig.$vux.loading.hide()
    switch (res.status)
    {
        case 401 :{        //登录过期
            window.console.log('登录过期')
            break;
        }
        default:
            window.console.log('服务器存在异常', 'middle')
            break;
    }
}

// 后台自定义 code错误处理
const checkCode = (res) => {
    if(res){
        if(res.data.code === 0 || moreCodeFn(res)){       //code为0成功
            return res
        }else {
            window.console.log(res.data.msg)  //统一处理错误
            return false
        }
    }else {
        return false
    }
}

//过滤 接口函数（这些接口可以自定义 错误处理）
const moreCodeFn = (res) =>{
    let url = res.config.url;    //当前URL
    let urlArray = ['/auth/morecode/token','/social/morecode/join'];
    for(let i=0;i<urlArray.length;i++){
        if(url.indexOf(urlArray[i])>=0){
            return true
        }
    }
    return false
}

// 解析参数
const formatParams =  (method = 'GET', params) => {
    //headers配置
    const headers = {
        'Content-Type': 'application/json;charset=utf-8',
    }
    switch (method) {
        case 'POST':
            return {
                headers,
                method,
                data:params
            }
        case 'PUT':
            return {
                headers,
                method,
                data:params
            }
        case 'DELETE':
            return {
                headers,
                method,
            }
        case 'GET':
            return {
                headers,
                method,
                params
            }
        default:
            return {
                headers,
                method,
                params
            }
    }
}

export {api,formatParams,config}