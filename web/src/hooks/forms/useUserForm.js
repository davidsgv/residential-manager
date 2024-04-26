import useInput from '../useInput'
import { types } from '../../helpers/validations'
import { useEffect, useState } from 'react';
import {useBoolean} from '@chakra-ui/react'
import { getUserByIdApi } from '../../apis/user';
import usePostUser from '../api/user/usePostUser';
import usePutUser from '../api/user/usePutUser';
//import useUser from '../api/useUser';
// import useNotify from './useNotify'
// import { loginApi } from '../apis/auth'
// import { useState } from 'react'
// import useToken from './useToken'

export default function useUserForm(onSubmitForm, userId) {
    const mutation = usePostUser()
    const updateMutate = usePutUser()

    const mail = useInput("", types.Mail, userId != undefined ? true : false);
    const [rol, setRol] = useState("")
    const block = useInput("", types.None)
    const apartment = useInput("", types.None)
    const [button, setButton] = useBoolean(false); //loading button
    const [message, setMessage] = useState(null);

    useEffect(()=>{
        if(!userId)
            return
        
        getUserByIdApi(userId).then((data)=>{
            if(data.status == "success"){
                mail.setValue(data?.data?.mail)
                setRol(data?.data?.rol)
                block.setValue(data?.data?.aparment?.block)
                apartment.setValue(data?.data?.aparment?.number)
            }
        })
    }, [])


    const handleChangeRol = (event)=>{
        setRol(event.target.value)
    }

    const createUser = async ()=>{
        return await mutation.mutateAsync({
            mail:mail.value,
            rol : rol,
            apartment:{
                block: block.value,
                number: apartment.value
            }
        })
    }

    const updateUser = async ()=>{
        return await updateMutate.mutateAsync({
            userId: userId,
            rol : rol,
            apartment:{
                block: block.value,
                number: apartment.value
            }
        })
    }

    const submitForm = ()=>{
        if(mail.error || mail.empty){
            setMessage("Complete the mail input");
            return
        }
        if(rol == ""){
            setMessage("Select a Rol");
            return
        }
        if(block.empty || apartment.empty){
            setMessage("Complete the apartment info");
            return
        }

        setButton.on()

        var submitPromise
        if(userId){
            submitPromise = updateUser()
        }
        else{
            submitPromise = createUser()
        }
        
        submitPromise.then((data)=>{
            console.log(data)
            setButton.off()
            onSubmitForm()
            return data
        })
    }

    return {
        mail,
        block,
        apartment,
        rol: {
            onChange: handleChangeRol,
            value: rol
        },
        submitForm,
        message,
        button
    }
}