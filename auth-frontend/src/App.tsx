import React from 'react';
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import ClientUserConfirmation from "./user/confirmation/ClientUserConfirmation";
import ClientUserAuthentication from "./user/authentication/ClientUserAuthentication";
import {ChakraProvider} from "@chakra-ui/react";
import ClientRegistration from "./client/registration/ClientRegistration";

function App() {
    return (
        <ChakraProvider>
            <Router>
                <Routes>
                    <Route path="/authentication" element={<ClientUserAuthentication/>}/>
                    <Route path="/confirmation" element={<ClientUserConfirmation/>}/>
                    <Route path="/client-registration" element={<ClientRegistration/>}/>
                </Routes>
            </Router>
        </ChakraProvider>
    )
}

export default App;
