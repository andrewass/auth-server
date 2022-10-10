import React, {useState} from "react";
import {signUpUser} from "../api/userApi";
import {Button, FormControl, FormLabel, Input, Stack} from "@chakra-ui/react";

interface Props {
    setDisplaySignUp: (value: boolean) => void
    redirectToPage: () => void
}

const SignUp = ({setDisplaySignUp, redirectToPage}: Props) => {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [confirmedPassword, setConfirmedPassword] = useState("")

    const submitSignUp = (event: React.FormEvent<HTMLElement>) => {
        event.preventDefault()
        if (password !== confirmedPassword) {
            console.log("Password mismatch")
        } else {
            signUpUser(email, password)
                .then(() => redirectToPage())
                .catch(error => console.log(error))
        }
    }

    return (
        <form onSubmit={submitSignUp}>
            <Stack maxWidth={500} margin="auto" spacing={5} marginTop={90}>
                <FormControl isRequired>
                    <FormLabel>Email</FormLabel>
                    <Input type="email" placeholder="test@test.com" size="lg" autoComplete="off"
                           onChange={event => setEmail(event.currentTarget.value)}
                    />
                </FormControl>
                <FormControl isRequired mt={6}>
                    <FormLabel>Password</FormLabel>
                    <Input type="password" placeholder="*******" size="lg" autoComplete="off"
                           onChange={event => setPassword(event.currentTarget.value)}/>
                </FormControl>
                <FormControl isRequired mt={6}>
                    <FormLabel>Confirm Password</FormLabel>
                    <Input type="password" placeholder="*******" size="lg" autoComplete="off"
                           onChange={event => setConfirmedPassword(event.currentTarget.value)}/>
                </FormControl>
                <Button variant="outline" colorScheme="teal" type="submit" mt={4}>
                    Sign Up
                </Button>
                <Button variant="outline" colorScheme="teal" onClick={() => setDisplaySignUp(false)} mt={4}>
                    Sign in to existing account
                </Button>
            </Stack>
        </form>
    )
}

export default SignUp