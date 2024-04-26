const types = {
    "None": "None",
    "Mail": "Mail",
    "Password": "Password",
    "Number": "Number"
}
Object.freeze(types)

function validateEmail(mail){
    var reg = /^[a-zA-Z0-9._-]+@[a-zA-Z0–9.-]+\.[a-zA-Z]{2,4}$/
    return reg.test(mail)
}

function validatePassword(password){
    return password.length > 2
}

export default function validate(string, type){
    var validation = false;
    switch(type){
        case types.Mail:
            validation = validateEmail(string);
            break
        case types.Password:
            validation = validatePassword(string);
            break
        case types.None:
            validation = true
            break
        default:
            throw new Error("type " + type + " does not exist")
    }
    return validation;
}

export {types}