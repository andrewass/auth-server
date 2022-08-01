import React from 'react';
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import Confirmation from "./pages/Confirmation";
import Authentication from "./pages/Authentication";

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/authentication" element={<Authentication/>}/>
                <Route path="/confirmation" element={<Confirmation/>}/>
            </Routes>
        </Router>
    )
}

export default App;
