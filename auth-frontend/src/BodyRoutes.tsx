import {createBrowserRouter, RouterProvider} from "react-router-dom";
import {UserAuthorization} from "./authorization/UserAuthorization";


const router = createBrowserRouter([
    {path: "/authorize", element: <UserAuthorization/>},
],);

export const BodyRoutes = () => {
    return (
        <RouterProvider router={router}/>
    );
}

/*
const BodyRoutes = () => {
    return (
        <Routes>
            <Route path="/account"
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
        </Routes>
    )
}
 */