import {Button, FormControl, FormLabel, Input, Stack} from "@chakra-ui/react";
import {useState} from "react";
import {useAxiosWrapper} from "../../config/axiosWrapper";
import {useMutation} from "@tanstack/react-query";
import {registerClientConfig} from "../api/clientApi";
import {useAuth} from "react-oidc-context";
import {useNavigate} from "react-router-dom";
import {mapToClientDetails} from "../clientTypes";


export const ClientRegistrationForm = () => {
    const [applicationName, setApplicationName] = useState<string>("");
    const [applicationDescription, setApplicationDescription] = useState<string>("");
    const [homepageUrl, setHomepageUrl] = useState<string>("");
    const [callbackUrl, setCallbackUrl] = useState<string>("");

    const {axiosPost} = useAxiosWrapper();
    const {user} = useAuth();
    const navigate = useNavigate();

    const registerClient = () => {
        return axiosPost(registerClientConfig({
                userEmail: user!.profile.email!,
                applicationName: applicationName,
                applicationDescription: applicationDescription,
                homepageUrl: homepageUrl,
                callbackUrl: callbackUrl
            }
        ));
    }

    const mutation = useMutation(registerClient, {
        onSuccess: (data) => {
            navigate("/client/details", {state: mapToClientDetails(data)})
        },
        onError: (error) => console.log("Unable to submit new client : " + error)
    })

    return (
        <form>
            <Stack maxWidth={1000} margin="auto" spacing={5} marginTop={15}>
                <FormControl isRequired>
                    <FormLabel>Application name</FormLabel>
                    <Input type="text" placeholder="Application name" autoComplete="off"
                           onChange={event => setApplicationName(event.currentTarget.value)}
                    />
                </FormControl>

                <FormControl isRequired>
                    <FormLabel>Homepage URL</FormLabel>
                    <Input type="text" placeholder="Homepage URL" autoComplete="off"
                           onChange={event => setHomepageUrl(event.currentTarget.value)}
                    />
                </FormControl>

                <FormControl>
                    <FormLabel>Application description</FormLabel>
                    <Input type="text" placeholder="Application description" autoComplete="off"
                           onChange={event => setApplicationDescription(event.currentTarget.value)}
                    />
                </FormControl>

                <FormControl isRequired>
                    <FormLabel>Authorization Callback URL</FormLabel>
                    <Input type="text" placeholder="Authorization callback URL" autoComplete="off"
                           onChange={event => setCallbackUrl(event.currentTarget.value)}
                    />
                </FormControl>

                <FormControl>
                    <Button variant="outline" type="button" colorScheme="teal"
                            onClick={() => mutation.mutate()}>
                        Register client
                    </Button>
                </FormControl>
            </Stack>
        </form>
    );
}