import {Flex, IconButton, Spacer, Tab, TabList, TabPanel, TabPanels, Tabs} from "@chakra-ui/react";
import {BiLogOut} from "react-icons/bi";
import {useAuth} from "react-oidc-context";
import Clients from "../../client/Clients";

const UserAccount = () => {

    const auth = useAuth()

    return (
        <Flex maxWidth="1500px" margin="auto">
            <Tabs>
                <TabList>
                    <Tab>Clients</Tab>
                    <Tab>User settings</Tab>
                </TabList>
                <TabPanels>
                    <TabPanel>
                        <Clients/>
                    </TabPanel>
                    <TabPanel></TabPanel>
                </TabPanels>
            </Tabs>
            <Spacer/>
            <IconButton aria-label="Sign out" as={BiLogOut} variant="ghost"
            onClick={() => auth.removeUser()}/>
        </Flex>
    )
}

export default UserAccount