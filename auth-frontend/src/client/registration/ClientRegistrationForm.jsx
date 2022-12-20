import {Button, FormControl, FormLabel, HStack, Input, Select, Stack, Tag, TagCloseButton} from "@chakra-ui/react";
import React, {useState} from "react";
import {useAxiosWrapper} from "../../config/axiosWrapper";
import {registerClientConfig} from "../api/clientApi";
import {getSessionEmail} from "../../config/oidcConfig";


const ClientRegistrationForm = () => {
    const [clientName, setClientName] = useState("")
    const [applicationType] = useState([])
    const [redirectUris, setRedirectUris] = useState([])
    const [tokenEndpointAuthMethod] = useState()
    const {post} = useAxiosWrapper()


    const registerClient = () => {
        post(registerClientConfig({
                userEmail: getSessionEmail(),
                redirectUris: redirectUris,
                clientName: clientName
            }
        )).catch(error => console.log(error))
    }

    const removeRedirectUri = (index) => {
        redirectUris.splice(index, 1)
        setRedirectUris([...redirectUris])
    }

    const addUriIfSpaceIsPressed = (event) => {
        if (event.key === " " && event.target.value !== "") {
            setRedirectUris([...redirectUris].concat([event.target.value.trim()]))
            event.target.value = "";
        }
    }

    return (
        <form onSubmit={registerClient}>
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
                    <Button variant="outline" type="submit" colorScheme="teal">
                        Register client
                    </Button>
                </FormControl>
            </Stack>
        </form>
    )
}

export default ClientRegistrationForm