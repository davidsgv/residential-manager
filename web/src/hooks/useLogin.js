import useInput from './useInput'
import { types } from '../helpers/validations'
import { useBoolean } from '@chakra-ui/react'
import useNotify from './useNotify'
import { loginApi } from '../apis/auth'
import { useState } from 'react'
import useToken from './useToken'

export default function useLogin() {
    const mail = useInput("", types.Mail);
    const password = useInput("", types.Password);
    const [show, setShow] = useBoolean(false); //show password

    const [button, setButton] = useBoolean(false); //loading button
    const [message, setMessage] = useState(null);

    const [, setToken] = useToken()


    const {notify, updateSuccess, updateFail} = useNotify("login-toast");

    const login = () => {
        if(mail.error || mail.empty){
            setMessage("Complete the login form");
            return
        }
        if(password.error || password.empty){
            setMessage("Complete the login form");
            return
        } 

        if (button == false) {
            notify('Login pending', "please wait")

            loginApi(mail.value, password.value)
            .then(data => {
                if(data.status != "success"){
                    loginFail('Login fail', data.message)
                    setMessage(data.message)
                    return
                }
                updateSuccess('Login success', "succesfuly login")
                setToken(data.data.token)
            })
            .catch(reason =>{
                console.log(reason)
                loginFail("Something went wrong")
            });
        }
        setButton.on()
    }

    const loginFail = (message) =>{
        updateFail('Login fail', message)
        setButton.off()
    }

    return {
        mail,
        password: {
            show: show,
            toogle: setShow.toggle,
            error: password.error,
            handleChange: password.handleChange,
            value: password.value,
        },
        login:{
            button,
            message,
            loginFunc: login
        }
    }
}