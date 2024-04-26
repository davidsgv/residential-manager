import { post } from "./fetch/fetch"

export async function loginApi(mail, password){
    const response =  await post("/v1/auth/login", {
        mail: mail,
        password: password
    })
    return response
}