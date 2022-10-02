import {Route, Routes} from "react-router-dom";
import ProtectedRoute from "./config/ProtectedRoute";
import UserAccount from "./user/account/UserAccount";
import ClientRegistration from "./client/registration/ClientRegistration";
import ClientUserAuthentication from "./user/authentication/ClientUserAuthentication";
import ClientUserConfirmation from "./user/confirmation/ClientUserConfirmation";
import UserAuthentication from "./user/authentication/UserAuthentication";

const BodyRoutes = () => {

    return (
        <Routes>
            <Route
                path="*"
                element={
                    <ProtectedRoute>
                        <UserAccount/>
                    </ProtectedRoute>
                }
            />

            <Route
                path="/client-registration"
                element={
                    <ProtectedRoute>
                        <ClientRegistration/>
                    </ProtectedRoute>
                }
            />

            <Route path="/client-authentication" element={<ClientUserAuthentication/>}/>
            <Route path="/confirmation" element={<ClientUserConfirmation/>}/>
            <Route path="/user-authentication" element={<UserAuthentication/>}/>
        </Routes>
    )


}

export default BodyRoutes