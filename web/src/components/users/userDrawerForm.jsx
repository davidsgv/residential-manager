import {
    Drawer,
    DrawerOverlay,
    DrawerContent,
    DrawerCloseButton,
    DrawerHeader,
    DrawerBody,
} from "@chakra-ui/react"
import UserForm from "./userForm"

export default function UserDrawerForm(props) {
    const { isOpen, onClose, onCreateUser, userId } = props
    // const btnRef = React.useRef()

    return (
        <>
            <Drawer
                isOpen={isOpen}
                placement='right'
                onClose={onClose}
                size="md"
            // finalFocusRef={btnRef}
            >
                <DrawerOverlay />
                <DrawerContent>
                    <DrawerCloseButton />
                    <DrawerHeader>Create your account</DrawerHeader>
                    <DrawerBody>
                        <UserForm onCreateUser={onCreateUser} userId={userId} />
                    </DrawerBody>
                </DrawerContent>
            </Drawer>
        </>
    )
}