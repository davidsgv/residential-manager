import { useToast } from '@chakra-ui/react'
import { useState } from 'react';

export default function useNotify(id){
    //const [idState] = useState(id)
    const toast = useToast();

    const updateToast = (options)=>{
        if(toast.isActive(id)){
            toast.update(id, options)
            return 
        }
        toast(options);
    }

    const notify = (title, description)=>{
        updateToast({
            id,
            title,
            description,
            status: "loading",
        })
    }

    const updateSuccess = (title, description)=>{
        updateToast({
            title,
            description,
            status: "success"
        })
    }

    const updateFail = (title, description) =>{
        updateToast({
            title,
            description,
            status: "error",
        })
    }

    const close = () =>{
        toast.close(id)
    }

    return {notify, updateSuccess, updateFail, close}
}