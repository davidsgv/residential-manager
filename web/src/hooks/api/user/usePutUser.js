import { useMutation } from '@tanstack/react-query'
import useNotify from '../../useNotify';
import { updateUserApi } from '../../../apis/user';
import { useQueryClient } from '@tanstack/react-query';


export default function usePutUser(){
    const {notify, updateSuccess, updateFail} = useNotify("users-update-toast");

    const queryClient = useQueryClient()

    return useMutation({
        mutationFn: (data) => {
            console.log(data)
            notify('Updating user')
            return updateUserApi(data.userId, data).then((data)=>{
                if(data.status == "success"){
                    updateSuccess("User updated")
                    return data.data
                }
                console.log(data)
                updateFail("Failed updating user", data?.data?.message)
                Promise.reject(data.message)
            })
        },
        onSuccess: () => {
            queryClient.invalidateQueries("users")
        },
    })
}