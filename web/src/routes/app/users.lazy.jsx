import { createLazyFileRoute } from '@tanstack/react-router'
import UsersTable from '../../components/users/usersTable'
import { Heading, useDisclosure, Button, Stack } from '@chakra-ui/react'
import UserDrawerForm from '../../components/users/userDrawerForm'
import { useState } from 'react'


export const Route = createLazyFileRoute('/app/users')({
  component: Users,
})

function Users() {
  const { isOpen, onOpen, onClose } = useDisclosure()
  const [userId, setUserId] = useState()

  const onEditClick = (id) => {
    setUserId(id)
    onOpen()
  }

  const onCreateClick = () => {
    setUserId(undefined)
    onOpen()
  }

  const onCreateUser = () => {
    onClose()
  }

  return (
    <>
      <Stack direction='row' spacing={4} align='center' justifyContent="space-between">
        <Heading as='h1' size='xl' mt={5} mb={5}>
          Users
        </Heading>
        <Button colorScheme='teal' onClick={onCreateClick} size="sm" variant='solid'>Create User</Button>
      </Stack>
      <UsersTable onEditClick={onEditClick} />
      <UserDrawerForm isOpen={isOpen} onOpen={onOpen} onClose={onClose} onCreateUser={onCreateUser} userId={userId} />
    </>
  )
}