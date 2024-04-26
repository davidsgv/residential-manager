import Config from "../../config";
import { getToken } from "../../helpers/jwt";

function apiFetch(path, data = undefined, authorized, method = 'GET'){
    const url = Config.API_URL + path
    const options = {
        method,
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        }
    };
    if(authorized){
        options.headers["Authorization"] = getToken()
    }
    if (data) {
        options.body = JSON.stringify(data);
    }
    return fetch(url, options).then(response => response.json());
};

const get = (path, data, authorized) => apiFetch(path, data, authorized);
const post = (path, data, authorized) => apiFetch(path, data, authorized, "POST");
const put = (path, data, authorized) => apiFetch(path, data, authorized, "PUT");
const del = (path, data, authorized) => apiFetch(path, data, authorized, "DELETE");

export {get, post, put, del}