import { useMutation } from '@tanstack/react-query'
import useNotify from '../../useNotify';
import { postUserApi } from '../../../apis/user';
import { useQueryClient } from '@tanstack/react-query';


export default function usePostUser(){
    const {notify, updateSuccess, updateFail} = useNotify("users-post-toast");

    const queryClient = useQueryClient()

    return useMutation({
        mutationFn: (data) => {
            notify('Posting user')
            return postUserApi(data).then((data)=>{
                if(data.status == "success"){
                    updateSuccess("User created")
                    return data.data
                }
                console.log(data)
                updateFail("Failed creating user", data?.data?.message)
                Promise.reject(data.message)
            })
        },
        onSuccess: () => {
            queryClient.invalidateQueries("users")
        },
    })
}