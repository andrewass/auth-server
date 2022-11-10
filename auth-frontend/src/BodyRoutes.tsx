import {Route, Routes} from "react-router-dom";
import ProtectedRoute from "./common/ProtectedRoute";
import UserAccount from "./user/account/UserAccount";
import ClientRegistration from "./client/registration/ClientRegistration";
import InternalAuthentication from "./user/authentication/InternalAuthentication";
import ExternalAuthentication from "./user/authentication/ExternalAuthentication";

const BodyRoutes = () => {

    return (
        <Routes>
            <Route path="*"
                element={
                    <ProtectedRoute>
                        <UserAccount/>
                    </ProtectedRoute>
                }
            />

            <Route path="/client-registration"
                element={
                    <ProtectedRoute>
                        <ClientRegistration/>
                    </ProtectedRoute>
                }
            />
            <Route path="/authentication/internal" element={<InternalAuthentication/>}/>
            <Route path="/authentication/external" element={<ExternalAuthentication/>}/>
        </Routes>
    )
}

export default BodyRoutes