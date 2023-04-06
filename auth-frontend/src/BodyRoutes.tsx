import {Route, Routes} from "react-router-dom";
import ProtectedRoute from "./config/ProtectedRoute";
import {Account} from "./user/Account";
import {InternalAuthentication} from "./authentication/InternalAuthentication";
import {ExternalAuthentication} from "./authentication/ExternalAuthentication";
import {Clients} from "./client/Clients";
import {ClientInformation} from "./client/presentation/ClientInformation";

export const BodyRoutes = () => {
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
                           <ClientInformation/>
                       </ProtectedRoute>
                   }
            />

            <Route path="/authentication/internal" element={<InternalAuthentication/>}/>
            <Route path="/authentication/external" element={<ExternalAuthentication/>}/>
        </Routes>
    );
}