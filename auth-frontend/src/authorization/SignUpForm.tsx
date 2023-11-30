import React, {useState} from "react";
import {Button, FormControl, FormErrorMessage, FormLabel, Input, Stack} from "@chakra-ui/react";
import axios from "axios";
import toast, {Toaster} from "react-hot-toast";
import {useAxiosWrapper} from "../config/axiosWrapper";
import {signUpUserConfig} from "./api/authenticationApi";


export const SignUpForm = ({setDisplaySignUp, postProcessSignUp}: {
    setDisplaySignUp: (value: boolean) => void,
    postProcessSignUp: (email: string) => void
}) => {
    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [confirmedPassword, setConfirmedPassword] = useState<string>("");
    const {axiosPost} = useAxiosWrapper();

    const submitSignUp = async (event: React.FormEvent<HTMLElement>) => {
        event.preventDefault();
        try {
            await axiosPost(signUpUserConfig(email, password));
            postProcessSignUp(email);
        } catch (error) {
            if (axios.isAxiosError(error) && error.response) {
                if (error.response.status === 409) {
                    toast.error("Email already registered to an account");
                    setEmail("");
                    setPassword("");
                    setConfirmedPassword("");
                }
            }
        }
    }

    return (
        <form onSubmit={submitSignUp}>
            <Stack maxWidth={500} margin="auto" spacing={5} marginTop={90}>
                <FormControl isRequired>
                    <FormLabel>Email</FormLabel>
                    <Input type="email" value={email}
                           placeholder="test@test.com" size="lg" autoComplete="off"
                           onChange={event => setEmail(event.currentTarget.value)}
                    />
                </FormControl>
                <FormControl isRequired mt={6}>
                    <FormLabel>Password</FormLabel>
                    <Input type="password" value={password}
                           placeholder="*******" size="lg" autoComplete="off"
                           onChange={event => setPassword(event.currentTarget.value)}/>
                </FormControl>
                <FormControl isRequired mt={6} isInvalid={password !== confirmedPassword}>
                    <FormLabel>Confirm Password</FormLabel>
                    <Input type="password" value={confirmedPassword}
                           placeholder="*******" size="lg" autoComplete="off"
                           onChange={event => setConfirmedPassword(event.currentTarget.value)}/>
                    {password !== confirmedPassword &&
                        <FormErrorMessage>Confirmed password does not match</FormErrorMessage>
                    }
                </FormControl>
                <Button variant="outline" colorScheme="teal"
                        type="submit" mt={4} isDisabled={password !== confirmedPassword}>
                    Sign Up
                </Button>
                <Button variant="outline" colorScheme="teal"
                        onClick={() => setDisplaySignUp(false)} mt={4}>
                    Sign in to existing account
                </Button>
            </Stack>
            <Toaster/>
        </form>
    );
}