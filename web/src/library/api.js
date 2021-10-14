import {api,formatParams,config} from './axios'

export function getActivityApi(params){
    return api(Object.assign(formatParams('POST',params),{
        url:`${config.proxy}/social/promotReward/list`
    }))
}