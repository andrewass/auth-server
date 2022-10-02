import React, {useState} from "react";
import {signInUser} from "../userApi";
import {Button, FormControl, FormLabel, Input} from "@chakra-ui/react";

interface Props {
    setDisplaySignUp: (value: boolean) => void
    redirectToPage: () => void
}

const SignIn = ({setDisplaySignUp, redirectToPage}: Props) => {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")

    const submitSignIn = (event: React.FormEvent<HTMLElement>) => {
        event.preventDefault()
        signInUser({email, password})
            .then(() => redirectToPage())
            .catch(error => console.log(error))
    }

    return (
        <>
            <form onSubmit={submitSignIn}>
                <FormControl isRequired>
                    <FormLabel>Email</FormLabel>
                    <Input type="email" placeholder="test@test.com" size="lg"
                           onChange={event => setEmail(event.currentTarget.value)}
                    />
                </FormControl>
                <FormControl isRequired mt={6}>
                    <FormLabel>Password</FormLabel>
                    <Input type="password" placeholder="*******" size="lg"
                           onChange={event => setPassword(event.currentTarget.value)}/>
                </FormControl>
                <Button variant="outline" colorScheme="teal" type="submit" width="full" mt={4}>
                    Sign In
                </Button>
                <Button variant="outline" colorScheme="teal" width="full"
                        onClick={() => setDisplaySignUp(true)} mt={4}>
                    Create new account
                </Button>
            </form>
        </>
    )
}

export default SignIn