import React, {useState} from "react";
import {signInUserConfig} from "./api/authenticationApi";
import {Button, FormControl, FormLabel, Input, Stack} from "@chakra-ui/react";
import {useAxiosWrapper} from "../config/axiosWrapper";
import axios from "axios";
import toast, {Toaster} from "react-hot-toast";


export const SignInForm = ({setDisplaySignUp, postProcessSignIn}: {
    setDisplaySignUp: (value: boolean) => void,
    postProcessSignIn: (email: string) => void
}) => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const {axiosPost} = useAxiosWrapper();

    const submitSignIn = async (event: React.FormEvent<HTMLElement>) => {
        event.preventDefault();
        try {
            await axiosPost(signInUserConfig(email, password));
            postProcessSignIn(email);
        } catch (error){
            if(axios.isAxiosError(error) && error.response){
                if(error.response.status === 401){
                    toast.error("Incorrect email or password");
                    setEmail("");
                    setPassword("");
                }
            }
        }
    }

    return (
        <form onSubmit={submitSignIn}>
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
                <Button variant="outline" colorScheme="teal" type="submit" width="full" mt={4}>
                    Sign In
                </Button>
                <Button variant="outline" colorScheme="teal" onClick={() => setDisplaySignUp(true)} mt={4}>
                    Create new account
                </Button>
            </Stack>
            <Toaster/>
        </form>
    );
}