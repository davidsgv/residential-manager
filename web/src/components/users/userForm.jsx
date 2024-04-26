import {
    FormControl,
    FormLabel,
    InputGroup,
    InputLeftElement,
    Input,
    FormErrorMessage,
    HStack,
    Button,
    Alert,
    AlertIcon,
    AlertTitle
} from "@chakra-ui/react"
import { EmailIcon } from "@chakra-ui/icons"
import useUserForm from "../../hooks/forms/useUserForm"
import RolSelect from "../form/rolSelect";

const errorColor = 'red.100';
const grayColor = 'gray.300';

export default function UserForm(props) {
    const { onCreateUser, userId } = props
    const { mail, rol, block, apartment, submitForm, message, button } = useUserForm(onCreateUser, userId)

    return (
        <>
            <FormControl isRequired isInvalid={mail.error} isDisabled={mail.disabled}>
                <FormLabel>Email address</FormLabel>
                <InputGroup size="sm" variant='filled'>
                    <InputLeftElement pointerEvents='none'>
                        <EmailIcon color={mail.error ? errorColor : grayColor} />
                    </InputLeftElement>
                    <Input
                        type='email'
                        // errorBorderColor={errorColor}
                        placeholder='User Email'
                        onChange={mail.handleChange}
                        value={mail.value}
                    />
                </InputGroup>
                {mail.error && <FormErrorMessage>Invalid Email Address</FormErrorMessage>}
            </FormControl >

            <RolSelect mt={10} onChange={rol.onChange} value={rol.value} />

            <HStack mt={10}>
                <FormControl isRequired isInvalid={block.error}>
                    <FormLabel>Apartment Block</FormLabel>
                    <Input
                        type='text'
                        // errorBorderColor={errorColor}
                        placeholder='User Block'
                        onChange={block.handleChange}
                        value={block.value}
                        size="sm"
                        variant='filled'
                    />
                </FormControl >
                <FormControl isRequired isInvalid={apartment.error}>
                    <FormLabel>Apartment Number</FormLabel>
                    <Input
                        type='text'
                        // errorBorderColor={errorColor}
                        placeholder='User Block'
                        onChange={apartment.handleChange}
                        value={apartment.value}
                        size="sm"
                        variant='filled'
                    />
                </FormControl >
            </HStack>


            <HStack justifyContent="end">
                <Button
                    colorScheme='teal'
                    size="sm"
                    mt={10}
                    onClick={submitForm}
                    isLoading={button}
                    loadingText='Loading'
                    variant='outline'
                    spinnerPlacement='end'>Save</Button>
            </HStack>
            {message &&
                <Alert status='error' size="sm" variant="solid" mt={10}>
                    <AlertIcon />
                    <AlertTitle>{message}</AlertTitle>
                </Alert>
            }
        </>
    )
}