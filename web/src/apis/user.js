import { del, get, post, put } from "./fetch/fetch"

export async function getUsersApi(){
    return await get("/v1/users", undefined, true)
}

export async function getUserByIdApi(userId){
    return await get(`/v1/users/${userId}`, undefined, true)
}

export async function postUserApi(data){
    return await post("/v1/users", data, true)
}

export async function updateUserApi(userId, data){
    return await put(`/v1/users/${userId}`, data, true)
}

export async function deleteUserApi(userId){
    return await del(`/v1/users/${userId}`, null, true)
}