import { useQuery } from '@tanstack/react-query'
import { getRolesApi } from "../../apis/rol"
import useNotify from '../useNotify';

export default function useRoles(){
    const {updateFail} = useNotify("roles-fetch-toast");

    const query = useQuery({ 
        queryKey: ['roles'], 
        queryFn: async ()=>{
            const response = await getRolesApi()
                .then((data)=>{
                    if(data.status != "success")
                        updateFail("Failed to fetch", data.message)
                    return data
                })
            return response
        }
    })

    return {query}
}