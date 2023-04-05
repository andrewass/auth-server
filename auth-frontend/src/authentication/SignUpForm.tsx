import React, {useState} from "react";
import {signUpUser} from "./api/authenticationApi";
import {Button, FormControl, FormLabel, Input, Stack} from "@chakra-ui/react";


export const SignUpForm = ({setDisplaySignUp, postProcessSignUp}: {
    setDisplaySignUp: (value: boolean) => void,
    postProcessSignUp: (email: string) => void
}) => {
    const [email, setEmail] = useState<string>("")
    const [password, setPassword] = useState<string>("")
    const [confirmedPassword, setConfirmedPassword] = useState<string>("")

    const submitSignUp = (event: React.FormEvent<HTMLElement>) => {
        event.preventDefault()
        if (password !== confirmedPassword) {
            console.log("Password mismatch")
        } else {
            signUpUser(email, password)
                .then(() => postProcessSignUp(email))
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
    );
}