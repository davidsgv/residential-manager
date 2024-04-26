import {
    Table,
    Thead,
    Tbody,
    Tr,
    Th,
    TableContainer,
} from '@chakra-ui/react'
import useQueryUsers from '../../hooks/api/user/useQueryUsers'
import UsersTableItem from './usersTableItem'
import TableSkeleton from './tableSkeleton'

export default function UsersTable(props) {
    const { onEditClick } = props

    const { isPending, data } = useQueryUsers()

    if (isPending) {
        return <TableSkeleton />
    }

    return (
        <TableContainer>
            <Table variant='simple'>
                <Thead>
                    <Tr>
                        <Th>Mail</Th>
                        <Th>Rol</Th>
                        <Th>Block</Th>
                        <Th isNumeric>Apartment Number</Th>
                        <Th>Actions</Th>
                    </Tr>
                </Thead>
                <Tbody>


                    {data?.map((item) => {
                        let props = {
                            id: item.id,
                            mail: item.mail,
                            rol: item.rol,
                            apartment: item.aparment
                        }
                        return <UsersTableItem key={props.id} {...props} onEditClick={onEditClick} />
                    })}
                </Tbody>
            </Table>
        </TableContainer>
    )
}
