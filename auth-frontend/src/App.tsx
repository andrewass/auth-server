import React from 'react';
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import Confirmation from "./user/confirmation/Confirmation";
import Authentication from "./user/authentication/Authentication";
import {ChakraProvider} from "@chakra-ui/react";

function App() {
    return (
        <ChakraProvider>
            <Router>
                <Routes>
                    <Route path="/authentication" element={<Authentication/>}/>
                    <Route path="/confirmation" element={<Confirmation/>}/>
                </Routes>
            </Router>
        </ChakraProvider>
    )
}

export default App;
