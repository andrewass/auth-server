import React from 'react';
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import SignIn from "./pages/SignIn";
import SignUp from "./pages/SignUp";
import Confirmation from "./pages/Confirmation";

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/sign-in" element={<SignIn/>}/>
                <Route path="/sign-up" element={<SignUp/>}/>
                <Route path="/confirmation" element={<Confirmation/>}/>
            </Routes>
        </Router>
    )
}

export default App;
