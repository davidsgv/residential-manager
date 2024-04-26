import { Wrap, WrapItem, Center, Text } from '@chakra-ui/react'
import { Link } from '@tanstack/react-router'



export default function Nav() {
    return (
        <nav>
            <Wrap justify="end" mt={5} mb={5} spacing={5}>
                <WrapItem>
                    <Center>
                        <Link to="/app/users">
                            {({ isActive }) => {
                                return (
                                    <Text as={isActive ? 'u' : 'i'}>Users</Text>
                                )
                            }}
                        </Link>
                    </Center>
                </WrapItem>
                <WrapItem>
                    <Center>
                        <Link to="/app/about/test">
                            {({ isActive }) => {
                                return (
                                    <Text as={isActive ? 'u' : 'i'}>Roles</Text>
                                )
                            }}
                        </Link>
                    </Center>
                </WrapItem>
            </Wrap>
        </nav>
    )
}