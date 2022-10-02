import {BrowserRouter as Router} from "react-router-dom";
import {ChakraProvider} from "@chakra-ui/react";
import {oidcConfig} from "./config/oidcConfig";
import {AuthProvider} from 'react-oidc-context';
import BodyRoutes from "./BodyRoutes";

const App = () => {
    return (
        <AuthProvider {...oidcConfig}>
            <ChakraProvider>
                <Router>
                    <BodyRoutes/>
                </Router>
            </ChakraProvider>
        </AuthProvider>
    )
}

export default App
