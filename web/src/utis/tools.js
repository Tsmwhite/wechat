const Tools = {}
export const getUrlParam = (variable) => {
    let query = window.location.search.substring(1);
    if (query == '') {
        let href = window.location.href
        query = href.substring(href.indexOf('?') + 1)
    }
    let vars = query.split("&");
    for (let i = 0; i < vars.length; i++) {
        let pair = vars[i].split("=");
        if (pair[0] == variable) {
            return pair[1]
        }
    }
    return (false);
}

export const replaceParamVal = function (paramName, replaceWith, oUrl = "") {
    oUrl = oUrl ? oUrl : window.location.href.toString()
    let re = eval('/(' + paramName + '=)([^&]*)/gi')
    return oUrl.replace(re, paramName + '=' + replaceWith)
}

export const isWxBrowser = function () {
    let ua = window.navigator.userAgent.toLowerCase();
    if (ua.match(/MicroMessenger/i) == 'micromessenger') {
        return true;
    } else {
        return false;
    }
}

export const getUrlQuery = function () {
    return window.location.href.substring(window.location.href.indexOf('?'))
}

export const queryStringToObject = function (str) {
    if (str.indexOf('?') > -1) {
        str = str.substring(str.indexOf('?') + 1)
    }
    const data = {}
    const strArr = str.split('&')
    for (let i = 0; i < strArr.length; i++) {
        const paramArr = strArr[i].split('=')
        data[paramArr[0]] = paramArr[1] || null
    }

    return data
}

//判断是否电子邮件
export const isEmail = function (str) {
    if (/^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$/.test(str)) {
        return true
    }
    return false;
}

export const copyObject = (obj) => {
    return JSON.parse(JSON.stringify(obj))
}

export const renderAvatar = (path) => {
    return 'http://whiteme2020testcnrunrun.oss-cn-beijing.aliyuncs.com/wechat/avatar/' + path
}

export default {
    getUrlParam,
    replaceParamVal,
    isWxBrowser,
    getUrlQuery,
    isEmail,
    copyObject,
}
