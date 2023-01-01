import {BrowserRouter as Router} from "react-router-dom";
import {ChakraProvider} from "@chakra-ui/react";
import {oidcConfig, onSignInCallback} from "./config/oidcConfig";
import {AuthProvider} from 'react-oidc-context';
import BodyRoutes from "./BodyRoutes";
import {QueryClient, QueryClientProvider} from "react-query";
import {ReactQueryDevtools} from "react-query/devtools";

export const queryClient = new QueryClient()

const App = () => {
    return (
        <Router>
            <QueryClientProvider client={queryClient}>
                <ReactQueryDevtools initialIsOpen={false} />
                <AuthProvider {...oidcConfig} onSigninCallback={onSignInCallback}>
                    <ChakraProvider>
                        <BodyRoutes/>
                    </ChakraProvider>
                </AuthProvider>
            </QueryClientProvider>
        </Router>
    )
}

export default App
