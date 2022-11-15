import {Route, Routes} from "react-router-dom";
import ProtectedRoute from "./common/ProtectedRoute";
import Account from "./user/account/Account";
import InternalAuthentication from "./user/authentication/InternalAuthentication";
import ExternalAuthentication from "./user/authentication/ExternalAuthentication";
import Clients from "./client/Clients";
import ClientDetails from "./client/presentation/ClientDetails";

const BodyRoutes = () => {

    return (
        <Routes>
            <Route path="/*"
                   element={
                       <ProtectedRoute>
                           <Account/>
                       </ProtectedRoute>
                   }
            />

            <Route path="/clients"
                   element={
                       <ProtectedRoute>
                           <Clients/>
                       </ProtectedRoute>
                   }
            />

            <Route path="/client/details"
                   element={
                       <ProtectedRoute>
                           <ClientDetails/>
                       </ProtectedRoute>
                   }
            />

            <Route path="/authentication/internal" element={<InternalAuthentication/>}/>
            <Route path="/authentication/external" element={<ExternalAuthentication/>}/>
        </Routes>
    )
}

export default BodyRoutes