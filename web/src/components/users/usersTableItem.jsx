import {
    Tr,
    Td,
    IconButton
} from '@chakra-ui/react'
import { EditIcon, DeleteIcon } from '@chakra-ui/icons'
import useDeleteUser from '../../hooks/api/user/useDeleteUser'

export default function UsersTableItem(props) {
    const { id, mail, rol, apartment, onEditClick } = props

    const deleteMutation = useDeleteUser()

    const handleEdit = () => {
        onEditClick(id)
    }

    const handleDelete = () => {
        deleteMutation.mutateAsync(id)
    }

    return (
        <Tr>
            <Td>{mail}</Td>
            <Td>{rol}</Td>
            <Td>{apartment?.block}</Td>
            <Td isNumeric>{apartment?.number}</Td>
            <Td>
                <IconButton
                    colorScheme='teal'
                    aria-label='Edit user'
                    icon={<EditIcon />}
                    isRound
                    onClick={handleEdit}
                />
                <IconButton
                    colorScheme='red'
                    aria-label='Delete user'
                    icon={<DeleteIcon />}
                    isRound
                    ml={5}
                    onClick={handleDelete}
                />
            </Td>
        </Tr>
    )
}