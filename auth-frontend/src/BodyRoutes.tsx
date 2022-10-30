import {Route, Routes} from "react-router-dom";
import ProtectedRoute from "./config/ProtectedRoute";
import UserAccount from "./user/account/UserAccount";
import ClientRegistration from "./client/registration/ClientRegistration";
import UserAuthentication from "./user/authentication/UserAuthentication";
import UserConfirmation from "./user/confirmation/UserConfirmation";

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

            <Route path="/authentication" element={<UserAuthentication/>}/>
            <Route path="/confirmation" element={<UserConfirmation/>}/>
        </Routes>
    )
}

export default BodyRoutes