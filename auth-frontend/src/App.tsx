import {BrowserRouter as Router} from "react-router-dom";
import {ChakraProvider} from "@chakra-ui/react";
import {oidcConfig, onSignInCallback} from "./common/oidcConfig";
import {AuthProvider} from 'react-oidc-context';
import BodyRoutes from "./BodyRoutes";

const App = () => {
    return (
        <Router>
            <AuthProvider {...oidcConfig} onSigninCallback={onSignInCallback}>
                <ChakraProvider>
                    <BodyRoutes/>
                </ChakraProvider>
            </AuthProvider>
        </Router>
    )
}

export default App
