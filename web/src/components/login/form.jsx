import {
    FormControl,
    FormLabel,
    FormErrorMessage,
    Input,
    InputGroup,
    InputRightElement,
    Button,
    InputLeftElement,
    Alert,
    AlertIcon,
    AlertTitle,
} from '@chakra-ui/react'
import { ArrowForwardIcon, EmailIcon, LockIcon } from '@chakra-ui/icons'
import useLogin from '../../hooks/useLogin'

const errorColor = 'red.100';
const grayColor = 'gray.300';

export default function Login() {
    const { mail, password, login } = useLogin()

    return (
        <>
            <FormControl isRequired isInvalid={mail.error}>
                <FormLabel>Email address</FormLabel>
                <InputGroup size="sm" variant='filled'>
                    <InputLeftElement pointerEvents='none'>
                        <EmailIcon color={mail.error ? errorColor : grayColor} />
                    </InputLeftElement>
                    <Input
                        type='email'
                        errorBorderColor={errorColor}
                        placeholder='User Email'
                        onChange={mail.handleChange}
                        value={mail.value}
                    />
                </InputGroup>
                {mail.error && <FormErrorMessage>Invalid Email Address</FormErrorMessage>}
            </FormControl>

            <FormControl mt={5} isRequired={true} isInvalid={password.error}>
                <FormLabel>Password</FormLabel>
                <InputGroup size='sm' variant='filled'>
                    <InputLeftElement pointerEvents='none'>
                        <LockIcon color={password.error ? errorColor : grayColor} />
                    </InputLeftElement>
                    <Input
                        type={password.show ? 'text' : 'password'}
                        placeholder='Password'
                        value={password.value}
                        onChange={password.handleChange}
                    />
                    <InputRightElement width='4.5rem'>
                        <Button h='1.5rem' size='md' onClick={password.toogle}>
                            {password.show ? 'Hide' : 'Show'}
                        </Button>
                    </InputRightElement>
                </InputGroup>
                {password.error && <FormErrorMessage>Invalid Password</FormErrorMessage>}
            </FormControl>

            <Button
                rightIcon={<ArrowForwardIcon />}
                colorScheme='teal'
                mt={7}
                isLoading={login.button}
                size="sm"
                loadingText='Loading'
                variant='outline'
                spinnerPlacement='end'
                onClick={login.loginFunc}
            >
                Log-in
            </Button>

            {login.message &&
                <Alert status='error' size="sm" variant="solid" mt={10}>
                    <AlertIcon />
                    <AlertTitle>{login.message}</AlertTitle>
                </Alert>
            }
        </>
    )
}