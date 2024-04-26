import { get } from "./fetch/fetch"

export async function getRolesApi(){
    const response =  await get("/v1/roles", undefined, true)
    return response
}