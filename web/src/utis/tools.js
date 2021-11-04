const Tools = {}
Tools.getUrlParam = function (variable) {
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

Tools.replaceParamVal = function (paramName, replaceWith, oUrl = "") {
    oUrl = oUrl ? oUrl : window.location.href.toString()
    let re = eval('/(' + paramName + '=)([^&]*)/gi')
    return oUrl.replace(re, paramName + '=' + replaceWith)
}

Tools.isWxBrowser = function () {
    let ua = window.navigator.userAgent.toLowerCase();
    if (ua.match(/MicroMessenger/i) == 'micromessenger') {
        return true;
    } else {
        return false;
    }
}

Tools.linkXKF = function (source = 'h5') {
    window.open('https://xiaokefu.com.cn?source=' + source)
}

Tools.getUrlQuery = function () {
    return window.location.href.substring(window.location.href.indexOf('?'))
}

Tools.link = function (path) {
    window.location.href = window.location.href.substring(0, window.location.href.indexOf('/#/') + 2) + path + Tools.getUrlQuery()
}

Tools.queryStringToObject = function(str){
    if(str.indexOf('?') > -1){
        str = str.substring(str.indexOf('?') + 1)
    }
    const data = {}
    const strArr = str.split('&')
    for(let i=0;i<strArr.length;i++){
        const paramArr = strArr[i].split('=')
        data[paramArr[0]] = paramArr[1] || null
    }

    return data
}

export default Tools
