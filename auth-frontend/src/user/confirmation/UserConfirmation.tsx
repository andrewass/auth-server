import {createSearchParams, useNavigate, useSearchParams} from "react-router-dom";

const UserConfirmation = () => {
    const navigate = useNavigate()
    const [searchParams] = useSearchParams();

    const postConfirmationToServer = () => {
        navigate({
            pathname: "/authentication",
            search: `?${createSearchParams(searchParams)}`
        })
    }

    return (
        <>
            <button onClick={() => postConfirmationToServer()}>Confirm</button>
        </>
    )
}

export default UserConfirmation