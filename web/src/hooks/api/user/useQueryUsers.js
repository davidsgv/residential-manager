import { useQuery } from '@tanstack/react-query'
import useNotify from '../../useNotify';
import { getUsersApi } from '../../../apis/user';

export default function useQueryUsers(){
    const {notify, updateFail, close} = useNotify("users-fetch-toast");

    return useQuery({ 
        queryKey: ['users'],
        queryFn: async ()=>{
            notify('Fetching users', "fetching...")
            return await getUsersApi().then((data)=>{
                if(data.status == "success"){
                    close()
                    return data.data
                }
                else{
                    updateFail("Failed to fetch", data.message)
                    Promise.reject(data)
                }
            });
        }
    });
}