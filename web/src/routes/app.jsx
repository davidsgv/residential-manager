import { Outlet, createFileRoute, redirect } from "@tanstack/react-router"
import { isAuthenticate } from "../helpers/jwt"
import { Container } from '@chakra-ui/react'
import Nav from "../components/navigation/nav"

export const Route = createFileRoute('/app')({
    beforeLoad: async ({ location }) => {
        if (!isAuthenticate()) {
            throw redirect({
                to: '/',
            })
        }
    },
    component: App
})

function App({ }) {
    return (
        <Container maxW='container.lg'>
            <Nav />
            <Outlet />
        </Container>
    )
}