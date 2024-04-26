import { useQuery } from '@tanstack/react-query'
import { getUserByIdApi } from '../../../apis/user';
// import useNotify from '../../useNotify';
// import { useState } from 'react';

export default function useQueryUserById(idParam){
    // const [id, setId] = useState(idParam)
    //const {notify, updateFail, close} = useNotify("users-fetch-id-toast");

    return useQuery({ 
        queryKey: ['users', idParam],
        queryFn: async ()=>{
            return await getUserByIdApi(idParam).then((data)=>{
                if(data.status == "success"){
                    return data.data
                }
                else{
                    Promise.reject(data)
                }
            });
        }
    });
}