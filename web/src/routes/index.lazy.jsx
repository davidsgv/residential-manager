import { createLazyFileRoute } from '@tanstack/react-router'
import Config from '../config'
import LoginForm from '../components/login/form'
import { Center, Container, Heading, StackDivider, VStack } from '@chakra-ui/react'

export const Route = createLazyFileRoute('/')({
  component: Index,
})

function Index() {
  return (
    <Center height="100%">
      <VStack>
        <Heading as="h1" size='xl'>Residential Manager</Heading>
        <Container mt={10} centerContent={true} maxW="400px">
          <LoginForm />
        </Container>
      </VStack>
    </Center>
  )
}