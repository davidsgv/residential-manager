import { useMutation } from '@tanstack/react-query'
import useNotify from '../../useNotify';
import { deleteUserApi } from '../../../apis/user';
import { useQueryClient } from '@tanstack/react-query';


export default function useDeleteUser(){
    const {updateSuccess, updateFail} = useNotify("users-delete-toast");

    const queryClient = useQueryClient()

    return useMutation({
        mutationFn: (userId) => {
            return deleteUserApi(userId).then((data)=>{
                if(data.status == "success"){
                    updateSuccess("User deleted")
                    return data
                }
                updateFail("Failed deleting user", data?.data?.message)
                Promise.reject(data.message)
            })
        },
        onSuccess: () => {
            queryClient.invalidateQueries("users")
        },
    })
}