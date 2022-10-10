import {Button, FormControl, FormLabel, HStack, Input, Select, Stack, Tag, TagCloseButton} from "@chakra-ui/react";
import React, {useState} from "react";
import {useAxiosWrapper} from "../../config/axiosWrapper";
import {registerClientConfig} from "../api/clientApi";

export interface RegisterClientData {

}

const ClientRegistrationForm = () => {
    const [clientName, setClientName] = useState("")
    const [applicationType] = useState([])
    const [redirectUris, setRedirectUris] = useState<string[]>([])
    const [tokenEndpointAuthMethod] = useState()
    const {post} = useAxiosWrapper()


    const registerClient = () => {
        post(registerClientConfig({
            "client_name": clientName,
            "redirect_uris": redirectUris
        })).catch(error => console.log(error))
    }

    const removeRedirectUri = (index: number) => {
        redirectUris.splice(index, 1)
        setRedirectUris([...redirectUris])
    }

    const addUriIfSpaceIsPressed = (event: React.KeyboardEvent<HTMLInputElement>) => {
        if (event.key === " " && event.target.value !== "") {
            setRedirectUris([...redirectUris].concat([event.target.value]))
            event.target.value = "";
        }
    }

    return (
        <form>
            <Stack maxWidth={1000} margin="auto" spacing={5} marginTop={15}>
                <FormControl>
                    <FormLabel>Client name</FormLabel>
                    <Input type="text" placeholder="client name" autoComplete="off"
                           onChange={event => setClientName(event.currentTarget.value)}
                    />
                </FormControl>

                <FormControl>
                    <FormLabel>Application types</FormLabel>
                    <Select placeholder='Select application types'>
                        <option value='option1'>Option 1</option>
                        <option value='option2'>Option 2</option>
                        <option value='option3'>Option 3</option>
                    </Select>
                </FormControl>

                <FormControl>
                    <FormLabel>Redirect URIs</FormLabel>
                    <Input placeholder="Add redirect uri" onKeyUp={event => addUriIfSpaceIsPressed(event)}/>
                    <HStack spacing={4} marginTop={2}>
                        {redirectUris.map((uri, index) => (
                            <Tag key={index}>
                                {uri}
                                <TagCloseButton onClick={() => removeRedirectUri(index)}/>
                            </Tag>
                        ))}
                    </HStack>
                </FormControl>

                <FormControl>
                    <Button variant="outline" type="button" colorScheme="teal" onClick={registerClient}>
                        Register client
                    </Button>
                </FormControl>
            </Stack>
        </form>
    )
}

export default ClientRegistrationForm