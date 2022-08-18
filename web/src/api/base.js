import request from "./request";

export const sendCodeReq = (params) => {
    return request.post("/sendCode",params)
}

export const loginRegisterReq = (params) => {
    return request.post("/loginReg",params)
}
