import {
    Box,
    Button,
    Card,
    CardBody,
    CardHeader,
    Heading,
    HStack,
    IconButton,
    Input,
    InputGroup,
    InputRightElement,
    Stack,
    Text
} from "@chakra-ui/react";
import {NavigationBar} from "../../navigation/NavigationBar";
import {ClientTypes, mapToClient} from "./clientTypes";
import {useLocation} from "react-router-dom";
import {rotateSecretConfig} from "../api/clientApi";
import {useMutation} from "@tanstack/react-query";
import {useAxiosWrapper} from "../../config/axiosWrapper";
import React, {useState} from "react";
import {BiCopy} from "react-icons/bi";


export const ClientDetails = () => {
    const {state} = useLocation()
    const {axiosPost} = useAxiosWrapper()
    const [client, setClient] = useState(state as ClientTypes)

    const decideSecretDisplay = () => {
        return client.clientSecret ? client.clientSecret :
            "********************************"
    }

    const rotateSecret = () => {
        return axiosPost(rotateSecretConfig(client.clientId))
    }

    const mutation = useMutation(rotateSecret, {
        onSuccess: (data) => {
            setClient(mapToClient(data))
        },
        onError: (error) => console.log("Unable to rotate client secret : " + error)
    })

    return (
        <Box>
            <NavigationBar/>
            <Card marginTop="5%" marginLeft="15%">
                <CardHeader>
                    <Heading size='md'>{client.name}</Heading>
                </CardHeader>
                <CardBody>
                    <Stack spacing="4">
                        <Box>
                            <Heading size="xs">
                                Client ID :
                            </Heading>
                            <InputGroup width="350px">
                                <Input type="text" value={client.clientId} isDisabled/>
                                <InputRightElement
                                    children={<IconButton
                                        onClick={() => navigator.clipboard.writeText(client.clientId)}
                                        aria-label="Copy Client ID"
                                        size="sm"
                                        icon={<BiCopy/>}/>}
                                />
                            </InputGroup>
                        </Box>

                        <HStack spacing="25px">
                            <Box>
                                <Heading size="xs">
                                    Client Secret :
                                </Heading>
                                <InputGroup width="350px">
                                    <Input type='text' value={decideSecretDisplay()} isDisabled/>
                                    <InputRightElement
                                        children={<IconButton
                                            onClick={() => navigator.clipboard.writeText(client.clientSecret)}
                                            aria-label="Copy Client Secret"
                                            size="sm"
                                            icon={<BiCopy/>}/>}
                                    />
                                </InputGroup>
                            </Box>
                            <Box>
                                <Button size="sm" onClick={() => mutation.mutate()}>
                                    Rotate Secret
                                </Button>
                            </Box>
                        </HStack>

                        <Box>
                            <Heading size="xs">
                                Domain
                            </Heading>
                            <Text pt="2" fontSize="sm">
                                Placeholder text for domain
                            </Text>
                        </Box>
                    </Stack>
                </CardBody>
            </Card>
        </Box>
    );
}