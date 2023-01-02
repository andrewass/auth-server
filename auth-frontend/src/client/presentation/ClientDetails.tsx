import {Box, Button, Card, CardBody, CardHeader, Heading, HStack, Stack, Text} from "@chakra-ui/react";
import NavigationBar from "../../navigation/NavigationBar";
import {Client} from "./client";
import {useLocation} from "react-router-dom";


const ClientDetails = () => {
    const {state} = useLocation()
    const client = state as Client

    const decideSecretDisplay = () => {
        return client.clientSecret ? client.clientSecret :
            "********************************"
    }

    return (
        <Box>
            <NavigationBar/>
            <Card marginTop="5%" marginLeft="15%">
                <CardHeader>
                    <Heading size='md'>{client.name}</Heading>
                </CardHeader>
                <CardBody>
                    <Stack spacing='4'>
                        <Box>
                            <Heading size='xs' textTransform='uppercase'>
                                Client ID
                            </Heading>
                            <Text pt='2' fontSize='sm'>
                                {client.clientId}
                            </Text>
                        </Box>

                        <HStack spacing="25px">
                            <Box>
                                <Heading size='xs' textTransform='uppercase'>
                                    Client Secret
                                </Heading>
                                <Text pt='2' fontSize='sm'>
                                    {decideSecretDisplay()}
                                </Text>
                            </Box>
                            <Box>
                                <Button size="sm">Rotate Secret</Button>
                            </Box>
                        </HStack>

                        <Box>
                            <Heading size='xs' textTransform='uppercase'>
                                Domain
                            </Heading>
                            <Text pt='2' fontSize='sm'>
                                Placeholder text for domain
                            </Text>
                        </Box>
                    </Stack>
                </CardBody>
            </Card>
        </Box>
    )
}

export default ClientDetails