function getToken(){
    return localStorage.getItem("token")
}

function setToken(token){
    localStorage.setItem("token", token)
}

function parseJwt (token) {
    var base64Url = token.split('.')[1];
    var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    var jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
}

function getClaims(){
    let token = getToken()
    if (!token){
        return
    }

    let unParsedClaims = parseJwt(token)
    let claims = {...unParsedClaims}

    if (claims.exp){
        claims.exp = new Date(unParsedClaims.exp * 1000);
    }

    return claims
}

function isAuthenticate(){
    let claims = getClaims()

    if (!claims){
        return false
    }

    if (claims.exp < Date.now()){
        return false
    }

    return true
}

export {getToken, getClaims, setToken, isAuthenticate}