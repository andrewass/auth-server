import {Button, FormControl, FormLabel, Input, Select, Stack} from "@chakra-ui/react";
import React, {useState} from "react";

const ClientRegistrationForm = () => {
    const [clientName, setClientName] = useState("")
    const [applicationType] = useState([])
    const [redirectUris] = useState([])
    const [tokenEndpointAuthMethod] = useState()

    const submitSignIn = (event: React.FormEvent<HTMLElement>) => {

    }

    return (
        <form onSubmit={submitSignIn} method="POST">
            <Stack maxWidth={500} margin="auto" spacing={5} marginTop={15}>
                <FormControl>
                    <FormLabel>Application types</FormLabel>
                    <Select placeholder='Select application types'>
                        <option value='option1'>Option 1</option>
                        <option value='option2'>Option 2</option>
                        <option value='option3'>Option 3</option>
                    </Select>
                </FormControl>

                <FormControl>
                    <Button variant="outline" colorScheme="teal" type="submit">
                        Sign In
                    </Button>
                </FormControl>
            </Stack>
        </form>
    )
}

export default ClientRegistrationForm